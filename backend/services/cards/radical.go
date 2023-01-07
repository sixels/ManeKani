package cards

import (
	"context"
	"fmt"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/services/cards/ent"
	"sixels.io/manekani/services/cards/ent/radical"
	"sixels.io/manekani/services/cards/util"
)

var PARTIAL_RADICAL_FIELDS = [...]string{
	radical.FieldID,
	radical.FieldName,
	radical.FieldLevel,
	radical.FieldSymbol,
}

func (repo CardsRepository) CreateRadical(ctx context.Context, req cards.CreateRadicalRequest) (*cards.Radical, error) {
	query := repo.client.Radical.Create().
		SetName(req.Name).
		SetLevel(req.Level).
		SetMeaningMnemonic(req.MeaningMnemonic)

	if req.Symbol != nil {
		query.SetSymbol(*req.Symbol)
	}

	created, err := query.Save(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return radicalFromEnt(created), nil
}

func (repo CardsRepository) QueryRadical(ctx context.Context, name string) (*cards.Radical, error) {
	queried, err := repo.client.Radical.Query().
		Where(radical.NameEQ(name)).
		Only(ctx)

	if err != nil {
		return nil, parseRadicalEntError(err, name)
	}

	return radicalFromEnt(queried), nil
}

func (repo CardsRepository) UpdateRadical(ctx context.Context, name string, req cards.UpdateRadicalRequest) (*cards.Radical, error) {
	rad, err := repo.client.Radical.Query().
		Where(radical.NameEQ(name)).
		Select(radical.FieldID).
		Only(ctx)

	if err != nil {
		return nil, parseRadicalEntError(err, name)
	}

	query := rad.Update()

	util.UpdateValue(req.Level, query.SetLevel)
	util.UpdateValue(req.MeaningMnemonic, query.SetMeaningMnemonic)
	util.UpdateValue(req.Symbol, query.SetSymbol)

	updated, err := query.Save(ctx)
	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return radicalFromEnt(updated), nil
}

func (repo CardsRepository) DeleteRadical(ctx context.Context, name string) error {
	deletes, err := repo.client.Radical.Delete().
		Where(radical.NameEQ(name)).
		Exec(ctx)

	if err != nil {
		return util.ParseEntError(err)
	}

	if deletes == 0 {
		return errors.NotFound(fmt.Sprintf("no such radical: '%s'", name))
	}

	return nil
}

func (repo CardsRepository) AllRadicals(ctx context.Context) ([]*cards.PartialRadicalResponse, error) {
	queried, err := repo.client.Radical.Query().
		Select(PARTIAL_RADICAL_FIELDS[:]...).
		All(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return util.MapArray(queried, PartialRadicalFromEnt), nil
}

func (repo CardsRepository) QueryRadicalKanjis(ctx context.Context, name string) ([]*cards.PartialKanjiResponse, error) {
	queried, err := repo.client.Kanji.Query().
		WithRadicals(func(rq *ent.RadicalQuery) {
			rq.Select(radical.FieldID).
				Where(radical.NameEQ(name))
		}).
		Select(PARTIAL_KANJI_FIELDS[:]...).
		All(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return util.MapArray(queried, PartialKanjiFromEnt), nil
}

func radicalFromEnt(e *ent.Radical) *cards.Radical {
	return &cards.Radical{
		Id:              e.ID,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		Name:            e.Name,
		Level:           e.Level,
		Symbol:          e.Symbol,
		MeaningMnemonic: e.MeaningMnemonic,
	}
}

func PartialRadicalFromEnt(e *ent.Radical) *cards.PartialRadicalResponse {
	return &cards.PartialRadicalResponse{
		Id:     e.ID,
		Name:   e.Name,
		Level:  e.Level,
		Symbol: e.Symbol,
	}
}

func parseRadicalEntError(err error, name string) *errors.Error {
	switch err.(type) {
	case *ent.NotFoundError:
		return errors.NotFound(fmt.Sprintf("no such radical: '%s'", name))
	default:
		return util.ParseEntError(err)
	}
}
