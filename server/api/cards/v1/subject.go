package cards

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	domain_cards "github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/core/domain/files"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/server/api/cards/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Query all subjects
// (GET /api/v1/subjects)
func (api *CardsApiV1) GetSubjects(c *gin.Context, params GetSubjectsParams) {
	filters := new(domain_cards.QueryManySubjectsRequest)
	if err := c.BindQuery(filters); err != nil {
		c.Error(err)
		return
	}

	ctx := c.Request.Context()
	subjects, err := api.Cards.AllSubjects(ctx, *filters)
	if err != nil {
		c.Error(err)
		c.JSON(err.(*errors.Error).Status, err)
	} else {
		c.JSON(http.StatusOK, subjects)
	}
}

// Create a new subject
// (POST /api/v1/subjects)
func (api *CardsApiV1) CreateSubject(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	println(string(body))

	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	var req CreateSubjectMultipartRequestBody
	if err := c.Bind(&req); err != nil {
		c.Error(fmt.Errorf("create-subject bind error: %w", err))
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// upload value image
	subjectImage, err := uploadFile(ctx, api.Cards.FilesRepo, uploadFileReq{
		File: req.ValueImage,
		Kind: "value",
		Name: req.Slug,
	})
	if err != nil {
		c.Error(fmt.Errorf("could not upload the subject image: %w", err))
		c.Status(http.StatusInternalServerError)
		return
	}

	// upload resources
	var (
		subjectResources = make([]domain_cards.Resource, 0)
	)
	if req.Resource != nil {
		for i, resource := range *req.Resource {
			resourcePath, err := uploadFile(ctx, api.Cards.FilesRepo, uploadFileReq{
				File: &resource,
				Kind: fmt.Sprintf("resource-%d", i),
				Name: req.Slug,
			})
			if err != nil {
				c.Error(fmt.Errorf("could not upload the subject resource: %w", err))
				c.Status(http.StatusInternalServerError)
				return
			}

			var metas map[string]string
			if req.ResourcesMeta != nil && i < len(*req.ResourcesMeta) {
				metas = (*req.ResourcesMeta)[i]
			}

			if resourcePath != nil {
				subjectResources = append(subjectResources, domain_cards.Resource{
					URL:      *resourcePath,
					Metadata: metas,
				})
			}
		}
	}

	var studyData []domain_cards.StudyData
	if req.StudyData != nil {
		studyData = make([]domain_cards.StudyData, 0)
		for i, sd := range *req.StudyData {
			studyData[i] = domainFromStudyData(sd)
		}
	}

	subj := domain_cards.CreateSubjectRequest{
		Kind:                req.Kind,
		Level:               req.Level,
		Name:                req.Name,
		Value:               req.Value,
		ValueImage:          subjectImage,
		Slug:                req.Slug,
		Priority:            req.Priority,
		StudyData:           studyData,
		Resources:           &subjectResources,
		AdditionalStudyData: req.AdditionalStudyData,
		Dependencies:        omitEmpty(req.Dependencies),
		Dependents:          omitEmpty(req.Dependents),
		Similars:            omitEmpty(req.Similars),
		Deck:                req.Deck,
	}
	subj.ValueImage = subjectImage
	subj.Resources = &subjectResources

	created, err := api.Cards.CreateSubject(ctx, userID, subj)
	if err != nil {
		c.Error(fmt.Errorf("could not create the subject: %w", err))
		c.Status(http.StatusInternalServerError)
		return
	}

	log.Printf("subject created by user: %s\n", userID)
	c.JSON(http.StatusCreated, created)
}

// Delete a subject
// (DELETE /api/v1/subjects/{id})
func (api *CardsApiV1) DeleteSubject(c *gin.Context, id string) {
	ctx := c.Request.Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	subjectID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			errors.NotFound("subject not found"))
		return
	}

	if err := api.Cards.DeleteSubject(ctx, subjectID, userID); err != nil {
		c.Error(err)
		c.JSON(err.(*errors.Error).Status, err)
	} else {
		c.Status(http.StatusNoContent)
	}
}

// Query a subject
// (GET /api/v1/subjects/{id})
func (api *CardsApiV1) GetSubject(c *gin.Context, id string) {
	subjectID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			errors.NotFound("subject not found"))
		return
	}

	queried, err := api.Cards.QuerySubject(c.Request.Context(), subjectID)
	if err != nil {
		c.Error(err)
		c.JSON(err.(*errors.Error).Status, err)
	} else {
		c.JSON(http.StatusOK, queried)
	}
}

// Update a subject
// (PATCH /api/v1/subjects/{id})
func (api *CardsApiV1) UpdateSubject(c *gin.Context, id string) {
	ctx := c.Request.Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusUnauthorized)
		return
	}

	subjectID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			errors.NotFound("subject not found"))
		return
	}

	var subject UpdateSubjectJSONRequestBody
	if err := c.Bind(&subject); err != nil {
		c.Error(err)
		return
	}

	subj := domain_cards.UpdateSubjectRequest{
		Kind:       subject.Kind,
		Level:      subject.Level,
		Name:       subject.Name,
		Value:      subject.Value,
		ValueImage: subject.ValueImage,
		Slug:       subject.Slug,
		Priority:   subject.Priority,
		// StudyData:           subject.StudyData,
		// Resources:           subject.Resources,
		AdditionalStudyData: subject.AdditionalStudyData,
		Dependencies:        subject.Dependencies,
		Dependents:          subject.Dependents,
		Similars:            subject.Similars,
	}
	updated, err := api.Cards.UpdateSubject(ctx, subjectID, userID, subj)
	if err != nil {
		c.Error(err)
		c.JSON(err.(*errors.Error).Status, err)
	} else {
		c.JSON(http.StatusOK, updated)
	}
}

func domainFromStudyData(sd SubjectStudyData) domain_cards.StudyData {
	items := make([]domain_cards.StudyItem, len(sd.Items))
	for i, si := range sd.Items {
		items[i] = domainFromStudyItem(si)
	}
	return domain_cards.StudyData{
		Kind:     sd.Kind,
		Items:    items,
		Mnemonic: sd.Mnemonic,
	}
}
func domainFromStudyItem(si SubjectStudyItem) domain_cards.StudyItem {
	return domain_cards.StudyItem{
		Value:         si.Value,
		IsPrimary:     si.IsPrimary,
		IsValidAnswer: si.IsValidAnswer,
		IsHidden:      si.IsHidden,
		Category:      si.Category,
		Resource:      si.Resource,
	}
}
func omitEmpty[T any](t *T) T {
	var empty T
	if t == nil {
		return empty
	}
	return *t
}

// func (api *CardsApiV1) CreateSubject() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()

// 		userID, err := util.CtxUserID(c)
// 		if err != nil {
// 			c.Error(err)
// 			c.Status(http.StatusUnauthorized)
// 			return
// 		}

// 		body, _ := io.ReadAll(c.Request.Body)
// 		println(string(body))

// 		c.Request.Body = io.NopCloser(bytes.NewReader(body))

// 		var req CreateSubjectAPIRequest
// 		if err := c.Bind(&req); err != nil {
// 			c.Error(fmt.Errorf("create-subject bind error: %w", err))
// 			c.String(http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		// upload value image
// 		subjectImage, err := uploadFile(ctx, api.Cards.FilesRepo, uploadFileReq{
// 			File: req.ValueImage,
// 			Kind: "value",
// 			Name: req.Slug,
// 		})
// 		if err != nil {
// 			c.Error(fmt.Errorf("could not upload the subject image: %w", err))
// 			c.Status(http.StatusInternalServerError)
// 			return
// 		}

// 		// upload resources
// 		var (
// 			subjectResources = make([]domain_cards.Resource, 0)
// 		)
// 		for i, metas := range req.ResourcesMeta {
// 			resourcePath, err := uploadFile(ctx, api.Cards.FilesRepo, uploadFileReq{
// 				File: req.Resources[i],
// 				Kind: fmt.Sprintf("resource-%d", i),
// 				Name: req.Slug,
// 			})
// 			if err != nil {
// 				c.Error(fmt.Errorf("could not upload the subject resource: %w", err))
// 				c.Status(http.StatusInternalServerError)
// 				return
// 			}
// 			if resourcePath != nil {
// 				subjectResources = append(subjectResources, domain_cards.Resource{
// 					URL:      *resourcePath,
// 					Metadata: metas,
// 				})
// 			}
// 		}

// 		subj := req.CreateSubjectRequest
// 		subj.ValueImage = subjectImage
// 		subj.Resources = &subjectResources

// 		created, err := api.Cards.CreateSubject(ctx, userID, subj)
// 		if err != nil {
// 			c.Error(fmt.Errorf("could not create the subject: %w", err))
// 			c.Status(http.StatusInternalServerError)
// 			return
// 		}

// 		log.Printf("subject created by user: %s\n", userID)
// 		c.JSON(http.StatusCreated, created)
// 	}
// }

// func (api *CardsApiV1) QuerySubject() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id, err := uuid.Parse(c.Param("id"))
// 		if err != nil {
// 			c.JSON(
// 				http.StatusNotFound,
// 				errors.NotFound("subject not found"))
// 			return
// 		}

// 		queried, err := api.Cards.QuerySubject(c.Request.Context(), id)
// 		if err != nil {
// 			c.Error(err)
// 			c.JSON(err.(*errors.Error).Status, err)
// 		} else {
// 			c.JSON(http.StatusOK, queried)
// 		}
// 	}
// }

// func (api *CardsApiV1) UpdateSubject() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()

// 		userID, err := util.CtxUserID(c)
// 		if err != nil {
// 			c.Error(err)
// 			c.Status(http.StatusUnauthorized)
// 			return
// 		}

// 		subjectID, err := uuid.Parse(c.Param("id"))
// 		if err != nil {
// 			c.JSON(
// 				http.StatusNotFound,
// 				errors.NotFound("subject not found"))
// 			return
// 		}

// 		var subject domain_cards.UpdateSubjectRequest
// 		if err := c.Bind(&subject); err != nil {
// 			c.Error(err)
// 			return
// 		}

// 		updated, err := api.Cards.UpdateSubject(ctx, subjectID, userID, subject)
// 		if err != nil {
// 			c.Error(err)
// 			c.JSON(err.(*errors.Error).Status, err)
// 		} else {
// 			c.JSON(http.StatusOK, updated)
// 		}
// 	}

// }

// func (api *CardsApiV1) DeleteSubject() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()

// 		userID, err := util.CtxUserID(c)
// 		if err != nil {
// 			c.Error(err)
// 			c.Status(http.StatusUnauthorized)
// 			return
// 		}

// 		subjectID, err := uuid.Parse(c.Param("id"))
// 		if err != nil {
// 			c.JSON(
// 				http.StatusNotFound,
// 				errors.NotFound("subject not found"))
// 			return
// 		}

// 		if err := api.Cards.DeleteSubject(ctx, subjectID, userID); err != nil {
// 			c.Error(err)
// 			c.JSON(err.(*errors.Error).Status, err)
// 		} else {
// 			c.Status(http.StatusNoContent)
// 		}
// 	}
// }

// func (api *CardsApiV1) AllSubjects() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		filters := new(domain_cards.QueryManySubjectsRequest)
// 		if err := c.BindQuery(filters); err != nil {
// 			c.Error(err)
// 			return
// 		}

// 		ctx := c.Request.Context()
// 		subjects, err := api.Cards.AllSubjects(ctx, *filters)
// 		if err != nil {
// 			c.Error(err)
// 			c.JSON(err.(*errors.Error).Status, err)
// 		} else {
// 			c.JSON(http.StatusOK, subjects)
// 		}
// 	}
// }

// func (api *CardsApiV1) startRemoteResourceUpload(
// 	ctx context.Context,
// 	meta *ContentMeta,
// 	name string,
// 	kind string,
// ) (remoteContent *cards.RemoteContent, uploadURL files.UploadURL, err error) {
// 	namespace := "cards"
// 	upname := strings.Trim(fmt.Sprintf("%s-%s", ulid.Make(), name), " \n\t\r")

// 	var contentType string = "application/octet-stream"
// 	if meta.ContentType != nil {
// 		contentType = *meta.ContentType
// 	}

// 	log.Printf("file content-type: %q\n", contentType)

// 	key, uploadURL, err := startFileUpload(ctx, api.Files, files.FileInfo{
// 		Name:        upname,
// 		Namespace:   namespace,
// 		Kind:        kind,
// 		ContentType: contentType,
// 	})
// 	if err != nil {
// 		return nil, files.UploadURL{}, err
// 	}

// 	if meta.Group != nil {
// 		uploadURL.Resource = *meta.Group
// 	}

// 	return &cards.RemoteContent{
// 		// TODO: return the full url to the resource instead of its storage key
// 		// e.g: https://files.manekani.com/{key}
// 		URL:         key,
// 		ContentType: contentType,
// 		Metadata:    meta.Metadata,
// 	}, uploadURL, nil
// }

// func startFileUpload(ctx context.Context, filesService *files_service.FilesRepository, info files.FileInfo) (objectKey string, uploadURL files.UploadURL, err error) {
// 	return filesService.UploadFileURL(ctx, info)
// }

type uploadFileReq struct {
	File *types.File
	Kind string
	Name string
}

func uploadFile(ctx context.Context, fileService ports.FilesRepository, req uploadFileReq) (*string, error) {
	if req.File == nil {
		return nil, nil
	}

	fileHandle, err := req.File.Reader()
	if err != nil {
		return nil, fmt.Errorf("could not open the form file: %w", err)
	}

	uploadedURL, err := fileService.CreateFile(ctx, files.CreateFileRequest{
		Handle: fileHandle,
		FileInfo: files.FileInfo{
			Size:      req.File.FileSize(),
			Namespace: "subject",
			Kind:      req.Kind,
			Name:      req.Name,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("could not upload the file: %w", err)
	}
	return &uploadedURL, nil
}
