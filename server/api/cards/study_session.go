package cards

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/cards/filters"
	"github.com/sixels/manekani/server/api/cards/util"
)

func (api *CardsApi) StudySession() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userID, err := util.CtxUserID(c)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusUnauthorized)
			return
		}
		var req cards.CreateStudySessionRequest
		if err := c.BindQuery(&req); err != nil {
			c.Error(err)
			c.Status(http.StatusBadRequest)
			return
		}

		availableCards, err := api.V1.Cards.AllCards(
			ctx,
			userID,
			cards.QueryManyCardsRequest{
				FilterDecks: filters.FilterDecks{
					Decks: util.Ptr((filters.CommaSeparatedUUID)(req.DeckID)),
				},
				AvailableBefore: util.Ptr(time.Now()),
				IsStarted:       util.Ptr(req.SessionType == cards.SessionReview),
			},
		)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		sessionSize := 8
		if sessionSize > len(availableCards) {
			sessionSize = len(availableCards)
		}
		sessionQueue := createQueue(availableCards[:sessionSize])

		c.JSON(http.StatusOK, sessionQueue)
	}
}

func createQueue(cards_ []cards.Card) cards.SessionQueue {
	sessionCards := make([]cards.QueueItem, 0, len(cards_))

	for _, card := range cards_ {
		answers := make([]cards.QueueItemAnswers, len(card.Subject.StudyData))
		for i, sdata := range card.Subject.StudyData {
			answers[i] = cards.QueueItemAnswers{
				StudyItemType: sdata.Kind,
				Expected: FilterMap(sdata.Items, func(item *cards.StudyItem) *string {
					if item.IsValidAnswer {
						return &item.Value
					}
					return nil
				}),
				Blacklisted: FilterMap(sdata.Items, func(item *cards.StudyItem) *string {
					if !item.IsValidAnswer {
						return &item.Value
					}
					return nil
				}),
			}
		}
		sessionCards = append(sessionCards, cards.QueueItem{
			CardID:  card.ID,
			Answers: answers,
			Subject: minimalSubject(card.Subject),
		})
	}

	return cards.SessionQueue{Cards: sessionCards}
}

func minimalSubject(p cards.PartialSubject) cards.MinimalSubject {
	return cards.MinimalSubject{
		ID:         p.ID,
		Kind:       p.Kind,
		Level:      p.Level,
		Name:       p.Name,
		Value:      p.Value,
		ValueImage: p.ValueImage,
		Slug:       p.Slug,
		Priority:   p.Priority,
		Deck:       p.Deck,
	}
}

func FilterMap[T any, U any](values []T, mapper func(*T) *U) []U {
	slice := make([]U, len(values))

	for i, val := range values {
		ret := mapper(&val)
		if ret != nil {
			slice[i] = *ret
		}
	}

	return slice
}