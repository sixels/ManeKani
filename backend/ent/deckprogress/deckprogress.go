// Code generated by ent, DO NOT EDIT.

package deckprogress

const (
	// Label holds the string label denoting the deckprogress type in the database.
	Label = "deck_progress"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeDeck holds the string denoting the deck edge name in mutations.
	EdgeDeck = "deck"
	// Table holds the table name of the deckprogress in the database.
	Table = "deck_progresses"
	// CardsTable is the table that holds the cards relation/edge.
	CardsTable = "cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
	// CardsColumn is the table column denoting the cards relation/edge.
	CardsColumn = "deck_progress_cards"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "deck_progresses"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_decks_progress"
	// DeckTable is the table that holds the deck relation/edge.
	DeckTable = "deck_progresses"
	// DeckInverseTable is the table name for the Deck entity.
	// It exists in this package in order to avoid circular dependency with the "deck" package.
	DeckInverseTable = "decks"
	// DeckColumn is the table column denoting the deck relation/edge.
	DeckColumn = "deck_users_progress"
)

// Columns holds all SQL columns for deckprogress fields.
var Columns = []string{
	FieldID,
	FieldLevel,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "deck_progresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"deck_users_progress",
	"user_decks_progress",
}

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
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(uint32) error
)