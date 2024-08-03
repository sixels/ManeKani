package db

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sixels/manekani/services/user-management/domain"
)

type CreateUserInput struct {
	Email      string
	Username   *string
	IsVerified bool
	IsComplete bool
}

var (
	ErrCreateUserDuplicateEmail    = errors.New("email already exists")
	ErrCreateUserDuplicateUsername = errors.New("username already exists")
)

type createUserResult struct {
	ID          uuid.UUID `db:"id"`
	Email       string    `db:"email"`
	Username    *string   `db:"username"`
	DisplayName *string   `db:"display_name"`
	IsVerified  bool      `db:"is_verified"`
	IsComplete  bool      `db:"is_complete"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (db *DB) CreateUser(ctx context.Context, user CreateUserInput) (*domain.User, error) {
	rows, err := db.pool.Query(ctx,
		`INSERT INTO users (email, username, is_verified, is_complete)
	VALUES ($1, $2, $3, $4)
	RETURNING id, email, username, is_verified, is_complete, display_name, created_at, updated_at`,
		user.Email, user.Username, user.IsVerified, user.IsComplete,
	)

	if err != nil {
		return nil, unknownError(err)
	}

	created, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[createUserResult])

	var pgxErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgxErr) {
			log.Printf("pgx error: %s\n", pgxErr.Code)
			if pgxErr.Code == "23505" {
				switch pgxErr.ConstraintName {
				case "user_username_unique":
					return nil, ErrCreateUserDuplicateUsername
				case "user_email_unique":
					return nil, ErrCreateUserDuplicateEmail
				}
			}
		}
		return nil, unknownError(err)
	}

	return &domain.User{
		ID:         created.ID,
		Email:      created.Email,
		Username:   created.Username,
		IsVerified: created.IsVerified,
		IsComplete: created.IsComplete,
		CreatedAt:  created.CreatedAt,
		UpdatedAt:  created.UpdatedAt,
		DeckIds:    nil,
	}, nil
}
