// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	ulid "github.com/oklog/ulid/v2"
	"github.com/sixels/manekani/core/domain/tokens"
	"github.com/sixels/manekani/ent/apitoken"
	"github.com/sixels/manekani/ent/user"
)

// ApiTokenCreate is the builder for creating a ApiToken entity.
type ApiTokenCreate struct {
	config
	mutation *ApiTokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (atc *ApiTokenCreate) SetName(s string) *ApiTokenCreate {
	atc.mutation.SetName(s)
	return atc
}

// SetStatus sets the "status" field.
func (atc *ApiTokenCreate) SetStatus(tts tokens.APITokenStatus) *ApiTokenCreate {
	atc.mutation.SetStatus(tts)
	return atc
}

// SetUsedAt sets the "used_at" field.
func (atc *ApiTokenCreate) SetUsedAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetUsedAt(t)
	return atc
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableUsedAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetUsedAt(*t)
	}
	return atc
}

// SetToken sets the "token" field.
func (atc *ApiTokenCreate) SetToken(s string) *ApiTokenCreate {
	atc.mutation.SetToken(s)
	return atc
}

// SetPrefix sets the "prefix" field.
func (atc *ApiTokenCreate) SetPrefix(s string) *ApiTokenCreate {
	atc.mutation.SetPrefix(s)
	return atc
}

// SetClaims sets the "claims" field.
func (atc *ApiTokenCreate) SetClaims(ttc tokens.APITokenClaims) *ApiTokenCreate {
	atc.mutation.SetClaims(ttc)
	return atc
}

// SetID sets the "id" field.
func (atc *ApiTokenCreate) SetID(u ulid.ULID) *ApiTokenCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableID(u *ulid.ULID) *ApiTokenCreate {
	if u != nil {
		atc.SetID(*u)
	}
	return atc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (atc *ApiTokenCreate) SetUserID(id string) *ApiTokenCreate {
	atc.mutation.SetUserID(id)
	return atc
}

// SetUser sets the "user" edge to the User entity.
func (atc *ApiTokenCreate) SetUser(u *User) *ApiTokenCreate {
	return atc.SetUserID(u.ID)
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atc *ApiTokenCreate) Mutation() *ApiTokenMutation {
	return atc.mutation
}

// Save creates the ApiToken in the database.
func (atc *ApiTokenCreate) Save(ctx context.Context) (*ApiToken, error) {
	atc.defaults()
	return withHooks(ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *ApiTokenCreate) SaveX(ctx context.Context) *ApiToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *ApiTokenCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *ApiTokenCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *ApiTokenCreate) defaults() {
	if _, ok := atc.mutation.ID(); !ok {
		v := apitoken.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *ApiTokenCreate) check() error {
	if _, ok := atc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ApiToken.name"`)}
	}
	if v, ok := atc.mutation.Name(); ok {
		if err := apitoken.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ApiToken.name": %w`, err)}
		}
	}
	if _, ok := atc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "ApiToken.status"`)}
	}
	if v, ok := atc.mutation.Status(); ok {
		if err := apitoken.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ApiToken.status": %w`, err)}
		}
	}
	if _, ok := atc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "ApiToken.token"`)}
	}
	if _, ok := atc.mutation.Prefix(); !ok {
		return &ValidationError{Name: "prefix", err: errors.New(`ent: missing required field "ApiToken.prefix"`)}
	}
	if _, ok := atc.mutation.Claims(); !ok {
		return &ValidationError{Name: "claims", err: errors.New(`ent: missing required field "ApiToken.claims"`)}
	}
	if _, ok := atc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ApiToken.user"`)}
	}
	return nil
}

func (atc *ApiTokenCreate) sqlSave(ctx context.Context) (*ApiToken, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*ulid.ULID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *ApiTokenCreate) createSpec() (*ApiToken, *sqlgraph.CreateSpec) {
	var (
		_node = &ApiToken{config: atc.config}
		_spec = sqlgraph.NewCreateSpec(apitoken.Table, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeBytes))
	)
	_spec.OnConflict = atc.conflict
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := atc.mutation.Status(); ok {
		_spec.SetField(apitoken.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := atc.mutation.UsedAt(); ok {
		_spec.SetField(apitoken.FieldUsedAt, field.TypeTime, value)
		_node.UsedAt = &value
	}
	if value, ok := atc.mutation.Token(); ok {
		_spec.SetField(apitoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := atc.mutation.Prefix(); ok {
		_spec.SetField(apitoken.FieldPrefix, field.TypeString, value)
		_node.Prefix = value
	}
	if value, ok := atc.mutation.Claims(); ok {
		_spec.SetField(apitoken.FieldClaims, field.TypeJSON, value)
		_node.Claims = value
	}
	if nodes := atc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_api_tokens = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ApiToken.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApiTokenUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (atc *ApiTokenCreate) OnConflict(opts ...sql.ConflictOption) *ApiTokenUpsertOne {
	atc.conflict = opts
	return &ApiTokenUpsertOne{
		create: atc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atc *ApiTokenCreate) OnConflictColumns(columns ...string) *ApiTokenUpsertOne {
	atc.conflict = append(atc.conflict, sql.ConflictColumns(columns...))
	return &ApiTokenUpsertOne{
		create: atc,
	}
}

type (
	// ApiTokenUpsertOne is the builder for "upsert"-ing
	//  one ApiToken node.
	ApiTokenUpsertOne struct {
		create *ApiTokenCreate
	}

	// ApiTokenUpsert is the "OnConflict" setter.
	ApiTokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetStatus sets the "status" field.
func (u *ApiTokenUpsert) SetStatus(v tokens.APITokenStatus) *ApiTokenUpsert {
	u.Set(apitoken.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateStatus() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldStatus)
	return u
}

// SetUsedAt sets the "used_at" field.
func (u *ApiTokenUpsert) SetUsedAt(v time.Time) *ApiTokenUpsert {
	u.Set(apitoken.FieldUsedAt, v)
	return u
}

// UpdateUsedAt sets the "used_at" field to the value that was provided on create.
func (u *ApiTokenUpsert) UpdateUsedAt() *ApiTokenUpsert {
	u.SetExcluded(apitoken.FieldUsedAt)
	return u
}

// ClearUsedAt clears the value of the "used_at" field.
func (u *ApiTokenUpsert) ClearUsedAt() *ApiTokenUpsert {
	u.SetNull(apitoken.FieldUsedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApiTokenUpsertOne) UpdateNewValues() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(apitoken.FieldID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(apitoken.FieldName)
		}
		if _, exists := u.create.mutation.Token(); exists {
			s.SetIgnore(apitoken.FieldToken)
		}
		if _, exists := u.create.mutation.Prefix(); exists {
			s.SetIgnore(apitoken.FieldPrefix)
		}
		if _, exists := u.create.mutation.Claims(); exists {
			s.SetIgnore(apitoken.FieldClaims)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ApiTokenUpsertOne) Ignore() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiTokenUpsertOne) DoNothing() *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApiTokenCreate.OnConflict
// documentation for more info.
func (u *ApiTokenUpsertOne) Update(set func(*ApiTokenUpsert)) *ApiTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *ApiTokenUpsertOne) SetStatus(v tokens.APITokenStatus) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateStatus() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateStatus()
	})
}

// SetUsedAt sets the "used_at" field.
func (u *ApiTokenUpsertOne) SetUsedAt(v time.Time) *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUsedAt(v)
	})
}

// UpdateUsedAt sets the "used_at" field to the value that was provided on create.
func (u *ApiTokenUpsertOne) UpdateUsedAt() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUsedAt()
	})
}

// ClearUsedAt clears the value of the "used_at" field.
func (u *ApiTokenUpsertOne) ClearUsedAt() *ApiTokenUpsertOne {
	return u.Update(func(s *ApiTokenUpsert) {
		s.ClearUsedAt()
	})
}

// Exec executes the query.
func (u *ApiTokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ApiTokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiTokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ApiTokenUpsertOne) ID(ctx context.Context) (id ulid.ULID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ApiTokenUpsertOne.ID is not supported by MySQL driver. Use ApiTokenUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ApiTokenUpsertOne) IDX(ctx context.Context) ulid.ULID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ApiTokenCreateBulk is the builder for creating many ApiToken entities in bulk.
type ApiTokenCreateBulk struct {
	config
	builders []*ApiTokenCreate
	conflict []sql.ConflictOption
}

// Save creates the ApiToken entities in the database.
func (atcb *ApiTokenCreateBulk) Save(ctx context.Context) ([]*ApiToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*ApiToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApiTokenMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = atcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) SaveX(ctx context.Context) []*ApiToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *ApiTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ApiToken.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ApiTokenUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (atcb *ApiTokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *ApiTokenUpsertBulk {
	atcb.conflict = opts
	return &ApiTokenUpsertBulk{
		create: atcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (atcb *ApiTokenCreateBulk) OnConflictColumns(columns ...string) *ApiTokenUpsertBulk {
	atcb.conflict = append(atcb.conflict, sql.ConflictColumns(columns...))
	return &ApiTokenUpsertBulk{
		create: atcb,
	}
}

// ApiTokenUpsertBulk is the builder for "upsert"-ing
// a bulk of ApiToken nodes.
type ApiTokenUpsertBulk struct {
	create *ApiTokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apitoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ApiTokenUpsertBulk) UpdateNewValues() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(apitoken.FieldID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(apitoken.FieldName)
			}
			if _, exists := b.mutation.Token(); exists {
				s.SetIgnore(apitoken.FieldToken)
			}
			if _, exists := b.mutation.Prefix(); exists {
				s.SetIgnore(apitoken.FieldPrefix)
			}
			if _, exists := b.mutation.Claims(); exists {
				s.SetIgnore(apitoken.FieldClaims)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ApiToken.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ApiTokenUpsertBulk) Ignore() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiTokenUpsertBulk) DoNothing() *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ApiTokenCreateBulk.OnConflict
// documentation for more info.
func (u *ApiTokenUpsertBulk) Update(set func(*ApiTokenUpsert)) *ApiTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetStatus sets the "status" field.
func (u *ApiTokenUpsertBulk) SetStatus(v tokens.APITokenStatus) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateStatus() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateStatus()
	})
}

// SetUsedAt sets the "used_at" field.
func (u *ApiTokenUpsertBulk) SetUsedAt(v time.Time) *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.SetUsedAt(v)
	})
}

// UpdateUsedAt sets the "used_at" field to the value that was provided on create.
func (u *ApiTokenUpsertBulk) UpdateUsedAt() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.UpdateUsedAt()
	})
}

// ClearUsedAt clears the value of the "used_at" field.
func (u *ApiTokenUpsertBulk) ClearUsedAt() *ApiTokenUpsertBulk {
	return u.Update(func(s *ApiTokenUpsert) {
		s.ClearUsedAt()
	})
}

// Exec executes the query.
func (u *ApiTokenUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ApiTokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ApiTokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiTokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
