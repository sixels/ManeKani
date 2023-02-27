package cards

import (
	"time"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards/filters"
)

type (
	Card struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		Subject PartialSubject `json:"subject"`

		Progress    uint8 `json:"progress"`
		TotalErrors int32 `json:"total_errors"`

		UnlockedAt  *time.Time `json:"unlocked_at"`
		StartedAt   *time.Time `json:"started_at"`
		PassedAt    *time.Time `json:"passed_at"`
		AvailableAt *time.Time `json:"available_at"`
		BurnedAt    *time.Time `json:"burned_at"`
	}

	PartialCard struct {
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

	CreateCardRequest struct {
		SubjectID uuid.UUID `json:"subject_id"`

		Progress    *uint8 `json:"progress"`
		TotalErrors *int32 `json:"total_errors"`

		UnlockedAt  *time.Time `json:"unlocked_at"`
		AvailableAt *time.Time `json:"available_at"`
	}

	UpdateCardRequest struct {
		Progress    *uint8 `json:"progress"`
		TotalErrors *int32 `json:"total_errors"`

		UnlockedAt  **time.Time `json:"unlocked_at"`
		StartedAt   **time.Time `json:"started_at"`
		PassedAt    **time.Time `json:"passed_at"`
		AvailableAt **time.Time `json:"available_at"`
		BurnedAt    **time.Time `json:"burned_at"`
	}

	QueryManyCardsRequest struct {
		AvailableAfter  *time.Time `query:"available_after"`
		AvailableBefore *time.Time `query:"available_before"`

		IsUnlocked *bool `query:"is_unlocked" binding:"-"`
		IsBurned   *bool `query:"is_burned" binding:"-"`
		IsStarted  *bool `query:"is_started" binding:"-"`
		IsPassed   *bool `query:"is_started" binding:"-"`

		ProgressAfter  *uint8 `query:"progress_after"`
		ProgressBefore *uint8 `query:"progress_before"`

		filters.FilterIDs
		filters.FilterDecks
		filters.FilterLevels
		filters.FilterPagination

		// internal use filters
		WithDependencies *filters.CommaSeparatedUUID `query:"-"`
		WithDependents   *filters.CommaSeparatedUUID `query:"-"`
	}
)
