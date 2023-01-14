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
		log.Println("uploading radical image", radicalName)

		// upload image if any
		radicalSymbolImage, err := c.FormFile("symbol_image")
		if err == nil && radicalName != "" {
			radicalSymbolFile, err := radicalSymbolImage.Open()
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, errors.InvalidRequest("invalid file in symbol_image"))
				return
			}
			defer radicalSymbolFile.Close()
			log.Println("image size:", radicalSymbolImage.Size)

			objectKey, err := uploadFile(c.Request.Context(), api.files, radicalSymbolFile, files.FileInfo{
				Size:      radicalSymbolImage.Size,
				Name:      radicalName,
				Namespace: "radical",
				Kind:      "image",
			})
			if err != nil {
				log.Println("failed")
				log.Println(err)
				c.AbortWithError(err.(*errors.Error).Status, err)
				return
			}

			form := c.Request.Form

			form.Set("symbol", objectKey)
			form.Del("symbol_image")
			c.Request.MultipartForm.Value["symbol"] = []string{objectKey}
			c.Request.MultipartForm.Value["symbol_image"] = []string{}

			log.Println("image uploaded")
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
