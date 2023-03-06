package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sixels/manekani/ent/schema/common"
)

// Deck holds the schema definition for the Deck entity.
type Deck struct {
	ent.Schema
}

func (Deck) Mixin() []ent.Mixin {
	return []ent.Mixin{
		common.TimeMixin{},
	}
}

// Fields of the Deck.
func (Deck) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("name").MaxLen(50),
		field.Text("description").MaxLen(300),
	}
}

// Edges of the Deck.
func (Deck) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscribers", User.Type),

		edge.From("owner", User.Type).
			Ref("decks").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),

		edge.To("subjects", Subject.Type),

		edge.To("users_progress", DeckProgress.Type),
	}
}
