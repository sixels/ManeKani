// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"sixels.io/manekani/ent/deck"
	"sixels.io/manekani/ent/deckprogress"
	"sixels.io/manekani/ent/predicate"
	"sixels.io/manekani/ent/subject"
	"sixels.io/manekani/ent/user"
)

// DeckQuery is the builder for querying Deck entities.
type DeckQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.Deck
	withSubscribers   *UserQuery
	withOwner         *UserQuery
	withSubjects      *SubjectQuery
	withUsersProgress *DeckProgressQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeckQuery builder.
func (dq *DeckQuery) Where(ps ...predicate.Deck) *DeckQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DeckQuery) Limit(limit int) *DeckQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DeckQuery) Offset(offset int) *DeckQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeckQuery) Unique(unique bool) *DeckQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DeckQuery) Order(o ...OrderFunc) *DeckQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QuerySubscribers chains the current query on the "subscribers" edge.
func (dq *DeckQuery) QuerySubscribers() *UserQuery {
	query := &UserQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deck.Table, deck.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, deck.SubscribersTable, deck.SubscribersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (dq *DeckQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deck.Table, deck.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deck.OwnerTable, deck.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySubjects chains the current query on the "subjects" edge.
func (dq *DeckQuery) QuerySubjects() *SubjectQuery {
	query := &SubjectQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deck.Table, deck.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, deck.SubjectsTable, deck.SubjectsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsersProgress chains the current query on the "users_progress" edge.
func (dq *DeckQuery) QueryUsersProgress() *DeckProgressQuery {
	query := &DeckProgressQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deck.Table, deck.FieldID, selector),
			sqlgraph.To(deckprogress.Table, deckprogress.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deck.UsersProgressTable, deck.UsersProgressColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Deck entity from the query.
// Returns a *NotFoundError when no Deck was found.
func (dq *DeckQuery) First(ctx context.Context) (*Deck, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deck.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeckQuery) FirstX(ctx context.Context) *Deck {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Deck ID from the query.
// Returns a *NotFoundError when no Deck ID was found.
func (dq *DeckQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deck.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeckQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Deck entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Deck entity is found.
// Returns a *NotFoundError when no Deck entities are found.
func (dq *DeckQuery) Only(ctx context.Context) (*Deck, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deck.Label}
	default:
		return nil, &NotSingularError{deck.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeckQuery) OnlyX(ctx context.Context) *Deck {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Deck ID in the query.
// Returns a *NotSingularError when more than one Deck ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeckQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deck.Label}
	default:
		err = &NotSingularError{deck.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeckQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Decks.
func (dq *DeckQuery) All(ctx context.Context) ([]*Deck, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeckQuery) AllX(ctx context.Context) []*Deck {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Deck IDs.
func (dq *DeckQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dq.Select(deck.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeckQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeckQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeckQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeckQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeckQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeckQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeckQuery) Clone() *DeckQuery {
	if dq == nil {
		return nil
	}
	return &DeckQuery{
		config:            dq.config,
		limit:             dq.limit,
		offset:            dq.offset,
		order:             append([]OrderFunc{}, dq.order...),
		predicates:        append([]predicate.Deck{}, dq.predicates...),
		withSubscribers:   dq.withSubscribers.Clone(),
		withOwner:         dq.withOwner.Clone(),
		withSubjects:      dq.withSubjects.Clone(),
		withUsersProgress: dq.withUsersProgress.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithSubscribers tells the query-builder to eager-load the nodes that are connected to
// the "subscribers" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeckQuery) WithSubscribers(opts ...func(*UserQuery)) *DeckQuery {
	query := &UserQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withSubscribers = query
	return dq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeckQuery) WithOwner(opts ...func(*UserQuery)) *DeckQuery {
	query := &UserQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withOwner = query
	return dq
}

// WithSubjects tells the query-builder to eager-load the nodes that are connected to
// the "subjects" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeckQuery) WithSubjects(opts ...func(*SubjectQuery)) *DeckQuery {
	query := &SubjectQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withSubjects = query
	return dq
}

// WithUsersProgress tells the query-builder to eager-load the nodes that are connected to
// the "users_progress" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeckQuery) WithUsersProgress(opts ...func(*DeckProgressQuery)) *DeckQuery {
	query := &DeckProgressQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withUsersProgress = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Deck.Query().
//		GroupBy(deck.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DeckQuery) GroupBy(field string, fields ...string) *DeckGroupBy {
	grbuild := &DeckGroupBy{config: dq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = deck.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Deck.Query().
//		Select(deck.FieldCreatedAt).
//		Scan(ctx, &v)
func (dq *DeckQuery) Select(fields ...string) *DeckSelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DeckSelect{DeckQuery: dq}
	selbuild.label = deck.Label
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a DeckSelect configured with the given aggregations.
func (dq *DeckQuery) Aggregate(fns ...AggregateFunc) *DeckSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DeckQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !deck.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeckQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Deck, error) {
	var (
		nodes       = []*Deck{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [4]bool{
			dq.withSubscribers != nil,
			dq.withOwner != nil,
			dq.withSubjects != nil,
			dq.withUsersProgress != nil,
		}
	)
	if dq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, deck.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Deck).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Deck{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withSubscribers; query != nil {
		if err := dq.loadSubscribers(ctx, query, nodes,
			func(n *Deck) { n.Edges.Subscribers = []*User{} },
			func(n *Deck, e *User) { n.Edges.Subscribers = append(n.Edges.Subscribers, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withOwner; query != nil {
		if err := dq.loadOwner(ctx, query, nodes, nil,
			func(n *Deck, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withSubjects; query != nil {
		if err := dq.loadSubjects(ctx, query, nodes,
			func(n *Deck) { n.Edges.Subjects = []*Subject{} },
			func(n *Deck, e *Subject) { n.Edges.Subjects = append(n.Edges.Subjects, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withUsersProgress; query != nil {
		if err := dq.loadUsersProgress(ctx, query, nodes,
			func(n *Deck) { n.Edges.UsersProgress = []*DeckProgress{} },
			func(n *Deck, e *DeckProgress) { n.Edges.UsersProgress = append(n.Edges.UsersProgress, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DeckQuery) loadSubscribers(ctx context.Context, query *UserQuery, nodes []*Deck, init func(*Deck), assign func(*Deck, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Deck)
	nids := make(map[string]map[*Deck]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(deck.SubscribersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(deck.SubscribersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(deck.SubscribersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(deck.SubscribersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := values[1].(*sql.NullString).String
			if nids[inValue] == nil {
				nids[inValue] = map[*Deck]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "subscribers" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DeckQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Deck, init func(*Deck), assign func(*Deck, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Deck)
	for i := range nodes {
		if nodes[i].user_decks == nil {
			continue
		}
		fk := *nodes[i].user_decks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_decks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DeckQuery) loadSubjects(ctx context.Context, query *SubjectQuery, nodes []*Deck, init func(*Deck), assign func(*Deck, *Subject)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Deck)
	nids := make(map[uuid.UUID]map[*Deck]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(deck.SubjectsTable)
		s.Join(joinT).On(s.C(subject.FieldID), joinT.C(deck.SubjectsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(deck.SubjectsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(deck.SubjectsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(uuid.UUID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := *values[0].(*uuid.UUID)
			inValue := *values[1].(*uuid.UUID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Deck]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "subjects" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DeckQuery) loadUsersProgress(ctx context.Context, query *DeckProgressQuery, nodes []*Deck, init func(*Deck), assign func(*Deck, *DeckProgress)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Deck)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DeckProgress(func(s *sql.Selector) {
		s.Where(sql.InValues(deck.UsersProgressColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.deck_users_progress
		if fk == nil {
			return fmt.Errorf(`foreign-key "deck_users_progress" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "deck_users_progress" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DeckQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeckQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (dq *DeckQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deck.Table,
			Columns: deck.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deck.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deck.FieldID)
		for i := range fields {
			if fields[i] != deck.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeckQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(deck.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = deck.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeckGroupBy is the group-by builder for Deck entities.
type DeckGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeckGroupBy) Aggregate(fns ...AggregateFunc) *DeckGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DeckGroupBy) Scan(ctx context.Context, v any) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DeckGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range dgb.fields {
		if !deck.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DeckGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DeckSelect is the builder for selecting fields of Deck entities.
type DeckSelect struct {
	*DeckQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DeckSelect) Aggregate(fns ...AggregateFunc) *DeckSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeckSelect) Scan(ctx context.Context, v any) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DeckQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DeckSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(ds.sql))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ds.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ds.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
