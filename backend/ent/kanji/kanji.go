// Code generated by ent, DO NOT EDIT.

package kanji

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the kanji type in the database.
	Label = "kanji"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSymbol holds the string denoting the symbol field in the database.
	FieldSymbol = "symbol"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAltNames holds the string denoting the alt_names field in the database.
	FieldAltNames = "alt_names"
	// FieldSimilar holds the string denoting the similar field in the database.
	FieldSimilar = "similar"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldReading holds the string denoting the reading field in the database.
	FieldReading = "reading"
	// FieldOnyomi holds the string denoting the onyomi field in the database.
	FieldOnyomi = "onyomi"
	// FieldKunyomi holds the string denoting the kunyomi field in the database.
	FieldKunyomi = "kunyomi"
	// FieldNanori holds the string denoting the nanori field in the database.
	FieldNanori = "nanori"
	// FieldMeaningMnemonic holds the string denoting the meaning_mnemonic field in the database.
	FieldMeaningMnemonic = "meaning_mnemonic"
	// FieldReadingMnemonic holds the string denoting the reading_mnemonic field in the database.
	FieldReadingMnemonic = "reading_mnemonic"
	// EdgeVocabularies holds the string denoting the vocabularies edge name in mutations.
	EdgeVocabularies = "vocabularies"
	// EdgeRadicals holds the string denoting the radicals edge name in mutations.
	EdgeRadicals = "radicals"
	// Table holds the table name of the kanji in the database.
	Table = "kanjis"
	// VocabulariesTable is the table that holds the vocabularies relation/edge. The primary key declared below.
	VocabulariesTable = "vocabulary_kanjis"
	// VocabulariesInverseTable is the table name for the Vocabulary entity.
	// It exists in this package in order to avoid circular dependency with the "vocabulary" package.
	VocabulariesInverseTable = "vocabularies"
	// RadicalsTable is the table that holds the radicals relation/edge. The primary key declared below.
	RadicalsTable = "kanji_radicals"
	// RadicalsInverseTable is the table name for the Radical entity.
	// It exists in this package in order to avoid circular dependency with the "radical" package.
	RadicalsInverseTable = "radicals"
)

// Columns holds all SQL columns for kanji fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSymbol,
	FieldName,
	FieldAltNames,
	FieldSimilar,
	FieldLevel,
	FieldReading,
	FieldOnyomi,
	FieldKunyomi,
	FieldNanori,
	FieldMeaningMnemonic,
	FieldReadingMnemonic,
}

var (
	// VocabulariesPrimaryKey and VocabulariesColumn2 are the table columns denoting the
	// primary key for the vocabularies relation (M2M).
	VocabulariesPrimaryKey = []string{"vocabulary_id", "kanji_id"}
	// RadicalsPrimaryKey and RadicalsColumn2 are the table columns denoting the
	// primary key for the radicals relation (M2M).
	RadicalsPrimaryKey = []string{"kanji_id", "radical_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// SymbolValidator is a validator for the "symbol" field. It is called by the builders before save.
	SymbolValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(int32) error
	// ReadingValidator is a validator for the "reading" field. It is called by the builders before save.
	ReadingValidator func(string) error
	// MeaningMnemonicValidator is a validator for the "meaning_mnemonic" field. It is called by the builders before save.
	MeaningMnemonicValidator func(string) error
	// ReadingMnemonicValidator is a validator for the "reading_mnemonic" field. It is called by the builders before save.
	ReadingMnemonicValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)