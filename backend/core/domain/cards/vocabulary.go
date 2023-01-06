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
		Name             string            `json:"name" form:"name"`
		Level            int32             `json:"level" form:"level"`
		AltNames         []string          `json:"alt_names,omitempty" form:"alt_names"`
		Word             string            `json:"word" form:"word"`
		WordType         []string          `json:"word_type" form:"word_type"`
		Reading          string            `json:"reading" form:"reading"`
		AltReadings      []string          `json:"alt_readings,omitempty" form:"alt_readings"`
		MeaningMnemonic  string            `json:"meaning_mnemonic" form:"meaning_mnemonic"`
		ReadingMnemonic  string            `json:"reading_mnemonic" form:"reading_mnemonic"`
		Patterns         []schema.Pattern  `json:"patterns" form:"patterns"`
		Sentences        []schema.Sentence `json:"sentences" form:"sentences"`
		KanjiComposition []string          `json:"kanji_composition,omitempty" form:"kanji_composition"`
	}

	UpdateVocabularyRequest struct {
		Name             *string            `json:"name,omitempty" form:"name"`
		Level            *int32             `json:"level,omitempty" form:"level"`
		AltNames         *[]string          `json:"alt_names,omitempty" form:"alt_names"`
		WordType         *[]string          `json:"word_type,omitempty" form:"word_type"`
		Reading          *string            `json:"reading,omitempty" form:"reading"`
		AltReadings      *[]string          `json:"alt_readings,omitempty" form:"alt_readings"`
		MeaningMnemonic  *string            `json:"meaning_mnemonic,omitempty" form:"meaning_mnemonic"`
		ReadingMnemonic  *string            `json:"reading_mnemonic,omitempty" form:"reading_mnemonic"`
		Patterns         *[]schema.Pattern  `json:"patterns,omitempty" form:"patterns"`
		Sentences        *[]schema.Sentence `json:"sentences,omitempty" form:"sentences"`
		KanjiComposition *[]string          `json:"kanji_composition,omitempty" form:"kanji_composition"`
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
