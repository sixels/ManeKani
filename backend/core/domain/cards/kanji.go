package cards

import (
	"time"

	"github.com/google/uuid"
)

type (
	Kanji struct {
		Id              uuid.UUID `json:"id"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		Name            string    `json:"name"`
		Level           int32     `json:"level"`
		AltNames        []string  `json:"alt_names,omitempty"`
		Symbol          string    `json:"symbol"`
		Reading         string    `json:"reading"`
		Onyomi          []string  `json:"onyomi"`
		Kunyomi         []string  `json:"kunyomi"`
		Nanori          []string  `json:"nanori"`
		MeaningMnemonic string    `json:"meaning_mnemonic"`
		ReadingMnemonic string    `json:"reading_mnemonic"`
	}

	CreateKanjiRequest struct {
		Name               string   `json:"name" form:"name"`
		Level              int32    `json:"level" form:"level"`
		AltNames           []string `json:"alt_names,omitempty" form:"alt_names"`
		Symbol             string   `json:"symbol" form:"symbol"`
		Reading            string   `json:"reading" form:"reading"`
		Onyomi             []string `json:"onyomi" form:"onyomi"`
		Kunyomi            []string `json:"kunyomi" form:"kunyomi"`
		Nanori             []string `json:"nanori" form:"nanori"`
		MeaningMnemonic    string   `json:"meaning_mnemonic" form:"meaning_mnemonic"`
		ReadingMnemonic    string   `json:"reading_mnemonic" form:"reading_mnemonic"`
		RadicalComposition []string `json:"radical_composition,omitempty" form:"radical_composition"`
	}

	UpdateKanjiRequest struct {
		Level              *int32    `json:"level,omitempty"`
		Name               *string   `json:"name,omitempty"`
		AltNames           *[]string `json:"alt_names,omitempty"`
		MeaningMnemonic    *string   `json:"meaning_mnemonic,omitempty"`
		Reading            *string   `json:"reading,omitempty"`
		ReadingMnemonic    *string   `json:"reading_mnemonic,omitempty"`
		Onyomi             *[]string `json:"onyomi,omitempty"`
		Kunyomi            *[]string `json:"kunyomi,omitempty"`
		Nanori             *[]string `json:"nanori,omitempty"`
		RadicalComposition *[]string `json:"radical_composition,omitempty"`
	}

	PartialKanjiResponse struct {
		Id       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		AltNames []string  `json:"alt_names,omitempty"`
		Reading  string    `json:"reading"`
		Symbol   string    `json:"symbol"`
		Level    int32     `json:"level"`
	}
)
