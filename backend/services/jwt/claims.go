package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	APITokenScope      string
	APITokenCapability string
)

const (
	// actions are performed without restrictions
	TokenScopeGlobal APITokenScope = "global"
	// actions are restricted to the user's properties
	TokenScopeUser APITokenScope = "user"

	TokenCapabiltyDeckCreate          APITokenCapability = "deck:create"
	TokenCapabiltyDeckDelete          APITokenCapability = "deck:delete"
	TokenCapabiltyDeckUpdate          APITokenCapability = "deck:update"
	TokenCababiltySubjectCreate       APITokenCapability = "subject:create"
	TokenCababiltySubjectUpdate       APITokenCapability = "subject:update"
	TokenCababiltySubjectDelete       APITokenCapability = "subject:delete"
	TokenCababiltyReviewCreate        APITokenCapability = "review:create"
	TokenCababiltyStudyMaterialCreate APITokenCapability = "study_material:create"
	TokenCababiltyStudyMaterialUpdate APITokenCapability = "study_material:update"
	TokenCababiltyUserUpdate          APITokenCapability = "user:update"
	TokenCababiltyUserDelete          APITokenCapability = "user:delete"
)

type APITokenOptions struct {
	UserID       string
	Scope        APITokenScope
	Capabilities []APITokenCapability
	ExpiresAt    *time.Time
}

type APITokenCapabilities struct {
	TokenCapabiltyDeckCreate          bool `json:"deck:create"`
	TokenCapabiltyDeckDelete          bool `json:"deck:delete"`
	TokenCapabiltyDeckUpdate          bool `json:"deck:update"`
	TokenCababiltySubjectCreate       bool `json:"subject:create"`
	TokenCababiltySubjectUpdate       bool `json:"subject:update"`
	TokenCababiltySubjectDelete       bool `json:"subject:delete"`
	TokenCababiltyReviewCreate        bool `json:"review:create"`
	TokenCababiltyStudyMaterialCreate bool `json:"study_material:create"`
	TokenCababiltyStudyMaterialUpdate bool `json:"study_material:update"`
	TokenCababiltyUserUpdate          bool `json:"user:update"`
	TokenCababiltyUserDelete          bool `json:"user:delete"`
}

type APITokenClaims struct {
	UserID string        `json:"user_id"`
	Scope  APITokenScope `json:"scope"`

	APITokenCapabilities

	jwt.StandardClaims
}

func MapCapabilities(claims APITokenClaims) map[APITokenCapability]bool {
	return map[APITokenCapability]bool{
		TokenCapabiltyDeckCreate:          claims.TokenCapabiltyDeckCreate,
		TokenCapabiltyDeckDelete:          claims.TokenCapabiltyDeckDelete,
		TokenCapabiltyDeckUpdate:          claims.TokenCapabiltyDeckUpdate,
		TokenCababiltySubjectCreate:       claims.TokenCababiltySubjectCreate,
		TokenCababiltySubjectUpdate:       claims.TokenCababiltySubjectUpdate,
		TokenCababiltySubjectDelete:       claims.TokenCababiltySubjectDelete,
		TokenCababiltyReviewCreate:        claims.TokenCababiltyReviewCreate,
		TokenCababiltyStudyMaterialCreate: claims.TokenCababiltyStudyMaterialCreate,
		TokenCababiltyStudyMaterialUpdate: claims.TokenCababiltyStudyMaterialUpdate,
		TokenCababiltyUserUpdate:          claims.TokenCababiltyUserUpdate,
		TokenCababiltyUserDelete:          claims.TokenCababiltyUserDelete,
	}
}
func MapCapabilitiesRef(claims *APITokenClaims) map[APITokenCapability]*bool {
	return map[APITokenCapability]*bool{
		TokenCapabiltyDeckCreate:          &claims.TokenCababiltyReviewCreate,
		TokenCapabiltyDeckDelete:          &claims.TokenCapabiltyDeckDelete,
		TokenCapabiltyDeckUpdate:          &claims.TokenCapabiltyDeckUpdate,
		TokenCababiltySubjectCreate:       &claims.TokenCababiltySubjectCreate,
		TokenCababiltySubjectUpdate:       &claims.TokenCababiltySubjectUpdate,
		TokenCababiltySubjectDelete:       &claims.TokenCababiltySubjectDelete,
		TokenCababiltyReviewCreate:        &claims.TokenCababiltyReviewCreate,
		TokenCababiltyStudyMaterialCreate: &claims.TokenCababiltyStudyMaterialCreate,
		TokenCababiltyStudyMaterialUpdate: &claims.TokenCababiltyStudyMaterialUpdate,
		TokenCababiltyUserUpdate:          &claims.TokenCababiltyUserUpdate,
		TokenCababiltyUserDelete:          &claims.TokenCababiltyUserDelete,
	}
}
