package cards

import (
	"sixels.io/manekani/core/services/cards"
	"sixels.io/manekani/services/ent/users"
	"sixels.io/manekani/services/files"
)

type CardsApiV1 struct {
	Cards cards.CardsService
	Users *users.UsersRepository
	Files *files.FilesRepository
}
