package cards

import "sixels.io/manekani/services/ent"

type CardsRepository struct {
	client *ent.EntClient
}

func NewRepository(client *ent.EntClient) *CardsRepository {
	return &CardsRepository{
		client: client,
	}
}
