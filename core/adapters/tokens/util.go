package tokens

import "github.com/sixels/manekani/core/domain/tokens"

func MapCapabilities(claims tokens.APITokenClaims) map[tokens.APITokenCapability]bool {
	return map[tokens.APITokenCapability]bool{
		tokens.TokenCapabiltyDeckCreate:           claims.Capabilities.TokenCapabiltyDeckCreate,
		tokens.TokenCapabiltyDeckDelete:           claims.Capabilities.TokenCapabiltyDeckDelete,
		tokens.TokenCapabiltyDeckUpdate:           claims.Capabilities.TokenCapabiltyDeckUpdate,
		tokens.TokenCapabilitySubjectCreate:       claims.Capabilities.TokenCapabilitySubjectCreate,
		tokens.TokenCapabilitySubjectUpdate:       claims.Capabilities.TokenCapabilitySubjectUpdate,
		tokens.TokenCapabilitySubjectDelete:       claims.Capabilities.TokenCapabilitySubjectDelete,
		tokens.TokenCapabilityReviewCreate:        claims.Capabilities.TokenCapabilityReviewCreate,
		tokens.TokenCapabilityStudyMaterialCreate: claims.Capabilities.TokenCapabilityStudyMaterialCreate,
		tokens.TokenCapabilityStudyMaterialUpdate: claims.Capabilities.TokenCapabilityStudyMaterialUpdate,
		tokens.TokenCapabilityUserUpdate:          claims.Capabilities.TokenCapabilityUserUpdate,
		tokens.TokenCapabilityUserDelete:          claims.Capabilities.TokenCapabilityUserDelete,
	}
}
