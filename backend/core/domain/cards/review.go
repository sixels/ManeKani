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

	CreateReviewRequest struct {
		Card        filters.CommaSeparatedUUID `json:"card" form:"card" binding:"required"`
		Errors      ReviewErrors               `json:"errors" form:"errors" binding:"required"`
		SessionType SessionType                `json:"session_type" form:"session_type" binding:"required"`
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
