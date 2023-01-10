// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// KanjisColumns holds the columns for the "kanjis" table.
	KanjisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "symbol", Type: field.TypeString, Unique: true, Size: 5},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "alt_names", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "similar", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "level", Type: field.TypeInt32},
		{Name: "reading", Type: field.TypeString, Size: 2147483647},
		{Name: "onyomi", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "kunyomi", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "nanori", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "meaning_mnemonic", Type: field.TypeString, Size: 2147483647},
		{Name: "reading_mnemonic", Type: field.TypeString, Size: 2147483647},
	}
	// KanjisTable holds the schema information for the "kanjis" table.
	KanjisTable = &schema.Table{
		Name:       "kanjis",
		Columns:    KanjisColumns,
		PrimaryKey: []*schema.Column{KanjisColumns[0]},
	}
	// RadicalsColumns holds the columns for the "radicals" table.
	RadicalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "level", Type: field.TypeInt32},
		{Name: "symbol", Type: field.TypeString, Size: 2147483647},
		{Name: "meaning_mnemonic", Type: field.TypeString, Size: 2147483647},
	}
	// RadicalsTable holds the schema information for the "radicals" table.
	RadicalsTable = &schema.Table{
		Name:       "radicals",
		Columns:    RadicalsColumns,
		PrimaryKey: []*schema.Column{RadicalsColumns[0]},
	}
	// VocabulariesColumns holds the columns for the "vocabularies" table.
	VocabulariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "alt_names", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "level", Type: field.TypeInt32},
		{Name: "word", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "word_type", Type: field.TypeOther, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "reading", Type: field.TypeString, Size: 2147483647},
		{Name: "alt_readings", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"postgres": "TEXT[]"}},
		{Name: "patterns", Type: field.TypeJSON},
		{Name: "sentences", Type: field.TypeJSON},
		{Name: "meaning_mnemonic", Type: field.TypeString, Size: 2147483647},
		{Name: "reading_mnemonic", Type: field.TypeString, Size: 2147483647},
	}
	// VocabulariesTable holds the schema information for the "vocabularies" table.
	VocabulariesTable = &schema.Table{
		Name:       "vocabularies",
		Columns:    VocabulariesColumns,
		PrimaryKey: []*schema.Column{VocabulariesColumns[0]},
	}
	// KanjiRadicalsColumns holds the columns for the "kanji_radicals" table.
	KanjiRadicalsColumns = []*schema.Column{
		{Name: "kanji_id", Type: field.TypeUUID},
		{Name: "radical_id", Type: field.TypeUUID},
	}
	// KanjiRadicalsTable holds the schema information for the "kanji_radicals" table.
	KanjiRadicalsTable = &schema.Table{
		Name:       "kanji_radicals",
		Columns:    KanjiRadicalsColumns,
		PrimaryKey: []*schema.Column{KanjiRadicalsColumns[0], KanjiRadicalsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "kanji_radicals_kanji_id",
				Columns:    []*schema.Column{KanjiRadicalsColumns[0]},
				RefColumns: []*schema.Column{KanjisColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "kanji_radicals_radical_id",
				Columns:    []*schema.Column{KanjiRadicalsColumns[1]},
				RefColumns: []*schema.Column{RadicalsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// VocabularyKanjisColumns holds the columns for the "vocabulary_kanjis" table.
	VocabularyKanjisColumns = []*schema.Column{
		{Name: "vocabulary_id", Type: field.TypeUUID},
		{Name: "kanji_id", Type: field.TypeUUID},
	}
	// VocabularyKanjisTable holds the schema information for the "vocabulary_kanjis" table.
	VocabularyKanjisTable = &schema.Table{
		Name:       "vocabulary_kanjis",
		Columns:    VocabularyKanjisColumns,
		PrimaryKey: []*schema.Column{VocabularyKanjisColumns[0], VocabularyKanjisColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vocabulary_kanjis_vocabulary_id",
				Columns:    []*schema.Column{VocabularyKanjisColumns[0]},
				RefColumns: []*schema.Column{VocabulariesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "vocabulary_kanjis_kanji_id",
				Columns:    []*schema.Column{VocabularyKanjisColumns[1]},
				RefColumns: []*schema.Column{KanjisColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		KanjisTable,
		RadicalsTable,
		VocabulariesTable,
		KanjiRadicalsTable,
		VocabularyKanjisTable,
	}
)

func init() {
	KanjiRadicalsTable.ForeignKeys[0].RefTable = KanjisTable
	KanjiRadicalsTable.ForeignKeys[1].RefTable = RadicalsTable
	VocabularyKanjisTable.ForeignKeys[0].RefTable = VocabulariesTable
	VocabularyKanjisTable.ForeignKeys[1].RefTable = KanjisTable
}