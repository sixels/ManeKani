// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sixels/manekani/ent/card"
	"github.com/sixels/manekani/ent/review"
)

// Review is the model entity for the Review schema.
type Review struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Errors holds the value of the "errors" field.
	Errors map[string]int32 `json:"errors,omitempty"`
	// StartProgress holds the value of the "start_progress" field.
	StartProgress uint8 `json:"start_progress,omitempty"`
	// EndProgress holds the value of the "end_progress" field.
	EndProgress uint8 `json:"end_progress,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewQuery when eager-loading is set.
	Edges        ReviewEdges `json:"edges"`
	card_reviews *uuid.UUID
}

// ReviewEdges holds the relations/edges for other nodes in the graph.
type ReviewEdges struct {
	// Card holds the value of the card edge.
	Card *Card `json:"card,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewEdges) CardOrErr() (*Card, error) {
	if e.loadedTypes[0] {
		if e.Card == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: card.Label}
		}
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Review) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case review.FieldErrors:
			values[i] = new([]byte)
		case review.FieldStartProgress, review.FieldEndProgress:
			values[i] = new(sql.NullInt64)
		case review.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case review.FieldID:
			values[i] = new(uuid.UUID)
		case review.ForeignKeys[0]: // card_reviews
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Review", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Review fields.
func (r *Review) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case review.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case review.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case review.FieldErrors:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field errors", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Errors); err != nil {
					return fmt.Errorf("unmarshal field errors: %w", err)
				}
			}
		case review.FieldStartProgress:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_progress", values[i])
			} else if value.Valid {
				r.StartProgress = uint8(value.Int64)
			}
		case review.FieldEndProgress:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_progress", values[i])
			} else if value.Valid {
				r.EndProgress = uint8(value.Int64)
			}
		case review.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field card_reviews", values[i])
			} else if value.Valid {
				r.card_reviews = new(uuid.UUID)
				*r.card_reviews = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCard queries the "card" edge of the Review entity.
func (r *Review) QueryCard() *CardQuery {
	return (&ReviewClient{config: r.config}).QueryCard(r)
}

// Update returns a builder for updating this Review.
// Note that you need to call Review.Unwrap() before calling this method if this Review
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Review) Update() *ReviewUpdateOne {
	return (&ReviewClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Review entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Review) Unwrap() *Review {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Review is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Review) String() string {
	var builder strings.Builder
	builder.WriteString("Review(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("errors=")
	builder.WriteString(fmt.Sprintf("%v", r.Errors))
	builder.WriteString(", ")
	builder.WriteString("start_progress=")
	builder.WriteString(fmt.Sprintf("%v", r.StartProgress))
	builder.WriteString(", ")
	builder.WriteString("end_progress=")
	builder.WriteString(fmt.Sprintf("%v", r.EndProgress))
	builder.WriteByte(')')
	return builder.String()
}

// Reviews is a parsable slice of Review.
type Reviews []*Review

func (r Reviews) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}