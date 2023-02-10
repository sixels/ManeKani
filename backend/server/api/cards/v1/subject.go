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

	ValueImageMeta *cards.ContentMeta   `json:"value_image_meta" form:"value_image_meta" binding:"-"`
	ResourcesMeta  []*cards.ContentMeta `json:"resources_meta" form:"resources_meta" binding:"-"`

	Data struct {
		// shadow ValueImage and Resources from CreateSubjectRequest
		ValueImage struct{} `json:"-" form:"-"`
		Resources  struct{} `json:"-" form:"-"`

		cards.CreateSubjectRequest
	} `json:"data" form:"data" binding:"required"`
}

// type UpdateSubjectForm struct {
// 	ValueImage *multipart.FileHeader    `json:"-" form:"value_image"`
// 	Resources  *[]*multipart.FileHeader `json:"-" form:"resources[]"`

// 	ValueImageMeta *cards.ContentMeta   `json:"value_image_meta,omitempty" form:"value_image_meta"`
// 	ResourcesMeta  *[]cards.ContentMeta `json:"resources_meta,omitempty" form:"resources_meta"`

// 	Data struct {
// 		// shadow ValueImage and Resources from UpdateSubjectRequest
// 		ValueImage struct{} `json:"-" form:"-"`
// 		Resources  struct{} `json:"-" form:"-"`

// 		cards.UpdateSubjectRequest
// 	} `json:"data" form:"data"`
// }

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

		log.Printf("subject create by user: %s\n", userID)

		var form CreateSubjectForm
		if err := c.Bind(&form); err != nil {
			c.Error(err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if status, err :=
			checkResourceOwner(ctx, userID, form.Data.Deck, api.cards.DeckOwner); err != nil {
			c.Error(err)
			c.Status(status)
			return
		}

		var subjectImage *cards.RemoteContent = nil
		var subjectResources *map[string][]cards.RemoteContent = nil

		// upload value image
		if form.ValueImage != nil && form.ValueImageMeta != nil {
			url, err := api.uploadRemoteResource(
				ctx,
				*form.ValueImage,
				fmt.Sprintf("L%d-%s\n", form.Data.Level, form.Data.Name),
				"resource",
			)
			if err != nil {
				c.Error(fmt.Errorf("could not upload the file: %w", err))
				c.Status(http.StatusInternalServerError)
				return
			}

			subjectImage = &cards.RemoteContent{
				URL:         url,
				ContentType: form.ValueImage.Header.Get("Content-Type"),
				Metadata:    form.ValueImageMeta.Metadata,
			}
		}

		// upload resources
		if form.ResourcesMeta != nil {
			resources := map[string][]cards.RemoteContent{}

			for _, resource := range form.ResourcesMeta {
				if resource.Group == nil || resource.Attachment == nil {
					continue
				}

				if (*resource.Attachment >= 0 && *resource.Attachment < len(form.Resources)) &&
					form.Resources[*resource.Attachment] != nil {

					resourceFile := form.Resources[*resource.Attachment]
					url, err := api.uploadRemoteResource(
						ctx,
						*resourceFile,
						fmt.Sprintf("%s-%s\n", form.Data.Name, *resource.Group),
						"resource",
					)
					if err != nil {
						c.Error(fmt.Errorf("could not upload the file: %w", err))
						c.Status(http.StatusInternalServerError)
						return
					}

					resourceContents, ok := resources[*resource.Group]
					if !ok {
						resourceContents = []cards.RemoteContent{}
					}
					resourceContents = append(resourceContents, cards.RemoteContent{
						URL:         url,
						ContentType: resourceFile.Header.Get("Content-Type"),
						Metadata:    resource.Metadata,
					})

					resources[*resource.Group] = resourceContents

				}
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
			c.Error(err)
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
// @Success 200 {array} cards.PartialSubjectResponse
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

func (api *CardsApi) uploadRemoteResource(ctx context.Context, file multipart.FileHeader, contentName string, contentKind string) (string, error) {
	contentType := file.Header.Get("Content-Type")
	fileHandle, err := file.Open()
	if err != nil {
		return "", err
	}

	log.Printf("file content-type: %q\n", contentType)

	return uploadFile(ctx, api.files, fileHandle, files.FileInfo{
		Size:        file.Size,
		Name:        fmt.Sprintf("%s-%s", ulid.Make(), contentName),
		Namespace:   "cards",
		Kind:        contentKind,
		ContentType: contentType,
	})

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
