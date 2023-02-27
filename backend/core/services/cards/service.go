package cards

import "sixels.io/manekani/core/ports"

type CardsService struct {
	repo      ports.CardsRepository
	filesRepo ports.FilesRepository
}

func NewService(cardsRepo ports.CardsRepository, filesRepo ports.FilesRepository) CardsService {
	return CardsService{
		repo:      cardsRepo,
		filesRepo: filesRepo,
	}
}
