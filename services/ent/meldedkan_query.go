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
	"github.com/kanade0404/tenhou-log/services/ent/call"
	"github.com/kanade0404/tenhou-log/services/ent/meldedkan"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// MeldedKanQuery is the builder for querying MeldedKan entities.
type MeldedKanQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.MeldedKan
	withCall   *CallQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MeldedKanQuery builder.
func (mkq *MeldedKanQuery) Where(ps ...predicate.MeldedKan) *MeldedKanQuery {
	mkq.predicates = append(mkq.predicates, ps...)
	return mkq
}

// Limit the number of records to be returned by this query.
func (mkq *MeldedKanQuery) Limit(limit int) *MeldedKanQuery {
	mkq.limit = &limit
	return mkq
}

// Offset to start from.
func (mkq *MeldedKanQuery) Offset(offset int) *MeldedKanQuery {
	mkq.offset = &offset
	return mkq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mkq *MeldedKanQuery) Unique(unique bool) *MeldedKanQuery {
	mkq.unique = &unique
	return mkq
}

// Order specifies how the records should be ordered.
func (mkq *MeldedKanQuery) Order(o ...OrderFunc) *MeldedKanQuery {
	mkq.order = append(mkq.order, o...)
	return mkq
}

// QueryCall chains the current query on the "call" edge.
func (mkq *MeldedKanQuery) QueryCall() *CallQuery {
	query := (&CallClient{config: mkq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mkq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mkq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(meldedkan.Table, meldedkan.FieldID, selector),
			sqlgraph.To(call.Table, call.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, meldedkan.CallTable, meldedkan.CallColumn),
		)
		fromU = sqlgraph.SetNeighbors(mkq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MeldedKan entity from the query.
// Returns a *NotFoundError when no MeldedKan was found.
func (mkq *MeldedKanQuery) First(ctx context.Context) (*MeldedKan, error) {
	nodes, err := mkq.Limit(1).All(newQueryContext(ctx, TypeMeldedKan, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{meldedkan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mkq *MeldedKanQuery) FirstX(ctx context.Context) *MeldedKan {
	node, err := mkq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MeldedKan ID from the query.
// Returns a *NotFoundError when no MeldedKan ID was found.
func (mkq *MeldedKanQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mkq.Limit(1).IDs(newQueryContext(ctx, TypeMeldedKan, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{meldedkan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mkq *MeldedKanQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mkq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MeldedKan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MeldedKan entity is found.
// Returns a *NotFoundError when no MeldedKan entities are found.
func (mkq *MeldedKanQuery) Only(ctx context.Context) (*MeldedKan, error) {
	nodes, err := mkq.Limit(2).All(newQueryContext(ctx, TypeMeldedKan, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{meldedkan.Label}
	default:
		return nil, &NotSingularError{meldedkan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mkq *MeldedKanQuery) OnlyX(ctx context.Context) *MeldedKan {
	node, err := mkq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MeldedKan ID in the query.
// Returns a *NotSingularError when more than one MeldedKan ID is found.
// Returns a *NotFoundError when no entities are found.
func (mkq *MeldedKanQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mkq.Limit(2).IDs(newQueryContext(ctx, TypeMeldedKan, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{meldedkan.Label}
	default:
		err = &NotSingularError{meldedkan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mkq *MeldedKanQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mkq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MeldedKans.
func (mkq *MeldedKanQuery) All(ctx context.Context) ([]*MeldedKan, error) {
	ctx = newQueryContext(ctx, TypeMeldedKan, "All")
	if err := mkq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MeldedKan, *MeldedKanQuery]()
	return withInterceptors[[]*MeldedKan](ctx, mkq, qr, mkq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mkq *MeldedKanQuery) AllX(ctx context.Context) []*MeldedKan {
	nodes, err := mkq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MeldedKan IDs.
func (mkq *MeldedKanQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeMeldedKan, "IDs")
	if err := mkq.Select(meldedkan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mkq *MeldedKanQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mkq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mkq *MeldedKanQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeMeldedKan, "Count")
	if err := mkq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mkq, querierCount[*MeldedKanQuery](), mkq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mkq *MeldedKanQuery) CountX(ctx context.Context) int {
	count, err := mkq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mkq *MeldedKanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeMeldedKan, "Exist")
	switch _, err := mkq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mkq *MeldedKanQuery) ExistX(ctx context.Context) bool {
	exist, err := mkq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MeldedKanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mkq *MeldedKanQuery) Clone() *MeldedKanQuery {
	if mkq == nil {
		return nil
	}
	return &MeldedKanQuery{
		config:     mkq.config,
		limit:      mkq.limit,
		offset:     mkq.offset,
		order:      append([]OrderFunc{}, mkq.order...),
		inters:     append([]Interceptor{}, mkq.inters...),
		predicates: append([]predicate.MeldedKan{}, mkq.predicates...),
		withCall:   mkq.withCall.Clone(),
		// clone intermediate query.
		sql:    mkq.sql.Clone(),
		path:   mkq.path,
		unique: mkq.unique,
	}
}

// WithCall tells the query-builder to eager-load the nodes that are connected to
// the "call" edge. The optional arguments are used to configure the query builder of the edge.
func (mkq *MeldedKanQuery) WithCall(opts ...func(*CallQuery)) *MeldedKanQuery {
	query := (&CallClient{config: mkq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mkq.withCall = query
	return mkq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (mkq *MeldedKanQuery) GroupBy(field string, fields ...string) *MeldedKanGroupBy {
	mkq.fields = append([]string{field}, fields...)
	grbuild := &MeldedKanGroupBy{build: mkq}
	grbuild.flds = &mkq.fields
	grbuild.label = meldedkan.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (mkq *MeldedKanQuery) Select(fields ...string) *MeldedKanSelect {
	mkq.fields = append(mkq.fields, fields...)
	sbuild := &MeldedKanSelect{MeldedKanQuery: mkq}
	sbuild.label = meldedkan.Label
	sbuild.flds, sbuild.scan = &mkq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MeldedKanSelect configured with the given aggregations.
func (mkq *MeldedKanQuery) Aggregate(fns ...AggregateFunc) *MeldedKanSelect {
	return mkq.Select().Aggregate(fns...)
}

func (mkq *MeldedKanQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mkq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mkq); err != nil {
				return err
			}
		}
	}
	for _, f := range mkq.fields {
		if !meldedkan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mkq.path != nil {
		prev, err := mkq.path(ctx)
		if err != nil {
			return err
		}
		mkq.sql = prev
	}
	return nil
}

func (mkq *MeldedKanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MeldedKan, error) {
	var (
		nodes       = []*MeldedKan{}
		_spec       = mkq.querySpec()
		loadedTypes = [1]bool{
			mkq.withCall != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MeldedKan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MeldedKan{config: mkq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mkq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mkq.withCall; query != nil {
		if err := mkq.loadCall(ctx, query, nodes, nil,
			func(n *MeldedKan, e *Call) { n.Edges.Call = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mkq *MeldedKanQuery) loadCall(ctx context.Context, query *CallQuery, nodes []*MeldedKan, init func(*MeldedKan), assign func(*MeldedKan, *Call)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*MeldedKan)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Call(func(s *sql.Selector) {
		s.Where(sql.InValues(meldedkan.CallColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.melded_kan_call
		if fk == nil {
			return fmt.Errorf(`foreign-key "melded_kan_call" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "melded_kan_call" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mkq *MeldedKanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mkq.querySpec()
	_spec.Node.Columns = mkq.fields
	if len(mkq.fields) > 0 {
		_spec.Unique = mkq.unique != nil && *mkq.unique
	}
	return sqlgraph.CountNodes(ctx, mkq.driver, _spec)
}

func (mkq *MeldedKanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   meldedkan.Table,
			Columns: meldedkan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: meldedkan.FieldID,
			},
		},
		From:   mkq.sql,
		Unique: true,
	}
	if unique := mkq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mkq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meldedkan.FieldID)
		for i := range fields {
			if fields[i] != meldedkan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mkq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mkq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mkq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mkq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mkq *MeldedKanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mkq.driver.Dialect())
	t1 := builder.Table(meldedkan.Table)
	columns := mkq.fields
	if len(columns) == 0 {
		columns = meldedkan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mkq.sql != nil {
		selector = mkq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mkq.unique != nil && *mkq.unique {
		selector.Distinct()
	}
	for _, p := range mkq.predicates {
		p(selector)
	}
	for _, p := range mkq.order {
		p(selector)
	}
	if offset := mkq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mkq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MeldedKanGroupBy is the group-by builder for MeldedKan entities.
type MeldedKanGroupBy struct {
	selector
	build *MeldedKanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mkgb *MeldedKanGroupBy) Aggregate(fns ...AggregateFunc) *MeldedKanGroupBy {
	mkgb.fns = append(mkgb.fns, fns...)
	return mkgb
}

// Scan applies the selector query and scans the result into the given value.
func (mkgb *MeldedKanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeMeldedKan, "GroupBy")
	if err := mkgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MeldedKanQuery, *MeldedKanGroupBy](ctx, mkgb.build, mkgb, mkgb.build.inters, v)
}

func (mkgb *MeldedKanGroupBy) sqlScan(ctx context.Context, root *MeldedKanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mkgb.fns))
	for _, fn := range mkgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mkgb.flds)+len(mkgb.fns))
		for _, f := range *mkgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mkgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mkgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MeldedKanSelect is the builder for selecting fields of MeldedKan entities.
type MeldedKanSelect struct {
	*MeldedKanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mks *MeldedKanSelect) Aggregate(fns ...AggregateFunc) *MeldedKanSelect {
	mks.fns = append(mks.fns, fns...)
	return mks
}

// Scan applies the selector query and scans the result into the given value.
func (mks *MeldedKanSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeMeldedKan, "Select")
	if err := mks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MeldedKanQuery, *MeldedKanSelect](ctx, mks.MeldedKanQuery, mks, mks.inters, v)
}

func (mks *MeldedKanSelect) sqlScan(ctx context.Context, root *MeldedKanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mks.fns))
	for _, fn := range mks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
