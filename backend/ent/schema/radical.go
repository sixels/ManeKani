package schema

import (
	"sixels.io/manekani/ent/schema/common"

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
		common.TimeMixin{},
	}
}

// Fields of the Radical.
func (Radical) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("name").Unique().NotEmpty(),
		field.Int32("level").Positive(),
		field.Text("symbol").NotEmpty().Nillable(),
		field.Text("meaning_mnemonic").NotEmpty(),
	}
}

// Edges of the Radical.
func (Radical) Edges() []ent.Edge {
	return []ent.Edge{
		// kanji <--kanjis--- radical
		edge.From("kanjis", Kanji.Type).
			Ref("radicals"),
	}
}
