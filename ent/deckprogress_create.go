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
	"github.com/sixels/manekani/ent/card"
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/user"
)

// DeckProgressCreate is the builder for creating a DeckProgress entity.
type DeckProgressCreate struct {
	config
	mutation *DeckProgressMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetLevel sets the "level" field.
func (dpc *DeckProgressCreate) SetLevel(u uint32) *DeckProgressCreate {
	dpc.mutation.SetLevel(u)
	return dpc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (dpc *DeckProgressCreate) SetNillableLevel(u *uint32) *DeckProgressCreate {
	if u != nil {
		dpc.SetLevel(*u)
	}
	return dpc
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (dpc *DeckProgressCreate) AddCardIDs(ids ...uuid.UUID) *DeckProgressCreate {
	dpc.mutation.AddCardIDs(ids...)
	return dpc
}

// AddCards adds the "cards" edges to the Card entity.
func (dpc *DeckProgressCreate) AddCards(c ...*Card) *DeckProgressCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dpc.AddCardIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (dpc *DeckProgressCreate) SetUserID(id string) *DeckProgressCreate {
	dpc.mutation.SetUserID(id)
	return dpc
}

// SetUser sets the "user" edge to the User entity.
func (dpc *DeckProgressCreate) SetUser(u *User) *DeckProgressCreate {
	return dpc.SetUserID(u.ID)
}

// SetDeckID sets the "deck" edge to the Deck entity by ID.
func (dpc *DeckProgressCreate) SetDeckID(id uuid.UUID) *DeckProgressCreate {
	dpc.mutation.SetDeckID(id)
	return dpc
}

// SetDeck sets the "deck" edge to the Deck entity.
func (dpc *DeckProgressCreate) SetDeck(d *Deck) *DeckProgressCreate {
	return dpc.SetDeckID(d.ID)
}

// Mutation returns the DeckProgressMutation object of the builder.
func (dpc *DeckProgressCreate) Mutation() *DeckProgressMutation {
	return dpc.mutation
}

// Save creates the DeckProgress in the database.
func (dpc *DeckProgressCreate) Save(ctx context.Context) (*DeckProgress, error) {
	var (
		err  error
		node *DeckProgress
	)
	dpc.defaults()
	if len(dpc.hooks) == 0 {
		if err = dpc.check(); err != nil {
			return nil, err
		}
		node, err = dpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeckProgressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dpc.check(); err != nil {
				return nil, err
			}
			dpc.mutation = mutation
			if node, err = dpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dpc.hooks) - 1; i >= 0; i-- {
			if dpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DeckProgress)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeckProgressMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dpc *DeckProgressCreate) SaveX(ctx context.Context) *DeckProgress {
	v, err := dpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dpc *DeckProgressCreate) Exec(ctx context.Context) error {
	_, err := dpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpc *DeckProgressCreate) ExecX(ctx context.Context) {
	if err := dpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dpc *DeckProgressCreate) defaults() {
	if _, ok := dpc.mutation.Level(); !ok {
		v := deckprogress.DefaultLevel
		dpc.mutation.SetLevel(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dpc *DeckProgressCreate) check() error {
	if _, ok := dpc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "DeckProgress.level"`)}
	}
	if v, ok := dpc.mutation.Level(); ok {
		if err := deckprogress.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "DeckProgress.level": %w`, err)}
		}
	}
	if _, ok := dpc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "DeckProgress.user"`)}
	}
	if _, ok := dpc.mutation.DeckID(); !ok {
		return &ValidationError{Name: "deck", err: errors.New(`ent: missing required edge "DeckProgress.deck"`)}
	}
	return nil
}

func (dpc *DeckProgressCreate) sqlSave(ctx context.Context) (*DeckProgress, error) {
	_node, _spec := dpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dpc *DeckProgressCreate) createSpec() (*DeckProgress, *sqlgraph.CreateSpec) {
	var (
		_node = &DeckProgress{config: dpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deckprogress.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deckprogress.FieldID,
			},
		}
	)
	_spec.OnConflict = dpc.conflict
	if value, ok := dpc.mutation.Level(); ok {
		_spec.SetField(deckprogress.FieldLevel, field.TypeUint32, value)
		_node.Level = value
	}
	if nodes := dpc.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deckprogress.CardsTable,
			Columns: []string{deckprogress.CardsColumn},
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
	if nodes := dpc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deckprogress.UserTable,
			Columns: []string{deckprogress.UserColumn},
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
		_node.user_decks_progress = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dpc.mutation.DeckIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deckprogress.DeckTable,
			Columns: []string{deckprogress.DeckColumn},
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
		_node.deck_users_progress = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeckProgress.Create().
//		SetLevel(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeckProgressUpsert) {
//			SetLevel(v+v).
//		}).
//		Exec(ctx)
func (dpc *DeckProgressCreate) OnConflict(opts ...sql.ConflictOption) *DeckProgressUpsertOne {
	dpc.conflict = opts
	return &DeckProgressUpsertOne{
		create: dpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeckProgress.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dpc *DeckProgressCreate) OnConflictColumns(columns ...string) *DeckProgressUpsertOne {
	dpc.conflict = append(dpc.conflict, sql.ConflictColumns(columns...))
	return &DeckProgressUpsertOne{
		create: dpc,
	}
}

type (
	// DeckProgressUpsertOne is the builder for "upsert"-ing
	//  one DeckProgress node.
	DeckProgressUpsertOne struct {
		create *DeckProgressCreate
	}

	// DeckProgressUpsert is the "OnConflict" setter.
	DeckProgressUpsert struct {
		*sql.UpdateSet
	}
)

// SetLevel sets the "level" field.
func (u *DeckProgressUpsert) SetLevel(v uint32) *DeckProgressUpsert {
	u.Set(deckprogress.FieldLevel, v)
	return u
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *DeckProgressUpsert) UpdateLevel() *DeckProgressUpsert {
	u.SetExcluded(deckprogress.FieldLevel)
	return u
}

// AddLevel adds v to the "level" field.
func (u *DeckProgressUpsert) AddLevel(v uint32) *DeckProgressUpsert {
	u.Add(deckprogress.FieldLevel, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DeckProgress.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DeckProgressUpsertOne) UpdateNewValues() *DeckProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeckProgress.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeckProgressUpsertOne) Ignore() *DeckProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeckProgressUpsertOne) DoNothing() *DeckProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeckProgressCreate.OnConflict
// documentation for more info.
func (u *DeckProgressUpsertOne) Update(set func(*DeckProgressUpsert)) *DeckProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeckProgressUpsert{UpdateSet: update})
	}))
	return u
}

// SetLevel sets the "level" field.
func (u *DeckProgressUpsertOne) SetLevel(v uint32) *DeckProgressUpsertOne {
	return u.Update(func(s *DeckProgressUpsert) {
		s.SetLevel(v)
	})
}

// AddLevel adds v to the "level" field.
func (u *DeckProgressUpsertOne) AddLevel(v uint32) *DeckProgressUpsertOne {
	return u.Update(func(s *DeckProgressUpsert) {
		s.AddLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *DeckProgressUpsertOne) UpdateLevel() *DeckProgressUpsertOne {
	return u.Update(func(s *DeckProgressUpsert) {
		s.UpdateLevel()
	})
}

// Exec executes the query.
func (u *DeckProgressUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeckProgressCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeckProgressUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeckProgressUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeckProgressUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeckProgressCreateBulk is the builder for creating many DeckProgress entities in bulk.
type DeckProgressCreateBulk struct {
	config
	builders []*DeckProgressCreate
	conflict []sql.ConflictOption
}

// Save creates the DeckProgress entities in the database.
func (dpcb *DeckProgressCreateBulk) Save(ctx context.Context) ([]*DeckProgress, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dpcb.builders))
	nodes := make([]*DeckProgress, len(dpcb.builders))
	mutators := make([]Mutator, len(dpcb.builders))
	for i := range dpcb.builders {
		func(i int, root context.Context) {
			builder := dpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeckProgressMutation)
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
					_, err = mutators[i+1].Mutate(root, dpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, dpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dpcb *DeckProgressCreateBulk) SaveX(ctx context.Context) []*DeckProgress {
	v, err := dpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dpcb *DeckProgressCreateBulk) Exec(ctx context.Context) error {
	_, err := dpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpcb *DeckProgressCreateBulk) ExecX(ctx context.Context) {
	if err := dpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeckProgress.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeckProgressUpsert) {
//			SetLevel(v+v).
//		}).
//		Exec(ctx)
func (dpcb *DeckProgressCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeckProgressUpsertBulk {
	dpcb.conflict = opts
	return &DeckProgressUpsertBulk{
		create: dpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeckProgress.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dpcb *DeckProgressCreateBulk) OnConflictColumns(columns ...string) *DeckProgressUpsertBulk {
	dpcb.conflict = append(dpcb.conflict, sql.ConflictColumns(columns...))
	return &DeckProgressUpsertBulk{
		create: dpcb,
	}
}

// DeckProgressUpsertBulk is the builder for "upsert"-ing
// a bulk of DeckProgress nodes.
type DeckProgressUpsertBulk struct {
	create *DeckProgressCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeckProgress.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DeckProgressUpsertBulk) UpdateNewValues() *DeckProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeckProgress.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeckProgressUpsertBulk) Ignore() *DeckProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeckProgressUpsertBulk) DoNothing() *DeckProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeckProgressCreateBulk.OnConflict
// documentation for more info.
func (u *DeckProgressUpsertBulk) Update(set func(*DeckProgressUpsert)) *DeckProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeckProgressUpsert{UpdateSet: update})
	}))
	return u
}

// SetLevel sets the "level" field.
func (u *DeckProgressUpsertBulk) SetLevel(v uint32) *DeckProgressUpsertBulk {
	return u.Update(func(s *DeckProgressUpsert) {
		s.SetLevel(v)
	})
}

// AddLevel adds v to the "level" field.
func (u *DeckProgressUpsertBulk) AddLevel(v uint32) *DeckProgressUpsertBulk {
	return u.Update(func(s *DeckProgressUpsert) {
		s.AddLevel(v)
	})
}

// UpdateLevel sets the "level" field to the value that was provided on create.
func (u *DeckProgressUpsertBulk) UpdateLevel() *DeckProgressUpsertBulk {
	return u.Update(func(s *DeckProgressUpsert) {
		s.UpdateLevel()
	})
}

// Exec executes the query.
func (u *DeckProgressUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeckProgressCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeckProgressCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeckProgressUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
