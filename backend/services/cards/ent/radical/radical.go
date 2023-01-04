// Code generated by ent, DO NOT EDIT.

package radical

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the radical type in the database.
	Label = "radical"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldSymbol holds the string denoting the symbol field in the database.
	FieldSymbol = "symbol"
	// FieldMeaningMnemonic holds the string denoting the meaning_mnemonic field in the database.
	FieldMeaningMnemonic = "meaning_mnemonic"
	// EdgeKanjis holds the string denoting the kanjis edge name in mutations.
	EdgeKanjis = "kanjis"
	// Table holds the table name of the radical in the database.
	Table = "radicals"
	// KanjisTable is the table that holds the kanjis relation/edge. The primary key declared below.
	KanjisTable = "kanji_radicals"
	// KanjisInverseTable is the table name for the Kanji entity.
	// It exists in this package in order to avoid circular dependency with the "kanji" package.
	KanjisInverseTable = "kanjis"
)

// Columns holds all SQL columns for radical fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldLevel,
	FieldSymbol,
	FieldMeaningMnemonic,
}

var (
	// KanjisPrimaryKey and KanjisColumn2 are the table columns denoting the
	// primary key for the kanjis relation (M2M).
	KanjisPrimaryKey = []string{"kanji_id", "radical_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(int32) error
	// SymbolValidator is a validator for the "symbol" field. It is called by the builders before save.
	SymbolValidator func(string) error
	// MeaningMnemonicValidator is a validator for the "meaning_mnemonic" field. It is called by the builders before save.
	MeaningMnemonicValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
