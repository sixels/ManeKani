package tokens

import (
	"github.com/google/uuid"
)

type GenerateTokenRequest struct {
	Name        string              `json:"name"`
	Permissions APITokenPermissions `json:"permissions"`
}

type CreateTokenRequest struct {
	TokenHash string         `json:"token"`
	Prefix    string         `json:"prefix"`
	Claims    APITokenClaims `json:"claims"`
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

type UserToken struct {
	ID     uuid.UUID      `json:"id"`
	Claims APITokenClaims `json:"claims"`
	Prefix string         `json:"prefix"`
	UserID string         `json:"user_id"`
}

type UserTokenPartial struct {
	ID     uuid.UUID      `json:"id"`
	Prefix string         `json:"prefix"`
	Claims APITokenClaims `json:"claims"`
}

type APITokenClaims struct {
	Permissions APITokenPermissions `json:"permissions"`
}

type APITokenPermissions struct {
	TokenPermissionDeckCreate          bool `json:"token_permission_deck_create,omitempty"`
	TokenPermissionDeckDelete          bool `json:"token_permission_deck_delete,omitempty"`
	TokenPermissionDeckUpdate          bool `json:"token_permission_deck_update,omitempty"`
	TokenPermissionSubjectCreate       bool `json:"token_permission_subject_create,omitempty"`
	TokenPermissionSubjectUpdate       bool `json:"token_permission_subject_update,omitempty"`
	TokenPermissionSubjectDelete       bool `json:"token_permission_subject_delete,omitempty"`
	TokenPermissionReviewCreate        bool `json:"token_permission_review_create,omitempty"`
	TokenPermissionStudyMaterialCreate bool `json:"token_permission_study_material_create,omitempty"`
	TokenPermissionStudyMaterialUpdate bool `json:"token_permission_study_material_update,omitempty"`
	TokenPermissionUserUpdate          bool `json:"token_permission_user_update,omitempty"`
	TokenPermissionUserDelete          bool `json:"token_permission_user_delete,omitempty"`
}
