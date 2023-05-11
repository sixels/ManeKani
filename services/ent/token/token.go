package token

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/apitoken"
	"github.com/sixels/manekani/ent/user"
	"github.com/sixels/manekani/services/ent/util"
)

var _ ports.TokenRepository = (*TokenRepository)(nil)

// GetToken gets the claims from the given token and the owner's id.
func (repo *TokenRepository) GetToken(ctx context.Context, tokenHash string) (*tokens.UserToken, error) {
	token, err := repo.client.ApiTokenClient().Query().
		WithUser().
		Where(apitoken.TokenEQ(tokenHash)).
		Only(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return util.Ptr(UserTokenFromEnt(token)), nil
}

// QueryTokens returns the list of token prefixes owned by the given user.
func (repo *TokenRepository) QueryTokens(ctx context.Context, userID string) ([]tokens.UserTokenPartial, error) {
	tokens, err := repo.client.ApiTokenClient().Query().
		Where(
			apitoken.HasUserWith(user.IDEQ(userID)),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return util.MapArray(tokens, UserTokenPartialFromEnt), nil
}

// CreateToken creates a new token for the given user.
func (repo *TokenRepository) CreateToken(ctx context.Context, userID string, req tokens.CreateTokenRequest) error {
	err := repo.client.ApiTokenClient().Create().
		SetToken(req.TokenHash).
		SetPrefix(req.Prefix).
		SetClaims(req.Claims).
		SetUserID(userID).
		Exec(ctx)
	return err
}

// DeleteToken deletes the given user token by prefix.
func (repo *TokenRepository) DeleteToken(ctx context.Context, tokenID uuid.UUID) error {
	err := repo.client.ApiTokenClient().
		DeleteOneID(tokenID).
		Exec(ctx)
	return err
}

func (repo *TokenRepository) TokenOwner(ctx context.Context, userID string, tokenID uuid.UUID) (string, error) {
	token, err := repo.client.ApiTokenClient().Query().
		WithUser().
		Where(apitoken.IDEQ(tokenID)).
		Only(ctx)
	if err != nil {
		return "", err
	}
	return token.Edges.User.ID, nil

}

func UserTokenFromEnt(e *ent.ApiToken) tokens.UserToken {
	return tokens.UserToken{
		ID:     e.ID,
		Claims: e.Claims,
		UserID: e.Edges.User.ID,
		Prefix: e.Prefix,
	}
}

func UserTokenPartialFromEnt(e *ent.ApiToken) tokens.UserTokenPartial {
	return tokens.UserTokenPartial{
		ID:     e.ID,
		Prefix: e.Prefix,
		Claims: e.Claims,
	}
}
