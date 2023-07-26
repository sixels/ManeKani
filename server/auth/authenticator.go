package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	ory "github.com/ory/client-go"
	adapter "github.com/sixels/manekani/core/adapters/tokens"
	"github.com/sixels/manekani/core/domain/tokens"
)

var (
	ErrNoAuthHeader      = errors.New("authorization header is missing")
	ErrInvalidAuthHeader = errors.New("authorization header is malformed")
	ErrNoLoginSession    = errors.New("login session cookies are missing")
)

type IdentityContext string

const (
	UserIDContext      IdentityContext = "userID"
	UserTokenContext   IdentityContext = "userToken"
	UserSessionContext IdentityContext = "userSession"
)

type TokenAuthenticator interface {
	ValidateToken(ctx context.Context, tk string) (tokens.UserToken, error)
	ValidateLoginSession(ctx context.Context, sessionCookies string) (*ory.Session, error)
}

func NewOAPIAuthenticator(auth TokenAuthenticator) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
		if ai.SecuritySchemeName == "ApiKey" {
			return AuthenticateApiKey(ctx, auth, ai)
		}
		if ai.SecuritySchemeName == "Login" {
			return AuthenticateLogin(ctx, auth, ai)
		}
		log.Println("request authentication failed")
		return fmt.Errorf("unknown security scheme name: %s", ai.SecuritySchemeName)
	}
}

func AuthenticateLogin(ctx context.Context, auth TokenAuthenticator, ai *openapi3filter.AuthenticationInput) error {
	sessionValue, err := GetLoginSession(ai.RequestValidationInput.Request)
	if err != nil {
		log.Println(err)
		return err
	}

	session, err := auth.ValidateLoginSession(ctx, sessionValue)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("could not get the login session: %w", err)
	}

	c := middleware.GetGinContext(ctx)
	c.Set(string(UserSessionContext), session)
	c.Set(string(UserIDContext), session.Identity.Id)

	return nil
}

func GetLoginSession(req *http.Request) (string, error) {
	// TODO: get only relevant cookies
	cookies := req.Header.Get("Cookie")
	if cookies == "" {
		return "", ErrNoLoginSession
	}
	return cookies, nil
}

func AuthenticateApiKey(ctx context.Context, auth TokenAuthenticator, ai *openapi3filter.AuthenticationInput) error {
	requestToken, err := GetAuthTokenHeader(ai.RequestValidationInput.Request)
	if err != nil {
		return err
	}

	tk, err := auth.ValidateToken(ctx, requestToken)
	if err != nil {
		return err
	}

	err = CheckTokenClaims(tk, ai.Scopes)
	if err != nil {
		return err
	}

	c := middleware.GetGinContext(ctx)
	c.Set(string(UserTokenContext), tk)
	c.Set(string(UserIDContext), tk.UserID)

	return nil
}

func GetAuthTokenHeader(req *http.Request) (string, error) {
	bearerToken := req.Header.Get("Authorization")
	if bearerToken == "" {
		return "", ErrNoAuthHeader
	}
	tokenPrefix := "Bearer "
	if !strings.HasPrefix(bearerToken, tokenPrefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(bearerToken, tokenPrefix), nil
}

func CheckTokenClaims(tk tokens.UserToken, expectedClaims []string) error {
	claimsMap := adapter.MapPermissions(tk.Claims)
	var missing []string
	for _, expectedClaim := range expectedClaims {
		if isSet, exists := claimsMap[tokens.APITokenPermission(expectedClaim)]; !exists || !isSet {
			missing = append(missing, expectedClaim)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing claims: %v", missing)
	}
	return nil
}
