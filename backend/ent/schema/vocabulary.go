package schema

import (
	"sixels.io/manekani/ent/schema/util"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Vocabulary holds the schema definition for the Vocabulary entity.
type Vocabulary struct {
	ent.Schema
}

func (Vocabulary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		util.TimeMixin{},
	}
}

// Fields of the Vocabulary.
func (Vocabulary) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("name").NotEmpty(),
		util.TextArray("alt_names", true),
		field.Int32("level").Positive(),

		field.Text("word").NotEmpty().Unique(),
		util.TextArray("word_type", false),

		field.Text("reading").NotEmpty(),
		util.TextArray("alt_readings", true),
		field.JSON("patterns", []Pattern{}),
		field.JSON("sentences", []Sentence{}),

		field.Text("meaning_mnemonic").NotEmpty(),
		field.Text("reading_mnemonic").NotEmpty(),
	}
}

// Edges of the Vocabulary.
func (Vocabulary) Edges() []ent.Edge {
	// vocabulary ---kanjis---> kanji
	return []ent.Edge{
		edge.To("kanjis", Kanji.Type),
	}
}

type Sentence struct {
	Sentence string `json:"sentence"`
	Meaning  string `json:"meaning"`
}

type Pattern struct {
	Name      string     `json:"name"`
	Sentences []Sentence `json:"sentences"`
}