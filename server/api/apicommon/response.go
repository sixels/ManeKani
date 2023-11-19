package apicommon

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type APIResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response creates an APIResponse with the given data
func Response[T any](code int, data T) APIResponse[T] {
	return APIResponse[T]{
		Code: code,
		Data: data,
	}
}

// Error creates an APIResponse with the given error
func Error(code int, err error) APIError {
	return APIError{
		Code:    code,
		Message: err.Error(),
	}
}

func Respond[T any](c echo.Context, res APIResponse[T]) error {
	return c.JSON(res.Code, res)
}

func (e APIError) Error() string {
	return fmt.Sprintf("%v", echo.NewHTTPError(e.Code, e.Message).Message)
}
