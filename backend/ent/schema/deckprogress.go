package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeckProgress holds the schema definition for the DeckProgress entity.
type DeckProgress struct {
	ent.Schema
}

// Fields of the DeckProgress.
func (DeckProgress) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("level").Min(1),
	}
}

// Edges of the DeckProgress.
func (DeckProgress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),

		edge.From("user", User.Type).
			Ref("decks_progress").
			Unique().
			Required(),
		edge.From("deck", Deck.Type).
			Ref("users_progress").
			Unique().
			Required(),
	}
}
