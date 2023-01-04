



// Code generated by ent, DO NOT EDIT.



package ent



import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
		"sixels.io/manekani/services/cards/ent/predicate"
				"github.com/google/uuid"
			"github.com/jackc/pgtype"
			"github.com/jackc/pgtype"
			"github.com/jackc/pgtype"
			"github.com/google/uuid"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"

)


import (
		 "sixels.io/manekani/services/cards/ent/vocabulary"
		 "sixels.io/manekani/services/cards/ent/kanji"
)




// VocabularyQuery is the builder for querying Vocabulary entities.
type VocabularyQuery struct {
	config
	limit		*int
	offset		*int
	unique		*bool
	order		[]OrderFunc
	fields		[]string
	predicates 	[]predicate.Vocabulary
		withKanjis *KanjiQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VocabularyQuery builder.
func (vq *VocabularyQuery) Where(ps ...predicate.Vocabulary) *VocabularyQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit adds a limit step to the query.
func (vq *VocabularyQuery) Limit(limit int) *VocabularyQuery {
	vq.limit = &limit
	return vq
}

// Offset adds an offset step to the query.
func (vq *VocabularyQuery) Offset(offset int) *VocabularyQuery {
	vq.offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VocabularyQuery) Unique(unique bool) *VocabularyQuery {
	vq.unique = &unique
	return vq
}

// Order adds an order step to the query.
func (vq *VocabularyQuery) Order(o ...OrderFunc) *VocabularyQuery {
	vq.order = append(vq.order, o...)
	return vq
}



	
	// QueryKanjis chains the current query on the "kanjis" edge.
	func (vq *VocabularyQuery) QueryKanjis() *KanjiQuery {
		query := &KanjiQuery{config: vq.config}
		query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
			if err := vq.prepareQuery(ctx); err != nil {
				return nil, err
			}  
	selector := vq.sqlQuery(ctx)
	if err := selector.Err(); err != nil {
		return nil, err
	}
	step := sqlgraph.NewStep(
		sqlgraph.From(vocabulary.Table, vocabulary.FieldID, selector),
		sqlgraph.To(kanji.Table, kanji.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, vocabulary.KanjisTable,vocabulary.KanjisPrimaryKey...),
	)
	fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
return fromU, nil
		}
		return query
	}


// First returns the first Vocabulary entity from the query. 
// Returns a *NotFoundError when no Vocabulary was found.
func (vq *VocabularyQuery) First(ctx context.Context) (*Vocabulary, error) {
	nodes, err := vq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ vocabulary.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VocabularyQuery) FirstX(ctx context.Context) *Vocabulary {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}


	// FirstID returns the first Vocabulary ID from the query.
	// Returns a *NotFoundError when no Vocabulary ID was found.
	func (vq *VocabularyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
		var ids []uuid.UUID
		if ids, err = vq.Limit(1).IDs(ctx); err != nil {
			return
		}
		if len(ids) == 0 {
			err = &NotFoundError{ vocabulary.Label}
			return
		}
		return ids[0], nil
	}
	
	// FirstIDX is like FirstID, but panics if an error occurs.
	func (vq *VocabularyQuery) FirstIDX(ctx context.Context) uuid.UUID {
		id, err := vq.FirstID(ctx)
		if err != nil && !IsNotFound(err) {
			panic(err)
		}
		return id
	}


// Only returns a single Vocabulary entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vocabulary entity is found.
// Returns a *NotFoundError when no Vocabulary entities are found.
func (vq *VocabularyQuery) Only(ctx context.Context) (*Vocabulary, error) {
	nodes, err := vq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ vocabulary.Label}
	default:
		return nil, &NotSingularError{ vocabulary.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VocabularyQuery) OnlyX(ctx context.Context) *Vocabulary {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}


	// OnlyID is like Only, but returns the only Vocabulary ID in the query.
	// Returns a *NotSingularError when more than one Vocabulary ID is found.
	// Returns a *NotFoundError when no entities are found.
	func (vq *VocabularyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
		var ids []uuid.UUID
		if ids, err = vq.Limit(2).IDs(ctx); err != nil {
			return
		}
		switch len(ids) {
		case 1:
			id = ids[0]
		case 0:
			err = &NotFoundError{ vocabulary.Label}
		default:
			err = &NotSingularError{ vocabulary.Label}
		}
		return
	}
	
	// OnlyIDX is like OnlyID, but panics if an error occurs.
	func (vq *VocabularyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
		id, err := vq.OnlyID(ctx)
		if err != nil {
			panic(err)
		}
		return id
	}


// All executes the query and returns a list of Vocabularies.
func (vq *VocabularyQuery) All(ctx context.Context) ([]*Vocabulary, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vq *VocabularyQuery) AllX(ctx context.Context) []*Vocabulary {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}


	// IDs executes the query and returns a list of Vocabulary IDs.
	func (vq *VocabularyQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
		var ids []uuid.UUID
		if err := vq.Select(vocabulary.FieldID).Scan(ctx, &ids); err != nil {
			return nil, err
		}
		return ids, nil
	}
	
	// IDsX is like IDs, but panics if an error occurs.
	func (vq *VocabularyQuery) IDsX(ctx context.Context) []uuid.UUID {
		ids, err := vq.IDs(ctx)
		if err != nil {
			panic(err)
		}
		return ids
	}


// Count returns the count of the given query.
func (vq *VocabularyQuery) Count(ctx context.Context) (int, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VocabularyQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VocabularyQuery) Exist(ctx context.Context) (bool, error) {
	if err := vq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VocabularyQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VocabularyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VocabularyQuery) Clone() *VocabularyQuery {
	if vq == nil {
		return nil
	}
	return &VocabularyQuery{
		config: 	vq.config,
		limit: 		vq.limit,
		offset: 	vq.offset,
		order: 		append([]OrderFunc{}, vq.order...),
		predicates: append([]predicate.Vocabulary{}, vq.predicates...),
			withKanjis: vq.withKanjis.Clone(),
		// clone intermediate query.
		sql: vq.sql.Clone(),
		path: vq.path,
		unique: vq.unique,
	}
}
	
	
	// WithKanjis tells the query-builder to eager-load the nodes that are connected to
	// the "kanjis" edge. The optional arguments are used to configure the query builder of the edge.
	func (vq *VocabularyQuery) WithKanjis(opts ...func(*KanjiQuery)) *VocabularyQuery {
		query := &KanjiQuery{config: vq.config}
		for _, opt := range opts {
			opt(query)
		}
		vq.withKanjis = query
		return vq
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
//	client.Vocabulary.Query().
//		GroupBy(vocabulary.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *VocabularyQuery) GroupBy(field string, fields ...string) *VocabularyGroupBy {
	grbuild := &VocabularyGroupBy{config: vq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vq.sqlQuery(ctx), nil
	}
	grbuild.label = vocabulary.Label
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
//	client.Vocabulary.Query().
//		Select(vocabulary.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (vq *VocabularyQuery) Select(fields ...string) *VocabularySelect {
	vq.fields = append(vq.fields, fields...)
	selbuild := &VocabularySelect{ VocabularyQuery: vq }
	selbuild.label = vocabulary.Label
	selbuild.flds, selbuild.scan = &vq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a VocabularySelect configured with the given aggregations.
func (vq *VocabularyQuery) Aggregate(fns ...AggregateFunc) *VocabularySelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VocabularyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vq.fields {
		if !vocabulary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}


	
	




func (vq *VocabularyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vocabulary, error) {
	var (
		nodes = []*Vocabulary{}
		_spec = vq.querySpec()
			loadedTypes = [1]bool{
					vq.withKanjis != nil,
			}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Vocabulary).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Vocabulary{config: vq.config}
		nodes = append(nodes, node)
			node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
		if query := vq.withKanjis; query != nil {
			if err := vq.loadKanjis(ctx, query, nodes, 
				func(n *Vocabulary){ n.Edges.Kanjis = []*Kanji{} },
				func(n *Vocabulary, e *Kanji){ n.Edges.Kanjis = append(n.Edges.Kanjis, e) }); err != nil {
				return nil, err
			}
		}
	return nodes, nil
}


	func (vq *VocabularyQuery) loadKanjis(ctx context.Context, query *KanjiQuery, nodes []*Vocabulary, init func(*Vocabulary), assign func(*Vocabulary, *Kanji)) error {
			edgeIDs := make([]driver.Value, len(nodes))
			byID := make(map[uuid.UUID]*Vocabulary)
			nids := make(map[uuid.UUID]map[*Vocabulary]struct{})
			for i, node := range nodes {
				edgeIDs[i] = node.ID
				byID[node.ID] = node
				if init != nil {
					init(node)
				}
			}
			query.Where(func(s *sql.Selector) {
				joinT := sql.Table(vocabulary.KanjisTable)
				s.Join(joinT).On(s.C(kanji.FieldID), joinT.C(vocabulary.KanjisPrimaryKey[1]))
				s.Where(sql.InValues(joinT.C(vocabulary.KanjisPrimaryKey[0]), edgeIDs...))
				columns := s.SelectedColumns()
				s.Select(joinT.C(vocabulary.KanjisPrimaryKey[0]))
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
						nids[inValue] = map[*Vocabulary]struct{}{byID[outValue]: {}}
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
					return fmt.Errorf(`unexpected "kanjis" node returned %v`, n.ID)
				}
				for kn := range nodes {
					assign(kn, n)
				}
			}
		return nil
	}

func (vq *VocabularyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
		_spec.Node.Columns = vq.fields
		if len(vq.fields) > 0 {
			_spec.Unique = vq.unique != nil && *vq.unique
		}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VocabularyQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := vq.FirstID(ctx);{
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (vq *VocabularyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table: vocabulary.Table,
			Columns: vocabulary.Columns,
				ID: &sqlgraph.FieldSpec{
					Type: field.TypeUUID,
					Column: vocabulary.FieldID,
				},
		},
		From: vq.sql,
		Unique: true,
	}
	if unique := vq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
			_spec.Node.Columns = append(_spec.Node.Columns, vocabulary.FieldID)
			for i := range fields {
				if fields[i] != vocabulary.FieldID {
					_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
				}
			}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}





func (vq *VocabularyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vocabulary.Table)
	columns := vq.fields
	if len(columns) == 0 {
		columns = vocabulary.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.unique != nil && *vq.unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

    

    











// VocabularyGroupBy is the group-by builder for Vocabulary entities.
type VocabularyGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VocabularyGroupBy) Aggregate(fns ...AggregateFunc) *VocabularyGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vgb *VocabularyGroupBy) Scan(ctx context.Context, v any) error {
	query, err := vgb.path(ctx)
	if err != nil {
		return err
	}
	vgb.sql = query
	return vgb.sqlScan(ctx, v)
}


	
	



func (vgb *VocabularyGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range vgb.fields {
		if !vocabulary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}


func (vgb *VocabularyGroupBy) sqlQuery() *sql.Selector {
	selector := vgb.sql.Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vgb.fields) + len(vgb.fns))
		for _, f := range vgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vgb.fields...)...)
}







// VocabularySelect is the builder for selecting fields of Vocabulary entities.
type VocabularySelect struct {
	*VocabularyQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VocabularySelect) Aggregate(fns ...AggregateFunc) *VocabularySelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}


// Scan applies the selector query and scans the result into the given value.
func (vs *VocabularySelect) Scan(ctx context.Context, v any) error {
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	vs.sql = vs.VocabularyQuery.sqlQuery(ctx)
	return vs.sqlScan(ctx, v)
}


	
	



func (vs *VocabularySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(vs.sql))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		vs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		vs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := vs.sql.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}



    






