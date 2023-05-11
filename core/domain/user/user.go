package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/sixels/manekani/ent/schema"
)

type User struct {
	ID       string      `json:"-"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Decks    []uuid.UUID `json:"decks"`

	PendingActions []schema.PendingAction `json:"-"`
}

type UserBasic struct {
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Decks    []uuid.UUID `json:"decks"`
}

type CreateUserRequest struct {
	ID        string     `json:"id" form:"id"`
	Email     string     `json:"email" form:"email"`
	Username  string     `json:"username" form:"username"`
	CreatedAt *time.Time `json:"created_at,omitempty" form:"created_at"`
}
