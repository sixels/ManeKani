package auth

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func NewAuthenticator() (*Authenticator, error) {
	keycloakUrl := os.Getenv("KEYCLOAK_URL")
	clientId := os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")

	log.Println("keycloak client id:", clientId)

	realmUrl, err := url.JoinPath(keycloakUrl, "/realms/manekani")
	if err != nil {
		return nil, err
	}

	// TODO: set the server url or port on an env var
	callbackUrl := "http://localhost:8081/auth/callback"

	provider, err := oidc.NewProvider(context.Background(), realmUrl)
	if err != nil {
		return nil, err
	}

	config := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  callbackUrl,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Authenticator{Provider: provider, Config: config}, nil
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func (a *Authenticator) GetUserInfo(ctx context.Context, token *oauth2.Token) (*oidc.UserInfo, error) {
	return a.UserInfo(ctx, oauth2.ReuseTokenSource(token, a.TokenSource(ctx, token)))
}
