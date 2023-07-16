package ports

import (
	"context"

	"github.com/oklog/ulid/v2"
	"github.com/sixels/manekani/core/domain/tokens"
)

type TokenRepository interface {
	// GetToken gets the claims from the given token and the owner's id.
	GetToken(ctx context.Context, token string) (tokens.UserToken, error)
	// QueryTokens returns the list of ids and prefixes from tokens owned by the given user.
	QueryTokens(ctx context.Context, userID string) ([]tokens.UserTokenPartial, error)
	// CreateToken creates a new token for the given user.
	CreateToken(ctx context.Context, userID string, req tokens.CreateTokenRequest) error
	// DeleteToken deletes the given token
	DeleteToken(ctx context.Context, tokenID ulid.ULID) error
	// TokenOwner returns the owner's id of the given token
	TokenOwner(ctx context.Context, userID string, tokenID ulid.ULID) (string, error)
}
