package auth

import (
	"time"

	"golang.org/x/oauth2"
)

type StaticToken struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	Expiry       time.Time
}

func ReviveToken(token StaticToken) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		Expiry:       token.Expiry,
	}
}

func ToStaticToken(token *oauth2.Token) StaticToken {
	if token == nil {
		return StaticToken{}
	}

	return StaticToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		Expiry:       token.Expiry,
	}
}
