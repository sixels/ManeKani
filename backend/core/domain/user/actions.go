package user

import "sixels.io/manekani/ent/schema"

const (
	CREATE_LEVEL_CARDS schema.Action = iota
)

type (
	CreateLevelCardsMeta struct {
		Level int32
	}
)
