package token

import (
	"context"
	"log"

	"github.com/oklog/ulid/v2"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/core/ports"
	"github.com/sixels/manekani/ent"
	"github.com/sixels/manekani/ent/apitoken"
	"github.com/sixels/manekani/ent/user"
	"github.com/sixels/manekani/services/ent/util"
)

var _ ports.TokenRepository = (*TokenRepository)(nil)

// GetToken gets the claims from the given token and the owner's id.
func (repo *TokenRepository) GetToken(ctx context.Context, tokenHash string) (tokens.UserToken, error) {
	token, err := repo.client.ApiTokenClient().Query().
		WithUser().
		Where(apitoken.TokenEQ(tokenHash)).
		Only(ctx)
	if err != nil {
		log.Println(err)
		return tokens.UserToken{}, err
	}
	return UserTokenFromEnt(token), nil
}

// QueryTokens returns the list of token prefixes owned by the given user.
func (repo *TokenRepository) QueryTokens(ctx context.Context, userID string) ([]tokens.UserTokenPartial, error) {
	apiTokens, err := repo.client.ApiTokenClient().Query().
		Where(apitoken.HasUserWith(user.IDEQ(userID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return util.MapArray(apiTokens, UserTokenPartialFromEnt), nil
}

// CreateToken creates a new token for the given user.
func (repo *TokenRepository) CreateToken(ctx context.Context, userID string, req tokens.CreateTokenRequest) (tokens.UserToken, error) {
	log.Printf("----\n%s\n%##v\n----", userID, req)
	token, err := repo.client.ApiTokenClient().Create().
		SetName(req.Name).
		SetToken(req.TokenHash).
		SetPrefix(req.Prefix).
		SetClaims(req.Claims).
		SetStatus(req.Status).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return tokens.UserToken{}, err
	}

	token.Edges.User = &ent.User{ID: userID}

	return UserTokenFromEnt(token), nil
}

// DeleteToken deletes the given user token by prefix.
func (repo *TokenRepository) DeleteToken(ctx context.Context, tokenID ulid.ULID) error {
	err := repo.client.ApiTokenClient().
		DeleteOneID(tokenID).
		Exec(ctx)
	return err
}

func (repo *TokenRepository) TokenOwner(ctx context.Context, userID string, tokenID ulid.ULID) (string, error) {
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
		UserID: e.Edges.User.ID,
		Name:   e.Name,
		Claims: e.Claims,
		Prefix: e.Prefix,
		Status: e.Status,
		UsedAt: e.UsedAt,
	}
}

func UserTokenPartialFromEnt(e *ent.ApiToken) tokens.UserTokenPartial {
	return tokens.UserTokenPartial{
		ID:     e.ID,
		Name:   e.Name,
		Prefix: e.Prefix,
		Claims: e.Claims,
		Status: e.Status,
		UsedAt: e.UsedAt,
	}
}
