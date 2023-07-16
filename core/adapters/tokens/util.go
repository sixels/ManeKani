package tokens

import "github.com/sixels/manekani/core/domain/tokens"

func MapPermissions(perms tokens.APITokenPermissions) map[tokens.APITokenPermission]bool {
	return map[tokens.APITokenPermission]bool{
		tokens.TokenPermissionDeckCreate:          perms.DeckCreate,
		tokens.TokenPermissionDeckDelete:          perms.DeckDelete,
		tokens.TokenPermissionDeckUpdate:          perms.DeckUpdate,
		tokens.TokenPermissionSubjectCreate:       perms.SubjectCreate,
		tokens.TokenPermissionSubjectUpdate:       perms.SubjectUpdate,
		tokens.TokenPermissionSubjectDelete:       perms.SubjectDelete,
		tokens.TokenPermissionReviewCreate:        perms.ReviewCreate,
		tokens.TokenPermissionStudyMaterialCreate: perms.StudyMaterialCreate,
		tokens.TokenPermissionStudyMaterialUpdate: perms.StudyMaterialUpdate,
		tokens.TokenPermissionUserUpdate:          perms.UserUpdate,
		tokens.TokenPermissionUserDelete:          perms.UserDelete,
	}
}
