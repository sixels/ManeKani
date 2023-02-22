package cards

import (
	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards/filters"
)

type (
	SessionType string
)

const (
	SessionLesson SessionType = "lesson"
	SessionReview SessionType = "review"
)

type (
	CreateStudySessionRequest struct {
		DeckID      filters.CommaSeparatedUUID `json:"deck" query:"deck" form:"deck"`
		SessionType SessionType                `json:"session_type" query:"session_type" form:"session_type"`
	}

	SessionQueue struct {
		Cards []*QueueItem `json:"cards"`
	}
	QueueItem struct {
		CardID  uuid.UUID          `json:"card_id"`
		Answers []QueueItemAnswers `json:"answers"`
		Subject MinimalSubject     `json:"subject"`
	}
	QueueItemAnswers struct {
		StudyItemType string   `json:"study_item_type"`
		Expected      []string `json:"expected"`
		Blacklisted   []string `json:"blacklisted"`
	}
)
