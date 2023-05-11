package tokens

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	"github.com/sixels/manekani/core/adapters/tokens/hash"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/services/ent/users/crypto"
)

type (
	GetTokenError    string
	CreateTokenError string
)

const (
	GetTokenErrorInvalidToken GetTokenError    = "invalid token"
	CreateTokenErrorHash      CreateTokenError = "could not hash the api token"
	CreateTokenErrorStore     CreateTokenError = "could not store the api token"
)

const (
	PREFIX_LEN       int    = 8
	PREFIX_SEPARATOR string = "-"
)

func (service *TokensAdapter) GetToken(ctx context.Context, key string) (*tokens.UserToken, error) {
	tokenParts := strings.SplitN(key, PREFIX_SEPARATOR, 2)
	if len(tokenParts) != 2 {
		return nil, errors.New(string(GetTokenErrorInvalidToken))
	}

	prefix, tokenEncoded := tokenParts[0], tokenParts[1]

	log.Println("prefix:", prefix)

	tokenBytes := base58.Decode(tokenEncoded)
	prefixBytes, err := hex.DecodeString(prefix)
	if err != nil {
		log.Println(err)
		return nil, errors.New(string(GetTokenErrorInvalidToken))
	}

	tokenHash := hash.Argon2IDHash(tokenBytes, prefixBytes)

	return service.repo.GetToken(ctx, tokenHash)
}

func (service *TokensAdapter) QueryTokens(ctx context.Context, userID string) ([]tokens.UserTokenPartial, error) {
	tokens, err := service.repo.QueryTokens(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (service *TokensAdapter) CreateToken(ctx context.Context, userID string, capabilities tokens.APITokenCapabilities) (string, error) {
	tokenClaims := tokens.APITokenClaims{
		Capabilities: capabilities,
	}

	prefixBytes, err := crypto.GenerateRandomBytes(PREFIX_LEN / 2)
	if err != nil {
		return "", err
	}
	prefix := hex.EncodeToString(prefixBytes)
	// TODO: check if prefix is unique for the given user

	// generate a random token
	tokenBytes := genUUIDBytes()
	token := base58.Encode(tokenBytes)

	// hash the token and store it safely
	tokenHash := hash.Argon2IDHash(tokenBytes, prefixBytes)
	if err := service.repo.CreateToken(ctx, userID, tokens.CreateTokenRequest{
		TokenHash: tokenHash,
		Prefix:    prefix,
		Claims:    tokenClaims,
	}); err != nil {
		log.Println(err)
		return "", errors.New(string(CreateTokenErrorStore))
	}

	return prefix + PREFIX_SEPARATOR + token, nil
}

func (service *TokensAdapter) DeleteToken(ctx context.Context, userID string, tokenID uuid.UUID) error {
	owner, err := service.repo.TokenOwner(ctx, userID, tokenID)
	if err != nil {
		log.Println(err)
		return errors.New("could not delete the token")
	}
	if userID != owner {
		return errors.New("token not found")
	}
	return service.repo.DeleteToken(ctx, tokenID)
}

func genUUIDBytes() []byte {
	uid := uuid.New()
	return uid[:]
}
