package tokens

import "github.com/sixels/manekani/core/domain/tokens"

func MapPermissions(claims tokens.APITokenClaims) map[tokens.APITokenPermission]bool {
	return map[tokens.APITokenPermission]bool{
		tokens.TokenPermissionDeckCreate:          claims.Permissions.TokenPermissionDeckCreate,
		tokens.TokenPermissionDeckDelete:          claims.Permissions.TokenPermissionDeckDelete,
		tokens.TokenPermissionDeckUpdate:          claims.Permissions.TokenPermissionDeckUpdate,
		tokens.TokenPermissionSubjectCreate:       claims.Permissions.TokenPermissionSubjectCreate,
		tokens.TokenPermissionSubjectUpdate:       claims.Permissions.TokenPermissionSubjectUpdate,
		tokens.TokenPermissionSubjectDelete:       claims.Permissions.TokenPermissionSubjectDelete,
		tokens.TokenPermissionReviewCreate:        claims.Permissions.TokenPermissionReviewCreate,
		tokens.TokenPermissionStudyMaterialCreate: claims.Permissions.TokenPermissionStudyMaterialCreate,
		tokens.TokenPermissionStudyMaterialUpdate: claims.Permissions.TokenPermissionStudyMaterialUpdate,
		tokens.TokenPermissionUserUpdate:          claims.Permissions.TokenPermissionUserUpdate,
		tokens.TokenPermissionUserDelete:          claims.Permissions.TokenPermissionUserDelete,
	}
}
