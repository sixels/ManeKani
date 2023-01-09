//-dsadasd go:build exclude

package schema

import (
	"time"

	"sixels.io/manekani/ent/schema/util"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Radical holds the schema definition for the Radical entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		util.TimeMixin{},
	}
}

// Fields of the Radical.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("username").MinLen(4).MaxLen(30).Immutable().Unique(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.String("email").MinLen(5).MaxLen(320),
		field.Bool("verified").Default(false),
		field.Text("symbol").NotEmpty().Nillable(),
		field.Text("meaning_mnemonic").NotEmpty(),
	}
}

// Edges of the Radical.
func (User) Edges() []ent.Edge {
	// kanji <--kanjis--- radical
	return []ent.Edge{
		edge.From("kanjis", Kanji.Type).
			Ref("radicals"),
	}
}
