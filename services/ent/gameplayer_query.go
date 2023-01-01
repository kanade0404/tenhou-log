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

// GamePlayerQuery is the builder for querying GamePlayer entities.
type GamePlayerQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.GamePlayer
	withPlayers *PlayerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GamePlayerQuery builder.
func (gpq *GamePlayerQuery) Where(ps ...predicate.GamePlayer) *GamePlayerQuery {
	gpq.predicates = append(gpq.predicates, ps...)
	return gpq
}

// Limit adds a limit step to the query.
func (gpq *GamePlayerQuery) Limit(limit int) *GamePlayerQuery {
	gpq.limit = &limit
	return gpq
}

// Offset adds an offset step to the query.
func (gpq *GamePlayerQuery) Offset(offset int) *GamePlayerQuery {
	gpq.offset = &offset
	return gpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gpq *GamePlayerQuery) Unique(unique bool) *GamePlayerQuery {
	gpq.unique = &unique
	return gpq
}

// Order adds an order step to the query.
func (gpq *GamePlayerQuery) Order(o ...OrderFunc) *GamePlayerQuery {
	gpq.order = append(gpq.order, o...)
	return gpq
}

// QueryPlayers chains the current query on the "players" edge.
func (gpq *GamePlayerQuery) QueryPlayers() *PlayerQuery {
	query := &PlayerQuery{config: gpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gameplayer.Table, gameplayer.FieldID, selector),
			sqlgraph.To(player.Table, player.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, gameplayer.PlayersTable, gameplayer.PlayersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GamePlayer entity from the query.
// Returns a *NotFoundError when no GamePlayer was found.
func (gpq *GamePlayerQuery) First(ctx context.Context) (*GamePlayer, error) {
	nodes, err := gpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gameplayer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gpq *GamePlayerQuery) FirstX(ctx context.Context) *GamePlayer {
	node, err := gpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GamePlayer ID from the query.
// Returns a *NotFoundError when no GamePlayer ID was found.
func (gpq *GamePlayerQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gameplayer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gpq *GamePlayerQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GamePlayer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GamePlayer entity is found.
// Returns a *NotFoundError when no GamePlayer entities are found.
func (gpq *GamePlayerQuery) Only(ctx context.Context) (*GamePlayer, error) {
	nodes, err := gpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gameplayer.Label}
	default:
		return nil, &NotSingularError{gameplayer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gpq *GamePlayerQuery) OnlyX(ctx context.Context) *GamePlayer {
	node, err := gpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GamePlayer ID in the query.
// Returns a *NotSingularError when more than one GamePlayer ID is found.
// Returns a *NotFoundError when no entities are found.
func (gpq *GamePlayerQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gameplayer.Label}
	default:
		err = &NotSingularError{gameplayer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gpq *GamePlayerQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GamePlayers.
func (gpq *GamePlayerQuery) All(ctx context.Context) ([]*GamePlayer, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gpq *GamePlayerQuery) AllX(ctx context.Context) []*GamePlayer {
	nodes, err := gpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GamePlayer IDs.
func (gpq *GamePlayerQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gpq.Select(gameplayer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gpq *GamePlayerQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gpq *GamePlayerQuery) Count(ctx context.Context) (int, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gpq *GamePlayerQuery) CountX(ctx context.Context) int {
	count, err := gpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gpq *GamePlayerQuery) Exist(ctx context.Context) (bool, error) {
	if err := gpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gpq *GamePlayerQuery) ExistX(ctx context.Context) bool {
	exist, err := gpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GamePlayerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gpq *GamePlayerQuery) Clone() *GamePlayerQuery {
	if gpq == nil {
		return nil
	}
	return &GamePlayerQuery{
		config:      gpq.config,
		limit:       gpq.limit,
		offset:      gpq.offset,
		order:       append([]OrderFunc{}, gpq.order...),
		predicates:  append([]predicate.GamePlayer{}, gpq.predicates...),
		withPlayers: gpq.withPlayers.Clone(),
		// clone intermediate query.
		sql:    gpq.sql.Clone(),
		path:   gpq.path,
		unique: gpq.unique,
	}
}

// WithPlayers tells the query-builder to eager-load the nodes that are connected to
// the "players" edge. The optional arguments are used to configure the query builder of the edge.
func (gpq *GamePlayerQuery) WithPlayers(opts ...func(*PlayerQuery)) *GamePlayerQuery {
	query := &PlayerQuery{config: gpq.config}
	for _, opt := range opts {
		opt(query)
	}
	gpq.withPlayers = query
	return gpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rate float64 `json:"rate,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GamePlayer.Query().
//		GroupBy(gameplayer.FieldRate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gpq *GamePlayerQuery) GroupBy(field string, fields ...string) *GamePlayerGroupBy {
	grbuild := &GamePlayerGroupBy{config: gpq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gpq.sqlQuery(ctx), nil
	}
	grbuild.label = gameplayer.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Rate float64 `json:"rate,omitempty"`
//	}
//
//	client.GamePlayer.Query().
//		Select(gameplayer.FieldRate).
//		Scan(ctx, &v)
//
func (gpq *GamePlayerQuery) Select(fields ...string) *GamePlayerSelect {
	gpq.fields = append(gpq.fields, fields...)
	selbuild := &GamePlayerSelect{GamePlayerQuery: gpq}
	selbuild.label = gameplayer.Label
	selbuild.flds, selbuild.scan = &gpq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a GamePlayerSelect configured with the given aggregations.
func (gpq *GamePlayerQuery) Aggregate(fns ...AggregateFunc) *GamePlayerSelect {
	return gpq.Select().Aggregate(fns...)
}

func (gpq *GamePlayerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gpq.fields {
		if !gameplayer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gpq.path != nil {
		prev, err := gpq.path(ctx)
		if err != nil {
			return err
		}
		gpq.sql = prev
	}
	return nil
}

func (gpq *GamePlayerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GamePlayer, error) {
	var (
		nodes       = []*GamePlayer{}
		_spec       = gpq.querySpec()
		loadedTypes = [1]bool{
			gpq.withPlayers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GamePlayer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GamePlayer{config: gpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gpq.withPlayers; query != nil {
		if err := gpq.loadPlayers(ctx, query, nodes,
			func(n *GamePlayer) { n.Edges.Players = []*Player{} },
			func(n *GamePlayer, e *Player) { n.Edges.Players = append(n.Edges.Players, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gpq *GamePlayerQuery) loadPlayers(ctx context.Context, query *PlayerQuery, nodes []*GamePlayer, init func(*GamePlayer), assign func(*GamePlayer, *Player)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*GamePlayer)
	nids := make(map[uuid.UUID]map[*GamePlayer]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(gameplayer.PlayersTable)
		s.Join(joinT).On(s.C(player.FieldID), joinT.C(gameplayer.PlayersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(gameplayer.PlayersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(gameplayer.PlayersPrimaryKey[1]))
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
				nids[inValue] = map[*GamePlayer]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "players" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (gpq *GamePlayerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gpq.querySpec()
	_spec.Node.Columns = gpq.fields
	if len(gpq.fields) > 0 {
		_spec.Unique = gpq.unique != nil && *gpq.unique
	}
	return sqlgraph.CountNodes(ctx, gpq.driver, _spec)
}

func (gpq *GamePlayerQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := gpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (gpq *GamePlayerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gameplayer.Table,
			Columns: gameplayer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gameplayer.FieldID,
			},
		},
		From:   gpq.sql,
		Unique: true,
	}
	if unique := gpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gameplayer.FieldID)
		for i := range fields {
			if fields[i] != gameplayer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gpq *GamePlayerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gpq.driver.Dialect())
	t1 := builder.Table(gameplayer.Table)
	columns := gpq.fields
	if len(columns) == 0 {
		columns = gameplayer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gpq.sql != nil {
		selector = gpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gpq.unique != nil && *gpq.unique {
		selector.Distinct()
	}
	for _, p := range gpq.predicates {
		p(selector)
	}
	for _, p := range gpq.order {
		p(selector)
	}
	if offset := gpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GamePlayerGroupBy is the group-by builder for GamePlayer entities.
type GamePlayerGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gpgb *GamePlayerGroupBy) Aggregate(fns ...AggregateFunc) *GamePlayerGroupBy {
	gpgb.fns = append(gpgb.fns, fns...)
	return gpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gpgb *GamePlayerGroupBy) Scan(ctx context.Context, v any) error {
	query, err := gpgb.path(ctx)
	if err != nil {
		return err
	}
	gpgb.sql = query
	return gpgb.sqlScan(ctx, v)
}

func (gpgb *GamePlayerGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range gpgb.fields {
		if !gameplayer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gpgb *GamePlayerGroupBy) sqlQuery() *sql.Selector {
	selector := gpgb.sql.Select()
	aggregation := make([]string, 0, len(gpgb.fns))
	for _, fn := range gpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gpgb.fields)+len(gpgb.fns))
		for _, f := range gpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gpgb.fields...)...)
}

// GamePlayerSelect is the builder for selecting fields of GamePlayer entities.
type GamePlayerSelect struct {
	*GamePlayerQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gps *GamePlayerSelect) Aggregate(fns ...AggregateFunc) *GamePlayerSelect {
	gps.fns = append(gps.fns, fns...)
	return gps
}

// Scan applies the selector query and scans the result into the given value.
func (gps *GamePlayerSelect) Scan(ctx context.Context, v any) error {
	if err := gps.prepareQuery(ctx); err != nil {
		return err
	}
	gps.sql = gps.GamePlayerQuery.sqlQuery(ctx)
	return gps.sqlScan(ctx, v)
}

func (gps *GamePlayerSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(gps.fns))
	for _, fn := range gps.fns {
		aggregation = append(aggregation, fn(gps.sql))
	}
	switch n := len(*gps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		gps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		gps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := gps.sql.Query()
	if err := gps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
