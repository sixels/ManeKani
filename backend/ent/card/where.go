// Code generated by ent, DO NOT EDIT.

package card

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"sixels.io/manekani/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Progress applies equality check predicate on the "progress" field. It's identical to ProgressEQ.
func Progress(v uint8) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProgress), v))
	})
}

// TotalErrors applies equality check predicate on the "total_errors" field. It's identical to TotalErrorsEQ.
func TotalErrors(v int32) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalErrors), v))
	})
}

// UnlockedAt applies equality check predicate on the "unlocked_at" field. It's identical to UnlockedAtEQ.
func UnlockedAt(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnlockedAt), v))
	})
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// PassedAt applies equality check predicate on the "passed_at" field. It's identical to PassedAtEQ.
func PassedAt(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassedAt), v))
	})
}

// AvailableAt applies equality check predicate on the "available_at" field. It's identical to AvailableAtEQ.
func AvailableAt(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvailableAt), v))
	})
}

// BurnedAt applies equality check predicate on the "burned_at" field. It's identical to BurnedAtEQ.
func BurnedAt(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBurnedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ProgressEQ applies the EQ predicate on the "progress" field.
func ProgressEQ(v uint8) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProgress), v))
	})
}

// ProgressNEQ applies the NEQ predicate on the "progress" field.
func ProgressNEQ(v uint8) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProgress), v))
	})
}

// ProgressIn applies the In predicate on the "progress" field.
func ProgressIn(vs ...uint8) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProgress), v...))
	})
}

// ProgressNotIn applies the NotIn predicate on the "progress" field.
func ProgressNotIn(vs ...uint8) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProgress), v...))
	})
}

// ProgressGT applies the GT predicate on the "progress" field.
func ProgressGT(v uint8) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProgress), v))
	})
}

// ProgressGTE applies the GTE predicate on the "progress" field.
func ProgressGTE(v uint8) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProgress), v))
	})
}

// ProgressLT applies the LT predicate on the "progress" field.
func ProgressLT(v uint8) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProgress), v))
	})
}

// ProgressLTE applies the LTE predicate on the "progress" field.
func ProgressLTE(v uint8) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProgress), v))
	})
}

// TotalErrorsEQ applies the EQ predicate on the "total_errors" field.
func TotalErrorsEQ(v int32) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalErrors), v))
	})
}

// TotalErrorsNEQ applies the NEQ predicate on the "total_errors" field.
func TotalErrorsNEQ(v int32) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalErrors), v))
	})
}

// TotalErrorsIn applies the In predicate on the "total_errors" field.
func TotalErrorsIn(vs ...int32) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalErrors), v...))
	})
}

// TotalErrorsNotIn applies the NotIn predicate on the "total_errors" field.
func TotalErrorsNotIn(vs ...int32) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalErrors), v...))
	})
}

// TotalErrorsGT applies the GT predicate on the "total_errors" field.
func TotalErrorsGT(v int32) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalErrors), v))
	})
}

// TotalErrorsGTE applies the GTE predicate on the "total_errors" field.
func TotalErrorsGTE(v int32) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalErrors), v))
	})
}

// TotalErrorsLT applies the LT predicate on the "total_errors" field.
func TotalErrorsLT(v int32) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalErrors), v))
	})
}

// TotalErrorsLTE applies the LTE predicate on the "total_errors" field.
func TotalErrorsLTE(v int32) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalErrors), v))
	})
}

// UnlockedAtEQ applies the EQ predicate on the "unlocked_at" field.
func UnlockedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnlockedAt), v))
	})
}

// UnlockedAtNEQ applies the NEQ predicate on the "unlocked_at" field.
func UnlockedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnlockedAt), v))
	})
}

// UnlockedAtIn applies the In predicate on the "unlocked_at" field.
func UnlockedAtIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnlockedAt), v...))
	})
}

// UnlockedAtNotIn applies the NotIn predicate on the "unlocked_at" field.
func UnlockedAtNotIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnlockedAt), v...))
	})
}

// UnlockedAtGT applies the GT predicate on the "unlocked_at" field.
func UnlockedAtGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnlockedAt), v))
	})
}

// UnlockedAtGTE applies the GTE predicate on the "unlocked_at" field.
func UnlockedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnlockedAt), v))
	})
}

// UnlockedAtLT applies the LT predicate on the "unlocked_at" field.
func UnlockedAtLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnlockedAt), v))
	})
}

// UnlockedAtLTE applies the LTE predicate on the "unlocked_at" field.
func UnlockedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnlockedAt), v))
	})
}

// UnlockedAtIsNil applies the IsNil predicate on the "unlocked_at" field.
func UnlockedAtIsNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnlockedAt)))
	})
}

// UnlockedAtNotNil applies the NotNil predicate on the "unlocked_at" field.
func UnlockedAtNotNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnlockedAt)))
	})
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartedAt)))
	})
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartedAt)))
	})
}

// PassedAtEQ applies the EQ predicate on the "passed_at" field.
func PassedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassedAt), v))
	})
}

// PassedAtNEQ applies the NEQ predicate on the "passed_at" field.
func PassedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassedAt), v))
	})
}

// PassedAtIn applies the In predicate on the "passed_at" field.
func PassedAtIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPassedAt), v...))
	})
}

// PassedAtNotIn applies the NotIn predicate on the "passed_at" field.
func PassedAtNotIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPassedAt), v...))
	})
}

// PassedAtGT applies the GT predicate on the "passed_at" field.
func PassedAtGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassedAt), v))
	})
}

// PassedAtGTE applies the GTE predicate on the "passed_at" field.
func PassedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassedAt), v))
	})
}

// PassedAtLT applies the LT predicate on the "passed_at" field.
func PassedAtLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassedAt), v))
	})
}

// PassedAtLTE applies the LTE predicate on the "passed_at" field.
func PassedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassedAt), v))
	})
}

// PassedAtIsNil applies the IsNil predicate on the "passed_at" field.
func PassedAtIsNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPassedAt)))
	})
}

// PassedAtNotNil applies the NotNil predicate on the "passed_at" field.
func PassedAtNotNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPassedAt)))
	})
}

// AvailableAtEQ applies the EQ predicate on the "available_at" field.
func AvailableAtEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtNEQ applies the NEQ predicate on the "available_at" field.
func AvailableAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtIn applies the In predicate on the "available_at" field.
func AvailableAtIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAvailableAt), v...))
	})
}

// AvailableAtNotIn applies the NotIn predicate on the "available_at" field.
func AvailableAtNotIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAvailableAt), v...))
	})
}

// AvailableAtGT applies the GT predicate on the "available_at" field.
func AvailableAtGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtGTE applies the GTE predicate on the "available_at" field.
func AvailableAtGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtLT applies the LT predicate on the "available_at" field.
func AvailableAtLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtLTE applies the LTE predicate on the "available_at" field.
func AvailableAtLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvailableAt), v))
	})
}

// AvailableAtIsNil applies the IsNil predicate on the "available_at" field.
func AvailableAtIsNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAvailableAt)))
	})
}

// AvailableAtNotNil applies the NotNil predicate on the "available_at" field.
func AvailableAtNotNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAvailableAt)))
	})
}

// BurnedAtEQ applies the EQ predicate on the "burned_at" field.
func BurnedAtEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBurnedAt), v))
	})
}

// BurnedAtNEQ applies the NEQ predicate on the "burned_at" field.
func BurnedAtNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBurnedAt), v))
	})
}

// BurnedAtIn applies the In predicate on the "burned_at" field.
func BurnedAtIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBurnedAt), v...))
	})
}

// BurnedAtNotIn applies the NotIn predicate on the "burned_at" field.
func BurnedAtNotIn(vs ...time.Time) predicate.Card {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBurnedAt), v...))
	})
}

// BurnedAtGT applies the GT predicate on the "burned_at" field.
func BurnedAtGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBurnedAt), v))
	})
}

// BurnedAtGTE applies the GTE predicate on the "burned_at" field.
func BurnedAtGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBurnedAt), v))
	})
}

// BurnedAtLT applies the LT predicate on the "burned_at" field.
func BurnedAtLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBurnedAt), v))
	})
}

// BurnedAtLTE applies the LTE predicate on the "burned_at" field.
func BurnedAtLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBurnedAt), v))
	})
}

// BurnedAtIsNil applies the IsNil predicate on the "burned_at" field.
func BurnedAtIsNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBurnedAt)))
	})
}

// BurnedAtNotNil applies the NotNil predicate on the "burned_at" field.
func BurnedAtNotNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBurnedAt)))
	})
}

// HasDeckProgress applies the HasEdge predicate on the "deck_progress" edge.
func HasDeckProgress() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeckProgressTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeckProgressTable, DeckProgressColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeckProgressWith applies the HasEdge predicate on the "deck_progress" edge with a given conditions (other predicates).
func HasDeckProgressWith(preds ...predicate.DeckProgress) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DeckProgressInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeckProgressTable, DeckProgressColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Subject) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReviews applies the HasEdge predicate on the "reviews" edge.
func HasReviews() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewsWith applies the HasEdge predicate on the "reviews" edge with a given conditions (other predicates).
func HasReviewsWith(preds ...predicate.Review) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		p(s.Not())
	})
}
