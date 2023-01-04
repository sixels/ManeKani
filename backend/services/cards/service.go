package cards

import (
	ports "sixels.io/manekani/core/ports/cards"
)

type CardsService struct {
	ports.CardsRepository
}

func NewService(repo ports.CardsRepository) *CardsService {
	return &CardsService{
		CardsRepository: repo,
	}
}
