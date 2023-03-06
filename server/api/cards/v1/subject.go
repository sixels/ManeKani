package cards

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/errors"
	"github.com/sixels/manekani/core/domain/files"
	"github.com/sixels/manekani/server/api/cards/util"
	files_service "github.com/sixels/manekani/services/files"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type ContentMeta struct {
	Group       *string         `json:"group,omitempty"`
	ContentType *string         `json:"content_type,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

type CreateSubjectAPIRequest struct {
	ValueImageMeta *ContentMeta  `json:"value_image_meta,omitempty" form:"value_image_meta" binding:"-"`
	ResourcesMeta  []ContentMeta `json:"resources_meta,omitempty" form:"resources_meta" binding:"-"`

	// shadow ValueImage and Resources from CreateSubjectRequest
	ValueImage struct{} `json:"-" form:"-"`
	Resources  struct{} `json:"-" form:"-"`

	cards.CreateSubjectRequest
}

type CreateSubjectAPIResponse struct {
	URLs Urls `json:"urls"`

	cards.Subject
}
type Urls struct {
	ValueImageURL files.UploadURL   `json:"value_image_url"`
	ResourcesURL  []files.UploadURL `json:"resources_url"`
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
		log.Printf("create subject request received\n")

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
		var (
			subjectImage       *cards.RemoteContent
			subjectImageUpload files.UploadURL
		)
		if req.ValueImageMeta != nil {
			subjectImage, subjectImageUpload, err = api.startRemoteResourceUpload(
				ctx,
				req.ValueImageMeta,
				fmt.Sprintf("L%d-%s\n", req.Level, req.Slug),
				"value",
			)
			if err != nil {
				c.Error(fmt.Errorf("could not upload the file: %w", err))
				c.Status(http.StatusInternalServerError)
				return
			}
		}

		// upload resources
		var (
			subjectResources       map[string][]cards.RemoteContent
			subjectResourcesUpload []files.UploadURL
		)
		if req.ResourcesMeta != nil {
			for _, resourceMeta := range req.ResourcesMeta {
				resource, uploadURL, err := api.startRemoteResourceUpload(
					ctx,
					&resourceMeta,
					fmt.Sprintf("%s-%s\n", req.Name, *resourceMeta.Group),
					"resource",
				)
				if err != nil {
					c.Error(fmt.Errorf("could not upload the file: %w", err))
					c.Status(http.StatusInternalServerError)
					return
				}

				resourceContents, ok := subjectResources[*resourceMeta.Group]
				if !ok {
					resourceContents = []cards.RemoteContent{}
				}
				resourceContents = append(resourceContents, *resource)
				subjectResourcesUpload = append(subjectResourcesUpload, uploadURL)

				subjectResources[*resourceMeta.Group] = resourceContents
			}
		}

		subj := req.CreateSubjectRequest
		subj.ValueImage = subjectImage
		subj.Resources = &subjectResources

		created, err := api.Cards.CreateSubject(ctx, userID, subj)
		if err != nil {
			c.Error(fmt.Errorf("could not create the subject: %w", err))
			c.JSON(err.(*errors.Error).Status, err)
			return
		}

		log.Printf("subject created by user: %s\n", userID)
		c.JSON(http.StatusCreated, CreateSubjectAPIResponse{
			Subject: *created,
			URLs: Urls{
				ValueImageURL: subjectImageUpload,
				ResourcesURL:  subjectResourcesUpload,
			},
		})

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

func (api *CardsApiV1) startRemoteResourceUpload(
	ctx context.Context,
	meta *ContentMeta,
	name string,
	kind string,
) (remoteContent *cards.RemoteContent, uploadURL files.UploadURL, err error) {

	namespace := "cards"
	upname := strings.Trim(fmt.Sprintf("%s-%s", ulid.Make(), name), " \n\t\r")

	var contentType string
	if meta.ContentType != nil {
		contentType = *meta.ContentType
	} else {
		contentType = "application/octet-stream"
	}

	log.Printf("file content-type: %q\n", contentType)

	key, uploadURL, err := startFileUpload(ctx, api.Files, files.FileInfo{
		Name:        upname,
		Namespace:   namespace,
		Kind:        kind,
		ContentType: contentType,
	})
	if err != nil {
		return nil, files.UploadURL{}, err
	}

	if meta.Group != nil {
		uploadURL.Resource = *meta.Group
	}

	return &cards.RemoteContent{
		// TODO: return the full url to the resource instead of its storage key
		// e.g: https://files.manekani.com/{key}
		URL:         key,
		ContentType: contentType,
		Metadata:    meta.Metadata,
	}, uploadURL, nil
}

func startFileUpload(ctx context.Context, filesService *files_service.FilesRepository, info files.FileInfo) (objectKey string, uploadURL files.UploadURL, err error) {
	return filesService.UploadFileURL(ctx, info)
}

func checkResourceOwner[T any](
	ctx context.Context,
	userID string, resourceID T,
	queryOwner func(context.Context, T) (string, error),
) (status int, err error) {
	if owner, err := queryOwner(ctx, resourceID); err != nil || owner != userID {
		if owner != userID {
			log.Printf("%s is not %s\n", userID, owner)
			err = fmt.Errorf("%q is not the owner of '%v'", userID, resourceID)
			status = http.StatusForbidden
		} else {
			err = fmt.Errorf(
				"could not check wether the specified user owns the resource: %w", err,
			)
			status = http.StatusNotFound
		}
		return status, err
	}
	return 0, nil
}
