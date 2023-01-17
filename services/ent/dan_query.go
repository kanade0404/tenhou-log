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
	"github.com/kanade0404/tenhou-log/services/ent/dan"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayer"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// DanQuery is the builder for querying Dan entities.
type DanQuery struct {
	config
	limit           *int
	offset          *int
	unique          *bool
	order           []OrderFunc
	fields          []string
	inters          []Interceptor
	predicates      []predicate.Dan
	withGamePlayers *GamePlayerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DanQuery builder.
func (dq *DanQuery) Where(ps ...predicate.Dan) *DanQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DanQuery) Limit(limit int) *DanQuery {
	dq.limit = &limit
	return dq
}

// Offset to start from.
func (dq *DanQuery) Offset(offset int) *DanQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DanQuery) Unique(unique bool) *DanQuery {
	dq.unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DanQuery) Order(o ...OrderFunc) *DanQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryGamePlayers chains the current query on the "game_players" edge.
func (dq *DanQuery) QueryGamePlayers() *GamePlayerQuery {
	query := (&GamePlayerClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dan.Table, dan.FieldID, selector),
			sqlgraph.To(gameplayer.Table, gameplayer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dan.GamePlayersTable, dan.GamePlayersColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Dan entity from the query.
// Returns a *NotFoundError when no Dan was found.
func (dq *DanQuery) First(ctx context.Context) (*Dan, error) {
	nodes, err := dq.Limit(1).All(newQueryContext(ctx, TypeDan, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DanQuery) FirstX(ctx context.Context) *Dan {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dan ID from the query.
// Returns a *NotFoundError when no Dan ID was found.
func (dq *DanQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(newQueryContext(ctx, TypeDan, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DanQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Dan entity is found.
// Returns a *NotFoundError when no Dan entities are found.
func (dq *DanQuery) Only(ctx context.Context) (*Dan, error) {
	nodes, err := dq.Limit(2).All(newQueryContext(ctx, TypeDan, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dan.Label}
	default:
		return nil, &NotSingularError{dan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DanQuery) OnlyX(ctx context.Context) *Dan {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dan ID in the query.
// Returns a *NotSingularError when more than one Dan ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DanQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(newQueryContext(ctx, TypeDan, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dan.Label}
	default:
		err = &NotSingularError{dan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DanQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Dans.
func (dq *DanQuery) All(ctx context.Context) ([]*Dan, error) {
	ctx = newQueryContext(ctx, TypeDan, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Dan, *DanQuery]()
	return withInterceptors[[]*Dan](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DanQuery) AllX(ctx context.Context) []*Dan {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dan IDs.
func (dq *DanQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeDan, "IDs")
	if err := dq.Select(dan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DanQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DanQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeDan, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DanQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DanQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeDan, "Exist")
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
func (dq *DanQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DanQuery) Clone() *DanQuery {
	if dq == nil {
		return nil
	}
	return &DanQuery{
		config:          dq.config,
		limit:           dq.limit,
		offset:          dq.offset,
		order:           append([]OrderFunc{}, dq.order...),
		inters:          append([]Interceptor{}, dq.inters...),
		predicates:      append([]predicate.Dan{}, dq.predicates...),
		withGamePlayers: dq.withGamePlayers.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithGamePlayers tells the query-builder to eager-load the nodes that are connected to
// the "game_players" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DanQuery) WithGamePlayers(opts ...func(*GamePlayerQuery)) *DanQuery {
	query := (&GamePlayerClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withGamePlayers = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Dan.Query().
//		GroupBy(dan.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DanQuery) GroupBy(field string, fields ...string) *DanGroupBy {
	dq.fields = append([]string{field}, fields...)
	grbuild := &DanGroupBy{build: dq}
	grbuild.flds = &dq.fields
	grbuild.label = dan.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Dan.Query().
//		Select(dan.FieldName).
//		Scan(ctx, &v)
func (dq *DanQuery) Select(fields ...string) *DanSelect {
	dq.fields = append(dq.fields, fields...)
	sbuild := &DanSelect{DanQuery: dq}
	sbuild.label = dan.Label
	sbuild.flds, sbuild.scan = &dq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DanSelect configured with the given aggregations.
func (dq *DanQuery) Aggregate(fns ...AggregateFunc) *DanSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DanQuery) prepareQuery(ctx context.Context) error {
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
		if !dan.ValidColumn(f) {
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

func (dq *DanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Dan, error) {
	var (
		nodes       = []*Dan{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withGamePlayers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Dan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Dan{config: dq.config}
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
	if query := dq.withGamePlayers; query != nil {
		if err := dq.loadGamePlayers(ctx, query, nodes,
			func(n *Dan) { n.Edges.GamePlayers = []*GamePlayer{} },
			func(n *Dan, e *GamePlayer) { n.Edges.GamePlayers = append(n.Edges.GamePlayers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DanQuery) loadGamePlayers(ctx context.Context, query *GamePlayerQuery, nodes []*Dan, init func(*Dan), assign func(*Dan, *GamePlayer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Dan)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GamePlayer(func(s *sql.Selector) {
		s.Where(sql.InValues(dan.GamePlayersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dan_game_players
		if fk == nil {
			return fmt.Errorf(`foreign-key "dan_game_players" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dan_game_players" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dan.Table,
			Columns: dan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dan.FieldID,
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
		_spec.Node.Columns = append(_spec.Node.Columns, dan.FieldID)
		for i := range fields {
			if fields[i] != dan.FieldID {
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

func (dq *DanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(dan.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = dan.Columns
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

// DanGroupBy is the group-by builder for Dan entities.
type DanGroupBy struct {
	selector
	build *DanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DanGroupBy) Aggregate(fns ...AggregateFunc) *DanGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeDan, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DanQuery, *DanGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DanGroupBy) sqlScan(ctx context.Context, root *DanQuery, v any) error {
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

// DanSelect is the builder for selecting fields of Dan entities.
type DanSelect struct {
	*DanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DanSelect) Aggregate(fns ...AggregateFunc) *DanSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DanSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeDan, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DanQuery, *DanSelect](ctx, ds.DanQuery, ds, ds.inters, v)
}

func (ds *DanSelect) sqlScan(ctx context.Context, root *DanQuery, v any) error {
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
