package users

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/users/crypto"
	"sixels.io/manekani/services/ent/util"
	mkjwt "sixels.io/manekani/services/jwt"
)

func (repo *UsersRepository) CreateUserAPITokenTX(ctx context.Context, tx *ent.Tx, userID string, options mkjwt.APITokenOptions) (string, error) {
	signedToken, err := repo.jwt.CreateAPIToken(options)
	if err != nil {
		return "", fmt.Errorf("could not generate the api token: %w", err)
	}

	encryptedToken, err := crypto.Aes256Encode([]byte(signedToken), repo.tokenEncryptionKey)
	if err != nil {
		return "", fmt.Errorf("could not encrypt the api token: %w", err)
	}

	if err := tx.ApiToken.Create().
		SetToken(encryptedToken).
		SetUserID(userID).
		Exec(ctx); err != nil {
		return "", fmt.Errorf("could not store the api token: %w", err)
	}

	return signedToken, nil
}

func (repo *UsersRepository) DumpUserAPITokens(ctx context.Context, userID string) ([]*jwt.Token, error) {
	tokens, err := repo.client.User.Query().
		Where(user.IDEQ(userID)).
		QueryAPITokens().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return util.MapArray(tokens, func(encryptedToken *ent.ApiToken) *jwt.Token {
		token, err := crypto.Aes256Decode(encryptedToken.Token, repo.tokenEncryptionKey)
		if err != nil {
			return nil
		}

		// TODO: REMOVE ME
		log.Println(string(token))

		tk, _ := repo.jwt.ValidateToken(string(token), mkjwt.APITokenClaims{})
		return tk
	}), nil
}

func getDefault[T any](value *T, def T) T {
	return getWithDefault(value, func() T { return def })
}

func getWithDefault[T any](value *T, def func() T) T {
	if value != nil {
		return *value
	}
	return def()
}
