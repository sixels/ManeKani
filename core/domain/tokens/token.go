package tokens

import (
	"encoding/json"

	"github.com/google/uuid"
)

type GenerateTokenRequest struct {
	Name        string               `json:"name"`
	Permissions APITokenCapabilities `json:"permissions"`
}

type CreateTokenRequest struct {
	TokenHash string         `json:"token"`
	Prefix    string         `json:"prefix"`
	Claims    APITokenClaims `json:"claims"`
}

type APITokenCapability string

const (
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
	Capabilities APITokenCapabilities `json:"capabilities"`
}

type APITokenCapabilities struct {
	TokenCapabiltyDeckCreate           bool
	TokenCapabiltyDeckDelete           bool
	TokenCapabiltyDeckUpdate           bool
	TokenCapabilitySubjectCreate       bool
	TokenCapabilitySubjectUpdate       bool
	TokenCapabilitySubjectDelete       bool
	TokenCapabilityReviewCreate        bool
	TokenCapabilityStudyMaterialCreate bool
	TokenCapabilityStudyMaterialUpdate bool
	TokenCapabilityUserUpdate          bool
	TokenCapabilityUserDelete          bool
}

// implement json marshal to APITokenCapabilities
func (c APITokenCapabilities) MarshalJSON() ([]byte, error) {
	capsMap := MapTokenCapabilities(&c)
	caps := make([]APITokenCapability, 0, len(capsMap))
	for cap, isDefined := range capsMap {
		if *isDefined {
			caps = append(caps, cap)
		}
	}

	return json.Marshal(caps)
}

func (c *APITokenCapabilities) UnmarshalJSON(data []byte) error {
	caps := make([]APITokenCapability, 0)
	if err := json.Unmarshal(data, &caps); err != nil {
		return err
	}

	capsMap := MapTokenCapabilities(c)
	for _, cap := range caps {
		if claim, ok := capsMap[cap]; ok {
			*claim = true
		}
	}

	return nil
}

func MapTokenCapabilities(c *APITokenCapabilities) map[APITokenCapability]*bool {
	return map[APITokenCapability]*bool{
		TokenCapabiltyDeckCreate:           &c.TokenCapabilityReviewCreate,
		TokenCapabiltyDeckDelete:           &c.TokenCapabiltyDeckDelete,
		TokenCapabiltyDeckUpdate:           &c.TokenCapabiltyDeckUpdate,
		TokenCapabilitySubjectCreate:       &c.TokenCapabilitySubjectCreate,
		TokenCapabilitySubjectUpdate:       &c.TokenCapabilitySubjectUpdate,
		TokenCapabilitySubjectDelete:       &c.TokenCapabilitySubjectDelete,
		TokenCapabilityReviewCreate:        &c.TokenCapabilityReviewCreate,
		TokenCapabilityStudyMaterialCreate: &c.TokenCapabilityStudyMaterialCreate,
		TokenCapabilityStudyMaterialUpdate: &c.TokenCapabilityStudyMaterialUpdate,
		TokenCapabilityUserUpdate:          &c.TokenCapabilityUserUpdate,
		TokenCapabilityUserDelete:          &c.TokenCapabilityUserDelete,
	}
}
