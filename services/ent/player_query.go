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
	"github.com/kanade0404/tenhou-log/services/ent/gameplayer"
	"github.com/kanade0404/tenhou-log/services/ent/player"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// PlayerQuery is the builder for querying Player entities.
type PlayerQuery struct {
	config
	ctx             *QueryContext
	order           []OrderFunc
	inters          []Interceptor
	predicates      []predicate.Player
	withGamePlayers *GamePlayerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlayerQuery builder.
func (pq *PlayerQuery) Where(ps ...predicate.Player) *PlayerQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PlayerQuery) Limit(limit int) *PlayerQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PlayerQuery) Offset(offset int) *PlayerQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PlayerQuery) Unique(unique bool) *PlayerQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PlayerQuery) Order(o ...OrderFunc) *PlayerQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryGamePlayers chains the current query on the "game_players" edge.
func (pq *PlayerQuery) QueryGamePlayers() *GamePlayerQuery {
	query := (&GamePlayerClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(player.Table, player.FieldID, selector),
			sqlgraph.To(gameplayer.Table, gameplayer.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, player.GamePlayersTable, player.GamePlayersColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Player entity from the query.
// Returns a *NotFoundError when no Player was found.
func (pq *PlayerQuery) First(ctx context.Context) (*Player, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{player.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PlayerQuery) FirstX(ctx context.Context) *Player {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Player ID from the query.
// Returns a *NotFoundError when no Player ID was found.
func (pq *PlayerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{player.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PlayerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Player entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Player entity is found.
// Returns a *NotFoundError when no Player entities are found.
func (pq *PlayerQuery) Only(ctx context.Context) (*Player, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{player.Label}
	default:
		return nil, &NotSingularError{player.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PlayerQuery) OnlyX(ctx context.Context) *Player {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Player ID in the query.
// Returns a *NotSingularError when more than one Player ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PlayerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{player.Label}
	default:
		err = &NotSingularError{player.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PlayerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Players.
func (pq *PlayerQuery) All(ctx context.Context) ([]*Player, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Player, *PlayerQuery]()
	return withInterceptors[[]*Player](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PlayerQuery) AllX(ctx context.Context) []*Player {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Player IDs.
func (pq *PlayerQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(player.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PlayerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PlayerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PlayerQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PlayerQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PlayerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
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
func (pq *PlayerQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlayerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PlayerQuery) Clone() *PlayerQuery {
	if pq == nil {
		return nil
	}
	return &PlayerQuery{
		config:          pq.config,
		ctx:             pq.ctx.Clone(),
		order:           append([]OrderFunc{}, pq.order...),
		inters:          append([]Interceptor{}, pq.inters...),
		predicates:      append([]predicate.Player{}, pq.predicates...),
		withGamePlayers: pq.withGamePlayers.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithGamePlayers tells the query-builder to eager-load the nodes that are connected to
// the "game_players" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PlayerQuery) WithGamePlayers(opts ...func(*GamePlayerQuery)) *PlayerQuery {
	query := (&GamePlayerClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withGamePlayers = query
	return pq
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
//	client.Player.Query().
//		GroupBy(player.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PlayerQuery) GroupBy(field string, fields ...string) *PlayerGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlayerGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = player.Label
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
//	client.Player.Query().
//		Select(player.FieldName).
//		Scan(ctx, &v)
func (pq *PlayerQuery) Select(fields ...string) *PlayerSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PlayerSelect{PlayerQuery: pq}
	sbuild.label = player.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlayerSelect configured with the given aggregations.
func (pq *PlayerQuery) Aggregate(fns ...AggregateFunc) *PlayerSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PlayerQuery) prepareQuery(ctx context.Context) error {
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
	for _, f := range pq.ctx.Fields {
		if !player.ValidColumn(f) {
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

func (pq *PlayerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Player, error) {
	var (
		nodes       = []*Player{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withGamePlayers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Player).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Player{config: pq.config}
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
	if query := pq.withGamePlayers; query != nil {
		if err := pq.loadGamePlayers(ctx, query, nodes,
			func(n *Player) { n.Edges.GamePlayers = []*GamePlayer{} },
			func(n *Player, e *GamePlayer) { n.Edges.GamePlayers = append(n.Edges.GamePlayers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PlayerQuery) loadGamePlayers(ctx context.Context, query *GamePlayerQuery, nodes []*Player, init func(*Player), assign func(*Player, *GamePlayer)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Player)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.GamePlayer(func(s *sql.Selector) {
		s.Where(sql.InValues(player.GamePlayersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.player_game_players
		if fk == nil {
			return fmt.Errorf(`foreign-key "player_game_players" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "player_game_players" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PlayerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PlayerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, player.FieldID)
		for i := range fields {
			if fields[i] != player.FieldID {
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
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
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

func (pq *PlayerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(player.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = player.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlayerGroupBy is the group-by builder for Player entities.
type PlayerGroupBy struct {
	selector
	build *PlayerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PlayerGroupBy) Aggregate(fns ...AggregateFunc) *PlayerGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PlayerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlayerQuery, *PlayerGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PlayerGroupBy) sqlScan(ctx context.Context, root *PlayerQuery, v any) error {
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

// PlayerSelect is the builder for selecting fields of Player entities.
type PlayerSelect struct {
	*PlayerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PlayerSelect) Aggregate(fns ...AggregateFunc) *PlayerSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PlayerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlayerQuery, *PlayerSelect](ctx, ps.PlayerQuery, ps, ps.inters, v)
}

func (ps *PlayerSelect) sqlScan(ctx context.Context, root *PlayerQuery, v any) error {
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
