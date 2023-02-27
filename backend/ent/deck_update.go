// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sixels/manekani/ent/deck"
	"github.com/sixels/manekani/ent/deckprogress"
	"github.com/sixels/manekani/ent/predicate"
	"github.com/sixels/manekani/ent/subject"
	"github.com/sixels/manekani/ent/user"
)

// DeckUpdate is the builder for updating Deck entities.
type DeckUpdate struct {
	config
	hooks    []Hook
	mutation *DeckMutation
}

// Where appends a list predicates to the DeckUpdate builder.
func (du *DeckUpdate) Where(ps ...predicate.Deck) *DeckUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DeckUpdate) SetUpdatedAt(t time.Time) *DeckUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetName sets the "name" field.
func (du *DeckUpdate) SetName(s string) *DeckUpdate {
	du.mutation.SetName(s)
	return du
}

// SetDescription sets the "description" field.
func (du *DeckUpdate) SetDescription(s string) *DeckUpdate {
	du.mutation.SetDescription(s)
	return du
}

// AddSubscriberIDs adds the "subscribers" edge to the User entity by IDs.
func (du *DeckUpdate) AddSubscriberIDs(ids ...string) *DeckUpdate {
	du.mutation.AddSubscriberIDs(ids...)
	return du
}

// AddSubscribers adds the "subscribers" edges to the User entity.
func (du *DeckUpdate) AddSubscribers(u ...*User) *DeckUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddSubscriberIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (du *DeckUpdate) SetOwnerID(id string) *DeckUpdate {
	du.mutation.SetOwnerID(id)
	return du
}

// SetOwner sets the "owner" edge to the User entity.
func (du *DeckUpdate) SetOwner(u *User) *DeckUpdate {
	return du.SetOwnerID(u.ID)
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (du *DeckUpdate) AddSubjectIDs(ids ...uuid.UUID) *DeckUpdate {
	du.mutation.AddSubjectIDs(ids...)
	return du
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (du *DeckUpdate) AddSubjects(s ...*Subject) *DeckUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddSubjectIDs(ids...)
}

// AddUsersProgresIDs adds the "users_progress" edge to the DeckProgress entity by IDs.
func (du *DeckUpdate) AddUsersProgresIDs(ids ...int) *DeckUpdate {
	du.mutation.AddUsersProgresIDs(ids...)
	return du
}

// AddUsersProgress adds the "users_progress" edges to the DeckProgress entity.
func (du *DeckUpdate) AddUsersProgress(d ...*DeckProgress) *DeckUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddUsersProgresIDs(ids...)
}

// Mutation returns the DeckMutation object of the builder.
func (du *DeckUpdate) Mutation() *DeckMutation {
	return du.mutation
}

// ClearSubscribers clears all "subscribers" edges to the User entity.
func (du *DeckUpdate) ClearSubscribers() *DeckUpdate {
	du.mutation.ClearSubscribers()
	return du
}

// RemoveSubscriberIDs removes the "subscribers" edge to User entities by IDs.
func (du *DeckUpdate) RemoveSubscriberIDs(ids ...string) *DeckUpdate {
	du.mutation.RemoveSubscriberIDs(ids...)
	return du
}

// RemoveSubscribers removes "subscribers" edges to User entities.
func (du *DeckUpdate) RemoveSubscribers(u ...*User) *DeckUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveSubscriberIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (du *DeckUpdate) ClearOwner() *DeckUpdate {
	du.mutation.ClearOwner()
	return du
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (du *DeckUpdate) ClearSubjects() *DeckUpdate {
	du.mutation.ClearSubjects()
	return du
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (du *DeckUpdate) RemoveSubjectIDs(ids ...uuid.UUID) *DeckUpdate {
	du.mutation.RemoveSubjectIDs(ids...)
	return du
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (du *DeckUpdate) RemoveSubjects(s ...*Subject) *DeckUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveSubjectIDs(ids...)
}

// ClearUsersProgress clears all "users_progress" edges to the DeckProgress entity.
func (du *DeckUpdate) ClearUsersProgress() *DeckUpdate {
	du.mutation.ClearUsersProgress()
	return du
}

// RemoveUsersProgresIDs removes the "users_progress" edge to DeckProgress entities by IDs.
func (du *DeckUpdate) RemoveUsersProgresIDs(ids ...int) *DeckUpdate {
	du.mutation.RemoveUsersProgresIDs(ids...)
	return du
}

// RemoveUsersProgress removes "users_progress" edges to DeckProgress entities.
func (du *DeckUpdate) RemoveUsersProgress(d ...*DeckProgress) *DeckUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveUsersProgresIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeckUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeckMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeckUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeckUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeckUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeckUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := deck.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeckUpdate) check() error {
	if v, ok := du.mutation.Name(); ok {
		if err := deck.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Deck.name": %w`, err)}
		}
	}
	if v, ok := du.mutation.Description(); ok {
		if err := deck.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Deck.description": %w`, err)}
		}
	}
	if _, ok := du.mutation.OwnerID(); du.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deck.owner"`)
	}
	return nil
}

func (du *DeckUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deck.Table,
			Columns: deck.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deck.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(deck.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(deck.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Description(); ok {
		_spec.SetField(deck.FieldDescription, field.TypeString, value)
	}
	if du.mutation.SubscribersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedSubscribersIDs(); len(nodes) > 0 && !du.mutation.SubscribersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SubscribersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !du.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SubjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.UsersProgressCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedUsersProgressIDs(); len(nodes) > 0 && !du.mutation.UsersProgressCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UsersProgressIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deck.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DeckUpdateOne is the builder for updating a single Deck entity.
type DeckUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeckMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DeckUpdateOne) SetUpdatedAt(t time.Time) *DeckUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetName sets the "name" field.
func (duo *DeckUpdateOne) SetName(s string) *DeckUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetDescription sets the "description" field.
func (duo *DeckUpdateOne) SetDescription(s string) *DeckUpdateOne {
	duo.mutation.SetDescription(s)
	return duo
}

// AddSubscriberIDs adds the "subscribers" edge to the User entity by IDs.
func (duo *DeckUpdateOne) AddSubscriberIDs(ids ...string) *DeckUpdateOne {
	duo.mutation.AddSubscriberIDs(ids...)
	return duo
}

// AddSubscribers adds the "subscribers" edges to the User entity.
func (duo *DeckUpdateOne) AddSubscribers(u ...*User) *DeckUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddSubscriberIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (duo *DeckUpdateOne) SetOwnerID(id string) *DeckUpdateOne {
	duo.mutation.SetOwnerID(id)
	return duo
}

// SetOwner sets the "owner" edge to the User entity.
func (duo *DeckUpdateOne) SetOwner(u *User) *DeckUpdateOne {
	return duo.SetOwnerID(u.ID)
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (duo *DeckUpdateOne) AddSubjectIDs(ids ...uuid.UUID) *DeckUpdateOne {
	duo.mutation.AddSubjectIDs(ids...)
	return duo
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (duo *DeckUpdateOne) AddSubjects(s ...*Subject) *DeckUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddSubjectIDs(ids...)
}

// AddUsersProgresIDs adds the "users_progress" edge to the DeckProgress entity by IDs.
func (duo *DeckUpdateOne) AddUsersProgresIDs(ids ...int) *DeckUpdateOne {
	duo.mutation.AddUsersProgresIDs(ids...)
	return duo
}

// AddUsersProgress adds the "users_progress" edges to the DeckProgress entity.
func (duo *DeckUpdateOne) AddUsersProgress(d ...*DeckProgress) *DeckUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddUsersProgresIDs(ids...)
}

// Mutation returns the DeckMutation object of the builder.
func (duo *DeckUpdateOne) Mutation() *DeckMutation {
	return duo.mutation
}

// ClearSubscribers clears all "subscribers" edges to the User entity.
func (duo *DeckUpdateOne) ClearSubscribers() *DeckUpdateOne {
	duo.mutation.ClearSubscribers()
	return duo
}

// RemoveSubscriberIDs removes the "subscribers" edge to User entities by IDs.
func (duo *DeckUpdateOne) RemoveSubscriberIDs(ids ...string) *DeckUpdateOne {
	duo.mutation.RemoveSubscriberIDs(ids...)
	return duo
}

// RemoveSubscribers removes "subscribers" edges to User entities.
func (duo *DeckUpdateOne) RemoveSubscribers(u ...*User) *DeckUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveSubscriberIDs(ids...)
}

// ClearOwner clears the "owner" edge to the User entity.
func (duo *DeckUpdateOne) ClearOwner() *DeckUpdateOne {
	duo.mutation.ClearOwner()
	return duo
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (duo *DeckUpdateOne) ClearSubjects() *DeckUpdateOne {
	duo.mutation.ClearSubjects()
	return duo
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (duo *DeckUpdateOne) RemoveSubjectIDs(ids ...uuid.UUID) *DeckUpdateOne {
	duo.mutation.RemoveSubjectIDs(ids...)
	return duo
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (duo *DeckUpdateOne) RemoveSubjects(s ...*Subject) *DeckUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveSubjectIDs(ids...)
}

// ClearUsersProgress clears all "users_progress" edges to the DeckProgress entity.
func (duo *DeckUpdateOne) ClearUsersProgress() *DeckUpdateOne {
	duo.mutation.ClearUsersProgress()
	return duo
}

// RemoveUsersProgresIDs removes the "users_progress" edge to DeckProgress entities by IDs.
func (duo *DeckUpdateOne) RemoveUsersProgresIDs(ids ...int) *DeckUpdateOne {
	duo.mutation.RemoveUsersProgresIDs(ids...)
	return duo
}

// RemoveUsersProgress removes "users_progress" edges to DeckProgress entities.
func (duo *DeckUpdateOne) RemoveUsersProgress(d ...*DeckProgress) *DeckUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveUsersProgresIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeckUpdateOne) Select(field string, fields ...string) *DeckUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Deck entity.
func (duo *DeckUpdateOne) Save(ctx context.Context) (*Deck, error) {
	var (
		err  error
		node *Deck
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeckMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (duo *DeckUpdateOne) SaveX(ctx context.Context) *Deck {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeckUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeckUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeckUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := deck.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeckUpdateOne) check() error {
	if v, ok := duo.mutation.Name(); ok {
		if err := deck.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Deck.name": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Description(); ok {
		if err := deck.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Deck.description": %w`, err)}
		}
	}
	if _, ok := duo.mutation.OwnerID(); duo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deck.owner"`)
	}
	return nil
}

func (duo *DeckUpdateOne) sqlSave(ctx context.Context) (_node *Deck, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deck.Table,
			Columns: deck.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deck.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Deck.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deck.FieldID)
		for _, f := range fields {
			if !deck.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deck.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(deck.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(deck.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Description(); ok {
		_spec.SetField(deck.FieldDescription, field.TypeString, value)
	}
	if duo.mutation.SubscribersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedSubscribersIDs(); len(nodes) > 0 && !duo.mutation.SubscribersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SubscribersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !duo.mutation.SubjectsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SubjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.UsersProgressCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedUsersProgressIDs(); len(nodes) > 0 && !duo.mutation.UsersProgressCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UsersProgressIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Deck{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deck.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
