package cards

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sixels/manekani/server/api/apicommon"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	domain "github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/core/domain/files"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/server/api/cards/util"
)

// GetSubjects gets all subjects
func (a *CardsApiV1) GetSubjects(c echo.Context, params GetSubjectsParams) error {
	filters := new(domain.QueryManySubjectsRequest)
	if err := c.Bind(filters); err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	subjects, err := a.Cards.AllSubjects(ctx, *filters)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusOK, subjects))
}

// CreateSubject creates a subject
func (a *CardsApiV1) CreateSubject(c echo.Context) error {
	log.Info("creating subject")
	ctx := c.Request().Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	//form, _ := c.FormParams()
	//dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "form"})
	//var req CreateSubjectMultipartRequestBody
	//if err := dec.Decode(form, &req); err != nil {
	//	log.Error("create-subject bind error: %w", err)
	//	return apicommon.Error(http.StatusBadRequest, err)
	//}
	//log.Debugf("%#v", req)

	//var req CreateSubjectMultipartRequestBody
	//if err := render.DecodeForm(c.Request().Body, &req); err != nil {
	//	log.Error("create-subject bind error: %w", err)
	//	return apicommon.Error(http.StatusBadRequest, err)
	//}
	//log.Debugf("%#v", req)

	//dec := form.NewDecoder(c.Request().Body)
	//params, err := c.FormParams()
	//var req CreateSubjectMultipartRequestBody
	//if err := dec.DecodeValues(&req, params); err != nil {
	//	log.Error("create-subject bind error: %w", err)
	//	return apicommon.Error(http.StatusBadRequest, err)
	//}
	//log.Debugf("%#v", req)

	var req CreateSubjectMultipartRequestBody
	if err := c.Bind(&req); err != nil {
		log.Error("create-subject bind error: %w", err)
		return apicommon.Error(http.StatusBadRequest, err)
	}
	log.Debugf("%#v", req)

	// upload value image
	subjectImage, err := uploadFile(ctx, a.Cards.FilesRepo, uploadFileReq{
		File: req.ValueImage,
		Kind: "value",
		Name: req.Slug,
	})
	if err != nil {
		log.Error("could not upload the subject image: %w", err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	// upload resources
	var (
		subjectResources = make([]domain.Resource, 0)
	)
	if req.Resources != nil {
		for i, resource := range *req.Resources {
			resourcePath, err := uploadFile(ctx, a.Cards.FilesRepo, uploadFileReq{
				File: &resource,
				Kind: fmt.Sprintf("resource-%d", i),
				Name: req.Slug,
			})
			if err != nil {
				log.Error("could not upload the subject resource: %w", err)
				return apicommon.Error(http.StatusInternalServerError, err)
			}

			var metas map[string]string
			if req.ResourcesMeta != nil && i < len(req.ResourcesMeta.List) {
				metas = req.ResourcesMeta.List[i]
			}

			if resourcePath != nil {
				subjectResources = append(subjectResources, domain.Resource{
					URL:      *resourcePath,
					Metadata: metas,
				})
			}
		}
	}

	var studyData []domain.StudyData
	if req.StudyData != nil {
		studyData = make([]domain.StudyData, 0)
		for i, sd := range req.StudyData.List {
			studyData[i] = domainFromStudyData(sd)
		}
	}

	var (
		Dependencies []uuid.UUID
		Dependents   []uuid.UUID
		Similars     []uuid.UUID
	)
	if req.Dependencies != nil {
		Dependencies = req.Dependencies.List
	}
	if req.Dependents != nil {
		for _, dep := range *req.Dependents {
			Dependents = append(Dependents, uuid.MustParse(dep))
		}
	}
	if req.Similar != nil {
		Similars = req.Similar.List
	}

	subj := domain.CreateSubjectRequest{
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
		Dependencies:        Dependencies,
		Dependents:          Dependents,
		Similars:            Similars,
		Deck:                req.Deck,
	}
	subj.ValueImage = subjectImage
	subj.Resources = &subjectResources

	created, err := a.Cards.CreateSubject(ctx, userID, subj)
	if err != nil {
		log.Errorf("could not create the subject: %w", err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	log.Infof("subject created by user: %s\n", userID)
	return apicommon.Respond(c, apicommon.Response(http.StatusCreated, created))
}

// DeleteSubject deletes a subject
func (a *CardsApiV1) DeleteSubject(c echo.Context, id string) error {
	ctx := c.Request().Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	subjectID, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusNotFound, errors.NotFound("subject not found"))
	}

	if err := a.Cards.DeleteSubject(ctx, subjectID, userID); err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response[any](http.StatusOK, nil))
}

// GetSubject gets a subject
func (a *CardsApiV1) GetSubject(c echo.Context, id string) error {
	subjectID, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusNotFound, errors.NotFound("subject not found"))
	}

	queried, err := a.Cards.QuerySubject(c.Request().Context(), subjectID)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}

	return apicommon.Respond(c, apicommon.Response(http.StatusOK, queried))
}

// UpdateSubject updates a subject
func (a *CardsApiV1) UpdateSubject(c echo.Context, id string) error {
	ctx := c.Request().Context()

	userID, err := util.CtxUserID(c)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusUnauthorized, err)
	}

	subjectID, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusNotFound, errors.NotFound("subject not found"))
	}

	var subject UpdateSubjectJSONRequestBody
	if err := c.Bind(&subject); err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusBadRequest, err)
	}

	subj := domain.UpdateSubjectRequest{
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
	updated, err := a.Cards.UpdateSubject(ctx, subjectID, userID, subj)
	if err != nil {
		log.Error(err)
		return apicommon.Error(http.StatusInternalServerError, err)
	}
	return apicommon.Respond(c, apicommon.Response(http.StatusOK, updated))
}

func domainFromStudyData(sd SubjectStudyData) domain.StudyData {
	items := make([]domain.StudyItem, len(sd.Items))
	for i, si := range sd.Items {
		items[i] = domainFromStudyItem(si)
	}
	return domain.StudyData{
		Kind:     sd.Kind,
		Items:    items,
		Mnemonic: sd.Mnemonic,
	}
}
func domainFromStudyItem(si SubjectStudyItem) domain.StudyItem {
	return domain.StudyItem{
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
