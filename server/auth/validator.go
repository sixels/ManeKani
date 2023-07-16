package auth

import (
	"context"

	ory "github.com/ory/client-go"
	adapter "github.com/sixels/manekani/core/adapters/tokens"
	"github.com/sixels/manekani/core/domain/tokens"
)

type OAPIValidator struct {
	tokens *adapter.TokensAdapter
	ory    *ory.APIClient
}

func (v *OAPIValidator) ValidateToken(ctx context.Context, token string) (tokens.UserToken, error) {
	return v.tokens.ValidateToken(ctx, token)
}

func (auth *OAPIValidator) ValidateLoginSession(ctx context.Context, sessionCookies string) (*ory.Session, error) {
	session, _, err := auth.ory.FrontendApi.ToSession(ctx).
		Cookie(sessionCookies).
		Execute()
	return session, err
}
