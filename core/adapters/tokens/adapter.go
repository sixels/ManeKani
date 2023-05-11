package tokens

import (
	"github.com/sixels/manekani/core/ports"
)

type TokensAdapter struct {
	repo ports.TokenRepository
}

func CreateAdapter(repo ports.TokenRepository) TokensAdapter {
	return TokensAdapter{
		repo: repo,
	}
}
