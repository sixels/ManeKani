package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	cards2 "github.com/sixels/manekani/core/adapters/cards"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/core/domain/files"
	"github.com/sixels/manekani/core/ports/transactions"
	"github.com/sixels/manekani/server/api/apicommon"
	cards3 "github.com/sixels/manekani/server/api/cards"
	"github.com/sixels/manekani/server/auth"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	subjectOkValue = "一"
	UserID         = "8d4286c9-3b48-431c-8f99-3e82682ee60d"
	SubjectOk      = cards.Subject{
		Deck:  uuid.MustParse("8d4286c9-3b48-431c-8f99-3e82682ee60c"),
		Name:  "Ground",
		Kind:  "radical",
		Level: 1,
		Slug:  "ground",
		Value: &subjectOkValue,
		Owner: UserID,
	}
)

func TestCreateSubject(t *testing.T) {

	t.Run("Create a simple subject", func(t *testing.T) {
		e := echo.New()

		cardsProvider := cards2.CreateAdapter(MockSubject{}, MockFile{})
		api := cards3.New(cardsProvider)

		body := &bytes.Buffer{}
		form := multipart.NewWriter(body)

		_ = form.WriteField("deck", SubjectOk.Deck.String())
		_ = form.WriteField("name", SubjectOk.Name)
		_ = form.WriteField("kind", SubjectOk.Kind)
		_ = form.WriteField("level", fmt.Sprintf("%v", SubjectOk.Level))
		_ = form.WriteField("slug", SubjectOk.Slug)
		_ = form.WriteField("value", *SubjectOk.Value)

		_ = form.Close()

		req := httptest.NewRequest(http.MethodGet, "/", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+form.Boundary())
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/api/tokens")
		c.Set(string(auth.UserIDContext), UserID)

		if assert.NoError(t, api.V1.CreateSubject(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var created apicommon.APIResponse[cards.Subject]
			assert.Nil(t, json.NewDecoder(rec.Body).Decode(&created))
			assert.Equal(t, SubjectOk, created.Data)
		}

	})
}

type MockSubject struct{}

func (s MockSubject) CreateSubject(ctx context.Context, ownerID string, req cards.CreateSubjectRequest) (*cards.Subject, error) {
	return &cards.Subject{
		ID:         uuid.UUID{},
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		Kind:       req.Kind,
		Level:      req.Level,
		Name:       req.Name,
		Value:      req.Value,
		ValueImage: req.ValueImage,
		Slug:       req.Slug,
		Priority:   req.Priority,
		//Resources:           req.Resources,
		StudyData:           req.StudyData,
		AdditionalStudyData: req.AdditionalStudyData,
		Dependencies:        req.Dependencies,
		Dependents:          req.Dependents,
		Similars:            req.Similars,
		Deck:                req.Deck,
		Owner:               ownerID,
	}, nil
}

func (s MockSubject) QuerySubject(ctx context.Context, id uuid.UUID) (*cards.Subject, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) UpdateSubject(ctx context.Context, id uuid.UUID, req cards.UpdateSubjectRequest) (*cards.Subject, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) DeleteSubject(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) AllSubjects(ctx context.Context, req cards.QueryManySubjectsRequest) ([]cards.PartialSubject, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) SubjectOwner(ctx context.Context, id uuid.UUID) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) CreateDeck(ctx context.Context, ownerID string, req cards.CreateDeckRequest) (*cards.Deck, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) QueryDeck(ctx context.Context, id uuid.UUID) (*cards.Deck, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) AllDecks(ctx context.Context, req cards.QueryManyDecksRequest) ([]cards.DeckPartial, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) DeckOwner(ctx context.Context, id uuid.UUID) (string, error) {
	return UserID, nil
}

func (s MockSubject) AddDeckSubscriber(ctx context.Context, id uuid.UUID, userID string) (deckProgressID int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) RemoveDeckSubscriber(ctx context.Context, id uuid.UUID, userID string) error {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) DeckSubscriberExists(ctx context.Context, id uuid.UUID, userID string) (deckProgressID int, exists bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) CreateReview(ctx context.Context, userID string, req cards.CreateReviewRequest) (*cards.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) AllReviews(ctx context.Context, userID string, req cards.QueryManyReviewsRequest) ([]cards.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) QueryCard(ctx context.Context, id uuid.UUID) (*cards.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) UpdateCard(ctx context.Context, id uuid.UUID, req cards.UpdateCardRequest) (*cards.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) AllCards(ctx context.Context, userID string, req cards.QueryManyCardsRequest) ([]cards.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) CreateManyCards(ctx context.Context, deckProgressID int, userID string, reqs []cards.CreateCardRequest) ([]cards.Card, error) {
	//TODO implement me
	panic("implement me")
}

func (s MockSubject) BeginTransaction(ctx context.Context) (transactions.TransactionalRepository, error) {
	return s, nil
}

func (s MockSubject) Rollback() error {
	return nil
}

func (s MockSubject) Commit() error {
	return nil
}

type MockFile struct{}

func (m MockFile) CreateFile(ctx context.Context, req files.CreateFileRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockFile) QueryFile(ctx context.Context, name string) (*files.ObjectWrapperResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockFile) DeleteFile(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
