// Package cards provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package cards

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
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

const (
	ApiKeyScopes = "ApiKey.Scopes"
	LoginScopes  = "Login.Scopes"
)

// CommonErrorResponse defines model for common.ErrorResponse.
type CommonErrorResponse struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// DeckCreateRequest defines model for deck.CreateRequest.
type DeckCreateRequest struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// DeckCreateResponse defines model for deck.CreateResponse.
type DeckCreateResponse = DeckModel

// DeckGetAllResponse defines model for deck.GetAllResponse.
type DeckGetAllResponse struct {
	Decks *[]struct {
		Description *string             `json:"description,omitempty"`
		Id          *openapi_types.UUID `json:"id,omitempty"`
		Name        *string             `json:"name,omitempty"`
		Owner       *string             `json:"owner,omitempty"`
	} `json:"decks,omitempty"`
}

// DeckGetResponse defines model for deck.GetResponse.
type DeckGetResponse = DeckModel

// DeckModel defines model for deck.Model.
type DeckModel struct {
	CreatedAt   *string             `json:"created_at,omitempty"`
	Description *string             `json:"description,omitempty"`
	Id          *openapi_types.UUID `json:"id,omitempty"`
	Name        *string             `json:"name,omitempty"`
	Owner       *string             `json:"owner,omitempty"`
	Subjects    *[]string           `json:"subjects,omitempty"`
	UpdatedAt   *string             `json:"updated_at,omitempty"`
}

// SubjectCreateRequest defines model for subject.CreateRequest.
type SubjectCreateRequest struct {
	AdditionalStudyData *map[string]interface{} `form:"additional_study_data" json:"additional_study_data,omitempty"`
	Deck                openapi_types.UUID      `form:"deck" json:"deck"`
	Dependencies        *[]openapi_types.UUID   `form:"dependencies" json:"dependencies,omitempty"`
	Dependents          *[]openapi_types.UUID   `form:"dependents" json:"dependents,omitempty"`
	Kind                string                  `form:"kind" json:"kind"`
	Level               int32                   `form:"level" json:"level"`
	Name                string                  `form:"name" json:"name"`
	Priority            uint8                   `form:"priority" json:"priority"`

	// Resource The subject resources
	Resource *[]openapi_types.File `json:"resource[],omitempty"`

	// ResourcesMeta Resources metadatas
	ResourcesMeta *[]map[string]string  `json:"resources_meta,omitempty"`
	Similars      *[]openapi_types.UUID `json:"similars,omitempty"`
	Slug          string                `json:"slug"`
	StudyData     *[]SubjectStudyData   `json:"study_data,omitempty"`
	Value         *string               `json:"value,omitempty"`

	// ValueImage The subject value image
	ValueImage *openapi_types.File `json:"value_image,omitempty"`
}

// SubjectCreateResponse defines model for subject.CreateResponse.
type SubjectCreateResponse = SubjectModel

// SubjectGetAllResponse defines model for subject.GetAllResponse.
type SubjectGetAllResponse = []struct {
	Deck         openapi_types.UUID   `json:"deck"`
	Dependencies []openapi_types.UUID `json:"dependencies"`
	Dependents   []openapi_types.UUID `json:"dependents"`
	Id           openapi_types.UUID   `json:"id"`
	Kind         string               `json:"kind"`
	Level        int32                `json:"level"`
	Name         string               `json:"name"`
	Owner        string               `json:"owner"`
	Priority     uint8                `json:"priority"`
	Similars     []openapi_types.UUID `json:"similars"`
	Slug         string               `json:"slug"`
	StudyData    []SubjectStudyData   `json:"study_data"`
	Value        *string              `json:"value,omitempty"`
	ValueImage   *string              `json:"value_image,omitempty"`
}

// SubjectGetResponse defines model for subject.GetResponse.
type SubjectGetResponse = SubjectModel

// SubjectModel defines model for subject.Model.
type SubjectModel struct {
	AdditionalStudyData *map[string]interface{} `json:"additional_study_data,omitempty"`
	CreatedAt           string                  `json:"created_at"`
	Deck                openapi_types.UUID      `json:"deck"`
	Dependencies        []openapi_types.UUID    `json:"dependencies"`
	Dependents          []openapi_types.UUID    `json:"dependents"`
	Id                  openapi_types.UUID      `json:"id"`
	Kind                string                  `json:"kind"`
	Level               int32                   `json:"level"`
	Name                string                  `json:"name"`
	Owner               string                  `json:"owner"`
	Priority            uint8                   `json:"priority"`
	Resources           []SubjectResource       `json:"resources"`
	Similars            []openapi_types.UUID    `json:"similars"`
	Slug                string                  `json:"slug"`
	StudyData           []SubjectStudyData      `json:"study_data"`
	UpdatedAt           string                  `json:"updated_at"`
	Value               *string                 `json:"value,omitempty"`
	ValueImage          *string                 `json:"value_image,omitempty"`
}

// SubjectResource defines model for subject.Resource.
type SubjectResource struct {
	Metadata *map[string]string `json:"metadata,omitempty"`
	Url      *string            `json:"url,omitempty"`
}

// SubjectStudyData defines model for subject.StudyData.
type SubjectStudyData struct {
	Items    []SubjectStudyItem `json:"items"`
	Kind     string             `json:"kind"`
	Mnemonic string             `json:"mnemonic"`
}

// SubjectStudyItem defines model for subject.StudyItem.
type SubjectStudyItem struct {
	Category      *string `json:"category,omitempty"`
	IsHidden      bool    `json:"is_hidden"`
	IsPrimary     bool    `json:"is_primary"`
	IsValidAnswer bool    `json:"is_valid_answer"`
	Resource      *string `json:"resource,omitempty"`
	Value         string  `json:"value"`
}

// SubjectUpdateRequest defines model for subject.UpdateRequest.
type SubjectUpdateRequest struct {
	AdditionalStudyData *map[string]interface{} `json:"additional_study_data,omitempty"`
	Dependencies        *[]openapi_types.UUID   `json:"dependencies,omitempty"`
	Dependents          *[]openapi_types.UUID   `json:"dependents,omitempty"`
	Kind                *string                 `json:"kind,omitempty"`
	Level               *int32                  `json:"level,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	Priority            *uint8                  `json:"priority,omitempty"`
	Resources           *[]SubjectResource      `json:"resources,omitempty"`
	Similars            *[]openapi_types.UUID   `json:"similars,omitempty"`
	Slug                *string                 `json:"slug,omitempty"`
	StudyData           *[]SubjectStudyData     `json:"study_data,omitempty"`
	Value               *string                 `json:"value,omitempty"`
	ValueImage          *string                 `json:"value_image,omitempty"`
}

// SubjectUpdateResponse defines model for subject.UpdateResponse.
type SubjectUpdateResponse = SubjectModel

// GetDecksParams defines parameters for GetDecks.
type GetDecksParams struct {
	Ids      *string `form:"ids,omitempty" json:"ids,omitempty"`
	Subjects *string `form:"subjects,omitempty" json:"subjects,omitempty"`
	Owners   *string `form:"owners,omitempty" json:"owners,omitempty"`
	Page     *int    `form:"page,omitempty" json:"page,omitempty"`
	Names    *string `form:"names,omitempty" json:"names,omitempty"`
}

// GetSubjectsParams defines parameters for GetSubjects.
type GetSubjectsParams struct {
	Decks  *string `form:"decks,omitempty" json:"decks,omitempty"`
	Ids    *string `form:"ids,omitempty" json:"ids,omitempty"`
	Kinds  *string `form:"kinds,omitempty" json:"kinds,omitempty"`
	Levels *string `form:"levels,omitempty" json:"levels,omitempty"`
	Owners *string `form:"owners,omitempty" json:"owners,omitempty"`
	Page   *int    `form:"page,omitempty" json:"page,omitempty"`
	Slugs  *string `form:"slugs,omitempty" json:"slugs,omitempty"`
}

// CreateDeckJSONRequestBody defines body for CreateDeck for application/json ContentType.
type CreateDeckJSONRequestBody = DeckCreateRequest

// CreateSubjectMultipartRequestBody defines body for CreateSubject for multipart/form-data ContentType.
type CreateSubjectMultipartRequestBody = SubjectCreateRequest

// UpdateSubjectJSONRequestBody defines body for UpdateSubject for application/json ContentType.
type UpdateSubjectJSONRequestBody = SubjectUpdateRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Query all decks
	// (GET /api/v1/decks)
	GetDecks(c *gin.Context, params GetDecksParams)
	// Create a new deck
	// (POST /api/v1/decks)
	CreateDeck(c *gin.Context)
	// Query a deck
	// (GET /api/v1/decks/{id})
	GetDeck(c *gin.Context, id string)
	// Query all subjects
	// (GET /api/v1/subjects)
	GetSubjects(c *gin.Context, params GetSubjectsParams)
	// Create a new subject
	// (POST /api/v1/subjects)
	CreateSubject(c *gin.Context)
	// Delete a subject
	// (DELETE /api/v1/subjects/{id})
	DeleteSubject(c *gin.Context, id string)
	// Query a subject
	// (GET /api/v1/subjects/{id})
	GetSubject(c *gin.Context, id string)
	// Update a subject
	// (PATCH /api/v1/subjects/{id})
	UpdateSubject(c *gin.Context, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetDecks operation middleware
func (siw *ServerInterfaceWrapper) GetDecks(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDecksParams

	// ------------- Optional query parameter "ids" -------------

	err = runtime.BindQueryParameter("form", true, false, "ids", c.Request.URL.Query(), &params.Ids)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ids: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "subjects" -------------

	err = runtime.BindQueryParameter("form", true, false, "subjects", c.Request.URL.Query(), &params.Subjects)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter subjects: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "owners" -------------

	err = runtime.BindQueryParameter("form", true, false, "owners", c.Request.URL.Query(), &params.Owners)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter owners: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "names" -------------

	err = runtime.BindQueryParameter("form", true, false, "names", c.Request.URL.Query(), &params.Names)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter names: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetDecks(c, params)
}

// CreateDeck operation middleware
func (siw *ServerInterfaceWrapper) CreateDeck(c *gin.Context) {

	c.Set(ApiKeyScopes, []string{"deck:create"})

	c.Set(LoginScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateDeck(c)
}

// GetDeck operation middleware
func (siw *ServerInterfaceWrapper) GetDeck(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetDeck(c, id)
}

// GetSubjects operation middleware
func (siw *ServerInterfaceWrapper) GetSubjects(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetSubjectsParams

	// ------------- Optional query parameter "decks" -------------

	err = runtime.BindQueryParameter("form", true, false, "decks", c.Request.URL.Query(), &params.Decks)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter decks: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "ids" -------------

	err = runtime.BindQueryParameter("form", true, false, "ids", c.Request.URL.Query(), &params.Ids)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter ids: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "kinds" -------------

	err = runtime.BindQueryParameter("form", true, false, "kinds", c.Request.URL.Query(), &params.Kinds)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter kinds: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "levels" -------------

	err = runtime.BindQueryParameter("form", true, false, "levels", c.Request.URL.Query(), &params.Levels)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter levels: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "owners" -------------

	err = runtime.BindQueryParameter("form", true, false, "owners", c.Request.URL.Query(), &params.Owners)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter owners: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "slugs" -------------

	err = runtime.BindQueryParameter("form", true, false, "slugs", c.Request.URL.Query(), &params.Slugs)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slugs: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetSubjects(c, params)
}

// CreateSubject operation middleware
func (siw *ServerInterfaceWrapper) CreateSubject(c *gin.Context) {

	c.Set(ApiKeyScopes, []string{"subject:create"})

	c.Set(LoginScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateSubject(c)
}

// DeleteSubject operation middleware
func (siw *ServerInterfaceWrapper) DeleteSubject(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(ApiKeyScopes, []string{"subject:delete"})

	c.Set(LoginScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteSubject(c, id)
}

// GetSubject operation middleware
func (siw *ServerInterfaceWrapper) GetSubject(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetSubject(c, id)
}

// UpdateSubject operation middleware
func (siw *ServerInterfaceWrapper) UpdateSubject(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	c.Set(ApiKeyScopes, []string{"subject:update"})

	c.Set(LoginScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateSubject(c, id)
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

	router.GET(options.BaseURL+"/api/v1/decks", wrapper.GetDecks)
	router.POST(options.BaseURL+"/api/v1/decks", wrapper.CreateDeck)
	router.GET(options.BaseURL+"/api/v1/decks/:id", wrapper.GetDeck)
	router.GET(options.BaseURL+"/api/v1/subjects", wrapper.GetSubjects)
	router.POST(options.BaseURL+"/api/v1/subjects", wrapper.CreateSubject)
	router.DELETE(options.BaseURL+"/api/v1/subjects/:id", wrapper.DeleteSubject)
	router.GET(options.BaseURL+"/api/v1/subjects/:id", wrapper.GetSubject)
	router.PATCH(options.BaseURL+"/api/v1/subjects/:id", wrapper.UpdateSubject)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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