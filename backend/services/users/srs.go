package users

import (
	"context"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/srs"
	domain "sixels.io/manekani/core/domain/user"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/card"
	"sixels.io/manekani/ent/schema"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"

	cards_svc "sixels.io/manekani/services/cards"
)

type subjectData struct {
	radical    *cards.PartialRadicalResponse
	kanji      *cards.PartialKanjiResponse
	vocabulary *cards.PartialVocabularyResponse
}

func (repo *UsersRepository) GetUserCards(ctx context.Context, userID string) ([]*srs.Card, error) {
	// start := time.Now()
	user, err := repo.client.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	queriedCards, err := user.QueryCards().
		WithSubject().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return util.MapArray(queriedCards, func(card *ent.Card) *srs.Card {
		var subjData subjectData

		subj := card.Edges.Subject

		switch subj.Kind {
		case subject.KindRadical:
			rad := repo.client.Radical.GetX(ctx, subj.ID)
			subjData = subjectData{radical: cards_svc.PartialRadicalFromEnt(rad)}
		case subject.KindKanji:
			kanj := repo.client.Kanji.GetX(ctx, subj.ID)
			subjData = subjectData{kanji: cards_svc.PartialKanjiFromEnt(kanj)}
		case subject.KindVocabulary:
			vocab := repo.client.Vocabulary.GetX(ctx, subj.ID)
			subjData = subjectData{vocabulary: cards_svc.PartialVocabularyFromEnt(vocab)}
		}
		// fmt.Println("-->", time.Since(start).Milliseconds())

		return cardFromEnt(card, subjectFromEnt(subj, subjData))
	}), nil
}

func (repo *UsersRepository) ResetSRSData(ctx context.Context, userID string) error {
	userCards, err := repo.client.User.Query().
		Where(user.IDEQ(userID)).
		QueryCards().
		IDs(ctx)
	if err != nil {
		return err
	}

	// TODO: put it in a single transaction

	_, err = repo.client.Card.Delete().
		Where(card.IDIn(userCards...)).
		Exec(ctx)
	if err != nil {
		return err
	}

	err = repo.client.User.UpdateOneID(userID).
		SetLevel(1).
		AppendPendingActions([]schema.PendingAction{
			{
				Action:   domain.CREATE_LEVEL_CARDS,
				Required: true,
				Metadata: domain.CreateLevelCardsMeta{
					Level: 1,
				},
			},
		}).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func cardFromEnt(e *ent.Card, s cards.Subject) *srs.Card {
	return &srs.Card{
		Id:          e.ID,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		Subject:     s,
		Progress:    e.Progress,
		TotalErrors: e.TotalErrors,
		UnlockedAt:  e.UnlockedAt,
		StartedAt:   e.StartedAt,
		PassedAt:    e.PassedAt,
		AvailableAt: e.AvailableAt,
		BurnedAt:    e.BurnedAt,
	}
}

func subjectFromEnt(e *ent.Subject, data subjectData) cards.Subject {
	return cards.Subject{
		Id:             e.ID,
		Kind:           e.Kind.String(),
		Level:          e.Level,
		RadicalData:    data.radical,
		KanjiData:      data.kanji,
		VocabularyData: data.vocabulary,
	}
}
