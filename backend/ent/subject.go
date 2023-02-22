// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/ent/deck"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// Level holds the value of the "level" field.
	Level int32 `json:"level,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Value holds the value of the "value" field.
	Value *string `json:"value,omitempty"`
	// ValueImage holds the value of the "value_image" field.
	ValueImage *cards.RemoteContent `json:"value_image,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// The priority to appear in lessons/reviews. The maximum priority is 0.
	Priority uint8 `json:"priority,omitempty"`
	// Resources holds the value of the "resources" field.
	Resources *map[string][]cards.RemoteContent `json:"resources,omitempty"`
	// StudyData holds the value of the "study_data" field.
	StudyData []cards.StudyData `json:"study_data,omitempty"`
	// ComplimentaryStudyData holds the value of the "complimentary_study_data" field.
	ComplimentaryStudyData *map[string]interface{} `json:"complimentary_study_data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges         SubjectEdges `json:"edges"`
	deck_subjects *uuid.UUID
	user_subjects *string
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// Similar holds the value of the similar edge.
	Similar []*Subject `json:"similar,omitempty"`
	// Dependencies holds the value of the dependencies edge.
	Dependencies []*Subject `json:"dependencies,omitempty"`
	// Dependents holds the value of the dependents edge.
	Dependents []*Subject `json:"dependents,omitempty"`
	// Deck holds the value of the deck edge.
	Deck *Deck `json:"deck,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// SimilarOrErr returns the Similar value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) SimilarOrErr() ([]*Subject, error) {
	if e.loadedTypes[1] {
		return e.Similar, nil
	}
	return nil, &NotLoadedError{edge: "similar"}
}

// DependenciesOrErr returns the Dependencies value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) DependenciesOrErr() ([]*Subject, error) {
	if e.loadedTypes[2] {
		return e.Dependencies, nil
	}
	return nil, &NotLoadedError{edge: "dependencies"}
}

// DependentsOrErr returns the Dependents value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) DependentsOrErr() ([]*Subject, error) {
	if e.loadedTypes[3] {
		return e.Dependents, nil
	}
	return nil, &NotLoadedError{edge: "dependents"}
}

// DeckOrErr returns the Deck value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectEdges) DeckOrErr() (*Deck, error) {
	if e.loadedTypes[4] {
		if e.Deck == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: deck.Label}
		}
		return e.Deck, nil
	}
	return nil, &NotLoadedError{edge: "deck"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[5] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subject.FieldValueImage, subject.FieldResources, subject.FieldStudyData, subject.FieldComplimentaryStudyData:
			values[i] = new([]byte)
		case subject.FieldLevel, subject.FieldPriority:
			values[i] = new(sql.NullInt64)
		case subject.FieldKind, subject.FieldName, subject.FieldValue, subject.FieldSlug:
			values[i] = new(sql.NullString)
		case subject.FieldCreatedAt, subject.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case subject.FieldID:
			values[i] = new(uuid.UUID)
		case subject.ForeignKeys[0]: // deck_subjects
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case subject.ForeignKeys[1]: // user_subjects
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subject", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case subject.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case subject.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case subject.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				s.Kind = value.String
			}
		case subject.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				s.Level = int32(value.Int64)
			}
		case subject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subject.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				s.Value = new(string)
				*s.Value = value.String
			}
		case subject.FieldValueImage:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field value_image", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.ValueImage); err != nil {
					return fmt.Errorf("unmarshal field value_image: %w", err)
				}
			}
		case subject.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				s.Slug = value.String
			}
		case subject.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				s.Priority = uint8(value.Int64)
			}
		case subject.FieldResources:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field resources", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Resources); err != nil {
					return fmt.Errorf("unmarshal field resources: %w", err)
				}
			}
		case subject.FieldStudyData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field study_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.StudyData); err != nil {
					return fmt.Errorf("unmarshal field study_data: %w", err)
				}
			}
		case subject.FieldComplimentaryStudyData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field complimentary_study_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.ComplimentaryStudyData); err != nil {
					return fmt.Errorf("unmarshal field complimentary_study_data: %w", err)
				}
			}
		case subject.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field deck_subjects", values[i])
			} else if value.Valid {
				s.deck_subjects = new(uuid.UUID)
				*s.deck_subjects = *value.S.(*uuid.UUID)
			}
		case subject.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_subjects", values[i])
			} else if value.Valid {
				s.user_subjects = new(string)
				*s.user_subjects = value.String
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the Subject entity.
func (s *Subject) QueryCards() *CardQuery {
	return (&SubjectClient{config: s.config}).QueryCards(s)
}

// QuerySimilar queries the "similar" edge of the Subject entity.
func (s *Subject) QuerySimilar() *SubjectQuery {
	return (&SubjectClient{config: s.config}).QuerySimilar(s)
}

// QueryDependencies queries the "dependencies" edge of the Subject entity.
func (s *Subject) QueryDependencies() *SubjectQuery {
	return (&SubjectClient{config: s.config}).QueryDependencies(s)
}

// QueryDependents queries the "dependents" edge of the Subject entity.
func (s *Subject) QueryDependents() *SubjectQuery {
	return (&SubjectClient{config: s.config}).QueryDependents(s)
}

// QueryDeck queries the "deck" edge of the Subject entity.
func (s *Subject) QueryDeck() *DeckQuery {
	return (&SubjectClient{config: s.config}).QueryDeck(s)
}

// QueryOwner queries the "owner" edge of the Subject entity.
func (s *Subject) QueryOwner() *UserQuery {
	return (&SubjectClient{config: s.config}).QueryOwner(s)
}

// Update returns a builder for updating this Subject.
// Note that you need to call Subject.Unwrap() before calling this method if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return (&SubjectClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(s.Kind)
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", s.Level))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	if v := s.Value; v != nil {
		builder.WriteString("value=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("value_image=")
	builder.WriteString(fmt.Sprintf("%v", s.ValueImage))
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(s.Slug)
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", s.Priority))
	builder.WriteString(", ")
	builder.WriteString("resources=")
	builder.WriteString(fmt.Sprintf("%v", s.Resources))
	builder.WriteString(", ")
	builder.WriteString("study_data=")
	builder.WriteString(fmt.Sprintf("%v", s.StudyData))
	builder.WriteString(", ")
	builder.WriteString("complimentary_study_data=")
	builder.WriteString(fmt.Sprintf("%v", s.ComplimentaryStudyData))
	builder.WriteByte(')')
	return builder.String()
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject

func (s Subjects) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
