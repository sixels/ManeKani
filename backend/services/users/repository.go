package users

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"sixels.io/manekani/services/ent"
)

type UsersRepository struct {
	client *ent.EntClient
}

func NewRepository(client *ent.EntClient) *UsersRepository {
	return &UsersRepository{
		client: client,
	}
}
