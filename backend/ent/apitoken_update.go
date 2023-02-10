// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"sixels.io/manekani/ent/apitoken"
	"sixels.io/manekani/ent/predicate"
	"sixels.io/manekani/ent/user"
)

// ApiTokenUpdate is the builder for updating ApiToken entities.
type ApiTokenUpdate struct {
	config
	hooks    []Hook
	mutation *ApiTokenMutation
}

// Where appends a list predicates to the ApiTokenUpdate builder.
func (atu *ApiTokenUpdate) Where(ps ...predicate.ApiToken) *ApiTokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (atu *ApiTokenUpdate) SetUserID(id string) *ApiTokenUpdate {
	atu.mutation.SetUserID(id)
	return atu
}

// SetUser sets the "user" edge to the User entity.
func (atu *ApiTokenUpdate) SetUser(u *User) *ApiTokenUpdate {
	return atu.SetUserID(u.ID)
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atu *ApiTokenUpdate) Mutation() *ApiTokenMutation {
	return atu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atu *ApiTokenUpdate) ClearUser() *ApiTokenUpdate {
	atu.mutation.ClearUser()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ApiTokenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atu.hooks) == 0 {
		if err = atu.check(); err != nil {
			return 0, err
		}
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApiTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atu.check(); err != nil {
				return 0, err
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ApiTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ApiTokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ApiTokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *ApiTokenUpdate) check() error {
	if _, ok := atu.mutation.UserID(); atu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiToken.user"`)
	}
	return nil
}

func (atu *ApiTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apitoken.Table,
			Columns: apitoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: apitoken.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ApiTokenUpdateOne is the builder for updating a single ApiToken entity.
type ApiTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApiTokenMutation
}

// SetUserID sets the "user" edge to the User entity by ID.
func (atuo *ApiTokenUpdateOne) SetUserID(id string) *ApiTokenUpdateOne {
	atuo.mutation.SetUserID(id)
	return atuo
}

// SetUser sets the "user" edge to the User entity.
func (atuo *ApiTokenUpdateOne) SetUser(u *User) *ApiTokenUpdateOne {
	return atuo.SetUserID(u.ID)
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atuo *ApiTokenUpdateOne) Mutation() *ApiTokenMutation {
	return atuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atuo *ApiTokenUpdateOne) ClearUser() *ApiTokenUpdateOne {
	atuo.mutation.ClearUser()
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ApiTokenUpdateOne) Select(field string, fields ...string) *ApiTokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ApiToken entity.
func (atuo *ApiTokenUpdateOne) Save(ctx context.Context) (*ApiToken, error) {
	var (
		err  error
		node *ApiToken
	)
	if len(atuo.hooks) == 0 {
		if err = atuo.check(); err != nil {
			return nil, err
		}
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApiTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atuo.check(); err != nil {
				return nil, err
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, atuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ApiToken)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ApiTokenMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ApiTokenUpdateOne) SaveX(ctx context.Context) *ApiToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ApiTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ApiTokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *ApiTokenUpdateOne) check() error {
	if _, ok := atuo.mutation.UserID(); atuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApiToken.user"`)
	}
	return nil
}

func (atuo *ApiTokenUpdateOne) sqlSave(ctx context.Context) (_node *ApiToken, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apitoken.Table,
			Columns: apitoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: apitoken.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApiToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apitoken.FieldID)
		for _, f := range fields {
			if !apitoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apitoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ApiToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
