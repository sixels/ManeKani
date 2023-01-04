package files

import (
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

// TODO: swagger comment
func (api *FilesApi) QueryFile(c echo.Context) error {
	key := strings.Join(c.ParamValues(), "/")

	object, err := api.files.QueryFile(c.Request().Context(), key)
	if err != nil {
		return err
	}
	defer object.Close()

	if _, err := io.Copy(c.Response(), object); err != nil {
		return err
	}
	c.Response().Header().Set("content-type", object.ContentType)

	return nil
}
