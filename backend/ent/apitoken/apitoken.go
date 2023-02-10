// Code generated by ent, DO NOT EDIT.

package apitoken

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the apitoken type in the database.
	Label = "api_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the apitoken in the database.
	Table = "api_tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "api_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_api_tokens"
)

// Columns holds all SQL columns for apitoken fields.
var Columns = []string{
	FieldID,
	FieldToken,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "api_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_api_tokens",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
