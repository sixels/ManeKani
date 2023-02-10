package cards

import "sixels.io/manekani/services/ent"

type CardsRepository struct {
	client *ent.EntRepository
}

func NewRepository(client *ent.EntRepository) *CardsRepository {
	return &CardsRepository{
		client: client,
	}
}
