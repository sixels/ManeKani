package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Annotations(entsql.Annotation{
			Size: 36,
		}).DefaultFunc(uuid.NewString),
		field.String("username").
			MinLen(4).MaxLen(20).
			Unique(),
		field.JSON("pending_actions", []PendingAction{}).
			Optional(),
		field.String("email").
			MinLen(3).MaxLen(255).
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("decks", Deck.Type),
		edge.To("subjects", Subject.Type),
		edge.From("subscribed_decks", Deck.Type).
			Ref("subscribers"),

		edge.To("api_tokens", ApiToken.Type),

		edge.To("decks_progress", DeckProgress.Type),
	}
}

type Action = uint64

type PendingAction struct {
	Action   Action `json:"action"`
	Required bool   `json:"required"`
	Metadata any    `json:"metadata"`
}
