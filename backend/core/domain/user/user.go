package user

import (
	"github.com/google/uuid"
	"sixels.io/manekani/ent/schema"
)

type User struct {
	ID       string      `json:"-"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Decks    []uuid.UUID `json:"decks"`

	PendingActions []schema.PendingAction `json:"-"`
}

type UserBasic struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CreateUserRequest struct {
	ID       string `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
}
