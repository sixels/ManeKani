package cards

import (
	"context"

	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/ent"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/services/ent/util"
)

func (repo *CardsRepository) AllSubjects(ctx context.Context, req cards.QueryAllSubjectsRequest) ([]*cards.PartialSubjectResponse, error) {
	var min int32 = 1
	if req.Level != nil && *req.Level >= 1 {
		min = *req.Level
	}
	var max int32 = 60
	if req.MaxLevel != nil && *req.MaxLevel >= 1 {
		max = *req.MaxLevel
	}

	queried, err := repo.client.Subject.Query().
		Where(subject.And(subject.LevelGTE(min), subject.LevelLTE(max))).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return util.MapArray(queried, PartialSubjectFromEnt), nil
}

func PartialSubjectFromEnt(e *ent.Subject) *cards.PartialSubjectResponse {
	return &cards.PartialSubjectResponse{
		Id:    e.ID,
		Kind:  e.Kind.String(),
		Level: e.Level,
	}
}
