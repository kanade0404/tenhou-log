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
	"github.com/kanade0404/tenhou-log/services/ent/chakan"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ChakanQuery is the builder for querying Chakan entities.
type ChakanQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Chakan
	withCall   *CallQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChakanQuery builder.
func (cq *ChakanQuery) Where(ps ...predicate.Chakan) *ChakanQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ChakanQuery) Limit(limit int) *ChakanQuery {
	cq.limit = &limit
	return cq
}

// Offset to start from.
func (cq *ChakanQuery) Offset(offset int) *ChakanQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ChakanQuery) Unique(unique bool) *ChakanQuery {
	cq.unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ChakanQuery) Order(o ...OrderFunc) *ChakanQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryCall chains the current query on the "call" edge.
func (cq *ChakanQuery) QueryCall() *CallQuery {
	query := (&CallClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chakan.Table, chakan.FieldID, selector),
			sqlgraph.To(call.Table, call.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, chakan.CallTable, chakan.CallColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Chakan entity from the query.
// Returns a *NotFoundError when no Chakan was found.
func (cq *ChakanQuery) First(ctx context.Context) (*Chakan, error) {
	nodes, err := cq.Limit(1).All(newQueryContext(ctx, TypeChakan, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chakan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChakanQuery) FirstX(ctx context.Context) *Chakan {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Chakan ID from the query.
// Returns a *NotFoundError when no Chakan ID was found.
func (cq *ChakanQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(newQueryContext(ctx, TypeChakan, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chakan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ChakanQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Chakan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Chakan entity is found.
// Returns a *NotFoundError when no Chakan entities are found.
func (cq *ChakanQuery) Only(ctx context.Context) (*Chakan, error) {
	nodes, err := cq.Limit(2).All(newQueryContext(ctx, TypeChakan, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chakan.Label}
	default:
		return nil, &NotSingularError{chakan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChakanQuery) OnlyX(ctx context.Context) *Chakan {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Chakan ID in the query.
// Returns a *NotSingularError when more than one Chakan ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ChakanQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(newQueryContext(ctx, TypeChakan, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chakan.Label}
	default:
		err = &NotSingularError{chakan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChakanQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Chakans.
func (cq *ChakanQuery) All(ctx context.Context) ([]*Chakan, error) {
	ctx = newQueryContext(ctx, TypeChakan, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Chakan, *ChakanQuery]()
	return withInterceptors[[]*Chakan](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChakanQuery) AllX(ctx context.Context) []*Chakan {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Chakan IDs.
func (cq *ChakanQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeChakan, "IDs")
	if err := cq.Select(chakan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChakanQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChakanQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeChakan, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ChakanQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChakanQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChakanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeChakan, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChakanQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChakanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChakanQuery) Clone() *ChakanQuery {
	if cq == nil {
		return nil
	}
	return &ChakanQuery{
		config:     cq.config,
		limit:      cq.limit,
		offset:     cq.offset,
		order:      append([]OrderFunc{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Chakan{}, cq.predicates...),
		withCall:   cq.withCall.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithCall tells the query-builder to eager-load the nodes that are connected to
// the "call" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChakanQuery) WithCall(opts ...func(*CallQuery)) *ChakanQuery {
	query := (&CallClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCall = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (cq *ChakanQuery) GroupBy(field string, fields ...string) *ChakanGroupBy {
	cq.fields = append([]string{field}, fields...)
	grbuild := &ChakanGroupBy{build: cq}
	grbuild.flds = &cq.fields
	grbuild.label = chakan.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (cq *ChakanQuery) Select(fields ...string) *ChakanSelect {
	cq.fields = append(cq.fields, fields...)
	sbuild := &ChakanSelect{ChakanQuery: cq}
	sbuild.label = chakan.Label
	sbuild.flds, sbuild.scan = &cq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChakanSelect configured with the given aggregations.
func (cq *ChakanQuery) Aggregate(fns ...AggregateFunc) *ChakanSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ChakanQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.fields {
		if !chakan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ChakanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Chakan, error) {
	var (
		nodes       = []*Chakan{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withCall != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Chakan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Chakan{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withCall; query != nil {
		if err := cq.loadCall(ctx, query, nodes, nil,
			func(n *Chakan, e *Call) { n.Edges.Call = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ChakanQuery) loadCall(ctx context.Context, query *CallQuery, nodes []*Chakan, init func(*Chakan), assign func(*Chakan, *Call)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Chakan)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Call(func(s *sql.Selector) {
		s.Where(sql.InValues(chakan.CallColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.chakan_call
		if fk == nil {
			return fmt.Errorf(`foreign-key "chakan_call" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chakan_call" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ChakanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChakanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chakan.Table,
			Columns: chakan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: chakan.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chakan.FieldID)
		for i := range fields {
			if fields[i] != chakan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ChakanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(chakan.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = chakan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChakanGroupBy is the group-by builder for Chakan entities.
type ChakanGroupBy struct {
	selector
	build *ChakanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChakanGroupBy) Aggregate(fns ...AggregateFunc) *ChakanGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ChakanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeChakan, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChakanQuery, *ChakanGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ChakanGroupBy) sqlScan(ctx context.Context, root *ChakanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChakanSelect is the builder for selecting fields of Chakan entities.
type ChakanSelect struct {
	*ChakanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ChakanSelect) Aggregate(fns ...AggregateFunc) *ChakanSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ChakanSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeChakan, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChakanQuery, *ChakanSelect](ctx, cs.ChakanQuery, cs, cs.inters, v)
}

func (cs *ChakanSelect) sqlScan(ctx context.Context, root *ChakanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
