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
	"github.com/kanade0404/tenhou-log/services/ent/discard"
	"github.com/kanade0404/tenhou-log/services/ent/drawn"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
	"github.com/kanade0404/tenhou-log/services/ent/reach"
)

// DiscardQuery is the builder for querying Discard entities.
type DiscardQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Discard
	withReach  *ReachQuery
	withCall   *CallQuery
	withDraw   *DrawnQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscardQuery builder.
func (dq *DiscardQuery) Where(ps ...predicate.Discard) *DiscardQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DiscardQuery) Limit(limit int) *DiscardQuery {
	dq.limit = &limit
	return dq
}

// Offset to start from.
func (dq *DiscardQuery) Offset(offset int) *DiscardQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DiscardQuery) Unique(unique bool) *DiscardQuery {
	dq.unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DiscardQuery) Order(o ...OrderFunc) *DiscardQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryReach chains the current query on the "reach" edge.
func (dq *DiscardQuery) QueryReach() *ReachQuery {
	query := (&ReachClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discard.Table, discard.FieldID, selector),
			sqlgraph.To(reach.Table, reach.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, discard.ReachTable, discard.ReachColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCall chains the current query on the "call" edge.
func (dq *DiscardQuery) QueryCall() *CallQuery {
	query := (&CallClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discard.Table, discard.FieldID, selector),
			sqlgraph.To(call.Table, call.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, discard.CallTable, discard.CallColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDraw chains the current query on the "draw" edge.
func (dq *DiscardQuery) QueryDraw() *DrawnQuery {
	query := (&DrawnClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discard.Table, discard.FieldID, selector),
			sqlgraph.To(drawn.Table, drawn.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, discard.DrawTable, discard.DrawColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Discard entity from the query.
// Returns a *NotFoundError when no Discard was found.
func (dq *DiscardQuery) First(ctx context.Context) (*Discard, error) {
	nodes, err := dq.Limit(1).All(newQueryContext(ctx, TypeDiscard, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discard.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DiscardQuery) FirstX(ctx context.Context) *Discard {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Discard ID from the query.
// Returns a *NotFoundError when no Discard ID was found.
func (dq *DiscardQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(newQueryContext(ctx, TypeDiscard, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discard.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DiscardQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Discard entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Discard entity is found.
// Returns a *NotFoundError when no Discard entities are found.
func (dq *DiscardQuery) Only(ctx context.Context) (*Discard, error) {
	nodes, err := dq.Limit(2).All(newQueryContext(ctx, TypeDiscard, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discard.Label}
	default:
		return nil, &NotSingularError{discard.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DiscardQuery) OnlyX(ctx context.Context) *Discard {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Discard ID in the query.
// Returns a *NotSingularError when more than one Discard ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DiscardQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(newQueryContext(ctx, TypeDiscard, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discard.Label}
	default:
		err = &NotSingularError{discard.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DiscardQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Discards.
func (dq *DiscardQuery) All(ctx context.Context) ([]*Discard, error) {
	ctx = newQueryContext(ctx, TypeDiscard, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Discard, *DiscardQuery]()
	return withInterceptors[[]*Discard](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DiscardQuery) AllX(ctx context.Context) []*Discard {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Discard IDs.
func (dq *DiscardQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeDiscard, "IDs")
	if err := dq.Select(discard.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DiscardQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DiscardQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeDiscard, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DiscardQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DiscardQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DiscardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeDiscard, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DiscardQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DiscardQuery) Clone() *DiscardQuery {
	if dq == nil {
		return nil
	}
	return &DiscardQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		inters:     append([]Interceptor{}, dq.inters...),
		predicates: append([]predicate.Discard{}, dq.predicates...),
		withReach:  dq.withReach.Clone(),
		withCall:   dq.withCall.Clone(),
		withDraw:   dq.withDraw.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithReach tells the query-builder to eager-load the nodes that are connected to
// the "reach" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DiscardQuery) WithReach(opts ...func(*ReachQuery)) *DiscardQuery {
	query := (&ReachClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withReach = query
	return dq
}

// WithCall tells the query-builder to eager-load the nodes that are connected to
// the "call" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DiscardQuery) WithCall(opts ...func(*CallQuery)) *DiscardQuery {
	query := (&CallClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withCall = query
	return dq
}

// WithDraw tells the query-builder to eager-load the nodes that are connected to
// the "draw" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DiscardQuery) WithDraw(opts ...func(*DrawnQuery)) *DiscardQuery {
	query := (&DrawnClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDraw = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (dq *DiscardQuery) GroupBy(field string, fields ...string) *DiscardGroupBy {
	dq.fields = append([]string{field}, fields...)
	grbuild := &DiscardGroupBy{build: dq}
	grbuild.flds = &dq.fields
	grbuild.label = discard.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (dq *DiscardQuery) Select(fields ...string) *DiscardSelect {
	dq.fields = append(dq.fields, fields...)
	sbuild := &DiscardSelect{DiscardQuery: dq}
	sbuild.label = discard.Label
	sbuild.flds, sbuild.scan = &dq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DiscardSelect configured with the given aggregations.
func (dq *DiscardQuery) Aggregate(fns ...AggregateFunc) *DiscardSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DiscardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.fields {
		if !discard.ValidColumn(f) {
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

func (dq *DiscardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Discard, error) {
	var (
		nodes       = []*Discard{}
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withReach != nil,
			dq.withCall != nil,
			dq.withDraw != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Discard).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Discard{config: dq.config}
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
	if query := dq.withReach; query != nil {
		if err := dq.loadReach(ctx, query, nodes, nil,
			func(n *Discard, e *Reach) { n.Edges.Reach = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withCall; query != nil {
		if err := dq.loadCall(ctx, query, nodes, nil,
			func(n *Discard, e *Call) { n.Edges.Call = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withDraw; query != nil {
		if err := dq.loadDraw(ctx, query, nodes, nil,
			func(n *Discard, e *Drawn) { n.Edges.Draw = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DiscardQuery) loadReach(ctx context.Context, query *ReachQuery, nodes []*Discard, init func(*Discard), assign func(*Discard, *Reach)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Discard)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Reach(func(s *sql.Selector) {
		s.Where(sql.InValues(discard.ReachColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.discard_reach
		if fk == nil {
			return fmt.Errorf(`foreign-key "discard_reach" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "discard_reach" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dq *DiscardQuery) loadCall(ctx context.Context, query *CallQuery, nodes []*Discard, init func(*Discard), assign func(*Discard, *Call)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Discard)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Call(func(s *sql.Selector) {
		s.Where(sql.InValues(discard.CallColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.discard_call
		if fk == nil {
			return fmt.Errorf(`foreign-key "discard_call" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "discard_call" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dq *DiscardQuery) loadDraw(ctx context.Context, query *DrawnQuery, nodes []*Discard, init func(*Discard), assign func(*Discard, *Drawn)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Discard)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Drawn(func(s *sql.Selector) {
		s.Where(sql.InValues(discard.DrawColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.discard_draw
		if fk == nil {
			return fmt.Errorf(`foreign-key "discard_draw" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "discard_draw" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DiscardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DiscardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discard.Table,
			Columns: discard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: discard.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, discard.FieldID)
		for i := range fields {
			if fields[i] != discard.FieldID {
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

func (dq *DiscardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(discard.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = discard.Columns
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

// DiscardGroupBy is the group-by builder for Discard entities.
type DiscardGroupBy struct {
	selector
	build *DiscardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DiscardGroupBy) Aggregate(fns ...AggregateFunc) *DiscardGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DiscardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeDiscard, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiscardQuery, *DiscardGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DiscardGroupBy) sqlScan(ctx context.Context, root *DiscardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DiscardSelect is the builder for selecting fields of Discard entities.
type DiscardSelect struct {
	*DiscardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DiscardSelect) Aggregate(fns ...AggregateFunc) *DiscardSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DiscardSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeDiscard, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiscardQuery, *DiscardSelect](ctx, ds.DiscardQuery, ds, ds.inters, v)
}

func (ds *DiscardSelect) sqlScan(ctx context.Context, root *DiscardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
