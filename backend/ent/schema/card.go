package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"sixels.io/manekani/ent/schema/common"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		common.TimeMixin{},
	}
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),

		field.Uint8("progress").Default(0),
		field.Int32("total_errors").Default(0),

		// timestamps
		field.Time("unlocked_at").Optional().Nillable(),
		field.Time("started_at").Optional().Nillable(),
		field.Time("passed_at").Optional().Nillable(),
		field.Time("available_at").Optional().Nillable(),
		field.Time("burned_at").Optional().Nillable(),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("cards").
			Required().
			Unique(),

		edge.From("subject", Subject.Type).
			Ref("cards").
			Required().
			Unique(),

		edge.To("reviews", Review.Type),
	}
}
