package v1

import (
	"context"
	"io"
	"log"
	"strings"

	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/core/domain/files"
	filesService "sixels.io/manekani/services/files"

	"github.com/labstack/echo/v4"
)

func (api *CardsApi) UploadRadicalImage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		radicalName := strings.TrimSpace(c.FormValue("name"))

		// upload image if any
		radicalSymbolImage, err := c.FormFile("symbol_image")
		if err == nil && radicalName != "" {
			log.Println("middleware: ", radicalName)
			radicalSymbolFile, err := radicalSymbolImage.Open()
			if err != nil {
				return errors.InvalidRequest("invalid file in symbol_image")
			}
			defer radicalSymbolFile.Close()

			storedKey, err := uploadFile(c.Request().Context(), api.files, radicalSymbolFile, files.FileInfo{
				Size:      radicalSymbolImage.Size,
				Name:      radicalName,
				Namespace: "radical",
				Kind:      "image",
			})
			if err != nil {
				return err
			}

			form := c.Request().Form
			form.Set("symbol", storedKey)
			form.Del("symbol_image")
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}

func uploadFile(ctx context.Context, filesService *filesService.FilesRepository, f io.Reader, info files.FileInfo) (string, error) {
	return filesService.CreateFile(ctx, files.CreateFileRequest{
		FileInfo: info,
		Handle:   f,
	})
}
