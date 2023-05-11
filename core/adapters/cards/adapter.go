package cards

import "github.com/sixels/manekani/core/ports"

type CardsAdapter struct {
	repo      ports.CardsRepository
	filesRepo ports.FilesRepository
}

func CreateAdapter(cardsRepo ports.CardsRepository, filesRepo ports.FilesRepository) CardsAdapter {
	return CardsAdapter{
		repo:      cardsRepo,
		filesRepo: filesRepo,
	}
}
