package cards

import (
	"time"

	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards/filters"
)

type (
	Review struct {
		ID            uuid.UUID    `json:"id"`
		CreatedAt     time.Time    `json:"created_at"`
		Errors        ReviewErrors `json:"errors"`
		StartProgress uint8        `json:"start_progress"`
		EndProgress   uint8        `json:"end_progress"`
		Card          uuid.UUID    `json:"card"`
	}

	CreateReviewAPIRequest struct {
		CardID      filters.CommaSeparatedUUID `json:"card" form:"card" binding:"required"`
		Errors      ReviewErrors               `json:"errors" form:"errors" binding:"required"`
		SessionType SessionType                `json:"session_type" query:"session_type" form:"session_type"`
	}
	CreateReviewRequest struct {
		CardID        uuid.UUID    `json:"card" form:"card" binding:"required"`
		Errors        ReviewErrors `json:"errors" form:"errors" binding:"required"`
		StartProgress uint8        `json:"start_progress" form:"start_progress" binding:"required"`
		EndProgress   uint8        `json:"end_progress" form:"end_progress" binding:"required"`
	}

	QueryManyReviewsRequest struct {
		Passed *bool `query:"failed"`

		filters.FilterIDs
		filters.FilterCreatedTime
		filters.FilterCards
		filters.FilterPagination
	}

	ReviewErrors = map[string]int32
)
