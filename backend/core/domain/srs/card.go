package srs

import (
	"time"

	"github.com/google/uuid"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/cards/filters"
)

type (
	Card struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		Subject cards.Subject `json:"subject"`

		Progress    uint8 `json:"progress"`
		TotalErrors int32 `json:"total_errors"`

		UnlockedAt  *time.Time `json:"unlocked_at"`
		StartedAt   *time.Time `json:"started_at"`
		PassedAt    *time.Time `json:"passed_at"`
		AvailableAt *time.Time `json:"available_at"`
		BurnedAt    *time.Time `json:"burned_at"`
	}

	CardRaw struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		Progress    uint8 `json:"progress"`
		TotalErrors int32 `json:"total_errors"`

		UnlockedAt  *time.Time `json:"unlocked_at,omitempty"`
		StartedAt   *time.Time `json:"started_at,omitempty"`
		PassedAt    *time.Time `json:"passed_at,omitempty"`
		AvailableAt *time.Time `json:"available_at,omitempty"`
		BurnedAt    *time.Time `json:"burned_at,omitempty"`
	}

	QueryManyCardsRequest struct {
		filters.FilterIDs
		filters.FilterLevels
		FilterUser
	}
)

type FilterUser struct {
	userID   []string
	username []string
}
