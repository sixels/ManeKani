package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const ISSUER = "ManeKani"

type JWTService struct {
	secret string
}

func CreateService(secret string) *JWTService {
	return &JWTService{
		secret: secret,
	}
}

func (service *JWTService) ValidateCapabilities(claims APITokenClaims, caps ...APITokenCapability) bool {
	capsMap := MapCapabilities(claims)
	for _, cap := range caps {
		if !capsMap[cap] {
			return false
		}
	}

	return true
}

func (service *JWTService) ValidateToken(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token: %v", t.Header["alg"])
		}
		return []byte(service.secret), nil
	})
}

func (service *JWTService) CreateAPIToken(options APITokenOptions) (string, error) {
	log.Printf("creating token with the following capabilities: %q\n", options.Capabilities)
	tokenClaims := APITokenClaims{
		UserID: options.UserID,
		Scope:  options.Scope,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: getDefault(options.ExpiresAt, time.Now().AddDate(0, 0, 7)).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    ISSUER,
		},
		APITokenCapabilities: APITokenCapabilities{},
	}
	capsMap := MapCapabilitiesRef(&tokenClaims)
	for _, cap := range options.Capabilities {
		if claim, ok := capsMap[cap]; ok {
			*claim = true
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenClaims)
	return token.SignedString([]byte(service.secret))
}

func getDefault[T any](value *T, def T) T {
	if value != nil {
		return *value
	}
	return def
}
