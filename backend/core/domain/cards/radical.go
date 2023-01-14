package cards

import (
	"time"

	"github.com/google/uuid"
)

type (
	Radical struct {
		Id              uuid.UUID `json:"id"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		Name            string    `json:"name"`
		Level           int32     `json:"level"`
		Symbol          *string   `json:"symbol,omitempty"`
		MeaningMnemonic string    `json:"meaning_mnemonic"`
	}

	CreateRadicalRequest struct {
		Name   string  `json:"name" form:"name"`
		Level  int32   `json:"level" form:"level"`
		Symbol *string `json:"symbol,omitempty" form:"symbol"`
		// SymbolImage     *[]byte `json:"symbol_image,omitempty" form:"symbol_image"`
		MeaningMnemonic string `json:"meaning_mnemonic" form:"meaning_mnemonic"`
	}

	UpdateRadicalRequest struct {
		Symbol          *string `json:"symbol,omitempty" form:"symbol"`
		Level           *int32  `json:"level,omitempty" form:"level"`
		MeaningMnemonic *string `json:"meaning_mnemonic,omitempty" form:"meaning_mnemonic"`
	}

	QueryAllRadicalRequest struct {
		FilterLevel
	}

	PartialRadicalResponse struct {
		Id     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Symbol *string   `json:"symbol,omitempty"`
		Level  int32     `json:"level"`
	}
)
