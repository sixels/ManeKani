package cards

import (
	"time"

	"sixels.io/manekani/services/cards/ent/schema"

	"github.com/google/uuid"
)

type (
	Vocabulary struct {
		Id              uuid.UUID         `json:"id"`
		CreatedAt       time.Time         `json:"created_at"`
		UpdatedAt       time.Time         `json:"updated_at"`
		Name            string            `json:"name"`
		Level           int32             `json:"level"`
		AltNames        []string          `json:"alt_names,omitempty"`
		Word            string            `json:"word"`
		WordType        []string          `json:"word_type"`
		Reading         string            `json:"reading"`
		AltReadings     []string          `json:"alt_readings,omitempty"`
		Patterns        []schema.Pattern  `json:"patterns"`
		Sentences       []schema.Sentence `json:"sentences"`
		MeaningMnemonic string            `json:"meaning_mnemonic"`
		ReadingMnemonic string            `json:"reading_mnemonic"`
	}

	CreateVocabularyRequest struct {
		Name             string            `json:"name"`
		Level            int32             `json:"level"`
		AltNames         []string          `json:"alt_names,omitempty"`
		Word             string            `json:"word"`
		WordType         []string          `json:"word_type"`
		Reading          string            `json:"reading"`
		AltReadings      []string          `json:"alt_readings,omitempty"`
		MeaningMnemonic  string            `json:"meaning_mnemonic"`
		ReadingMnemonic  string            `json:"reading_mnemonic"`
		Patterns         []schema.Pattern  `json:"patterns"`
		Sentences        []schema.Sentence `json:"sentences"`
		KanjiComposition []string          `json:"kanji_composition,omitempty"`
	}

	UpdateVocabularyRequest struct {
		Name             *string            `json:"name,omitempty"`
		Level            *int32             `json:"level,omitempty"`
		AltNames         *[]string          `json:"alt_names,omitempty"`
		WordType         *[]string          `json:"word_type,omitempty"`
		Reading          *string            `json:"reading,omitempty"`
		AltReadings      *[]string          `json:"alt_readings,omitempty"`
		MeaningMnemonic  *string            `json:"meaning_mnemonic,omitempty"`
		ReadingMnemonic  *string            `json:"reading_mnemonic,omitempty"`
		Patterns         *[]schema.Pattern  `json:"patterns,omitempty"`
		Sentences        *[]schema.Sentence `json:"sentences,omitempty"`
		KanjiComposition *[]string          `json:"kanji_composition,omitempty"`
	}

	PartialVocabularyResponse struct {
		Id       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		AltNames []string  `json:"alt_names"`
		Reading  string    `json:"reading"`
		Word     string    `json:"word"`
		Level    int32     `json:"level"`
	}
)
