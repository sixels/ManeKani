package cards

import (
	"context"
	"fmt"
	"github.com/sixels/manekani/server/api/apicommon"
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

// GetSubjects gets all subjects
func (api *CardsApiV1) GetSubjects(c *gin.Context, params GetSubjectsParams) {
	filters := new(domain_cards.QueryManySubjectsRequest)
	if err := c.BindQuery(filters); err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
		return
	}

	ctx := c.Request.Context()
	subjects, err := api.Cards.AllSubjects(ctx, *filters)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}

	apicommon.Respond(c, apicommon.Response(http.StatusOK, subjects))
}

// CreateSubject creates a subject
func (api *CardsApiV1) CreateSubject(c *gin.Context) {
	log.Println("creating subject")
	ctx := c.Request.Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
		return
	}

	var req CreateSubjectMultipartRequestBody
	if err := c.ShouldBind(&req); err != nil {
		c.Error(fmt.Errorf("create-subject bind error: %w", err))
		apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
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
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
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
				apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
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
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}

	log.Printf("subject created by user: %s\n", userID)
	apicommon.Respond(c, apicommon.Response(http.StatusCreated, created))
}

// DeleteSubject deletes a subject
func (api *CardsApiV1) DeleteSubject(c *gin.Context, id string) {
	ctx := c.Request.Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
		return
	}

	subjectID, err := uuid.Parse(id)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusNotFound, errors.NotFound("subject not found")))
		return
	}

	if err := api.Cards.DeleteSubject(ctx, subjectID, userID); err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}

	apicommon.Respond(c, apicommon.Response[any](http.StatusOK, nil))
}

// GetSubject gets a subject
func (api *CardsApiV1) GetSubject(c *gin.Context, id string) {
	subjectID, err := uuid.Parse(id)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusNotFound, errors.NotFound("subject not found")))
		return
	}

	queried, err := api.Cards.QuerySubject(c.Request.Context(), subjectID)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}

	apicommon.Respond(c, apicommon.Response(http.StatusOK, queried))
}

// UpdateSubject updates a subject
func (api *CardsApiV1) UpdateSubject(c *gin.Context, id string) {
	ctx := c.Request.Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusUnauthorized, err))
		return
	}

	subjectID, err := uuid.Parse(id)
	if err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusNotFound, errors.NotFound("subject not found")))
		return
	}

	var subject UpdateSubjectJSONRequestBody
	if err := c.Bind(&subject); err != nil {
		c.Error(err)
		apicommon.Respond(c, apicommon.Error(http.StatusBadRequest, err))
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
		apicommon.Respond(c, apicommon.Error(http.StatusInternalServerError, err))
		return
	}
	apicommon.Respond(c, apicommon.Response(http.StatusOK, updated))
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
