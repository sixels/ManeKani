// Code generated by ent, DO NOT EDIT.

package subject

import (
	"time"
)

const (
	// Label holds the string label denoting the subject type in the database.
	Label = "subject"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldValueImage holds the string denoting the value_image field in the database.
	FieldValueImage = "value_image"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldResources holds the string denoting the resources field in the database.
	FieldResources = "resources"
	// FieldStudyData holds the string denoting the study_data field in the database.
	FieldStudyData = "study_data"
	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// EdgeSimilar holds the string denoting the similar edge name in mutations.
	EdgeSimilar = "similar"
	// EdgeDependencies holds the string denoting the dependencies edge name in mutations.
	EdgeDependencies = "dependencies"
	// EdgeDependents holds the string denoting the dependents edge name in mutations.
	EdgeDependents = "dependents"
	// EdgeDecks holds the string denoting the decks edge name in mutations.
	EdgeDecks = "decks"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the subject in the database.
	Table = "subjects"
	// CardsTable is the table that holds the cards relation/edge.
	CardsTable = "cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
	// CardsColumn is the table column denoting the cards relation/edge.
	CardsColumn = "subject_cards"
	// SimilarTable is the table that holds the similar relation/edge. The primary key declared below.
	SimilarTable = "subject_similar"
	// DependenciesTable is the table that holds the dependencies relation/edge. The primary key declared below.
	DependenciesTable = "subject_dependents"
	// DependentsTable is the table that holds the dependents relation/edge. The primary key declared below.
	DependentsTable = "subject_dependents"
	// DecksTable is the table that holds the decks relation/edge. The primary key declared below.
	DecksTable = "deck_subjects"
	// DecksInverseTable is the table name for the Deck entity.
	// It exists in this package in order to avoid circular dependency with the "deck" package.
	DecksInverseTable = "decks"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "subjects"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_subjects"
)

// Columns holds all SQL columns for subject fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldKind,
	FieldLevel,
	FieldName,
	FieldValue,
	FieldValueImage,
	FieldSlug,
	FieldPriority,
	FieldResources,
	FieldStudyData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subjects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_subjects",
}

var (
	// SimilarPrimaryKey and SimilarColumn2 are the table columns denoting the
	// primary key for the similar relation (M2M).
	SimilarPrimaryKey = []string{"subject_id", "similar_id"}
	// DependenciesPrimaryKey and DependenciesColumn2 are the table columns denoting the
	// primary key for the dependencies relation (M2M).
	DependenciesPrimaryKey = []string{"subject_id", "dependency_id"}
	// DependentsPrimaryKey and DependentsColumn2 are the table columns denoting the
	// primary key for the dependents relation (M2M).
	DependentsPrimaryKey = []string{"subject_id", "dependency_id"}
	// DecksPrimaryKey and DecksColumn2 are the table columns denoting the
	// primary key for the decks relation (M2M).
	DecksPrimaryKey = []string{"deck_id", "subject_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(int32) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
)
