package tokens

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type UserToken struct {
	ID     ulid.ULID      `json:"id"`
	UserID string         `json:"user_id"`
	Name   string         `json:"name"`
	Claims APITokenClaims `json:"claims"`
	Prefix string         `json:"prefix"`
	Status APITokenStatus `json:"status"`
	UsedAt *time.Time     `json:"used_at"`
}

type UserTokenPartial struct {
	ID     ulid.ULID      `json:"id"`
	Name   string         `json:"name"`
	Claims APITokenClaims `json:"claims"`
	Prefix string         `json:"prefix"`
	Status APITokenStatus `json:"status"`
	UsedAt *time.Time     `json:"used_at"`
}

type GenerateTokenRequest struct {
	Name        string              `json:"name"`
	Permissions APITokenPermissions `json:"permissions"`
}

type CreateTokenRequest struct {
	Name      string         `json:"name"`
	Status    APITokenStatus `json:"status"`
	TokenHash string         `json:"token"`
	Prefix    string         `json:"prefix"`
	Claims    APITokenClaims `json:"claims"`
}

type APITokenClaims struct {
	Permissions APITokenPermissions `json:"permissions"`
}

type APITokenPermissions struct {
	DeckCreate          bool `json:"deck_create,omitempty"`
	DeckDelete          bool `json:"deck_delete,omitempty"`
	DeckUpdate          bool `json:"deck_update,omitempty"`
	SubjectCreate       bool `json:"subject_create,omitempty"`
	SubjectUpdate       bool `json:"subject_update,omitempty"`
	SubjectDelete       bool `json:"subject_delete,omitempty"`
	ReviewCreate        bool `json:"review_create,omitempty"`
	StudyMaterialCreate bool `json:"study_material_create,omitempty"`
	StudyMaterialUpdate bool `json:"study_material_update,omitempty"`
	UserUpdate          bool `json:"user_update,omitempty"`
	UserDelete          bool `json:"user_delete,omitempty"`
}

type APITokenPermission string

const (
	TokenPermissionDeckCreate          APITokenPermission = "deck:create"
	TokenPermissionDeckDelete          APITokenPermission = "deck:delete"
	TokenPermissionDeckUpdate          APITokenPermission = "deck:update"
	TokenPermissionSubjectCreate       APITokenPermission = "subject:create"
	TokenPermissionSubjectUpdate       APITokenPermission = "subject:update"
	TokenPermissionSubjectDelete       APITokenPermission = "subject:delete"
	TokenPermissionReviewCreate        APITokenPermission = "review:create"
	TokenPermissionStudyMaterialCreate APITokenPermission = "study_material:create"
	TokenPermissionStudyMaterialUpdate APITokenPermission = "study_material:update"
	TokenPermissionUserUpdate          APITokenPermission = "user:update"
	TokenPermissionUserDelete          APITokenPermission = "user:delete"
)

type APITokenStatus string

const (
	TokenStatusActive APITokenStatus = "active"
	TokenStatusFrozen APITokenStatus = "frozen"
)

// Values provides list valid values for Enum.
func (APITokenStatus) Values() (statuses []string) {
	return []string{
		string(TokenStatusActive),
		string(TokenStatusFrozen),
	}
}
