package cards

import (
	"context"

	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/core/domain/cards/filters"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/deck"
	"sixels.io/manekani/ent/predicate"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
	"sixels.io/manekani/services/ent/util"
)

func (repo *CardsRepository) QuerySubject(ctx context.Context, id uuid.UUID) (*cards.Subject, error) {
	result, err := repo.client.Subject.
		Query().
		Where(subject.IDEQ(id)).
		WithDependencies(func(sq *ent.SubjectQuery) {
			sq.Select(subject.FieldID)
		}).
		WithDependents(func(sq *ent.SubjectQuery) {
			sq.Select(subject.FieldID)
		}).
		WithSimilar(func(sq *ent.SubjectQuery) {
			sq.Select(subject.FieldID)
		}).
		WithDecks(func(dq *ent.DeckQuery) {
			dq.Select(deck.FieldID)
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return SubjectFromEnt(result), nil
}

func (repo *CardsRepository) CreateSubject(ctx context.Context, ownerID string, req cards.CreateSubjectRequest) (*cards.Subject, error) {
	query := repo.client.Subject.Create().
		SetID(uuid.New()).
		SetKind(req.Kind).
		SetLevel(req.Level).
		SetName(req.Name).
		SetNillableValue(req.Value).
		SetValueImage(req.ValueImage).
		SetSlug(req.Slug).
		SetPriority(req.Priority).
		SetStudyData(req.StudyData).
		SetResources(req.Resources).
		AddDependencyIDs(req.Dependencies...).
		AddDependentIDs(req.Dependents...).
		AddSimilarIDs(req.Similars...).
		SetOwnerID(ownerID)

	created, err := query.Save(ctx)
	if err != nil {
		return nil, util.ParseEntError(err)
	}
	return SubjectFromEnt(created), nil
}

func (repo *CardsRepository) UpdateSubject(ctx context.Context, id uuid.UUID, req cards.UpdateSubjectRequest) (*cards.Subject, error) {
	query := repo.client.Subject.UpdateOneID(id)

	util.UpdateValue(req.Kind, query.SetKind)
	util.UpdateValue(req.Level, query.SetLevel)
	util.UpdateValue(req.Name, query.SetName)

	query.SetNillableValue(req.Value)
	if req.ValueImage != nil {
		query.SetValueImage(req.ValueImage)
	}

	util.UpdateValue(req.Slug, query.SetSlug)
	util.UpdateValue(req.Priority, query.SetPriority)
	util.UpdateValue(req.StudyData, query.SetStudyData)
	if req.Resources != nil {
		query.SetResources(req.Resources)
	}

	util.UpdateValues(req.Dependencies, query.AddDependencyIDs)
	util.UpdateValues(req.Dependents, query.AddDependentIDs)
	util.UpdateValues(req.Similars, query.AddSimilarIDs)

	updated, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}
	return SubjectFromEnt(updated), nil
}

func (repo *CardsRepository) DeleteSubject(ctx context.Context, id uuid.UUID) error {
	return repo.client.Subject.DeleteOneID(id).Exec(ctx)
}

func (repo *CardsRepository) AllSubjects(ctx context.Context, req cards.QueryManySubjectsRequest) ([]*cards.PartialSubject, error) {
	var reqFilters []predicate.Subject
	reqFilters = filters.ApplyFilter(reqFilters, req.Levels.Separate(), subject.LevelIn)
	reqFilters = filters.ApplyFilter(reqFilters, req.IDs.Separate(), subject.IDIn)
	reqFilters = filters.ApplyFilter(reqFilters, req.Kinds.Separate(), subject.KindIn)
	reqFilters = filters.ApplyFilter(reqFilters, req.Slugs.Separate(), subject.SlugIn)
	reqFilters = filters.ApplyFilter(reqFilters, req.Decks.Separate(), func(ids ...uuid.UUID) predicate.Subject {
		return subject.HasDecksWith(deck.IDIn(ids...))
	})
	reqFilters = filters.ApplyFilter(reqFilters, req.Owners.Separate(), func(ids ...string) predicate.Subject {
		return subject.HasOwnerWith(user.IDIn(ids...))
	})

	page := 0
	if req.Page != nil {
		page = int(*req.Page)
	}

	queried, err := repo.client.Subject.Query().
		Where(subject.And(reqFilters...)).
		Limit(1000).
		Offset(page).
		Select(
			subject.FieldID,
			subject.FieldKind,
			subject.FieldLevel,
			subject.FieldPriority,
			subject.FieldName,
			subject.FieldValue,
			subject.FieldValueImage,
			subject.FieldSlug,
			subject.FieldStudyData,
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return util.MapArray(queried, PartialSubjectFromEnt), nil
}

func (repo *CardsRepository) SubjectOwner(ctx context.Context, subjectID uuid.UUID) (string, error) {
	return repo.client.Subject.Query().
		Where(subject.IDEQ(subjectID)).
		QueryOwner().
		OnlyID(ctx)
}

func SubjectFromEnt(e *ent.Subject) *cards.Subject {
	return &cards.Subject{
		ID:         e.ID,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		Kind:       e.Kind,
		Level:      e.Level,
		Name:       e.Name,
		Value:      e.Value,
		ValueImage: e.ValueImage,
		Slug:       e.Slug,
		Priority:   e.Priority,
		Resources:  e.Resources,
		StudyData:  e.StudyData,
		Dependencies: util.MapArray(e.Edges.Dependencies,
			func(s *ent.Subject) uuid.UUID { return s.ID },
		),
		Dependents: util.MapArray(e.Edges.Dependents,
			func(s *ent.Subject) uuid.UUID { return s.ID },
		),
		Similars: util.MapArray(e.Edges.Similar,
			func(s *ent.Subject) uuid.UUID { return s.ID },
		),
		Decks: util.MapArray(e.Edges.Decks,
			func(d *ent.Deck) uuid.UUID { return d.ID },
		),
	}
}

func PartialSubjectFromEnt(e *ent.Subject) *cards.PartialSubject {
	return &cards.PartialSubject{
		ID:         e.ID,
		Kind:       e.Kind,
		Level:      e.Level,
		Priority:   e.Priority,
		Name:       e.Name,
		Value:      e.Value,
		ValueImage: e.ValueImage,
		Slug:       e.Slug,
		StudyData:  e.StudyData,
	}
}
