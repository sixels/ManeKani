package user

import (
	"github.com/google/uuid"
	"sixels.io/manekani/ent/schema"
)

const (
	SUBSCRIBE_TO_DECK schema.Action = iota
)

type (
	CreateLevelCardsMeta struct {
		Deck uuid.UUID
	}
)
