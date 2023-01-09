package cards

import (
	"context"
	"fmt"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/errors"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/kanji"
	"sixels.io/manekani/ent/vocabulary"
	"sixels.io/manekani/services/cards/util"
)

var PARTIAL_VOCABULARY_FIELDS = [...]string{
	vocabulary.FieldID,
	vocabulary.FieldName,
	vocabulary.FieldLevel,
	vocabulary.FieldAltNames,
	vocabulary.FieldWord,
	vocabulary.FieldReading,
}

func (repo CardsRepository) CreateVocabulary(ctx context.Context, req cards.CreateVocabularyRequest) (*cards.Vocabulary, error) {
	return util.WithTx(ctx, repo.client, func(tx *ent.Tx) (*cards.Vocabulary, error) {
		kanjis, err := tx.Kanji.Query().Where(kanji.SymbolIn(req.KanjiComposition...)).All(ctx)
		if err != nil {
			return nil, util.ParseEntError(err)
		}

		if err := checkKanjis(kanjis, req.KanjiComposition); err != nil {
			return nil, err
		}

		created, err := tx.Vocabulary.Create().
			SetName(req.Name).
			SetLevel(req.Level).
			SetAltNames(util.ToPgTextArray(req.AltNames)).
			SetWord(req.Word).
			SetWordType(util.ToPgTextArray(req.WordType)).
			SetReading(req.Reading).
			SetAltReadings(util.ToPgTextArray(req.AltReadings)).
			SetMeaningMnemonic(req.MeaningMnemonic).
			SetReadingMnemonic(req.ReadingMnemonic).
			SetPatterns(req.Patterns).
			SetSentences(req.Sentences).
			AddKanjis(kanjis...).
			Save(ctx)

		if err != nil {
			fmt.Println(err)
			return nil, util.ParseEntError(err)
		}

		return vocabularyFromEnt(created), nil
	})
}

func (repo CardsRepository) QueryVocabulary(ctx context.Context, word string) (*cards.Vocabulary, error) {
	queried, err := repo.client.Vocabulary.Query().
		Where(vocabulary.WordEQ(word)).
		Only(ctx)

	if err != nil {
		return nil, parseVocabularyEntError(err, word)
	}

	return vocabularyFromEnt(queried), nil
}

func (repo CardsRepository) UpdateVocabulary(ctx context.Context, word string, req cards.UpdateVocabularyRequest) (*cards.Vocabulary, error) {
	return util.WithTx(ctx, repo.client, func(tx *ent.Tx) (*cards.Vocabulary, error) {

		vocab, err := repo.client.Vocabulary.Query().
			Where(vocabulary.WordEQ(word)).
			Select(vocabulary.FieldID).
			Only(ctx)

		if err != nil {
			return nil, parseVocabularyEntError(err, word)
		}

		query := vocab.Update()

		util.UpdateValue(req.Name, query.SetName)
		util.UpdateValue(req.Level, query.SetLevel)
		util.UpdateTextArray(req.AltNames, query.SetAltNames)
		util.UpdateTextArray(req.WordType, query.SetWordType)
		util.UpdateValue(req.Reading, query.SetReading)
		util.UpdateTextArray(req.AltReadings, query.SetAltReadings)
		util.UpdateValue(req.MeaningMnemonic, query.SetMeaningMnemonic)
		util.UpdateValue(req.ReadingMnemonic, query.SetReadingMnemonic)
		util.UpdateValue(req.Patterns, query.SetPatterns)
		util.UpdateValue(req.Sentences, query.SetSentences)

		if req.KanjiComposition != nil {
			kanjis, err := tx.Kanji.Query().
				Where(
					kanji.SymbolIn(*req.KanjiComposition...)).
				All(ctx)
			if err != nil {
				return nil, util.ParseEntError(err)
			}

			if err := checkKanjis(kanjis, *req.KanjiComposition); err != nil {
				return nil, err
			}

			query.ClearKanjis().AddKanjis(kanjis...)
		}

		updated, err := query.Save(ctx)
		if err != nil {
			return nil, util.ParseEntError(err)
		}

		return vocabularyFromEnt(updated), nil
	})
}

func (repo CardsRepository) DeleteVocabulary(ctx context.Context, word string) error {
	deletes, err := repo.client.Vocabulary.Delete().
		Where(vocabulary.WordEQ(word)).
		Exec(ctx)
	if err != nil {
		return util.ParseEntError(err)
	}

	if deletes == 0 {
		return errors.NotFound(fmt.Sprintf("no such vocabulary '%s'", word))
	}

	return nil
}

func (repo CardsRepository) AllVocabularies(ctx context.Context, req cards.QueryAllVocabularyRequest) ([]*cards.PartialVocabularyResponse, error) {
	var min int32 = 0
	if req.Level != nil && *req.Level >= 1 {
		min = *req.Level
	}
	var max int32 = min + 10
	if req.MaxLevel != nil && *req.MaxLevel >= 1 {
		max = *req.MaxLevel
	}

	queried, err := repo.client.Vocabulary.Query().
		Select(PARTIAL_VOCABULARY_FIELDS[:]...).
		Where(vocabulary.And(vocabulary.LevelGTE(min), vocabulary.LevelLTE(max))).
		All(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return util.MapArray(queried, PartialVocabularyFromEnt), nil
}

func (repo CardsRepository) QueryVocabularyKanjis(ctx context.Context, word string) ([]*cards.PartialKanjiResponse, error) {
	queried, err := repo.client.Kanji.Query().
		WithVocabularies(func(vq *ent.VocabularyQuery) {
			vq.Select(vocabulary.FieldID).
				Where(vocabulary.WordEQ(word))
		}).
		Select(PARTIAL_KANJI_FIELDS[:]...).
		All(ctx)

	if err != nil {
		return nil, util.ParseEntError(err)
	}

	return util.MapArray(queried, PartialKanjiFromEnt), nil
}

func vocabularyFromEnt(e *ent.Vocabulary) *cards.Vocabulary {
	return &cards.Vocabulary{
		Id:              e.ID,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		Name:            e.Name,
		Level:           e.Level,
		AltNames:        util.FromPgTextArray(e.AltNames),
		Word:            e.Word,
		WordType:        util.FromPgTextArray(e.WordType),
		Reading:         e.Reading,
		AltReadings:     util.FromPgTextArray(e.AltReadings),
		Patterns:        e.Patterns,
		Sentences:       e.Sentences,
		MeaningMnemonic: e.MeaningMnemonic,
		ReadingMnemonic: e.ReadingMnemonic,
	}
}

func PartialVocabularyFromEnt(e *ent.Vocabulary) *cards.PartialVocabularyResponse {
	return &cards.PartialVocabularyResponse{
		Id:       e.ID,
		Name:     e.Name,
		Level:    e.Level,
		AltNames: util.FromPgTextArray(e.AltNames),
		Word:     e.Word,
		Reading:  e.Reading,
	}
}

func checkKanjis(actual ent.Kanjis, req []string) error {
	kanjis := make([]string, len(actual))
	for i, kanj := range actual {
		kanjis[i] = kanj.Symbol
	}

	if diff := util.DiffStrings(kanjis, req); diff != nil {
		return errors.InvalidRequest(
			fmt.Sprintf("invalid kanji in kanji composition: %q", diff))
	}
	return nil
}

func parseVocabularyEntError(err error, word string) *errors.Error {
	switch err.(type) {
	case *ent.NotFoundError:
		return errors.NotFound(fmt.Sprintf("no such vocabulary: '%s'", word))
	default:
		return util.ParseEntError(err)
	}
}
