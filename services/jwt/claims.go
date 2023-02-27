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

	TokenCapabiltyDeckCreate           APITokenCapability = "deck:create"
	TokenCapabiltyDeckDelete           APITokenCapability = "deck:delete"
	TokenCapabiltyDeckUpdate           APITokenCapability = "deck:update"
	TokenCapabilitySubjectCreate       APITokenCapability = "subject:create"
	TokenCapabilitySubjectUpdate       APITokenCapability = "subject:update"
	TokenCapabilitySubjectDelete       APITokenCapability = "subject:delete"
	TokenCapabilityReviewCreate        APITokenCapability = "review:create"
	TokenCapabilityStudyMaterialCreate APITokenCapability = "study_material:create"
	TokenCapabilityStudyMaterialUpdate APITokenCapability = "study_material:update"
	TokenCapabilityUserUpdate          APITokenCapability = "user:update"
	TokenCapabilityUserDelete          APITokenCapability = "user:delete"
)

type APITokenOptions struct {
	UserID       string
	Scope        APITokenScope
	Capabilities []APITokenCapability
	ExpiresAt    *time.Time
}

type APITokenCapabilities struct {
	TokenCapabiltyDeckCreate           bool `json:"deck:create"`
	TokenCapabiltyDeckDelete           bool `json:"deck:delete"`
	TokenCapabiltyDeckUpdate           bool `json:"deck:update"`
	TokenCapabilitySubjectCreate       bool `json:"subject:create"`
	TokenCapabilitySubjectUpdate       bool `json:"subject:update"`
	TokenCapabilitySubjectDelete       bool `json:"subject:delete"`
	TokenCapabilityReviewCreate        bool `json:"review:create"`
	TokenCapabilityStudyMaterialCreate bool `json:"study_material:create"`
	TokenCapabilityStudyMaterialUpdate bool `json:"study_material:update"`
	TokenCapabilityUserUpdate          bool `json:"user:update"`
	TokenCapabilityUserDelete          bool `json:"user:delete"`
}

type APITokenClaims struct {
	UserID string        `json:"user_id"`
	Scope  APITokenScope `json:"scope"`

	APITokenCapabilities

	jwt.StandardClaims
}

func MapCapabilities(claims APITokenClaims) map[APITokenCapability]bool {
	return map[APITokenCapability]bool{
		TokenCapabiltyDeckCreate:           claims.TokenCapabiltyDeckCreate,
		TokenCapabiltyDeckDelete:           claims.TokenCapabiltyDeckDelete,
		TokenCapabiltyDeckUpdate:           claims.TokenCapabiltyDeckUpdate,
		TokenCapabilitySubjectCreate:       claims.TokenCapabilitySubjectCreate,
		TokenCapabilitySubjectUpdate:       claims.TokenCapabilitySubjectUpdate,
		TokenCapabilitySubjectDelete:       claims.TokenCapabilitySubjectDelete,
		TokenCapabilityReviewCreate:        claims.TokenCapabilityReviewCreate,
		TokenCapabilityStudyMaterialCreate: claims.TokenCapabilityStudyMaterialCreate,
		TokenCapabilityStudyMaterialUpdate: claims.TokenCapabilityStudyMaterialUpdate,
		TokenCapabilityUserUpdate:          claims.TokenCapabilityUserUpdate,
		TokenCapabilityUserDelete:          claims.TokenCapabilityUserDelete,
	}
}
func MapCapabilitiesRef(claims *APITokenClaims) map[APITokenCapability]*bool {
	return map[APITokenCapability]*bool{
		TokenCapabiltyDeckCreate:           &claims.TokenCapabilityReviewCreate,
		TokenCapabiltyDeckDelete:           &claims.TokenCapabiltyDeckDelete,
		TokenCapabiltyDeckUpdate:           &claims.TokenCapabiltyDeckUpdate,
		TokenCapabilitySubjectCreate:       &claims.TokenCapabilitySubjectCreate,
		TokenCapabilitySubjectUpdate:       &claims.TokenCapabilitySubjectUpdate,
		TokenCapabilitySubjectDelete:       &claims.TokenCapabilitySubjectDelete,
		TokenCapabilityReviewCreate:        &claims.TokenCapabilityReviewCreate,
		TokenCapabilityStudyMaterialCreate: &claims.TokenCapabilityStudyMaterialCreate,
		TokenCapabilityStudyMaterialUpdate: &claims.TokenCapabilityStudyMaterialUpdate,
		TokenCapabilityUserUpdate:          &claims.TokenCapabilityUserUpdate,
		TokenCapabilityUserDelete:          &claims.TokenCapabilityUserDelete,
	}
}
