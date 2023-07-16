// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sixels/manekani/ent/card"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/subject"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Progress holds the value of the "progress" field.
	Progress uint8 `json:"progress,omitempty"`
	// TotalErrors holds the value of the "total_errors" field.
	TotalErrors int32 `json:"total_errors,omitempty"`
	// UnlockedAt holds the value of the "unlocked_at" field.
	UnlockedAt *time.Time `json:"unlocked_at,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// PassedAt holds the value of the "passed_at" field.
	PassedAt *time.Time `json:"passed_at,omitempty"`
	// AvailableAt holds the value of the "available_at" field.
	AvailableAt *time.Time `json:"available_at,omitempty"`
	// BurnedAt holds the value of the "burned_at" field.
	BurnedAt *time.Time `json:"burned_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges               CardEdges `json:"edges"`
	deck_progress_cards *int
	subject_cards       *uuid.UUID
	selectValues        sql.SelectValues
}

// CardEdges holds the relations/edges for other nodes in the graph.
type CardEdges struct {
	// DeckProgress holds the value of the deck_progress edge.
	DeckProgress *DeckProgress `json:"deck_progress,omitempty"`
	// Subject holds the value of the subject edge.
	Subject *Subject `json:"subject,omitempty"`
	// Reviews holds the value of the reviews edge.
	Reviews []*Review `json:"reviews,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DeckProgressOrErr returns the DeckProgress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) DeckProgressOrErr() (*DeckProgress, error) {
	if e.loadedTypes[0] {
		if e.DeckProgress == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: deckprogress.Label}
		}
		return e.DeckProgress, nil
	}
	return nil, &NotLoadedError{edge: "deck_progress"}
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CardEdges) SubjectOrErr() (*Subject, error) {
	if e.loadedTypes[1] {
		if e.Subject == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: subject.Label}
		}
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e CardEdges) ReviewsOrErr() ([]*Review, error) {
	if e.loadedTypes[2] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case card.FieldProgress, card.FieldTotalErrors:
			values[i] = new(sql.NullInt64)
		case card.FieldCreatedAt, card.FieldUpdatedAt, card.FieldUnlockedAt, card.FieldStartedAt, card.FieldPassedAt, card.FieldAvailableAt, card.FieldBurnedAt:
			values[i] = new(sql.NullTime)
		case card.FieldID:
			values[i] = new(uuid.UUID)
		case card.ForeignKeys[0]: // deck_progress_cards
			values[i] = new(sql.NullInt64)
		case card.ForeignKeys[1]: // subject_cards
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case card.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case card.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case card.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case card.FieldProgress:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field progress", values[i])
			} else if value.Valid {
				c.Progress = uint8(value.Int64)
			}
		case card.FieldTotalErrors:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_errors", values[i])
			} else if value.Valid {
				c.TotalErrors = int32(value.Int64)
			}
		case card.FieldUnlockedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field unlocked_at", values[i])
			} else if value.Valid {
				c.UnlockedAt = new(time.Time)
				*c.UnlockedAt = value.Time
			}
		case card.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				c.StartedAt = new(time.Time)
				*c.StartedAt = value.Time
			}
		case card.FieldPassedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field passed_at", values[i])
			} else if value.Valid {
				c.PassedAt = new(time.Time)
				*c.PassedAt = value.Time
			}
		case card.FieldAvailableAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field available_at", values[i])
			} else if value.Valid {
				c.AvailableAt = new(time.Time)
				*c.AvailableAt = value.Time
			}
		case card.FieldBurnedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field burned_at", values[i])
			} else if value.Valid {
				c.BurnedAt = new(time.Time)
				*c.BurnedAt = value.Time
			}
		case card.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field deck_progress_cards", value)
			} else if value.Valid {
				c.deck_progress_cards = new(int)
				*c.deck_progress_cards = int(value.Int64)
			}
		case card.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field subject_cards", values[i])
			} else if value.Valid {
				c.subject_cards = new(uuid.UUID)
				*c.subject_cards = *value.S.(*uuid.UUID)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Card.
// This includes values selected through modifiers, order, etc.
func (c *Card) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryDeckProgress queries the "deck_progress" edge of the Card entity.
func (c *Card) QueryDeckProgress() *DeckProgressQuery {
	return NewCardClient(c.config).QueryDeckProgress(c)
}

// QuerySubject queries the "subject" edge of the Card entity.
func (c *Card) QuerySubject() *SubjectQuery {
	return NewCardClient(c.config).QuerySubject(c)
}

// QueryReviews queries the "reviews" edge of the Card entity.
func (c *Card) QueryReviews() *ReviewQuery {
	return NewCardClient(c.config).QueryReviews(c)
}

// Update returns a builder for updating this Card.
// Note that you need to call Card.Unwrap() before calling this method if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return NewCardClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Card entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("progress=")
	builder.WriteString(fmt.Sprintf("%v", c.Progress))
	builder.WriteString(", ")
	builder.WriteString("total_errors=")
	builder.WriteString(fmt.Sprintf("%v", c.TotalErrors))
	builder.WriteString(", ")
	if v := c.UnlockedAt; v != nil {
		builder.WriteString("unlocked_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.StartedAt; v != nil {
		builder.WriteString("started_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.PassedAt; v != nil {
		builder.WriteString("passed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.AvailableAt; v != nil {
		builder.WriteString("available_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.BurnedAt; v != nil {
		builder.WriteString("burned_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Cards is a parsable slice of Card.
type Cards []*Card
