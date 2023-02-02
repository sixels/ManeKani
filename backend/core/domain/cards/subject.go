package cards

import "github.com/google/uuid"

type (
	Subject struct {
		Id uuid.UUID `json:"id"`

		Kind  string `json:"kind"`
		Level int32  `json:"level"`

		RadicalData    *PartialRadicalResponse    `json:"radical,omitempty"`
		KanjiData      *PartialKanjiResponse      `json:"kanji,omitempty"`
		VocabularyData *PartialVocabularyResponse `json:"vocabulary,omitempty"`
	}
	QueryAllSubjectsRequest struct {
		FilterLevel
	}
	PartialSubjectResponse struct {
		Id uuid.UUID `json:"id"`

		Kind  string `json:"kind"`
		Level int32  `json:"level"`
	}
)
