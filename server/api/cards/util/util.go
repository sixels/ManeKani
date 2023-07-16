package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/server/auth"
)

func CtxUserID(c *gin.Context) (string, error) {
	ctxUser, ok := c.Get(string(auth.UserIDContext))
	if !ok {
		return "", fmt.Errorf("user is not set in the context")
	}
	return ctxUser.(string), nil
}

func Ptr[T any](c T) *T {
	return &c
}
