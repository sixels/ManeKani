package user

import (
	"github.com/google/uuid"
	"github.com/sixels/manekani/ent/schema"
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
