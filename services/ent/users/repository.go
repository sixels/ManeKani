package users

import (
	"context"

	"github.com/sixels/manekani/core/ports/transactions"
	ent_repo "github.com/sixels/manekani/services/ent"
)

type UsersRepository struct {
	client *ent_repo.EntRepository
}

func NewRepository(client *ent_repo.EntRepository) *UsersRepository {
	repo := UsersRepository{
		client: client,
	}
	return &repo
}

// Implement TransactionalRepository

func (repo *UsersRepository) BeginTransaction(ctx context.Context) (transactions.TransactionalRepository, error) {
	cli, err := repo.client.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}
	return &UsersRepository{
		client: cli.(*ent_repo.EntRepository),
	}, nil
}

func (repo *UsersRepository) Rollback() error {
	return repo.client.Rollback()
}
func (repo *UsersRepository) Commit() error {
	return repo.client.Commit()
}
