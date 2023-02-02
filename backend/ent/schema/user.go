package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		}),
		field.String("username").
			MinLen(4).MaxLen(20).
			Unique(),
		field.JSON("pending_actions", []PendingAction{}),
		field.String("email").
			MinLen(3).MaxLen(255).
			Unique(),
		field.Int32("level").
			Min(1).
			Default(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}

type Action = uint64

type PendingAction struct {
	Action   Action `json:"action"`
	Required bool   `json:"required"`
	Metadata any    `json:"metadata"`
}
