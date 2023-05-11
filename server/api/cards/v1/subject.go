package cards

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/core/domain/files"
	"github.com/sixels/manekani/server/api/cards/util"
	files_service "github.com/sixels/manekani/services/files"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateSubjectAPIRequest struct {
	ResourcesMeta []map[string]string `json:"resource_meta[],omitempty" form:"resource_meta[]" binding:"-"`

	// shadow ValueImage and Resources from CreateSubjectRequest
	ValueImage *multipart.FileHeader   `json:"value_image" form:"value_image" binding:"-"`
	Resources  []*multipart.FileHeader `json:"resource[]" form:"-" binding:"-"`

	cards.CreateSubjectRequest
}

// CreateSubject godoc
// @Id post-subject-create
// @Summary Create a new subject
// @Description Creates a subject with the given values
// @Tags cards, subject
// @Accept json
// @Produce json
// @Param subject body CreateSubjectAPIRequest true "The subject to be created"
// @Success 201 {object} CreateSubjectAPIResponse
// @Router /api/v1/subject [post]
func (api *CardsApiV1) CreateSubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		var req CreateSubjectAPIRequest
		if err := c.Bind(&req); err != nil {
			c.Error(fmt.Errorf("create-subject bind error: %w", err))
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// upload value image
		subjectImage, err := uploadFile(ctx, api.Files, uploadFileReq{
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
			subjectResources = make([]cards.Resource, 0)
		)
		for i, metas := range req.ResourcesMeta {
			resourcePath, err := uploadFile(ctx, api.Files, uploadFileReq{
				File: req.Resources[i],
				Kind: fmt.Sprintf("resource-%d", i),
				Name: req.Slug,
			})
			if err != nil {
				c.Error(fmt.Errorf("could not upload the subject resource: %w", err))
				c.Status(http.StatusInternalServerError)
				return
			}
			if resourcePath != nil {
				subjectResources = append(subjectResources, cards.Resource{
					URL:      *resourcePath,
					Metadata: metas,
				})
			}
		}

		subj := req.CreateSubjectRequest
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
}

// QuerySubject godoc
// @Id get-subject-query
// @Summary Query a subject
// @Description Search a subject by its name
// @Tags cards, subject
// @Accept */*
// @Produce json
// @Param name path string true "Subject name"
// @Success 200 {object} cards.Subject
// @Router /api/v1/subject/{name} [get]
func (api *CardsApiV1) QuerySubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				errors.NotFound("subject not found"))
			return
		}

		queried, err := api.Cards.QuerySubject(c.Request.Context(), id)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, queried)
		}
	}
}

// UpdateSubject godoc
// @Id patch-subject-update
// @Summary Update a subject
// @Description Update a subject with the given values
// @Tags cards, subject
// @Accept json
// @Produce json
// @Param name path string true "Subject name"
// @Param subject body cards.UpdateSubjectRequest true "Subject fields to update"
// @Success 200 {object} cards.Subject
// @Router /api/v1/subject/{name} [patch]
func (api *CardsApiV1) UpdateSubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		subjectID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				errors.NotFound("subject not found"))
			return
		}

		var subject cards.UpdateSubjectRequest
		if err := c.Bind(&subject); err != nil {
			c.Error(err)
			return
		}

		updated, err := api.Cards.UpdateSubject(ctx, subjectID, userID, subject)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, updated)
		}
	}

}

// DeleteSubject godoc
// @Id delete-subject-delete
// @Summary Delete a subject
// @Description Delete a subject by its name
// @Tags cards, subject
// @Accept */*
// @Produce json
// @Param name path string true "Subject name"
// @Success 200
// @Router /api/v1/subject/{name} [delete]
func (api *CardsApiV1) DeleteSubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}

		subjectID, err := uuid.Parse(c.Param("id"))
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
}

// AllSubjects godoc
// @Id get-subject-all
// @Summary Query all subjects
// @Description Return a list of all subjects
// @Tags cards, subject
// @Accept */*
// @Produce json
// @Success 200 {array} cards.PartialSubject
// @Router /api/v1/subject [get]
func (api *CardsApiV1) AllSubjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := new(cards.QueryManySubjectsRequest)
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
}

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
	File *multipart.FileHeader
	Kind string
	Name string
}

func uploadFile(ctx context.Context, fileService *files_service.FilesRepository, req uploadFileReq) (*string, error) {
	if req.File == nil {
		return nil, nil
	}

	fileHandle, err := req.File.Open()
	if err != nil {
		return nil, fmt.Errorf("Could not open the form file: %w", err)
	}

	uploadedURL, err := fileService.CreateFile(ctx, files.CreateFileRequest{
		Handle: fileHandle,
		FileInfo: files.FileInfo{
			Size:      req.File.Size,
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
