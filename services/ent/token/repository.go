package token

import (
	"context"
	"os"

	"github.com/sixels/manekani/core/ports/transactions"
	ent_repo "github.com/sixels/manekani/services/ent"
)

type TokenRepository struct {
	client             *ent_repo.EntRepository
	tokenEncryptionKey []byte
}

func NewRepository(client *ent_repo.EntRepository) *TokenRepository {
	encryptionKey := os.Getenv("TOKEN_ENCRYPTION_KEY")
	repo := TokenRepository{
		client:             client,
		tokenEncryptionKey: []byte(	encryptionKey),
	}
	return &repo
}

func (repo *TokenRepository) BeginTransaction(ctx context.Context) (transactions.TransactionalRepository, error) {
	cli, err := repo.client.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}
	return &TokenRepository{
		client:             cli.(*ent_repo.EntRepository),
		tokenEncryptionKey: repo.tokenEncryptionKey,
	}, nil
}

// Implement TransactionalRepository

func (repo *TokenRepository) Rollback() error {
	return repo.client.Rollback()
}
func (repo *TokenRepository) Commit() error {
	return repo.client.Commit()
}
