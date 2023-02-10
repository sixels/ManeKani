package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ApiToken holds the schema definition for the ApiToken entity.
type ApiToken struct {
	ent.Schema
}

// Fields of the ApiToken.
func (ApiToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bytes("token").
			Immutable().
			Sensitive(),
	}
}

// Edges of the ApiToken.
func (ApiToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("api_tokens").
			Required().
			Unique(),
	}
}
