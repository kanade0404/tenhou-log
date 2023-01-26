// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/discard"
	"github.com/kanade0404/tenhou-log/services/ent/drawn"
	"github.com/kanade0404/tenhou-log/services/ent/event"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// DrawnQuery is the builder for querying Drawn entities.
type DrawnQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.Drawn
	withEvent   *EventQuery
	withDiscard *DiscardQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DrawnQuery builder.
func (dq *DrawnQuery) Where(ps ...predicate.Drawn) *DrawnQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DrawnQuery) Limit(limit int) *DrawnQuery {
	dq.limit = &limit
	return dq
}

// Offset to start from.
func (dq *DrawnQuery) Offset(offset int) *DrawnQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DrawnQuery) Unique(unique bool) *DrawnQuery {
	dq.unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DrawnQuery) Order(o ...OrderFunc) *DrawnQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryEvent chains the current query on the "event" edge.
func (dq *DrawnQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(drawn.Table, drawn.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, drawn.EventTable, drawn.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDiscard chains the current query on the "discard" edge.
func (dq *DrawnQuery) QueryDiscard() *DiscardQuery {
	query := (&DiscardClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(drawn.Table, drawn.FieldID, selector),
			sqlgraph.To(discard.Table, discard.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, drawn.DiscardTable, drawn.DiscardColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Drawn entity from the query.
// Returns a *NotFoundError when no Drawn was found.
func (dq *DrawnQuery) First(ctx context.Context) (*Drawn, error) {
	nodes, err := dq.Limit(1).All(newQueryContext(ctx, TypeDrawn, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{drawn.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DrawnQuery) FirstX(ctx context.Context) *Drawn {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Drawn ID from the query.
// Returns a *NotFoundError when no Drawn ID was found.
func (dq *DrawnQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(newQueryContext(ctx, TypeDrawn, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{drawn.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DrawnQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Drawn entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Drawn entity is found.
// Returns a *NotFoundError when no Drawn entities are found.
func (dq *DrawnQuery) Only(ctx context.Context) (*Drawn, error) {
	nodes, err := dq.Limit(2).All(newQueryContext(ctx, TypeDrawn, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{drawn.Label}
	default:
		return nil, &NotSingularError{drawn.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DrawnQuery) OnlyX(ctx context.Context) *Drawn {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Drawn ID in the query.
// Returns a *NotSingularError when more than one Drawn ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DrawnQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(newQueryContext(ctx, TypeDrawn, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{drawn.Label}
	default:
		err = &NotSingularError{drawn.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DrawnQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Drawns.
func (dq *DrawnQuery) All(ctx context.Context) ([]*Drawn, error) {
	ctx = newQueryContext(ctx, TypeDrawn, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Drawn, *DrawnQuery]()
	return withInterceptors[[]*Drawn](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DrawnQuery) AllX(ctx context.Context) []*Drawn {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Drawn IDs.
func (dq *DrawnQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeDrawn, "IDs")
	if err := dq.Select(drawn.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DrawnQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DrawnQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeDrawn, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DrawnQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DrawnQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DrawnQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeDrawn, "Exist")
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
func (dq *DrawnQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DrawnQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DrawnQuery) Clone() *DrawnQuery {
	if dq == nil {
		return nil
	}
	return &DrawnQuery{
		config:      dq.config,
		limit:       dq.limit,
		offset:      dq.offset,
		order:       append([]OrderFunc{}, dq.order...),
		inters:      append([]Interceptor{}, dq.inters...),
		predicates:  append([]predicate.Drawn{}, dq.predicates...),
		withEvent:   dq.withEvent.Clone(),
		withDiscard: dq.withDiscard.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DrawnQuery) WithEvent(opts ...func(*EventQuery)) *DrawnQuery {
	query := (&EventClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withEvent = query
	return dq
}

// WithDiscard tells the query-builder to eager-load the nodes that are connected to
// the "discard" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DrawnQuery) WithDiscard(opts ...func(*DiscardQuery)) *DrawnQuery {
	query := (&DiscardClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDiscard = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (dq *DrawnQuery) GroupBy(field string, fields ...string) *DrawnGroupBy {
	dq.fields = append([]string{field}, fields...)
	grbuild := &DrawnGroupBy{build: dq}
	grbuild.flds = &dq.fields
	grbuild.label = drawn.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (dq *DrawnQuery) Select(fields ...string) *DrawnSelect {
	dq.fields = append(dq.fields, fields...)
	sbuild := &DrawnSelect{DrawnQuery: dq}
	sbuild.label = drawn.Label
	sbuild.flds, sbuild.scan = &dq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DrawnSelect configured with the given aggregations.
func (dq *DrawnQuery) Aggregate(fns ...AggregateFunc) *DrawnSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DrawnQuery) prepareQuery(ctx context.Context) error {
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
		if !drawn.ValidColumn(f) {
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

func (dq *DrawnQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Drawn, error) {
	var (
		nodes       = []*Drawn{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [2]bool{
			dq.withEvent != nil,
			dq.withDiscard != nil,
		}
	)
	if dq.withEvent != nil || dq.withDiscard != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, drawn.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Drawn).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Drawn{config: dq.config}
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
	if query := dq.withEvent; query != nil {
		if err := dq.loadEvent(ctx, query, nodes, nil,
			func(n *Drawn, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withDiscard; query != nil {
		if err := dq.loadDiscard(ctx, query, nodes, nil,
			func(n *Drawn, e *Discard) { n.Edges.Discard = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DrawnQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*Drawn, init func(*Drawn), assign func(*Drawn, *Event)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Drawn)
	for i := range nodes {
		if nodes[i].event_draw == nil {
			continue
		}
		fk := *nodes[i].event_draw
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(event.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "event_draw" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DrawnQuery) loadDiscard(ctx context.Context, query *DiscardQuery, nodes []*Drawn, init func(*Drawn), assign func(*Drawn, *Discard)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Drawn)
	for i := range nodes {
		if nodes[i].discard_draw == nil {
			continue
		}
		fk := *nodes[i].discard_draw
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(discard.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "discard_draw" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DrawnQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DrawnQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drawn.Table,
			Columns: drawn.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: drawn.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, drawn.FieldID)
		for i := range fields {
			if fields[i] != drawn.FieldID {
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

func (dq *DrawnQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(drawn.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = drawn.Columns
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

// DrawnGroupBy is the group-by builder for Drawn entities.
type DrawnGroupBy struct {
	selector
	build *DrawnQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DrawnGroupBy) Aggregate(fns ...AggregateFunc) *DrawnGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DrawnGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeDrawn, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DrawnQuery, *DrawnGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DrawnGroupBy) sqlScan(ctx context.Context, root *DrawnQuery, v any) error {
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

// DrawnSelect is the builder for selecting fields of Drawn entities.
type DrawnSelect struct {
	*DrawnQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DrawnSelect) Aggregate(fns ...AggregateFunc) *DrawnSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DrawnSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeDrawn, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DrawnQuery, *DrawnSelect](ctx, ds.DrawnQuery, ds, ds.inters, v)
}

func (ds *DrawnSelect) sqlScan(ctx context.Context, root *DrawnQuery, v any) error {
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