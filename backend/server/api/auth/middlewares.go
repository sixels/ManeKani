package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"sixels.io/manekani/server/api/cards/util"
	mkjwt "sixels.io/manekani/services/jwt"
)

func SupertokensMiddleware(c *gin.Context) {
	supertokens.Middleware(http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	c.Abort()
}

func VerifySession(options *sessmodels.VerifySessionOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		session.VerifySession(options, func(rw http.ResponseWriter, r *http.Request) {
			c.Request = c.Request.WithContext(r.Context())
			c.Next()
		})(c.Writer, c.Request)
		c.Abort()
	}
}

// Ensures that a user session exists on the request
func EnsureLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentSession, err := session.GetSession(c.Request, c.Writer, &sessmodels.VerifySessionOptions{SessionRequired: util.Ptr(true)})
		if err != nil {
			c.Error(err)
			err = supertokens.ErrorHandler(err, c.Request, c.Writer)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userID", currentSession.GetUserID())
	}
}

// Ensures that a user session or an API token exists and check wheter the
// authentication method have the given capabilties
func EnsureCapabilities(jwt *mkjwt.JWTService, caps ...mkjwt.APITokenCapability) gin.HandlerFunc {
	return func(c *gin.Context) {
		// check api key first
		bearerToken := c.GetHeader("Authorization")
		tokenPrefix := "Bearer"
		if strings.HasPrefix(bearerToken, tokenPrefix) {
			tokenString := bearerToken[len(tokenPrefix):]

			claims := mkjwt.APITokenClaims{}
			token, err := jwt.ValidateToken(tokenString, &claims)
			if err != nil || !token.Valid {
				c.AbortWithError(http.StatusUnauthorized, err)
				return
			}

			log.Println(claims)
			capMap := mkjwt.MapCapabilities(claims)
			for _, cap := range caps {
				if !capMap[cap] {
					c.AbortWithError(http.StatusForbidden,
						fmt.Errorf("this route requires capability: %v", cap))
					return
				}
			}

			c.Set("userID", claims.UserID)
			log.Println("user authorized by api key")
		} else {
			// check if user is logged in otherwise.
			// no need to check for capabilities for logged users
			sessionContainer, err := session.GetSession(c.Request, c.Writer, nil)
			if err != nil {
				c.Error(err)
				err = supertokens.ErrorHandler(err, c.Request, c.Writer)
				if err != nil {
					c.AbortWithError(http.StatusInternalServerError, err)
				}
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if sessionContainer == nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("userID", sessionContainer.GetUserID())
			log.Println("user authorized by login")
		}
	}
}
