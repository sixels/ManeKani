package ports

import (
	"context"

	domain "sixels.io/manekani/core/domain/cards"
)

type (
	KanjiRepository interface {
		CreateKanji(ctx context.Context, req domain.CreateKanjiRequest) (*domain.Kanji, error)
		QueryKanji(ctx context.Context, symbol string) (*domain.Kanji, error)
		UpdateKanji(ctx context.Context, symbol string, req domain.UpdateKanjiRequest) (*domain.Kanji, error)
		DeleteKanji(ctx context.Context, symbol string) error
		AllKanji(ctx context.Context, req domain.QueryAllKanjiRequest) ([]*domain.PartialKanjiResponse, error)

		QueryKanjiRadicals(ctx context.Context, symbol string) ([]*domain.PartialRadicalResponse, error)
		QueryKanjiVocabularies(ctx context.Context, symbol string) ([]*domain.PartialVocabularyResponse, error)
	}

	RadicalRepository interface {
		CreateRadical(ctx context.Context, req domain.CreateRadicalRequest) (*domain.Radical, error)
		QueryRadical(ctx context.Context, name string) (*domain.Radical, error)
		UpdateRadical(ctx context.Context, name string, req domain.UpdateRadicalRequest) (*domain.Radical, error)
		DeleteRadical(ctx context.Context, name string) error
		AllRadicals(ctx context.Context, req domain.QueryAllRadicalRequest) ([]*domain.PartialRadicalResponse, error)

		QueryRadicalKanjis(ctx context.Context, name string) ([]*domain.PartialKanjiResponse, error)
	}

	VocabularyRepository interface {
		CreateVocabulary(ctx context.Context, req domain.CreateVocabularyRequest) (*domain.Vocabulary, error)
		QueryVocabulary(ctx context.Context, word string) (*domain.Vocabulary, error)
		UpdateVocabulary(ctx context.Context, word string, req domain.UpdateVocabularyRequest) (*domain.Vocabulary, error)
		DeleteVocabulary(ctx context.Context, word string) error
		AllVocabularies(ctx context.Context, req domain.QueryAllVocabularyRequest) ([]*domain.PartialVocabularyResponse, error)

		QueryVocabularyKanjis(ctx context.Context, word string) ([]*domain.PartialKanjiResponse, error)
	}

	CardsRepository interface {
		KanjiRepository
		RadicalRepository
		VocabularyRepository
	}
)
