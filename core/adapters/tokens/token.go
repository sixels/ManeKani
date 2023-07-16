package tokens

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/sixels/manekani/core/adapters/tokens/hash"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/services/ent/users/crypto"
)

type (
	GetTokenError    error
	CreateTokenError error
)

var (
	ErrGetInvalidToken       GetTokenError    = errors.New("api token is malformed")
	ErrCreateTokenHash       CreateTokenError = errors.New("could not hash the api token")
	ErrCreateTokenStore      CreateTokenError = errors.New("could not store the api token")
	ErrValidateInactiveToken CreateTokenError = errors.New("api token is inactive")
)

const (
	PREFIX_LEN      int    = 8
	TOKEN_SEPARATOR string = "-"
)

func (adp *TokensAdapter) GetToken(ctx context.Context, key string) (tokens.UserToken, error) {
	tokenParts := strings.SplitN(key, TOKEN_SEPARATOR, 2)
	if len(tokenParts) != 2 {
		return tokens.UserToken{}, ErrGetInvalidToken
	}

	prefix, tokenEncoded := tokenParts[0], tokenParts[1]

	tokenBytes := base58.Decode(tokenEncoded)
	prefixBytes, err := hex.DecodeString(prefix)
	if err != nil {
		log.Println(err)
		return tokens.UserToken{}, ErrGetInvalidToken
	}

	tokenHash := hash.Argon2IDHash(tokenBytes, prefixBytes)

	return adp.repo.GetToken(ctx, tokenHash)
}

func (adp *TokensAdapter) QueryTokens(ctx context.Context, userID string) ([]tokens.UserTokenPartial, error) {
	tokens, err := adp.repo.QueryTokens(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (adp *TokensAdapter) CreateToken(ctx context.Context, userID string, req tokens.GenerateTokenRequest) (string, error) {
	tokenClaims := tokens.APITokenClaims{
		Permissions: req.Permissions,
	}

	prefixBytes, err := crypto.GenerateRandomBytes(PREFIX_LEN / 2)
	if err != nil {
		return "", err
	}
	prefix := hex.EncodeToString(prefixBytes)

	// generate a random token
	tokenBytes := genUUIDBytes()
	token := base58.Encode(tokenBytes)

	// hash the token and store it safely
	tokenHash := hash.Argon2IDHash(tokenBytes, prefixBytes)
	if err := adp.repo.CreateToken(ctx, userID, tokens.CreateTokenRequest{
		TokenHash: tokenHash,
		Prefix:    prefix,
		Claims:    tokenClaims,
		Name:      req.Name,
		Status:    tokens.TokenStatusActive,
	}); err != nil {
		log.Println(err)
		return "", ErrCreateTokenStore
	}

	return prefix + TOKEN_SEPARATOR + token, nil
}

func (adp *TokensAdapter) DeleteToken(ctx context.Context, userID string, tokenID ulid.ULID) error {
	owner, err := adp.repo.TokenOwner(ctx, userID, tokenID)
	if err != nil {
		log.Println(err)
		return errors.New("could not delete the token")
	}
	if userID != owner {
		return errors.New("token not found")
	}
	return adp.repo.DeleteToken(ctx, tokenID)
}

func (adp *TokensAdapter) ValidateToken(ctx context.Context, token string) (tokens.UserToken, error) {
	tk, err := adp.GetToken(ctx, token)
	if err != nil {
		return tokens.UserToken{}, err
	}

	if tk.Status != tokens.TokenStatusActive {
		return tokens.UserToken{}, ErrValidateInactiveToken
	}

	return tk, nil
}

func genUUIDBytes() []byte {
	uid := uuid.New()
	return uid[:]
}
