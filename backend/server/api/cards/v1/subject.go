package cards

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/core/domain/files"
	files_service "sixels.io/manekani/services/files"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type CreateSubjectForm struct {
	ValueImage *multipart.FileHeader   `json:"-" form:"value_image" binding:"-"`
	Resources  []*multipart.FileHeader `json:"-" form:"resources[]" binding:"-"`

	ValueImageMeta *cards.ContentMeta  `json:"value_image_meta" form:"value_image_meta" binding:"-"`
	ResourcesMeta  []cards.ContentMeta `json:"resource_meta[]" form:"resource_meta[]" binding:"-"`

	Data struct {
		// shadow ValueImage and Resources from CreateSubjectRequest
		ValueImage struct{} `json:"-" form:"-"`
		Resources  struct{} `json:"-" form:"-"`

		cards.CreateSubjectRequest
	} `json:"data" form:"data" binding:"required"`
}

// CreateSubject godoc
// @Id post-subject-create
// @Summary Create a new subject
// @Description Creates a subject with the given values
// @Tags cards, subject
// @Accept mpfd
// @Produce json
// @Param subject body CreateSubjectForm true "The subject to be created"
// @Success 201 {object} cards.Subject
// @Router /api/v1/subject [post]
func (api *CardsApi) CreateSubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userIDCtx, ok := c.Get("userID")
		if !ok {
			c.Error(fmt.Errorf("userID is not set"))
			c.Status(http.StatusUnauthorized)
		}
		userID := userIDCtx.(string)

		log.Printf("subject created by user: %s\n", userID)

		var form CreateSubjectForm
		if err := c.Bind(&form); err != nil {
			c.Error(fmt.Errorf("create-subject bind error: %w", err))
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if status, err :=
			checkResourceOwner(ctx, userID, form.Data.Deck, api.cards.DeckOwner); err != nil {
			c.Error(fmt.Errorf("authorization error: %w", err))
			c.Status(status)
			return
		}

		// check if the subject violates the uniqueness constraint inside the deck.
		// NOTE: this is only done now to prevent uploading resources of an invalid
		// card. A better approach would be:
		// TODO: return a pre-signed url to the client
		// so they can upload resources after the subject is successfully created.
		// That also removes the need of using `multipart/form-data`.
		exists, err := api.cards.SubjectExists(
			ctx, form.Data.Kind, form.Data.Name, form.Data.Deck,
		)
		if err != nil {
			c.Error(fmt.Errorf("could not check if subject exists: %w", err))
			c.Status(http.StatusInternalServerError)
			return
		}
		if exists {
			c.Status(http.StatusConflict)
			return
		}

		var subjectImage *cards.RemoteContent = nil
		var subjectResources *map[string][]cards.RemoteContent = nil

		// upload value image
		if form.ValueImage != nil && form.ValueImageMeta != nil {
			resource, err := api.uploadRemoteResource(
				ctx,
				*form.ValueImage,
				form.ValueImageMeta,
				fmt.Sprintf("L%d-%s\n", form.Data.Level, form.Data.Slug),
				"value",
			)
			if err != nil {
				c.Error(fmt.Errorf("could not upload the file: %w", err))
				c.Status(http.StatusInternalServerError)
				return
			}

			subjectImage = resource
		}

		// upload resources
		if form.ResourcesMeta != nil {
			resources := map[string][]cards.RemoteContent{}

			for _, resourceMeta := range form.ResourcesMeta {
				if (resourceMeta.Group == nil || resourceMeta.Attachment == nil) ||
					(*resourceMeta.Attachment < 0 || *resourceMeta.Attachment >= len(form.Resources)) ||
					(form.Resources[*resourceMeta.Attachment] == nil) {
					continue
				}

				resourceFile := form.Resources[*resourceMeta.Attachment]
				resource, err := api.uploadRemoteResource(
					ctx,
					*resourceFile,
					&resourceMeta,
					fmt.Sprintf("%s-%s\n", form.Data.Name, *resourceMeta.Group),
					"resource",
				)
				if err != nil {
					c.Error(fmt.Errorf("could not upload the file: %w", err))
					c.Status(http.StatusInternalServerError)
					return
				}

				resourceContents, ok := resources[*resourceMeta.Group]
				if !ok {
					resourceContents = []cards.RemoteContent{}
				}
				resourceContents = append(resourceContents, *resource)

				resources[*resourceMeta.Group] = resourceContents

			}

			if len(resources) > 0 {
				subjectResources = &resources
			}
		}

		subj := form.Data.CreateSubjectRequest
		subj.ValueImage = subjectImage
		subj.Resources = subjectResources

		created, err := api.cards.CreateSubject(ctx, userID, subj)
		if err != nil {
			c.Error(fmt.Errorf("could not create the subject: %w", err))
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusCreated, created)
		}
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
func (api *CardsApi) QuerySubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				errors.NotFound("subject not found"))
			return
		}

		ctx := c.Request.Context()
		queried, err := api.cards.QuerySubject(ctx, id)
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
func (api *CardsApi) UpdateSubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userIDCtx, ok := c.Get("userID")
		if !ok {
			c.Error(fmt.Errorf("userID is not set"))
			c.Status(http.StatusUnauthorized)
		}
		userID := userIDCtx.(string)

		subjectID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				errors.NotFound("subject not found"))
			return
		}

		if status, err :=
			checkResourceOwner(ctx, userID, subjectID, api.cards.SubjectOwner); err != nil {
			c.Error(err)
			c.Status(status)
			return
		}

		subject := new(cards.UpdateSubjectRequest)
		if err := c.Bind(subject); err != nil {
			c.Error(err)
			return
		}

		updated, err := api.cards.UpdateSubject(ctx, subjectID, *subject)
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
func (api *CardsApi) DeleteSubject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userIDCtx, ok := c.Get("userID")
		if !ok {
			c.Error(fmt.Errorf("userID is not set"))
			c.Status(http.StatusUnauthorized)
		}
		userID := userIDCtx.(string)

		subjectID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				errors.NotFound("subject not found"))
			return
		}

		if status, err :=
			checkResourceOwner(ctx, userID, subjectID, api.cards.SubjectOwner); err != nil {
			c.Error(err)
			c.Status(status)
			return
		}

		if err := api.cards.DeleteSubject(ctx, subjectID); err != nil {
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
func (api *CardsApi) AllSubjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		filters := new(cards.QueryManySubjectsRequest)
		if err := c.BindQuery(filters); err != nil {
			c.Error(err)
			return
		}

		ctx := c.Request.Context()
		subjects, err := api.cards.AllSubjects(ctx, *filters)
		if err != nil {
			c.Error(err)
			c.JSON(err.(*errors.Error).Status, err)
		} else {
			c.JSON(http.StatusOK, subjects)
		}
	}
}

func (api *CardsApi) uploadRemoteResource(
	ctx context.Context,
	file multipart.FileHeader,
	meta *cards.ContentMeta,
	name string,
	kind string,
) (*cards.RemoteContent, error) {

	fileHandle, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	namespace := "cards"
	upname := fmt.Sprintf("%s-%s", ulid.Make(), name)

	var contentType string
	if meta.ContentType != nil {
		contentType = *meta.ContentType
	} else {
		contentType = file.Header.Get("content-type")
	}

	log.Printf("file content-type: %q\n", contentType)

	url, err := uploadFile(ctx, api.files, fileHandle, files.FileInfo{
		Size:        file.Size,
		Name:        upname,
		Namespace:   namespace,
		Kind:        kind,
		ContentType: contentType,
	})
	if err != nil {
		return nil, err
	}

	return &cards.RemoteContent{
		URL:         url,
		ContentType: contentType,
		Metadata:    meta.Metadata,
	}, nil
}

func uploadFile(ctx context.Context, filesService *files_service.FilesRepository, f io.Reader, info files.FileInfo) (string, error) {
	return filesService.CreateFile(ctx, files.CreateFileRequest{
		FileInfo: info,
		Handle:   f,
	})
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
