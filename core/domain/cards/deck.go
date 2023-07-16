package cards

import (
	"time"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards/filters"
)

type (
	// TODO: deck image
	Deck struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		Name        string `json:"name"`
		Description string `json:"Description"`

		Subjects []uuid.UUID `json:"subjects"`
		Owner    string      `json:"owner"`
	}
	DeckPartial struct {
		ID uuid.UUID `json:"id"`

		Name        string `json:"name"`
		Description string `json:"description"`

		Owner string `json:"owner"`
	}

	CreateDeckRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	QueryManyDecksRequest struct {
		filters.FilterPagination
		filters.FilterIDs
		filters.FilterSubjects
		filters.FilterOwners
		filters.FilterNames
	}
)
