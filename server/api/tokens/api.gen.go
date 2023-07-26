// Package tokens provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package tokens

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

const (
	LoginScopes = "Login.Scopes"
)

// Defines values for TokenStatus.
const (
	TokenStatusActive TokenStatus = "active"
	TokenStatusFrozen TokenStatus = "frozen"
)

// CommonErrorResponse defines model for common.ErrorResponse.
type CommonErrorResponse struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// CommonResponse defines model for common.Response.
type CommonResponse struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
}

// TokenClaims defines model for token.Claims.
type TokenClaims struct {
	DeckCreate          *bool `json:"deck_create,omitempty"`
	DeckDelete          *bool `json:"deck_delete,omitempty"`
	DeckUpdate          *bool `json:"deck_update,omitempty"`
	ReviewCreate        *bool `json:"review_create,omitempty"`
	StudyMaterialCreate *bool `json:"study_material_create,omitempty"`
	StudyMaterialUpdate *bool `json:"study_material_update,omitempty"`
	SubjectCreate       *bool `json:"subject_create,omitempty"`
	SubjectDelete       *bool `json:"subject_delete,omitempty"`
	SubjectUpdate       *bool `json:"subject_update,omitempty"`
	UserDelete          *bool `json:"user_delete,omitempty"`
	UserUpdate          *bool `json:"user_update,omitempty"`
}

// TokenCreateRequest defines model for token.CreateRequest.
type TokenCreateRequest struct {
	Name        string       `json:"name"`
	Permissions *TokenClaims `json:"permissions,omitempty"`
}

// TokenCreateResponse defines model for token.CreateResponse.
type TokenCreateResponse struct {
	Claims *TokenClaims `json:"claims,omitempty"`
	Id     *string      `json:"id,omitempty"`
	Name   *string      `json:"name,omitempty"`
	Prefix *string      `json:"prefix,omitempty"`
	Status *TokenStatus `json:"status,omitempty"`
	Token  *string      `json:"token,omitempty"`
	UsedAt *string      `json:"used_at,omitempty"`
}

// TokenGetAllResponse defines model for token.GetAllResponse.
type TokenGetAllResponse = []struct {
	Claims *TokenClaims `json:"claims,omitempty"`
	Id     string       `json:"id"`
	Name   string       `json:"name"`
	Prefix string       `json:"prefix"`
	Status TokenStatus  `json:"status"`
	UsedAt *string      `json:"used_at,omitempty"`
}

// TokenStatus defines model for token.Status.
type TokenStatus string

// ValidateTokenParams defines parameters for ValidateToken.
type ValidateTokenParams struct {
	Authorization string `json:"Authorization"`
}

// CreateTokenJSONRequestBody defines body for CreateToken for application/json ContentType.
type CreateTokenJSONRequestBody = TokenCreateRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all user's API tokens
	// (GET /api/tokens)
	GetTokens(c *gin.Context)
	// Create a new API token
	// (POST /api/tokens)
	CreateToken(c *gin.Context)
	// Delete an API token
	// (DELETE /api/tokens/{id})
	DeleteToken(c *gin.Context, id string)
	// Validate a token
	// (GET /auth/validate-token)
	ValidateToken(c *gin.Context, params ValidateTokenParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetTokens operation middleware
func (siw *ServerInterfaceWrapper) GetTokens(c *gin.Context) {

	c.Set(LoginScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTokens(c)
}

// CreateToken operation middleware
func (siw *ServerInterfaceWrapper) CreateToken(c *gin.Context) {

	c.Set(LoginScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateToken(c)
}

// DeleteToken operation middleware
func (siw *ServerInterfaceWrapper) DeleteToken(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(LoginScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteToken(c, id)
}

// ValidateToken operation middleware
func (siw *ServerInterfaceWrapper) ValidateToken(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ValidateTokenParams

	headers := c.Request.Header

	// ------------- Required header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for Authorization, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter Authorization: %s", err), http.StatusBadRequest)
			return
		}

		params.Authorization = Authorization

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Header parameter Authorization is required, but not found"), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ValidateToken(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/api/tokens", wrapper.GetTokens)
	router.POST(options.BaseURL+"/api/tokens", wrapper.CreateToken)
	router.DELETE(options.BaseURL+"/api/tokens/:id", wrapper.DeleteToken)
	router.GET(options.BaseURL+"/auth/validate-token", wrapper.ValidateToken)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xY3W/bNhD/VwhuQF7kj3QJVujNXbYiyIYWTbqXwAgY8WyxlkiVPDlVA/3vAz9ky5as",
	"OG2xrUDfLPK+7353Rz/SROWFkiDR0PiRmiSFnLmficpzJce/a630OzCFkgbseaFVARoFBCruTrEqgMZU",
	"SIQlaFpHNAdj2LJ9aVALuaR1HTUn6v4DJGipg7YvUcQZstZNI7OOqIaPpdDAaXzr+QPxvMcAVCuQ498y",
	"JnLT1c4hWd0lGhi2jbhXKgMmnRGWgEMGgwRlwQ9K0LAW8DCoxGDJq7ucIWjBsueQDik2pQvCsLhAM+Rh",
	"QzOkqzSgB4U4gsMS6sOZc9a/g48lGOwmULK8rxQjWoDOhTFCSUf3s4YFjelPky0uJgEUk50S2a8vp2D+",
	"pHkHy3tTeMdbEFHBe3067KyGhfjUe2WQYXmkAdeetnGuV15pgN8xPBL+Xu5rwFmWtYMkEPrg+J1GazAm",
	"7WISnAa7NlZsdPbWmD9gWrNqG87rjZEgy9zKZQmKtRW60OozyJasYEhEP40s8WjNtNVvLNeNleaFzRoB",
	"rbM/gizrhIGk1AKraxsBn6tZIa6g8k3UJFoUKJSkMZ29vSQrqAgrMQWJImHuIqLC3qbAOOgmCjGdlZgq",
	"LT43RI3HXngd0T/VUsiuFndMGJIUsYgnk9MXv46n4+n4ND47Oz+fZP5acpKooiKYAkmURJtFohbu+0Tp",
	"6m6lGSpzZ8B1ihOSKLUS0Bi7+QrGdjm6FttoCblQXZP/YhKumBRk9vZybBkFZrB3TiO6Bm08w+l4aiOg",
	"CpCsEDSmv1gHbekwTF0OJqwQvhrd5xKwq/Y1IGFZRmwHPjFWCQkcTrR2kb/knvKmudEBrE7ui+nUj2oX",
	"QfuTFUUWMjv5YJTc7hjuNsveLGh8O4yi/dWgjjrTOWwAT2Nxr8P09KK5y8xuaN5c2fieTU+f5d0RPu0u",
	"Vz2a30sWCh+4teH8mRH+FjZcSgQtWUauQa9BE8dA22h3KQwAvJ3Xc7sL5DnT1RNlhWzpGoyfInPbfERQ",
	"RmPUJdgmrExPsfpxaggjEh62QsmDwNShdinWIEl7uO9XsRfhCpn67gsGXylefbMA9y0mPfG9SaHlgXLH",
	"hrYHQgjFHtZO/+dY21t5jsSa5+IecP9+sb9inGxy9QPzX4Z5n8N9cPYD3r7LliBHAYCje8WrUZij4Yz2",
	"NIY6ag+1yaPgte8SzdNi158Ld24Iky2o3VetVnF50ekQnqnpEAXTLAcEbZzrQxh2stxmYCfwdi9wa90u",
	"rKNW4vZ3wvlXjtdnAb1n6q1+gPA7BqGv3p2KP2rkOmSVmE7WLBP2HT7aPPLC3riLkr8D2QGcHLXNfzUm",
	"+ve1/7hy63Y6mjAR1snFJgF+HhqXbR+9Umc2fPtvl5fTl1NqYxFE9L2tUmAZpiS8GjehzwG1SNxgP1Bt",
	"lnlRysSes0y4wb9tYY2tXQHX/v8fMywg/EvUw38ByeoJZg7JqofTv0aGWX3Qu7zvDWiiIXNLx4babqt2",
	"P/knAAD//+7SiE4gFQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}