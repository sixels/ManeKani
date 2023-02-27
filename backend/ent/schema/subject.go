package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/sixels/manekani/core/domain/cards"
	"github.com/sixels/manekani/ent/schema/common"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

func (Subject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		common.TimeMixin{},
	}
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("kind"),
		field.Int32("level").Min(1),

		field.Text("name").NotEmpty(),
		field.Text("value").NotEmpty().Optional().Nillable(),
		field.JSON("value_image", &cards.RemoteContent{}).Optional(),
		field.Text("slug").NotEmpty().MaxLen(36),
		field.Uint8("priority").Comment("The priority to appear in lessons/reviews. The maximum priority is 0."),

		field.JSON("resources", &map[string][]cards.RemoteContent{}).Optional(),
		field.JSON("study_data", []cards.StudyData{}),
		field.JSON("complimentary_study_data", &map[string]any{}).Optional(),
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),

		edge.To("similar", Subject.Type),
		edge.To("dependents", Subject.Type).
			From("dependencies"),

		edge.From("deck", Deck.Type).
			Ref("subjects").
			Unique().
			Required(),

		edge.From("owner", User.Type).
			Ref("subjects").
			Unique().
			Required(),
	}
}

func (Subject) Indexes() []ent.Index {
	return []ent.Index{
		// prevent subject to have the same kind and slug in the same deck
		index.Fields("kind", "slug").
			Edges("deck").
			Unique(),
	}
}
