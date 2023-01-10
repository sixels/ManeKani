package v1

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/core/domain/files"
	filesService "sixels.io/manekani/services/files"

	"github.com/gin-gonic/gin"
)

func (api *CardsApi) UploadRadicalImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		radicalName := strings.TrimSpace(c.Request.FormValue("name"))

		// upload image if any
		radicalSymbolImage, err := c.FormFile("symbol_image")
		if err == nil && radicalName != "" {
			log.Println("middleware: ", radicalName)
			radicalSymbolFile, err := radicalSymbolImage.Open()
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, errors.InvalidRequest("invalid file in symbol_image"))
				return
			}
			defer radicalSymbolFile.Close()

			objectKey, err := uploadFile(c.Request.Context(), api.files, radicalSymbolFile, files.FileInfo{
				Size:      radicalSymbolImage.Size,
				Name:      radicalName,
				Namespace: "radical",
				Kind:      "image",
			})
			if err != nil {
				c.AbortWithError(err.(*errors.Error).Status, err)
				return
			}

			form := c.Request.Form
			form.Set("symbol", objectKey)
			form.Del("symbol_image")
		}

		c.Next()
	}
}

func uploadFile(ctx context.Context, filesService *filesService.FilesRepository, f io.Reader, info files.FileInfo) (string, error) {
	return filesService.CreateFile(ctx, files.CreateFileRequest{
		FileInfo: info,
		Handle:   f,
	})
}
