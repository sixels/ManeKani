package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/tokens"
)

// ApiToken holds the schema definition for the ApiToken entity.
type ApiToken struct {
	ent.Schema
}

// Fields of the ApiToken.
func (ApiToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("token").
			Sensitive().
			Immutable(),
		field.String("prefix").
			Immutable(),
		field.JSON("claims", tokens.APITokenClaims{}).
			Sensitive().
			Immutable(),
	}
}

// Edges of the ApiToken.
func (ApiToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("api_tokens").
			Required().
			Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (ApiToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("prefix").
			Edges("user").
			Unique(),
	}
}
