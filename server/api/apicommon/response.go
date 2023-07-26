package apicommon

import "github.com/gin-gonic/gin"

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data"`
}

// Response creates an APIResponse with the given data
func Response[T any](code int, data T) APIResponse[T] {
	return APIResponse[T]{
		Code: code,
		Data: data,
	}
}

// Error creates an APIResponse with the given error
func Error(code int, err error) APIResponse[any] {
	return APIResponse[any]{
		Code:    code,
		Message: err.Error(),
	}
}

func Respond[T any](c *gin.Context, res APIResponse[T]) {
	c.JSON(res.Code, res)
}
