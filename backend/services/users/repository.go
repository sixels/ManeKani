package users

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"sixels.io/manekani/services/cards"
	"sixels.io/manekani/services/ent"
)

type UsersRepository struct {
	client *ent.EntClient
	cards  *cards.CardsRepository
}

func NewRepository(client *ent.EntClient, cards *cards.CardsRepository) *UsersRepository {
	return &UsersRepository{
		client: client,
		cards:  cards,
	}
}
