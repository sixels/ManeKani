package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.String("email").
			MinLen(3).MaxLen(255).
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
