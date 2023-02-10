package users

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"sixels.io/manekani/core/domain/user"
	mkjwt "sixels.io/manekani/services/jwt"
)

func (api *UserApi) RequiresUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionRequired := true
		currentSession, err := session.GetSession(c.Request, c.Writer, &sessmodels.VerifySessionOptions{SessionRequired: &sessionRequired})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		authUser, err := emailpassword.GetUserByID(currentSession.GetUserID())
		if err != nil {
			c.AbortWithError(http.StatusNotAcceptable, err)
			return
		}

		userInfo, err := api.users.QueryUserResolved(c.Request.Context(), authUser.ID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Set("user", userInfo)
	}
}

func EnsureCapabilities(jwt *mkjwt.JWTService, caps ...mkjwt.APITokenCapability) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		log.Println(bearerToken)

		if bearerToken != "" {
			tokenString := strings.Split(bearerToken, " ")[1]

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
			log.Println("user authorized")
			return
		}

		if ctxUser, ok := c.Get("user"); ok {
			loggedUser, ok := ctxUser.(*user.User)
			if !ok {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// TODO: check user roles

			c.Set("userID", loggedUser.ID)
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
