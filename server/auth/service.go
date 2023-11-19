package auth

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sixels/manekani/server/api/apicommon"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	ory "github.com/ory/client-go"
	adapter "github.com/sixels/manekani/core/adapters/tokens"
	"github.com/sixels/manekani/server/api"
)

// TODO: set env variables
const (
	KRATOS_HOSTNAME      string = "kratos"
	KRATOS_API_URL       string = "http://" + KRATOS_HOSTNAME + ":4433"
	KRATOS_ADMIN_API_URL string = "http://" + KRATOS_HOSTNAME + ":4434"
	KRATOS_UI_URL        string = "http://127.0.0.1:4455"
)

type AuthService struct {
	validator OAPIValidator
}

func NewAuthService(tokenProvider *adapter.TokensAdapter) *AuthService {
	oryConfig := ory.NewConfiguration()
	oryConfig.Servers = ory.ServerConfigurations{{URL: KRATOS_ADMIN_API_URL}}
	oryConfig.Host = strings.TrimPrefix(KRATOS_API_URL, "http://")
	client := ory.NewAPIClient(oryConfig)

	return &AuthService{
		validator: OAPIValidator{
			tokens: tokenProvider,
			ory:    client,
		},
	}
}

func (auth *AuthService) ServiceName() string {
	return "authenticator"
}

func (auth *AuthService) Middlewares() ([]echo.MiddlewareFunc, error) {
	spec, err := api.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	openapi3filter.RegisterBodyDecoder("image/svg+xml", openapi3filter.FileBodyDecoder)

	return []echo.MiddlewareFunc{
		middleware.OapiRequestValidatorWithOptions(
			spec,
			&middleware.Options{
				ErrorHandler: func(c echo.Context, err *echo.HTTPError) error {
					log.Debug(err.Code, err.Internal, err.Message)
					if err.Message == "no matching operation was found" {
						return nil
						//return c.JSON(http.StatusNotFound, apicommon.Error(http.StatusNotFound, errors.New("API route not found")))
					}
					return c.JSON(err.Code, apicommon.Error(err.Code, err))
				},
				Options: openapi3filter.Options{
					ExcludeRequestBody:  true,
					ExcludeResponseBody: true,
					AuthenticationFunc:  NewOAPIAuthenticator(&auth.validator),
				},
				SilenceServersWarning: true,
			},
		),
	}, nil
}
