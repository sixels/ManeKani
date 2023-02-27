package cards

import (
	"github.com/sixels/manekani/core/services/cards"
	"github.com/sixels/manekani/services/ent/users"
	"github.com/sixels/manekani/services/files"
)

type CardsApiV1 struct {
	Cards cards.CardsService
	Users *users.UsersRepository
	Files *files.FilesRepository
}
