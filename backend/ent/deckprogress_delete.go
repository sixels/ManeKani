// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/predicate"
)

// DeckProgressDelete is the builder for deleting a DeckProgress entity.
type DeckProgressDelete struct {
	config
	hooks    []Hook
	mutation *DeckProgressMutation
}

// Where appends a list predicates to the DeckProgressDelete builder.
func (dpd *DeckProgressDelete) Where(ps ...predicate.DeckProgress) *DeckProgressDelete {
	dpd.mutation.Where(ps...)
	return dpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dpd *DeckProgressDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dpd.hooks) == 0 {
		affected, err = dpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeckProgressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dpd.mutation = mutation
			affected, err = dpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dpd.hooks) - 1; i >= 0; i-- {
			if dpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpd *DeckProgressDelete) ExecX(ctx context.Context) int {
	n, err := dpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dpd *DeckProgressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: deckprogress.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deckprogress.FieldID,
			},
		},
	}
	if ps := dpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DeckProgressDeleteOne is the builder for deleting a single DeckProgress entity.
type DeckProgressDeleteOne struct {
	dpd *DeckProgressDelete
}

// Exec executes the deletion query.
func (dpdo *DeckProgressDeleteOne) Exec(ctx context.Context) error {
	n, err := dpdo.dpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deckprogress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dpdo *DeckProgressDeleteOne) ExecX(ctx context.Context) {
	dpdo.dpd.ExecX(ctx)
}
