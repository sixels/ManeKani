package auth

import (
	"errors"
	"fmt"
	"github.com/sixels/manekani/server/api/apicommon"
	"strings"

	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
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

func (auth *AuthService) Middlewares() ([]gin.HandlerFunc, error) {
	spec, err := api.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("loading spec: %w", err)
	}

	// validator := OapiRequestValidatorWithOptions(
	// 	spec,
	// 	&middleware.Options{
	// 		Options: openapi3filter.Options{
	// 			AuthenticationFunc: NewOAPIAuthenticator(&auth.validator),
	// 		},
	// 		MultiErrorHandler: func(me openapi3.MultiError) error {
	// 			log.Println(me)
	// 			if me.Is(openapi3.ErrURINotSupported) {
	// 				return nil
	// 			}
	// 			return me
	// 		},
	// 		ErrorHandler: func(c *gin.Context, message string, statusCode int) {
	// 			apicommon.Respond(c, apicommon.Error(statusCode, errors.New(message)))
	// 		},
	// 	},
	// )
	return []gin.HandlerFunc{
		middleware.OapiRequestValidatorWithOptions(
			spec,
			&middleware.Options{
				ErrorHandler: func(c *gin.Context, message string, statusCode int) {
					if message == "no matching operation was found" {
						c.Next()
						return
					}
					apicommon.Respond(c, apicommon.Error(statusCode, errors.New(message)))
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

// func OAPIValidatorMiddleware(spec *openapi3.T, authenticator openapi3filter.AuthenticationFunc) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		route, err := FindRoute(c, spec)
// 		if err != nil {
// 			c.AbortWithError(http.StatusBadRequest, err)
// 			return
// 		}
// 		if route == nil {
// 			c.Next()
// 			return
// 		}

// 		req := c.Request.WithContext(context.WithValue(c.Request.Context(), GinContextKey, c))
// 		c.Request = req

// 		// req := c.Request.WithContext(context.WithValue(c.Request.Context(), GinContextKey, c))
// 		validationInput := &openapi3filter.RequestValidationInput{
// 			Request: req,
// 			Route:   route,
// 			Options: &openapi3filter.Options{AuthenticationFunc: authenticator},
// 		}
// 		if err := ValidateRequest(c, validationInput); err != nil {
// 			c.AbortWithError(http.StatusUnauthorized, err)
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func FindRoute(c *gin.Context, spec *openapi3.T) (*routers.Route, error) {
// 	path := c.Request.URL.Path
// 	method := c.Request.Method

// 	route := spec.Paths.Find(path)
// 	if route == nil {
// 		log.Printf("route %s not found in the spec\n", path)
// 		c.Next()
// 		return nil, nil
// 	}
// 	log.Printf("route %s found in the spec\n", path)

// 	operations := route.Operations()
// 	operation, ok := operations[method]
// 	if !ok {
// 		log.Printf("no operations with this method: %s\n", method)
// 		return nil, nil
// 	}

// 	return &routers.Route{
// 		Spec: spec,
// 		// Server:    route.Servers[],
// 		Path:      path,
// 		PathItem:  route,
// 		Method:    method,
// 		Operation: operation,
// 	}, nil
// }

// ValidateRequest is used to validate the given input according to previous
// loaded OpenAPIv3 spec. If the input does not match the OpenAPIv3 spec, a
// non-nil error will be returned.
//
// Note: One can tune the behavior of uniqueItems: true verification
// by registering a custom function with openapi3.RegisterArrayUniqueItemsChecker
// func ValidateRequest(ctx context.Context, input *openapi3filter.RequestValidationInput) (err error) {
// 	route := input.Route
// 	operation := route.Operation
// 	operationParameters := operation.Parameters
// 	pathItemParameters := route.PathItem.Parameters

// 	// Security
// 	security := operation.Security
// 	// If there aren't any security requirements for the operation
// 	if security == nil {
// 		// Use the global security requirements.
// 		security = &route.Spec.Security
// 	}
// 	if security != nil {
// 		if err := openapi3filter.ValidateSecurityRequirements(ctx, input, *security); err != nil {
// 			return err
// 		}
// 	}

// 	// For each parameter of the PathItem
// 	for _, parameterRef := range pathItemParameters {
// 		parameter := parameterRef.Value
// 		if operationParameters != nil {
// 			if override := operationParameters.GetByInAndName(parameter.In, parameter.Name); override != nil {
// 				continue
// 			}
// 		}

// 		if err := openapi3filter.ValidateParameter(ctx, input, parameter); err != nil {
// 			return err
// 		}
// 	}

// 	// For each parameter of the Operation
// 	for _, parameter := range operationParameters {
// 		if err := openapi3filter.ValidateParameter(ctx, input, parameter.Value); err != nil {
// 			return err
// 		}
// 	}

// 	// RequestBody
// 	requestBody := operation.RequestBody
// 	if requestBody != nil && !input.Options.ExcludeRequestBody {
// 		if err := openapi3filter.ValidateRequestBody(ctx, input, requestBody.Value); err != nil {
// 			return err
// 		}
// 	}

// 	return
// }

// OapiRequestValidatorWithOptions creates a validator from a swagger object, with validation options
// func OapiRequestValidatorWithOptions(swagger *openapi3.T, options *middleware.Options) gin.HandlerFunc {
// 	if swagger.Servers != nil && (options == nil || options.SilenceServersWarning) {
// 		log.Println("WARN: OapiRequestValidatorWithOptions called with an OpenAPI spec that has `Servers` set. This may lead to an HTTP 400 with `no matching operation was found` when sending a valid request, as the validator performs `Host` header validation. If you're expecting `Host` header validation, you can silence this warning by setting `Options.SilenceServersWarning = true`. See https://github.com/deepmap/oapi-codegen/issues/882 for more information.")
// 	}

// 	router, err := gorillamux.NewRouter(swagger)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return func(c *gin.Context) {
// 		err := middleware.ValidateRequestFromContext(c, router, options)
// 		if err != nil {
// 			log.Println(err)
// 			if options != nil && options.ErrorHandler != nil {
// 				options.ErrorHandler(c, err.Error(), http.StatusBadRequest)
// 				// in case the handler didn't internally call Abort, stop the chain
// 				c.Abort()
// 			} else {
// 				// note: i am not sure if this is the best way to handle this
// 				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			}
// 		}
// 		c.Next()
// 	}
// }

//--------
// this is a completely shit hack because oapi-codegen isn't flexible enough
// to handle non documented endpoints. i copied the exact code from the original
// implementation and made some ugly changes to it.

// type ContextKey string

// const (
// 	GinContextKey ContextKey = "oapi-codegen/gin-context"
// 	UserDataKey   ContextKey = "oapi-codegen/user-data"
// )

// var (
// 	ErrRouteNotFound = errors.New("route not found")
// )

// func OapiRequestValidatorCustomWithOptions(swagger *openapi3.T, options *middleware.Options) gin.HandlerFunc {
// 	if swagger.Servers != nil && (options == nil || options.SilenceServersWarning) {
// 		log.Println("WARN: OapiRequestValidatorWithOptions called with an OpenAPI spec that has `Servers` set. This may lead to an HTTP 400 with `no matching operation was found` when sending a valid request, as the validator performs `Host` header validation. If you're expecting `Host` header validation, you can silence this warning by setting `Options.SilenceServersWarning = true`. See https://github.com/deepmap/oapi-codegen/issues/882 for more information.")
// 	}

// 	router, err := gorillamux.NewRouter(swagger)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return func(c *gin.Context) {
// 		err := ValidateRequestFromContextCustom(c, router, options)
// 		if err != nil {
// 			c.Error(err)
// 			if errors.Is(err, ErrRouteNotFound) {
// 				c.Next()
// 				return
// 			}

// 			log.Println(err)
// 			if options != nil && options.ErrorHandler != nil {
// 				options.ErrorHandler(c, err.Error(), http.StatusBadRequest)
// 				// in case the handler didn't internally call Abort, stop the chain
// 				c.Abort()
// 			} else {
// 				// note: i am not sure if this is the best way to handle this
// 				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			}
// 		}
// 		c.Next()
// 	}
// }

// // ValidateRequestFromContextCustom is called from the middleware above and actually does the work
// // of validating a request.
// func ValidateRequestFromContextCustom(c *gin.Context, router routers.Router, options *middleware.Options) error {
// 	// req := c.Request

// 	// we need to enforce that gin context will be available to non documented endpoints
// 	req := c.Request.WithContext(context.WithValue(context.Background(), GinContextKey, c))

// 	// remove extra garbage from the url
// 	r := req.Clone(context.Background())
// 	r.URL.Path = filepath.Clean(req.URL.Path)

// 	route, pathParams, err := router.FindRoute(r)

// 	// We failed to find a matching route for the request.
// 	if err != nil {
// 		switch err.(type) {
// 		case *routers.RouteError:
// 			// We've got a bad request, the path requested doesn't match
// 			// either server, or path, or something.
// 			return ErrRouteNotFound
// 		default:
// 			// This should never happen today, but if our upstream code changes,
// 			// we don't want to crash the server, so handle the unexpected error.
// 			return fmt.Errorf("error validating route: %s", err.Error())
// 		}
// 	}

// 	validationInput := &openapi3filter.RequestValidationInput{
// 		Request:    req,
// 		PathParams: pathParams,
// 		Route:      route,
// 	}

// 	// Pass the gin context into the request validator, so that any callbacks
// 	// which it invokes make it available.
// 	requestContext := context.WithValue(context.Background(), GinContextKey, c) //nolint:staticcheck

// 	if options != nil {
// 		validationInput.Options = &options.Options
// 		validationInput.ParamDecoder = options.ParamDecoder
// 		requestContext = context.WithValue(requestContext, UserDataKey, options.UserData) //nolint:staticcheck
// 	}

// 	err = openapi3filter.ValidateRequest(requestContext, validationInput)
// 	if err != nil {
// 		me := openapi3.MultiError{}
// 		if errors.As(err, &me) {
// 			// errFunc := getMultiErrorHandlerFromOptions(options)
// 			// return errFunc(me)
// 			return fmt.Errorf("whoa! %w", me)
// 		}

// 		switch e := err.(type) {
// 		case *openapi3filter.RequestError:
// 			// We've got a bad request
// 			// Split up the verbose error by lines and return the first one
// 			// openapi errors seem to be multi-line with a decent message on the first
// 			errorLines := strings.Split(e.Error(), "\n")
// 			return fmt.Errorf("error in openapi3filter.RequestError: %s", errorLines[0])
// 		case *openapi3filter.SecurityRequirementsError:
// 			return fmt.Errorf("error in openapi3filter.SecurityRequirementsError: %s", e.Error())
// 		default:
// 			// This should never happen today, but if our upstream code changes,
// 			// we don't want to crash the server, so handle the unexpected error.
// 			return fmt.Errorf("error validating request: %s", err)
// 		}
// 	}
// 	return nil
// }

// func GetGinContext(c context.Context) *gin.Context {
// 	iface := c.Value(GinContextKey)
// 	if iface == nil {
// 		log.Println("no gin context")
// 		return nil
// 	}
// 	ginCtx, ok := iface.(*gin.Context)
// 	if !ok {
// 		log.Println("invalid gin context")
// 		return nil
// 	}
// 	log.Println("ok gin context")
// 	return ginCtx
// }
