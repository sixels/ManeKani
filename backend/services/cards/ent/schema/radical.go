package schema

import (
	"sixels.io/manekani/services/cards/ent/schema/util"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Radical holds the schema definition for the Radical entity.
type Radical struct {
	ent.Schema
}

func (Radical) Mixin() []ent.Mixin {
	return []ent.Mixin{
		util.TimeMixin{},
	}
}

// Fields of the Radical.
func (Radical) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("name").Unique().NotEmpty(),
		field.Int32("level").NonNegative(),
		field.Text("symbol").NotEmpty().Nillable(),
		field.Text("meaning_mnemonic").NotEmpty(),
	}
}

// Edges of the Radical.
func (Radical) Edges() []ent.Edge {
	// kanji <--kanjis--- radical
	return []ent.Edge{
		edge.From("kanjis", Kanji.Type).
			Ref("radicals"),
	}
}
