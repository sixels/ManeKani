package files

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
)

func (api *FilesApi) QueryFile(c echo.Context) error {
	key := strings.Join(c.ParamValues(), "/")
	object, err := api.files.QueryFile(c.Request().Context(), key[1:])
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusNotFound)
	}

	defer object.Close()
	return c.Stream(http.StatusOK, object.ContentType, object)
}
