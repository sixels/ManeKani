package cards

import (
	"context"
	"fmt"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/kanji"
	"sixels.io/manekani/ent/radical"
	"sixels.io/manekani/services/cards/util"
)

var PARTIAL_KANJI_FIELDS = [...]string{
	kanji.FieldID,
	kanji.FieldName,
	kanji.FieldReading,
	kanji.FieldSymbol,
	kanji.FieldLevel,
}

func (repo CardsRepository) CreateKanji(ctx context.Context, req cards.CreateKanjiRequest) (*cards.Kanji, error) {
	return util.WithTx(ctx, repo.client, func(tx *ent.Tx) (*cards.Kanji, error) {
		radicals, err := tx.Radical.Query().Where(radical.NameIn(req.RadicalComposition...)).All(ctx)
		if err != nil {
			return nil, util.ParseEntError(err)
		}

		if err := checkRadicals(radicals, req.RadicalComposition); err != nil {
			return nil, err
		}

		created, err := tx.Kanji.Create().
			SetName(req.Name).
			SetLevel(req.Level).
			SetAltNames(util.ToPgTextArray(req.AltNames)).
			SetSymbol(req.Symbol).
			SetSimilar(util.ToPgTextArray(req.Similar)).
			SetReading(req.Reading).
			SetOnyomi(util.ToPgTextArray(req.Onyomi)).
			SetKunyomi(util.ToPgTextArray(req.Kunyomi)).
			SetNanori(util.ToPgTextArray(req.Nanori)).
			SetMeaningMnemonic(req.MeaningMnemonic).
			SetReadingMnemonic(req.ReadingMnemonic).
			AddRadicals(radicals...).
			Save(ctx)

		if err != nil {
			return nil, util.ParseEntError(err)
		}

		return kanjiFromEnt(created), nil
	})
}

func (repo CardsRepository) QueryKanji(ctx context.Context, symbol string) (*cards.Kanji, error) {
	queried, err := repo.client.Kanji.Query().
		Where(kanji.SymbolEQ(symbol)).
		Only(ctx)

	if err != nil {
		return nil, parseKanjiEntError(err, symbol)
	}

	return kanjiFromEnt(queried), nil
}

func (repo CardsRepository) UpdateKanji(ctx context.Context, symbol string, req cards.UpdateKanjiRequest) (*cards.Kanji, error) {
	return util.WithTx(ctx, repo.client, func(tx *ent.Tx) (*cards.Kanji, error) {

		kanj, err := repo.client.Kanji.Query().
			Where(kanji.SymbolEQ(symbol)).
			Select(kanji.FieldID).
			Only(ctx)

		if err != nil {
			parseKanjiEntError(err, symbol)
		}

		query := kanj.Update()

		util.UpdateValue(req.Level, query.SetLevel)
		util.UpdateValue(req.Name, query.SetName)
		util.UpdateTextArray(req.Similar, query.SetSimilar)
		util.UpdateTextArray(req.AltNames, query.SetAltNames)
		util.UpdateValue(req.MeaningMnemonic, query.SetMeaningMnemonic)
		util.UpdateValue(req.Reading, query.SetReading)
		util.UpdateValue(req.ReadingMnemonic, query.SetReadingMnemonic)
		util.UpdateTextArray(req.Onyomi, query.SetOnyomi)
		util.UpdateTextArray(req.Kunyomi, query.SetKunyomi)
		util.UpdateTextArray(req.Nanori, query.SetNanori)

		if req.RadicalComposition != nil {
			radicals, err := tx.Radical.Query().Where(radical.NameIn(*req.RadicalComposition...)).All(ctx)
			if err != nil {
				return nil, util.ParseEntError(err)
			}

			if err := checkRadicals(radicals, *req.RadicalComposition); err != nil {
				return nil, err
			}

			query.ClearRadicals().AddRadicals(radicals...)
		}

		updated, err := query.Save(ctx)
		if err != nil {
			return nil, util.ParseEntError(err)
		}

		return kanjiFromEnt(updated), nil
	})

}

func (repo CardsRepository) DeleteKanji(ctx context.Context, symbol string) error {
	deletes, err := repo.client.Kanji.Delete().
		Where(kanji.SymbolEQ(symbol)).
		Exec(ctx)

	if err != nil {
		return util.ParseEntError(err)
	}

	if deletes == 0 {
		return errors.NotFound(fmt.Sprintf("no such kanji: '%s'", symbol))
	}

	return nil
}

func (repo CardsRepository) AllKanji(ctx context.Context, req cards.QueryAllKanjiRequest) ([]*cards.PartialKanjiResponse, error) {
	var min int32 = 1
	if req.Level != nil && *req.Level >= 1 {
		min = *req.Level
	}
	var max int32 = min + 10
	if req.MaxLevel != nil && *req.MaxLevel >= 1 {
		max = *req.MaxLevel
	}

	queried, err := repo.client.Kanji.Query().
		Select(PARTIAL_KANJI_FIELDS[:]...).
		Where(kanji.And(kanji.LevelGTE(min), kanji.LevelLTE(max))).
		All(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return util.MapArray(queried, PartialKanjiFromEnt), nil
}

func (repo CardsRepository) QueryKanjiVocabularies(ctx context.Context, symbol string) ([]*cards.PartialVocabularyResponse, error) {
	queried, err := repo.client.Vocabulary.Query().
		WithKanjis(func(kq *ent.KanjiQuery) {
			kq.Select(kanji.FieldID).
				Where(kanji.SymbolEQ(symbol))
		}).
		Select(PARTIAL_VOCABULARY_FIELDS[:]...).
		All(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return util.MapArray(queried, PartialVocabularyFromEnt), nil
}

func (repo CardsRepository) QueryKanjiRadicals(ctx context.Context, symbol string) ([]*cards.PartialRadicalResponse, error) {
	queried, err := repo.client.Radical.Query().
		WithKanjis(func(kq *ent.KanjiQuery) {
			kq.Select(kanji.FieldID).
				Where(kanji.SymbolEQ(symbol))
		}).
		Select(PARTIAL_RADICAL_FIELDS[:]...).
		All(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return util.MapArray(queried, PartialRadicalFromEnt), nil
}

func kanjiFromEnt(e *ent.Kanji) *cards.Kanji {
	return &cards.Kanji{
		Id:              e.ID,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		Name:            e.Name,
		Level:           e.Level,
		AltNames:        util.FromPgTextArray(e.AltNames),
		Symbol:          e.Symbol,
		Similar:         util.FromPgTextArray(e.Similar),
		Reading:         e.Reading,
		Onyomi:          util.FromPgTextArray(e.Onyomi),
		Kunyomi:         util.FromPgTextArray(e.Kunyomi),
		Nanori:          util.FromPgTextArray(e.Nanori),
		MeaningMnemonic: e.MeaningMnemonic,
		ReadingMnemonic: e.ReadingMnemonic,
	}
}

func PartialKanjiFromEnt(e *ent.Kanji) *cards.PartialKanjiResponse {
	return &cards.PartialKanjiResponse{
		Id:       e.ID,
		Name:     e.Name,
		Level:    e.Level,
		AltNames: util.FromPgTextArray(e.AltNames),
		Symbol:   e.Symbol,
		Reading:  e.Reading,
	}
}

func checkRadicals(actual ent.Radicals, req []string) error {
	rads := make([]string, len(actual))
	for i, rad := range actual {
		rads[i] = rad.Name
	}

	if diff := util.DiffStrings(rads, req); diff != nil {
		return errors.InvalidRequest(
			fmt.Sprintf("invalid radical in radical composition: %q", diff))
	}
	return nil
}

func parseKanjiEntError(err error, symbol string) *errors.Error {
	switch err.(type) {
	case *ent.NotFoundError:
		return errors.NotFound(fmt.Sprintf("no such kanji: '%s'", symbol))
	default:
		return util.ParseEntError(err)
	}
}
