package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID          uuid.UUID  `json:"id"`
	Email       string     `json:"email"`
	Username    *string    `json:"username"`
	DisplayName *string    `json:"display_name"`
	IsVerified  bool       `json:"is_verified"`
	IsComplete  bool       `json:"is_complete"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeckIds     uuid.UUIDs `json:"deck_ids"`
}
