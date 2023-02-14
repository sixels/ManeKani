// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"sixels.io/manekani/core/domain/cards"
	"sixels.io/manekani/ent/card"
	"sixels.io/manekani/ent/deck"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
)

// SubjectCreate is the builder for creating a Subject entity.
type SubjectCreate struct {
	config
	mutation *SubjectMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubjectCreate) SetCreatedAt(t time.Time) *SubjectCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableCreatedAt(t *time.Time) *SubjectCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubjectCreate) SetUpdatedAt(t time.Time) *SubjectCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableUpdatedAt(t *time.Time) *SubjectCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetKind sets the "kind" field.
func (sc *SubjectCreate) SetKind(s string) *SubjectCreate {
	sc.mutation.SetKind(s)
	return sc
}

// SetLevel sets the "level" field.
func (sc *SubjectCreate) SetLevel(i int32) *SubjectCreate {
	sc.mutation.SetLevel(i)
	return sc
}

// SetName sets the "name" field.
func (sc *SubjectCreate) SetName(s string) *SubjectCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetValue sets the "value" field.
func (sc *SubjectCreate) SetValue(s string) *SubjectCreate {
	sc.mutation.SetValue(s)
	return sc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableValue(s *string) *SubjectCreate {
	if s != nil {
		sc.SetValue(*s)
	}
	return sc
}

// SetValueImage sets the "value_image" field.
func (sc *SubjectCreate) SetValueImage(cc *cards.RemoteContent) *SubjectCreate {
	sc.mutation.SetValueImage(cc)
	return sc
}

// SetSlug sets the "slug" field.
func (sc *SubjectCreate) SetSlug(s string) *SubjectCreate {
	sc.mutation.SetSlug(s)
	return sc
}

// SetPriority sets the "priority" field.
func (sc *SubjectCreate) SetPriority(u uint8) *SubjectCreate {
	sc.mutation.SetPriority(u)
	return sc
}

// SetResources sets the "resources" field.
func (sc *SubjectCreate) SetResources(mc *map[string][]cards.RemoteContent) *SubjectCreate {
	sc.mutation.SetResources(mc)
	return sc
}

// SetStudyData sets the "study_data" field.
func (sc *SubjectCreate) SetStudyData(cd []cards.StudyData) *SubjectCreate {
	sc.mutation.SetStudyData(cd)
	return sc
}

// SetComplimentaryStudyData sets the "complimentary_study_data" field.
func (sc *SubjectCreate) SetComplimentaryStudyData(m *[]map[string]string) *SubjectCreate {
	sc.mutation.SetComplimentaryStudyData(m)
	return sc
}

// SetID sets the "id" field.
func (sc *SubjectCreate) SetID(u uuid.UUID) *SubjectCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SubjectCreate) SetNillableID(u *uuid.UUID) *SubjectCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (sc *SubjectCreate) AddCardIDs(ids ...uuid.UUID) *SubjectCreate {
	sc.mutation.AddCardIDs(ids...)
	return sc
}

// AddCards adds the "cards" edges to the Card entity.
func (sc *SubjectCreate) AddCards(c ...*Card) *SubjectCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCardIDs(ids...)
}

// AddSimilarIDs adds the "similar" edge to the Subject entity by IDs.
func (sc *SubjectCreate) AddSimilarIDs(ids ...uuid.UUID) *SubjectCreate {
	sc.mutation.AddSimilarIDs(ids...)
	return sc
}

// AddSimilar adds the "similar" edges to the Subject entity.
func (sc *SubjectCreate) AddSimilar(s ...*Subject) *SubjectCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddSimilarIDs(ids...)
}

// AddDependencyIDs adds the "dependencies" edge to the Subject entity by IDs.
func (sc *SubjectCreate) AddDependencyIDs(ids ...uuid.UUID) *SubjectCreate {
	sc.mutation.AddDependencyIDs(ids...)
	return sc
}

// AddDependencies adds the "dependencies" edges to the Subject entity.
func (sc *SubjectCreate) AddDependencies(s ...*Subject) *SubjectCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddDependencyIDs(ids...)
}

// AddDependentIDs adds the "dependents" edge to the Subject entity by IDs.
func (sc *SubjectCreate) AddDependentIDs(ids ...uuid.UUID) *SubjectCreate {
	sc.mutation.AddDependentIDs(ids...)
	return sc
}

// AddDependents adds the "dependents" edges to the Subject entity.
func (sc *SubjectCreate) AddDependents(s ...*Subject) *SubjectCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddDependentIDs(ids...)
}

// SetDeckID sets the "deck" edge to the Deck entity by ID.
func (sc *SubjectCreate) SetDeckID(id uuid.UUID) *SubjectCreate {
	sc.mutation.SetDeckID(id)
	return sc
}

// SetDeck sets the "deck" edge to the Deck entity.
func (sc *SubjectCreate) SetDeck(d *Deck) *SubjectCreate {
	return sc.SetDeckID(d.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (sc *SubjectCreate) SetOwnerID(id string) *SubjectCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetOwner sets the "owner" edge to the User entity.
func (sc *SubjectCreate) SetOwner(u *User) *SubjectCreate {
	return sc.SetOwnerID(u.ID)
}

// Mutation returns the SubjectMutation object of the builder.
func (sc *SubjectCreate) Mutation() *SubjectMutation {
	return sc.mutation
}

// Save creates the Subject in the database.
func (sc *SubjectCreate) Save(ctx context.Context) (*Subject, error) {
	var (
		err  error
		node *Subject
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Subject)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubjectMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubjectCreate) SaveX(ctx context.Context) *Subject {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubjectCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubjectCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubjectCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := subject.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := subject.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := subject.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubjectCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Subject.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Subject.updated_at"`)}
	}
	if _, ok := sc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "Subject.kind"`)}
	}
	if _, ok := sc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Subject.level"`)}
	}
	if v, ok := sc.mutation.Level(); ok {
		if err := subject.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Subject.level": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Subject.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := subject.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subject.name": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Value(); ok {
		if err := subject.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Subject.value": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Subject.slug"`)}
	}
	if v, ok := sc.mutation.Slug(); ok {
		if err := subject.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Subject.slug": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Subject.priority"`)}
	}
	if _, ok := sc.mutation.StudyData(); !ok {
		return &ValidationError{Name: "study_data", err: errors.New(`ent: missing required field "Subject.study_data"`)}
	}
	if _, ok := sc.mutation.ComplimentaryStudyData(); !ok {
		return &ValidationError{Name: "complimentary_study_data", err: errors.New(`ent: missing required field "Subject.complimentary_study_data"`)}
	}
	if _, ok := sc.mutation.DeckID(); !ok {
		return &ValidationError{Name: "deck", err: errors.New(`ent: missing required edge "Subject.deck"`)}
	}
	if _, ok := sc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Subject.owner"`)}
	}
	return nil
}

func (sc *SubjectCreate) sqlSave(ctx context.Context) (*Subject, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (sc *SubjectCreate) createSpec() (*Subject, *sqlgraph.CreateSpec) {
	var (
		_node = &Subject{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subject.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subject.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(subject.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(subject.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Kind(); ok {
		_spec.SetField(subject.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := sc.mutation.Level(); ok {
		_spec.SetField(subject.FieldLevel, field.TypeInt32, value)
		_node.Level = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(subject.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Value(); ok {
		_spec.SetField(subject.FieldValue, field.TypeString, value)
		_node.Value = &value
	}
	if value, ok := sc.mutation.ValueImage(); ok {
		_spec.SetField(subject.FieldValueImage, field.TypeJSON, value)
		_node.ValueImage = value
	}
	if value, ok := sc.mutation.Slug(); ok {
		_spec.SetField(subject.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := sc.mutation.Priority(); ok {
		_spec.SetField(subject.FieldPriority, field.TypeUint8, value)
		_node.Priority = value
	}
	if value, ok := sc.mutation.Resources(); ok {
		_spec.SetField(subject.FieldResources, field.TypeJSON, value)
		_node.Resources = value
	}
	if value, ok := sc.mutation.StudyData(); ok {
		_spec.SetField(subject.FieldStudyData, field.TypeJSON, value)
		_node.StudyData = value
	}
	if value, ok := sc.mutation.ComplimentaryStudyData(); ok {
		_spec.SetField(subject.FieldComplimentaryStudyData, field.TypeJSON, value)
		_node.ComplimentaryStudyData = value
	}
	if nodes := sc.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.CardsTable,
			Columns: []string{subject.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.SimilarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subject.SimilarTable,
			Columns: subject.SimilarPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.DependenciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subject.DependenciesTable,
			Columns: subject.DependenciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.DependentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subject.DependentsTable,
			Columns: subject.DependentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.DeckIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.DeckTable,
			Columns: []string{subject.DeckColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deck.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deck_subjects = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.OwnerTable,
			Columns: []string{subject.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_subjects = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubjectCreateBulk is the builder for creating many Subject entities in bulk.
type SubjectCreateBulk struct {
	config
	builders []*SubjectCreate
}

// Save creates the Subject entities in the database.
func (scb *SubjectCreateBulk) Save(ctx context.Context) ([]*Subject, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subject, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubjectCreateBulk) SaveX(ctx context.Context) []*Subject {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubjectCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubjectCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
