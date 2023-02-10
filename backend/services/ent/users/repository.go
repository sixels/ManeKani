package users

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/schema"
	"sixels.io/manekani/ent/user"
	client "sixels.io/manekani/services/ent"
	"sixels.io/manekani/services/ent/util"
	"sixels.io/manekani/services/jwt"
	mkjwt "sixels.io/manekani/services/jwt"
)

type UsersRepository struct {
	client             *client.EntRepository
	tokenEncryptionKey []byte

	jwt *mkjwt.JWTService
}

func NewRepository(ctx context.Context, client *client.EntRepository, jwtService *mkjwt.JWTService) (*UsersRepository, error) {
	repo := UsersRepository{
		client:             client,
		jwt:                jwtService,
		tokenEncryptionKey: []byte(os.Getenv("TOKEN_ENCRYPTION_KEY")),
	}

	if err := createManeKaniUser(ctx, &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

func createManeKaniUser(ctx context.Context, repo *UsersRepository) error {
	if exists, err := repo.client.User.Query().
		Where(user.UsernameEQ("manekani")).
		Exist(ctx); err != nil || exists {
		return err
	}
	log.Println("creating admin user")

	_, err := util.WithTx(ctx, repo.client.Client, func(tx *ent.Tx) (*struct{}, error) {

		usr, err := tx.User.Create().
			SetEmail("admin@manekani.io").
			SetUsername("manekani").
			SetPendingActions([]schema.PendingAction{}).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not create the manekani user: %w", err)
		}

		tokenExpiration := time.Now().AddDate(10, 0, 0)
		_, err = repo.CreateUserAPITokenTX(ctx, tx, usr.ID, mkjwt.APITokenOptions{
			UserID: usr.ID,
			Scope:  jwt.TokenScopeGlobal,
			Capabilities: []jwt.APITokenCapability{
				mkjwt.TokenCapabiltyDeckCreate,
				mkjwt.TokenCapabiltyDeckDelete,
				mkjwt.TokenCapabiltyDeckUpdate,
				mkjwt.TokenCababiltySubjectCreate,
				mkjwt.TokenCababiltySubjectUpdate,
				mkjwt.TokenCababiltySubjectDelete,
				mkjwt.TokenCababiltyReviewCreate,
			},
			ExpiresAt: &tokenExpiration,
		})
		if err != nil {
			return nil, err
		}

		// create the default deck
		err = tx.Deck.Create().
			SetName("WaniKani").
			SetDescription("All WaniKani cards to help you learn japanese fast!").
			SetOwnerID(usr.ID).
			Exec(ctx)

		return nil, err
	})

	admin := repo.client.User.Query().
		Where(user.UsernameEQ("manekani")).
		OnlyX(ctx)

	_, _ = repo.DumpUserAPITokens(ctx, admin.ID)

	return err
}
