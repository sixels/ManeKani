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
		Name            string  `json:"name"`
		Level           int32   `json:"level"`
		Symbol          *string `json:"symbol,omitempty"`
		SymbolImage     *[]byte `json:"symbol_image,omitempty"`
		MeaningMnemonic string  `json:"meaning_mnemonic"`
	}

	UpdateRadicalRequest struct {
		Symbol          *string `json:"symbol,omitempty"`
		Level           *int32  `json:"level,omitempty"`
		MeaningMnemonic *string `json:"meaning_mnemonic,omitempty"`
	}

	PartialRadicalResponse struct {
		Id     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Symbol *string   `json:"symbol,omitempty"`
		Level  int32     `json:"level"`
	}
)
