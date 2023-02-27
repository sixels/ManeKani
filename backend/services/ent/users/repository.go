package users

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sixels/manekani/core/ports/transactions"
	"github.com/sixels/manekani/ent/schema"
	"github.com/sixels/manekani/ent/user"
	ent_repo "github.com/sixels/manekani/services/ent"
	"github.com/sixels/manekani/services/jwt"
	mkjwt "github.com/sixels/manekani/services/jwt"
)

type UsersRepository struct {
	client             *ent_repo.EntRepository
	tokenEncryptionKey []byte

	jwt *mkjwt.JWTService
}

func NewRepository(ctx context.Context, client *ent_repo.EntRepository, jwtService *mkjwt.JWTService) (*UsersRepository, error) {
	repo := UsersRepository{
		client:             client,
		jwt:                jwtService,
		tokenEncryptionKey: []byte(os.Getenv("TOKEN_ENCRYPTION_KEY")),
	}

	tx := transactions.Begin(ctx)
	txRepo, err := transactions.MakeTransactional(tx, &repo)
	if err != nil {
		return nil, err
	}

	if err := createAdminUser(ctx, txRepo); err != nil {
		return nil, err
	}

	return &repo, nil
}

func createAdminUser(ctx context.Context, repo *UsersRepository) error {
	if exists, err := repo.client.UserClient().Query().
		Where(user.UsernameEQ("manekani")).
		Exist(ctx); err != nil || exists {
		return err
	}
	log.Println("creating admin user")

	usr, err := repo.client.UserClient().Create().
		SetEmail("admin@manekani.io").
		SetUsername("manekani").
		SetPendingActions([]schema.PendingAction{}).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("could not create the manekani user: %w", err)
	}

	tokenExpiration := time.Now().AddDate(10, 0, 0)
	_, err = repo.CreateUserAPITokenTX(ctx, usr.ID, mkjwt.APITokenOptions{
		UserID: usr.ID,
		Scope:  jwt.TokenScopeGlobal,
		Capabilities: []jwt.APITokenCapability{
			mkjwt.TokenCapabiltyDeckCreate,
			mkjwt.TokenCapabiltyDeckDelete,
			mkjwt.TokenCapabiltyDeckUpdate,
			mkjwt.TokenCapabilitySubjectCreate,
			mkjwt.TokenCapabilitySubjectUpdate,
			mkjwt.TokenCapabilitySubjectDelete,
			mkjwt.TokenCapabilityReviewCreate,
		},
		ExpiresAt: &tokenExpiration,
	})
	if err != nil {
		return err
	}

	// create the default deck
	err = repo.client.DeckClient().Create().
		SetName("WaniKani").
		SetDescription("All WaniKani cards to help you learn japanese fast!").
		SetOwnerID(usr.ID).
		Exec(ctx)
	if err != nil {
		return err
	}

	admin := repo.client.UserClient().Query().
		Where(user.UsernameEQ("manekani")).
		OnlyX(ctx)

	_, _ = repo.DumpUserAPITokens(ctx, admin.ID)

	return err
}

func (repo *UsersRepository) BeginTransaction(ctx context.Context) (transactions.TransactionalRepository, error) {
	cli, err := repo.client.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}
	return &UsersRepository{
		client:             cli.(*ent_repo.EntRepository),
		tokenEncryptionKey: repo.tokenEncryptionKey,
		jwt:                repo.jwt,
	}, nil
}

func (repo *UsersRepository) Rollback() error {
	return repo.client.Rollback()
}
func (repo *UsersRepository) Commit() error {
	return repo.client.Commit()
}
