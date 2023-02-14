package files

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"sixels.io/manekani/services/ent/util"
)

// TODO: swagger comment
func (api *FilesApi) QueryFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := util.MapArray(c.Params, func(params gin.Param) string {
			return params.Value
		})

		key := strings.Join(params, "/")
		object, err := api.files.QueryFile(c.Request.Context(), key[1:])
		if err != nil {
			c.Error(err)
			c.Status(http.StatusNotFound)
			return
		}
		defer object.Close()

		// c.Header("Content-Type", object.ContentType)
		// c.Header("Content-Length", fmt.Sprintf("%d", object.Size))
		// c.Status(http.StatusOK)
		// io.Copy(c.Writer, object)
		c.DataFromReader(
			http.StatusOK, object.Size, object.ContentType, object, map[string]string{})
	}
}
