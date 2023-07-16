// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/user"
)

// Deck is the model entity for the Deck schema.
type Deck struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeckQuery when eager-loading is set.
	Edges        DeckEdges `json:"edges"`
	user_decks   *string
	selectValues sql.SelectValues
}

// DeckEdges holds the relations/edges for other nodes in the graph.
type DeckEdges struct {
	// Subscribers holds the value of the subscribers edge.
	Subscribers []*User `json:"subscribers,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Subjects holds the value of the subjects edge.
	Subjects []*Subject `json:"subjects,omitempty"`
	// UsersProgress holds the value of the users_progress edge.
	UsersProgress []*DeckProgress `json:"users_progress,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// SubscribersOrErr returns the Subscribers value or an error if the edge
// was not loaded in eager-loading.
func (e DeckEdges) SubscribersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Subscribers, nil
	}
	return nil, &NotLoadedError{edge: "subscribers"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeckEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// SubjectsOrErr returns the Subjects value or an error if the edge
// was not loaded in eager-loading.
func (e DeckEdges) SubjectsOrErr() ([]*Subject, error) {
	if e.loadedTypes[2] {
		return e.Subjects, nil
	}
	return nil, &NotLoadedError{edge: "subjects"}
}

// UsersProgressOrErr returns the UsersProgress value or an error if the edge
// was not loaded in eager-loading.
func (e DeckEdges) UsersProgressOrErr() ([]*DeckProgress, error) {
	if e.loadedTypes[3] {
		return e.UsersProgress, nil
	}
	return nil, &NotLoadedError{edge: "users_progress"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deck) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deck.FieldName, deck.FieldDescription:
			values[i] = new(sql.NullString)
		case deck.FieldCreatedAt, deck.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case deck.FieldID:
			values[i] = new(uuid.UUID)
		case deck.ForeignKeys[0]: // user_decks
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deck fields.
func (d *Deck) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deck.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case deck.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case deck.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case deck.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case deck.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		case deck.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_decks", values[i])
			} else if value.Valid {
				d.user_decks = new(string)
				*d.user_decks = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Deck.
// This includes values selected through modifiers, order, etc.
func (d *Deck) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QuerySubscribers queries the "subscribers" edge of the Deck entity.
func (d *Deck) QuerySubscribers() *UserQuery {
	return NewDeckClient(d.config).QuerySubscribers(d)
}

// QueryOwner queries the "owner" edge of the Deck entity.
func (d *Deck) QueryOwner() *UserQuery {
	return NewDeckClient(d.config).QueryOwner(d)
}

// QuerySubjects queries the "subjects" edge of the Deck entity.
func (d *Deck) QuerySubjects() *SubjectQuery {
	return NewDeckClient(d.config).QuerySubjects(d)
}

// QueryUsersProgress queries the "users_progress" edge of the Deck entity.
func (d *Deck) QueryUsersProgress() *DeckProgressQuery {
	return NewDeckClient(d.config).QueryUsersProgress(d)
}

// Update returns a builder for updating this Deck.
// Note that you need to call Deck.Unwrap() before calling this method if this Deck
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deck) Update() *DeckUpdateOne {
	return NewDeckClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Deck entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deck) Unwrap() *Deck {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deck is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deck) String() string {
	var builder strings.Builder
	builder.WriteString("Deck(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Decks is a parsable slice of Deck.
type Decks []*Deck
