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
	"github.com/kanade0404/tenhou-log/services/ent/pon"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// PonQuery is the builder for querying Pon entities.
type PonQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Pon
	withCall   *CallQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PonQuery builder.
func (pq *PonQuery) Where(ps ...predicate.Pon) *PonQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PonQuery) Limit(limit int) *PonQuery {
	pq.limit = &limit
	return pq
}

// Offset to start from.
func (pq *PonQuery) Offset(offset int) *PonQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PonQuery) Unique(unique bool) *PonQuery {
	pq.unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PonQuery) Order(o ...OrderFunc) *PonQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryCall chains the current query on the "call" edge.
func (pq *PonQuery) QueryCall() *CallQuery {
	query := (&CallClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pon.Table, pon.FieldID, selector),
			sqlgraph.To(call.Table, call.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, pon.CallTable, pon.CallColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pon entity from the query.
// Returns a *NotFoundError when no Pon was found.
func (pq *PonQuery) First(ctx context.Context) (*Pon, error) {
	nodes, err := pq.Limit(1).All(newQueryContext(ctx, TypePon, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pon.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PonQuery) FirstX(ctx context.Context) *Pon {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pon ID from the query.
// Returns a *NotFoundError when no Pon ID was found.
func (pq *PonQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(newQueryContext(ctx, TypePon, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pon.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PonQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pon entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pon entity is found.
// Returns a *NotFoundError when no Pon entities are found.
func (pq *PonQuery) Only(ctx context.Context) (*Pon, error) {
	nodes, err := pq.Limit(2).All(newQueryContext(ctx, TypePon, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pon.Label}
	default:
		return nil, &NotSingularError{pon.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PonQuery) OnlyX(ctx context.Context) *Pon {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pon ID in the query.
// Returns a *NotSingularError when more than one Pon ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PonQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(newQueryContext(ctx, TypePon, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pon.Label}
	default:
		err = &NotSingularError{pon.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PonQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pons.
func (pq *PonQuery) All(ctx context.Context) ([]*Pon, error) {
	ctx = newQueryContext(ctx, TypePon, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pon, *PonQuery]()
	return withInterceptors[[]*Pon](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PonQuery) AllX(ctx context.Context) []*Pon {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pon IDs.
func (pq *PonQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypePon, "IDs")
	if err := pq.Select(pon.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PonQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PonQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypePon, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PonQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PonQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PonQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypePon, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PonQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PonQuery) Clone() *PonQuery {
	if pq == nil {
		return nil
	}
	return &PonQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Pon{}, pq.predicates...),
		withCall:   pq.withCall.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithCall tells the query-builder to eager-load the nodes that are connected to
// the "call" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PonQuery) WithCall(opts ...func(*CallQuery)) *PonQuery {
	query := (&CallClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCall = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (pq *PonQuery) GroupBy(field string, fields ...string) *PonGroupBy {
	pq.fields = append([]string{field}, fields...)
	grbuild := &PonGroupBy{build: pq}
	grbuild.flds = &pq.fields
	grbuild.label = pon.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (pq *PonQuery) Select(fields ...string) *PonSelect {
	pq.fields = append(pq.fields, fields...)
	sbuild := &PonSelect{PonQuery: pq}
	sbuild.label = pon.Label
	sbuild.flds, sbuild.scan = &pq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PonSelect configured with the given aggregations.
func (pq *PonQuery) Aggregate(fns ...AggregateFunc) *PonSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PonQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.fields {
		if !pon.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PonQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pon, error) {
	var (
		nodes       = []*Pon{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withCall != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pon).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pon{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withCall; query != nil {
		if err := pq.loadCall(ctx, query, nodes, nil,
			func(n *Pon, e *Call) { n.Edges.Call = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PonQuery) loadCall(ctx context.Context, query *CallQuery, nodes []*Pon, init func(*Pon), assign func(*Pon, *Call)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Pon)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Call(func(s *sql.Selector) {
		s.Where(sql.InValues(pon.CallColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.pon_call
		if fk == nil {
			return fmt.Errorf(`foreign-key "pon_call" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pon_call" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pon.Table,
			Columns: pon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pon.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pon.FieldID)
		for i := range fields {
			if fields[i] != pon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pon.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = pon.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PonGroupBy is the group-by builder for Pon entities.
type PonGroupBy struct {
	selector
	build *PonQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PonGroupBy) Aggregate(fns ...AggregateFunc) *PonGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PonGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypePon, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PonQuery, *PonGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PonGroupBy) sqlScan(ctx context.Context, root *PonQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PonSelect is the builder for selecting fields of Pon entities.
type PonSelect struct {
	*PonQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PonSelect) Aggregate(fns ...AggregateFunc) *PonSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PonSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypePon, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PonQuery, *PonSelect](ctx, ps.PonQuery, ps, ps.inters, v)
}

func (ps *PonSelect) sqlScan(ctx context.Context, root *PonQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
