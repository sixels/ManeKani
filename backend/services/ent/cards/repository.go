package cards

import (
	"context"

	"sixels.io/manekani/core/ports"
	"sixels.io/manekani/core/ports/transactions"
	ent_repo "sixels.io/manekani/services/ent"
)

var _ ports.CardsRepository = (*CardsRepository)(nil)

type CardsRepository struct {
	client *ent_repo.EntRepository
}

func NewRepository(client *ent_repo.EntRepository) *CardsRepository {
	return &CardsRepository{
		client: client,
	}
}

func (repo *CardsRepository) BeginTransaction(ctx context.Context) (transactions.TransactionalRepository, error) {
	cli, err := repo.client.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}
	return &CardsRepository{
		client: cli.(*ent_repo.EntRepository),
	}, nil
}

func (repo *CardsRepository) Rollback() error {
	return repo.client.Rollback()
}
func (repo *CardsRepository) Commit() error {
	return repo.client.Commit()
}
