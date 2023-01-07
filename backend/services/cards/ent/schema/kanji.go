package schema

import (
	"sixels.io/manekani/services/cards/ent/schema/util"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4"
)

// Kanji holds the schema definition for the Kanji entity.
type Kanji struct {
	ent.Schema
}

func (Kanji) Mixin() []ent.Mixin {
	return []ent.Mixin{
		util.TimeMixin{},
	}
}

// Fields of the Kanji.
func (Kanji) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("symbol").
			MaxLen(5).
			NotEmpty().
			Immutable().
			Unique(),
		field.Text("name").NotEmpty(),
		util.TextArray("alt_names", true),
		util.TextArray("similar", true),
		field.Int32("level").Positive(),

		field.Text("reading").NotEmpty(),
		util.TextArray("onyomi", false),
		util.TextArray("kunyomi", false),
		util.TextArray("nanori", false),

		field.Text("meaning_mnemonic").NotEmpty(),
		field.Text("reading_mnemonic").NotEmpty(),
	}
}

// Edges of the Kanji.
func (Kanji) Edges() []ent.Edge {
	// vocabulary <--vocabularies--- kanji
	// kanji ---radicals--> radicals
	return []ent.Edge{
		edge.From("vocabularies", Vocabulary.Type).
			Ref("kanjis"),
		edge.To("radicals", Radical.Type),
	}
}
