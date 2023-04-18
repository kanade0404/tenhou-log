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
	"github.com/kanade0404/tenhou-log/services/ent/concealedkan"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ConcealedKanQuery is the builder for querying ConcealedKan entities.
type ConcealedKanQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.ConcealedKan
	withCall   *CallQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConcealedKanQuery builder.
func (ckq *ConcealedKanQuery) Where(ps ...predicate.ConcealedKan) *ConcealedKanQuery {
	ckq.predicates = append(ckq.predicates, ps...)
	return ckq
}

// Limit the number of records to be returned by this query.
func (ckq *ConcealedKanQuery) Limit(limit int) *ConcealedKanQuery {
	ckq.ctx.Limit = &limit
	return ckq
}

// Offset to start from.
func (ckq *ConcealedKanQuery) Offset(offset int) *ConcealedKanQuery {
	ckq.ctx.Offset = &offset
	return ckq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ckq *ConcealedKanQuery) Unique(unique bool) *ConcealedKanQuery {
	ckq.ctx.Unique = &unique
	return ckq
}

// Order specifies how the records should be ordered.
func (ckq *ConcealedKanQuery) Order(o ...OrderFunc) *ConcealedKanQuery {
	ckq.order = append(ckq.order, o...)
	return ckq
}

// QueryCall chains the current query on the "call" edge.
func (ckq *ConcealedKanQuery) QueryCall() *CallQuery {
	query := (&CallClient{config: ckq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ckq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ckq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(concealedkan.Table, concealedkan.FieldID, selector),
			sqlgraph.To(call.Table, call.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, concealedkan.CallTable, concealedkan.CallColumn),
		)
		fromU = sqlgraph.SetNeighbors(ckq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ConcealedKan entity from the query.
// Returns a *NotFoundError when no ConcealedKan was found.
func (ckq *ConcealedKanQuery) First(ctx context.Context) (*ConcealedKan, error) {
	nodes, err := ckq.Limit(1).All(setContextOp(ctx, ckq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{concealedkan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ckq *ConcealedKanQuery) FirstX(ctx context.Context) *ConcealedKan {
	node, err := ckq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ConcealedKan ID from the query.
// Returns a *NotFoundError when no ConcealedKan ID was found.
func (ckq *ConcealedKanQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ckq.Limit(1).IDs(setContextOp(ctx, ckq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{concealedkan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ckq *ConcealedKanQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ckq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ConcealedKan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ConcealedKan entity is found.
// Returns a *NotFoundError when no ConcealedKan entities are found.
func (ckq *ConcealedKanQuery) Only(ctx context.Context) (*ConcealedKan, error) {
	nodes, err := ckq.Limit(2).All(setContextOp(ctx, ckq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{concealedkan.Label}
	default:
		return nil, &NotSingularError{concealedkan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ckq *ConcealedKanQuery) OnlyX(ctx context.Context) *ConcealedKan {
	node, err := ckq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ConcealedKan ID in the query.
// Returns a *NotSingularError when more than one ConcealedKan ID is found.
// Returns a *NotFoundError when no entities are found.
func (ckq *ConcealedKanQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ckq.Limit(2).IDs(setContextOp(ctx, ckq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{concealedkan.Label}
	default:
		err = &NotSingularError{concealedkan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ckq *ConcealedKanQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ckq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ConcealedKans.
func (ckq *ConcealedKanQuery) All(ctx context.Context) ([]*ConcealedKan, error) {
	ctx = setContextOp(ctx, ckq.ctx, "All")
	if err := ckq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ConcealedKan, *ConcealedKanQuery]()
	return withInterceptors[[]*ConcealedKan](ctx, ckq, qr, ckq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ckq *ConcealedKanQuery) AllX(ctx context.Context) []*ConcealedKan {
	nodes, err := ckq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ConcealedKan IDs.
func (ckq *ConcealedKanQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ckq.ctx.Unique == nil && ckq.path != nil {
		ckq.Unique(true)
	}
	ctx = setContextOp(ctx, ckq.ctx, "IDs")
	if err = ckq.Select(concealedkan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ckq *ConcealedKanQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ckq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ckq *ConcealedKanQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ckq.ctx, "Count")
	if err := ckq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ckq, querierCount[*ConcealedKanQuery](), ckq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ckq *ConcealedKanQuery) CountX(ctx context.Context) int {
	count, err := ckq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ckq *ConcealedKanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ckq.ctx, "Exist")
	switch _, err := ckq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ckq *ConcealedKanQuery) ExistX(ctx context.Context) bool {
	exist, err := ckq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConcealedKanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ckq *ConcealedKanQuery) Clone() *ConcealedKanQuery {
	if ckq == nil {
		return nil
	}
	return &ConcealedKanQuery{
		config:     ckq.config,
		ctx:        ckq.ctx.Clone(),
		order:      append([]OrderFunc{}, ckq.order...),
		inters:     append([]Interceptor{}, ckq.inters...),
		predicates: append([]predicate.ConcealedKan{}, ckq.predicates...),
		withCall:   ckq.withCall.Clone(),
		// clone intermediate query.
		sql:  ckq.sql.Clone(),
		path: ckq.path,
	}
}

// WithCall tells the query-builder to eager-load the nodes that are connected to
// the "call" edge. The optional arguments are used to configure the query builder of the edge.
func (ckq *ConcealedKanQuery) WithCall(opts ...func(*CallQuery)) *ConcealedKanQuery {
	query := (&CallClient{config: ckq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ckq.withCall = query
	return ckq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (ckq *ConcealedKanQuery) GroupBy(field string, fields ...string) *ConcealedKanGroupBy {
	ckq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ConcealedKanGroupBy{build: ckq}
	grbuild.flds = &ckq.ctx.Fields
	grbuild.label = concealedkan.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ckq *ConcealedKanQuery) Select(fields ...string) *ConcealedKanSelect {
	ckq.ctx.Fields = append(ckq.ctx.Fields, fields...)
	sbuild := &ConcealedKanSelect{ConcealedKanQuery: ckq}
	sbuild.label = concealedkan.Label
	sbuild.flds, sbuild.scan = &ckq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConcealedKanSelect configured with the given aggregations.
func (ckq *ConcealedKanQuery) Aggregate(fns ...AggregateFunc) *ConcealedKanSelect {
	return ckq.Select().Aggregate(fns...)
}

func (ckq *ConcealedKanQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ckq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ckq); err != nil {
				return err
			}
		}
	}
	for _, f := range ckq.ctx.Fields {
		if !concealedkan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ckq.path != nil {
		prev, err := ckq.path(ctx)
		if err != nil {
			return err
		}
		ckq.sql = prev
	}
	return nil
}

func (ckq *ConcealedKanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ConcealedKan, error) {
	var (
		nodes       = []*ConcealedKan{}
		_spec       = ckq.querySpec()
		loadedTypes = [1]bool{
			ckq.withCall != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ConcealedKan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ConcealedKan{config: ckq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ckq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ckq.withCall; query != nil {
		if err := ckq.loadCall(ctx, query, nodes, nil,
			func(n *ConcealedKan, e *Call) { n.Edges.Call = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ckq *ConcealedKanQuery) loadCall(ctx context.Context, query *CallQuery, nodes []*ConcealedKan, init func(*ConcealedKan), assign func(*ConcealedKan, *Call)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ConcealedKan)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Call(func(s *sql.Selector) {
		s.Where(sql.InValues(concealedkan.CallColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.concealed_kan_call
		if fk == nil {
			return fmt.Errorf(`foreign-key "concealed_kan_call" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "concealed_kan_call" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ckq *ConcealedKanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ckq.querySpec()
	_spec.Node.Columns = ckq.ctx.Fields
	if len(ckq.ctx.Fields) > 0 {
		_spec.Unique = ckq.ctx.Unique != nil && *ckq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ckq.driver, _spec)
}

func (ckq *ConcealedKanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(concealedkan.Table, concealedkan.Columns, sqlgraph.NewFieldSpec(concealedkan.FieldID, field.TypeUUID))
	_spec.From = ckq.sql
	if unique := ckq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ckq.path != nil {
		_spec.Unique = true
	}
	if fields := ckq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, concealedkan.FieldID)
		for i := range fields {
			if fields[i] != concealedkan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ckq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ckq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ckq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ckq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ckq *ConcealedKanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ckq.driver.Dialect())
	t1 := builder.Table(concealedkan.Table)
	columns := ckq.ctx.Fields
	if len(columns) == 0 {
		columns = concealedkan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ckq.sql != nil {
		selector = ckq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ckq.ctx.Unique != nil && *ckq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ckq.predicates {
		p(selector)
	}
	for _, p := range ckq.order {
		p(selector)
	}
	if offset := ckq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ckq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConcealedKanGroupBy is the group-by builder for ConcealedKan entities.
type ConcealedKanGroupBy struct {
	selector
	build *ConcealedKanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ckgb *ConcealedKanGroupBy) Aggregate(fns ...AggregateFunc) *ConcealedKanGroupBy {
	ckgb.fns = append(ckgb.fns, fns...)
	return ckgb
}

// Scan applies the selector query and scans the result into the given value.
func (ckgb *ConcealedKanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ckgb.build.ctx, "GroupBy")
	if err := ckgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConcealedKanQuery, *ConcealedKanGroupBy](ctx, ckgb.build, ckgb, ckgb.build.inters, v)
}

func (ckgb *ConcealedKanGroupBy) sqlScan(ctx context.Context, root *ConcealedKanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ckgb.fns))
	for _, fn := range ckgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ckgb.flds)+len(ckgb.fns))
		for _, f := range *ckgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ckgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ckgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConcealedKanSelect is the builder for selecting fields of ConcealedKan entities.
type ConcealedKanSelect struct {
	*ConcealedKanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cks *ConcealedKanSelect) Aggregate(fns ...AggregateFunc) *ConcealedKanSelect {
	cks.fns = append(cks.fns, fns...)
	return cks
}

// Scan applies the selector query and scans the result into the given value.
func (cks *ConcealedKanSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cks.ctx, "Select")
	if err := cks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConcealedKanQuery, *ConcealedKanSelect](ctx, cks.ConcealedKanQuery, cks, cks.inters, v)
}

func (cks *ConcealedKanSelect) sqlScan(ctx context.Context, root *ConcealedKanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cks.fns))
	for _, fn := range cks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
