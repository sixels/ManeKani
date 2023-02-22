package cards

import (
	"sixels.io/manekani/services/ent/cards"
	"sixels.io/manekani/services/ent/users"
	"sixels.io/manekani/services/files"
)

type CardsApiV1 struct {
	Cards *cards.CardsRepository
	Users *users.UsersRepository
	Files *files.FilesRepository
}
