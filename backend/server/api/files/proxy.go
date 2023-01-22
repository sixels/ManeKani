package files

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TODO: swagger comment
func (api *FilesApi) QueryFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramKind := c.Param("kind")
		paramNamespace := c.Param("namespace")
		paramName := c.Param("name")

		key := strings.Join([]string{paramKind, paramNamespace, paramName}, "/")

		object, err := api.files.QueryFile(c.Request.Context(), key)
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
