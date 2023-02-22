package user

import (
	"github.com/google/uuid"
	"sixels.io/manekani/ent/schema"
)

const (
	ActionCheckCardUnlocks schema.Action = iota
	ActionCheckDeckLevelUp schema.Action = iota // TODO
)

type (
	CreateLevelCardsMeta struct {
		Deck uuid.UUID
	}
	CheckCardUnlocksMeta struct {
		Card uuid.UUID
	}
)
