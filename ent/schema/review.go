package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),

		field.Time("created_at").Default(time.Now).Immutable(),

		field.JSON("errors", cards.ReviewErrors{}).Immutable(),

		field.Uint8("start_progress").Immutable(),
		field.Uint8("end_progress").Immutable(),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("card", Card.Type).
			Ref("reviews").
			Unique().
			Required(),
	}
}
