// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"sixels.io/manekani/ent/card"
	"sixels.io/manekani/ent/predicate"
	"sixels.io/manekani/ent/subject"
)

// SubjectUpdate is the builder for updating Subject entities.
type SubjectUpdate struct {
	config
	hooks    []Hook
	mutation *SubjectMutation
}

// Where appends a list predicates to the SubjectUpdate builder.
func (su *SubjectUpdate) Where(ps ...predicate.Subject) *SubjectUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetKind sets the "kind" field.
func (su *SubjectUpdate) SetKind(s subject.Kind) *SubjectUpdate {
	su.mutation.SetKind(s)
	return su
}

// SetLevel sets the "level" field.
func (su *SubjectUpdate) SetLevel(i int32) *SubjectUpdate {
	su.mutation.ResetLevel()
	su.mutation.SetLevel(i)
	return su
}

// AddLevel adds i to the "level" field.
func (su *SubjectUpdate) AddLevel(i int32) *SubjectUpdate {
	su.mutation.AddLevel(i)
	return su
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (su *SubjectUpdate) AddCardIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.AddCardIDs(ids...)
	return su
}

// AddCards adds the "cards" edges to the Card entity.
func (su *SubjectUpdate) AddCards(c ...*Card) *SubjectUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCardIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (su *SubjectUpdate) Mutation() *SubjectMutation {
	return su.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (su *SubjectUpdate) ClearCards() *SubjectUpdate {
	su.mutation.ClearCards()
	return su
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (su *SubjectUpdate) RemoveCardIDs(ids ...uuid.UUID) *SubjectUpdate {
	su.mutation.RemoveCardIDs(ids...)
	return su
}

// RemoveCards removes "cards" edges to Card entities.
func (su *SubjectUpdate) RemoveCards(c ...*Card) *SubjectUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCardIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubjectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubjectUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubjectUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubjectUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SubjectUpdate) check() error {
	if v, ok := su.mutation.Kind(); ok {
		if err := subject.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "Subject.kind": %w`, err)}
		}
	}
	if v, ok := su.mutation.Level(); ok {
		if err := subject.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Subject.level": %w`, err)}
		}
	}
	return nil
}

func (su *SubjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subject.Table,
			Columns: subject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subject.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Kind(); ok {
		_spec.SetField(subject.FieldKind, field.TypeEnum, value)
	}
	if value, ok := su.mutation.Level(); ok {
		_spec.SetField(subject.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedLevel(); ok {
		_spec.AddField(subject.FieldLevel, field.TypeInt32, value)
	}
	if su.mutation.CardsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCardsIDs(); len(nodes) > 0 && !su.mutation.CardsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CardsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SubjectUpdateOne is the builder for updating a single Subject entity.
type SubjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubjectMutation
}

// SetKind sets the "kind" field.
func (suo *SubjectUpdateOne) SetKind(s subject.Kind) *SubjectUpdateOne {
	suo.mutation.SetKind(s)
	return suo
}

// SetLevel sets the "level" field.
func (suo *SubjectUpdateOne) SetLevel(i int32) *SubjectUpdateOne {
	suo.mutation.ResetLevel()
	suo.mutation.SetLevel(i)
	return suo
}

// AddLevel adds i to the "level" field.
func (suo *SubjectUpdateOne) AddLevel(i int32) *SubjectUpdateOne {
	suo.mutation.AddLevel(i)
	return suo
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (suo *SubjectUpdateOne) AddCardIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.AddCardIDs(ids...)
	return suo
}

// AddCards adds the "cards" edges to the Card entity.
func (suo *SubjectUpdateOne) AddCards(c ...*Card) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCardIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (suo *SubjectUpdateOne) Mutation() *SubjectMutation {
	return suo.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (suo *SubjectUpdateOne) ClearCards() *SubjectUpdateOne {
	suo.mutation.ClearCards()
	return suo
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (suo *SubjectUpdateOne) RemoveCardIDs(ids ...uuid.UUID) *SubjectUpdateOne {
	suo.mutation.RemoveCardIDs(ids...)
	return suo
}

// RemoveCards removes "cards" edges to Card entities.
func (suo *SubjectUpdateOne) RemoveCards(c ...*Card) *SubjectUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCardIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubjectUpdateOne) Select(field string, fields ...string) *SubjectUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subject entity.
func (suo *SubjectUpdateOne) Save(ctx context.Context) (*Subject, error) {
	var (
		err  error
		node *Subject
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (suo *SubjectUpdateOne) SaveX(ctx context.Context) *Subject {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubjectUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubjectUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SubjectUpdateOne) check() error {
	if v, ok := suo.mutation.Kind(); ok {
		if err := subject.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "Subject.kind": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Level(); ok {
		if err := subject.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Subject.level": %w`, err)}
		}
	}
	return nil
}

func (suo *SubjectUpdateOne) sqlSave(ctx context.Context) (_node *Subject, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subject.Table,
			Columns: subject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subject.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subject.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subject.FieldID)
		for _, f := range fields {
			if !subject.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Kind(); ok {
		_spec.SetField(subject.FieldKind, field.TypeEnum, value)
	}
	if value, ok := suo.mutation.Level(); ok {
		_spec.SetField(subject.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedLevel(); ok {
		_spec.AddField(subject.FieldLevel, field.TypeInt32, value)
	}
	if suo.mutation.CardsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCardsIDs(); len(nodes) > 0 && !suo.mutation.CardsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CardsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Subject{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
