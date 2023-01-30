package schema

import (
	"time"

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

		field.UUID("radical_id", uuid.UUID{}).Nillable(),
		field.UUID("kanji_id", uuid.UUID{}).Nillable(),
		field.UUID("vocabulary_id", uuid.UUID{}).Nillable(),

		field.Uint8("progress").Default(0),
		field.Uint("errors").Default(0),

		// timestamps
		field.Time("unlocked_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.Time("started_at").Nillable().
			Comment(`
				The time in which the user started progressing this card.
			`),
		field.Time("passed_at").Nillable().
			Comment(`
				The time when the card passed the apprendice stage.
			`),
		field.Time("available_at").Default(time.Now()).
			Comment(`
				The time when the card will be available to review.
			`),
		field.Time("burned_at").Nillable().
			Comment(`
				The time when the card reached the fluent stage.
			`),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Required(),

		edge.From("radical", Radical.Type).Unique().Field("radical_id"),
		edge.From("kanji", Kanji.Type).Unique().Field("kanji_id"),
		edge.From("vocabulary", Vocabulary.Type).Unique().Field("vocabulary_id"),
	}
}
