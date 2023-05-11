package cards

import (
	"github.com/sixels/manekani/core/adapters/cards"
	"github.com/sixels/manekani/services/ent/users"
	"github.com/sixels/manekani/services/files"
)

type CardsApiV1 struct {
	Cards cards.CardsAdapter
	Users *users.UsersRepository
	Files *files.FilesRepository
}
