package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/client-go"
	domain "github.com/sixels/manekani/core/domain/tokens"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
)

// TODO: set env variables
const (
	KRATOS_HOSTNAME      string = "kratos"
	KRATOS_API_URL       string = "http://" + KRATOS_HOSTNAME + ":4433"
	KRATOS_ADMIN_API_URL string = "http://" + KRATOS_HOSTNAME + ":4434"
	KRATOS_UI_URL        string = "http://127.0.0.1:4455"
)

func VerifySession(options *sessmodels.VerifySessionOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		session.VerifySession(options, func(rw http.ResponseWriter, r *http.Request) {
			c.Request = c.Request.WithContext(r.Context())
			c.Next()
		})(c.Writer, c.Request)
		c.Abort()
	}
}

func getLoginSession(ory *ory.APIClient, r *http.Request) (*ory.Session, error) {
	log.Println("Getting login session")
	cookies := r.Header.Get("Cookie")

	// check if we have a session
	session, _, err := ory.FrontendApi.ToSession(r.Context()).
		Cookie(cookies).
		Execute()

	return session, err
}

// Ensures that a user session exists on the request
func (mid *AuthService) EnsureLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := getLoginSession(mid.ory, c.Request)

		if err != nil || session == nil || !*session.Active {
			c.Error(err)
			c.Redirect(http.StatusFound,
				fmt.Sprintf("%s/self-service/login/browser?return_to=%s", "http://127.0.0.1:4433", c.Request.RequestURI))
			return
		}

		c.Set("userSession", session)
		c.Set("userID", session.Identity.Id)
	}
}

// Ensures that a user session or an API token exists and check wheter the
// authentication method have the given permissions
func (mid *AuthService) EnsurePermissions(caps ...domain.APITokenPermission) gin.HandlerFunc {
	return func(c *gin.Context) {
		// check api key first
		bearerToken := c.GetHeader("Authorization")
		tokenPrefix := "Bearer "
		if strings.HasPrefix(bearerToken, tokenPrefix) {
			tokenString := bearerToken[len(tokenPrefix):]

			token, err := mid.tokens.GetToken(c.Request.Context(), tokenString)
			if err != nil {
				log.Println("token validation error:", err)
				c.AbortWithError(http.StatusUnauthorized, err)
				return
			}

			if missingPerms := checkPermissions(&token.Claims.Permissions); len(missingPerms) > 0 {
				c.AbortWithError(
					http.StatusForbidden,
					fmt.Errorf("the token does not have some required permissions: %q", missingPerms),
				)
				return
			}

			c.Set("userID", token.UserID)
			log.Println("user authorized by api key")
		} else {
			session, err := getLoginSession(mid.ory, c.Request)

			if err != nil || session == nil || !*session.Active {
				c.Error(err)
				c.Redirect(http.StatusFound,
					fmt.Sprintf("%s/self-service/login/browser?return_to=%s", "http://127.0.0.1:4433", c.Request.RequestURI))
				return
			}

			userID := session.Identity.Id
			c.Set("userSession", session)
			c.Set("userID", userID)

			log.Printf("user %s authorized by login\n", userID)
		}
	}
}

func checkPermissions(c *domain.APITokenPermissions, perms ...domain.APITokenPermission) (missing []domain.APITokenPermission) {
	permissionMap := map[domain.APITokenPermission]bool{
		domain.TokenPermissionDeckCreate:          c.TokenPermissionReviewCreate,
		domain.TokenPermissionDeckDelete:          c.TokenPermissionDeckDelete,
		domain.TokenPermissionDeckUpdate:          c.TokenPermissionDeckUpdate,
		domain.TokenPermissionSubjectCreate:       c.TokenPermissionSubjectCreate,
		domain.TokenPermissionSubjectUpdate:       c.TokenPermissionSubjectUpdate,
		domain.TokenPermissionSubjectDelete:       c.TokenPermissionSubjectDelete,
		domain.TokenPermissionReviewCreate:        c.TokenPermissionReviewCreate,
		domain.TokenPermissionStudyMaterialCreate: c.TokenPermissionStudyMaterialCreate,
		domain.TokenPermissionStudyMaterialUpdate: c.TokenPermissionStudyMaterialUpdate,
		domain.TokenPermissionUserUpdate:          c.TokenPermissionUserUpdate,
		domain.TokenPermissionUserDelete:          c.TokenPermissionUserDelete,
	}

	for _, cap := range perms {
		if !permissionMap[cap] {
			missing = append(missing, cap)
		}
	}

	return missing
}
