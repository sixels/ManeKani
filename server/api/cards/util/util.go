package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CtxUserID(c *gin.Context) (string, error) {
	ctxUser, ok := c.Get("userID")
	if !ok {
		return "", fmt.Errorf("user is not set in the context")
	}
	return ctxUser.(string), nil
}

func Ptr[T any](c T) *T {
	return &c
}
