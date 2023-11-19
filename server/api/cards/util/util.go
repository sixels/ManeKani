package util

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/sixels/manekani/server/auth"
)

func CtxUserID(c echo.Context) (string, error) {
	id, ok := c.Get(string(auth.UserIDContext)).(string)
	if !ok {
		return "", fmt.Errorf("user is not set in the context")
	}
	return id, nil
}

func Ptr[T any](c T) *T {
	return &c
}
