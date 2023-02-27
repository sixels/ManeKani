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
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/subject"
	"github.com/sixels/manekani/ent/user"
)

// DeckCreate is the builder for creating a Deck entity.
type DeckCreate struct {
	config
	mutation *DeckMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeckCreate) SetCreatedAt(t time.Time) *DeckCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeckCreate) SetNillableCreatedAt(t *time.Time) *DeckCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeckCreate) SetUpdatedAt(t time.Time) *DeckCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeckCreate) SetNillableUpdatedAt(t *time.Time) *DeckCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DeckCreate) SetName(s string) *DeckCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DeckCreate) SetDescription(s string) *DeckCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DeckCreate) SetID(u uuid.UUID) *DeckCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DeckCreate) SetNillableID(u *uuid.UUID) *DeckCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// AddSubscriberIDs adds the "subscribers" edge to the User entity by IDs.
func (dc *DeckCreate) AddSubscriberIDs(ids ...string) *DeckCreate {
	dc.mutation.AddSubscriberIDs(ids...)
	return dc
}

// AddSubscribers adds the "subscribers" edges to the User entity.
func (dc *DeckCreate) AddSubscribers(u ...*User) *DeckCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddSubscriberIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (dc *DeckCreate) SetOwnerID(id string) *DeckCreate {
	dc.mutation.SetOwnerID(id)
	return dc
}

// SetOwner sets the "owner" edge to the User entity.
func (dc *DeckCreate) SetOwner(u *User) *DeckCreate {
	return dc.SetOwnerID(u.ID)
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (dc *DeckCreate) AddSubjectIDs(ids ...uuid.UUID) *DeckCreate {
	dc.mutation.AddSubjectIDs(ids...)
	return dc
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (dc *DeckCreate) AddSubjects(s ...*Subject) *DeckCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddSubjectIDs(ids...)
}

// AddUsersProgresIDs adds the "users_progress" edge to the DeckProgress entity by IDs.
func (dc *DeckCreate) AddUsersProgresIDs(ids ...int) *DeckCreate {
	dc.mutation.AddUsersProgresIDs(ids...)
	return dc
}

// AddUsersProgress adds the "users_progress" edges to the DeckProgress entity.
func (dc *DeckCreate) AddUsersProgress(d ...*DeckProgress) *DeckCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddUsersProgresIDs(ids...)
}

// Mutation returns the DeckMutation object of the builder.
func (dc *DeckCreate) Mutation() *DeckMutation {
	return dc.mutation
}

// Save creates the Deck in the database.
func (dc *DeckCreate) Save(ctx context.Context) (*Deck, error) {
	var (
		err  error
		node *Deck
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeckMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Deck)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeckMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeckCreate) SaveX(ctx context.Context) *Deck {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeckCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeckCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeckCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := deck.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := deck.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := deck.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeckCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Deck.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Deck.updated_at"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Deck.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := deck.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Deck.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Deck.description"`)}
	}
	if v, ok := dc.mutation.Description(); ok {
		if err := deck.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Deck.description": %w`, err)}
		}
	}
	if _, ok := dc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Deck.owner"`)}
	}
	return nil
}

func (dc *DeckCreate) sqlSave(ctx context.Context) (*Deck, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
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

func (dc *DeckCreate) createSpec() (*Deck, *sqlgraph.CreateSpec) {
	var (
		_node = &Deck{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deck.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deck.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(deck.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(deck.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(deck.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(deck.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := dc.mutation.SubscribersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deck.SubscribersTable,
			Columns: deck.SubscribersPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deck.OwnerTable,
			Columns: []string{deck.OwnerColumn},
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
		_node.user_decks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deck.SubjectsTable,
			Columns: []string{deck.SubjectsColumn},
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
	if nodes := dc.mutation.UsersProgressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deck.UsersProgressTable,
			Columns: []string{deck.UsersProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: deckprogress.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeckCreateBulk is the builder for creating many Deck entities in bulk.
type DeckCreateBulk struct {
	config
	builders []*DeckCreate
}

// Save creates the Deck entities in the database.
func (dcb *DeckCreateBulk) Save(ctx context.Context) ([]*Deck, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deck, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeckMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeckCreateBulk) SaveX(ctx context.Context) []*Deck {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeckCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeckCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
