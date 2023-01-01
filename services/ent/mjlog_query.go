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
	"github.com/kanade0404/tenhou-log/services/ent/game"
	"github.com/kanade0404/tenhou-log/services/ent/mjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// MJLogQuery is the builder for querying MJLog entities.
type MJLogQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.MJLog
	withMjlogFiles *MJLogFileQuery
	withGames      *GameQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MJLogQuery builder.
func (mlq *MJLogQuery) Where(ps ...predicate.MJLog) *MJLogQuery {
	mlq.predicates = append(mlq.predicates, ps...)
	return mlq
}

// Limit adds a limit step to the query.
func (mlq *MJLogQuery) Limit(limit int) *MJLogQuery {
	mlq.limit = &limit
	return mlq
}

// Offset adds an offset step to the query.
func (mlq *MJLogQuery) Offset(offset int) *MJLogQuery {
	mlq.offset = &offset
	return mlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mlq *MJLogQuery) Unique(unique bool) *MJLogQuery {
	mlq.unique = &unique
	return mlq
}

// Order adds an order step to the query.
func (mlq *MJLogQuery) Order(o ...OrderFunc) *MJLogQuery {
	mlq.order = append(mlq.order, o...)
	return mlq
}

// QueryMjlogFiles chains the current query on the "mjlog_files" edge.
func (mlq *MJLogQuery) QueryMjlogFiles() *MJLogFileQuery {
	query := &MJLogFileQuery{config: mlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mjlog.Table, mjlog.FieldID, selector),
			sqlgraph.To(mjlogfile.Table, mjlogfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, mjlog.MjlogFilesTable, mjlog.MjlogFilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGames chains the current query on the "games" edge.
func (mlq *MJLogQuery) QueryGames() *GameQuery {
	query := &GameQuery{config: mlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mjlog.Table, mjlog.FieldID, selector),
			sqlgraph.To(game.Table, game.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, mjlog.GamesTable, mjlog.GamesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MJLog entity from the query.
// Returns a *NotFoundError when no MJLog was found.
func (mlq *MJLogQuery) First(ctx context.Context) (*MJLog, error) {
	nodes, err := mlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mjlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mlq *MJLogQuery) FirstX(ctx context.Context) *MJLog {
	node, err := mlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MJLog ID from the query.
// Returns a *NotFoundError when no MJLog ID was found.
func (mlq *MJLogQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mjlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mlq *MJLogQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MJLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MJLog entity is found.
// Returns a *NotFoundError when no MJLog entities are found.
func (mlq *MJLogQuery) Only(ctx context.Context) (*MJLog, error) {
	nodes, err := mlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mjlog.Label}
	default:
		return nil, &NotSingularError{mjlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mlq *MJLogQuery) OnlyX(ctx context.Context) *MJLog {
	node, err := mlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MJLog ID in the query.
// Returns a *NotSingularError when more than one MJLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (mlq *MJLogQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mjlog.Label}
	default:
		err = &NotSingularError{mjlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mlq *MJLogQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MJLogs.
func (mlq *MJLogQuery) All(ctx context.Context) ([]*MJLog, error) {
	if err := mlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mlq *MJLogQuery) AllX(ctx context.Context) []*MJLog {
	nodes, err := mlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MJLog IDs.
func (mlq *MJLogQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := mlq.Select(mjlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mlq *MJLogQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mlq *MJLogQuery) Count(ctx context.Context) (int, error) {
	if err := mlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mlq *MJLogQuery) CountX(ctx context.Context) int {
	count, err := mlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mlq *MJLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := mlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mlq *MJLogQuery) ExistX(ctx context.Context) bool {
	exist, err := mlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MJLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mlq *MJLogQuery) Clone() *MJLogQuery {
	if mlq == nil {
		return nil
	}
	return &MJLogQuery{
		config:         mlq.config,
		limit:          mlq.limit,
		offset:         mlq.offset,
		order:          append([]OrderFunc{}, mlq.order...),
		predicates:     append([]predicate.MJLog{}, mlq.predicates...),
		withMjlogFiles: mlq.withMjlogFiles.Clone(),
		withGames:      mlq.withGames.Clone(),
		// clone intermediate query.
		sql:    mlq.sql.Clone(),
		path:   mlq.path,
		unique: mlq.unique,
	}
}

// WithMjlogFiles tells the query-builder to eager-load the nodes that are connected to
// the "mjlog_files" edge. The optional arguments are used to configure the query builder of the edge.
func (mlq *MJLogQuery) WithMjlogFiles(opts ...func(*MJLogFileQuery)) *MJLogQuery {
	query := &MJLogFileQuery{config: mlq.config}
	for _, opt := range opts {
		opt(query)
	}
	mlq.withMjlogFiles = query
	return mlq
}

// WithGames tells the query-builder to eager-load the nodes that are connected to
// the "games" edge. The optional arguments are used to configure the query builder of the edge.
func (mlq *MJLogQuery) WithGames(opts ...func(*GameQuery)) *MJLogQuery {
	query := &GameQuery{config: mlq.config}
	for _, opt := range opts {
		opt(query)
	}
	mlq.withGames = query
	return mlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Version float64 `json:"version,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MJLog.Query().
//		GroupBy(mjlog.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mlq *MJLogQuery) GroupBy(field string, fields ...string) *MJLogGroupBy {
	grbuild := &MJLogGroupBy{config: mlq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mlq.sqlQuery(ctx), nil
	}
	grbuild.label = mjlog.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Version float64 `json:"version,omitempty"`
//	}
//
//	client.MJLog.Query().
//		Select(mjlog.FieldVersion).
//		Scan(ctx, &v)
//
func (mlq *MJLogQuery) Select(fields ...string) *MJLogSelect {
	mlq.fields = append(mlq.fields, fields...)
	selbuild := &MJLogSelect{MJLogQuery: mlq}
	selbuild.label = mjlog.Label
	selbuild.flds, selbuild.scan = &mlq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a MJLogSelect configured with the given aggregations.
func (mlq *MJLogQuery) Aggregate(fns ...AggregateFunc) *MJLogSelect {
	return mlq.Select().Aggregate(fns...)
}

func (mlq *MJLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mlq.fields {
		if !mjlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mlq.path != nil {
		prev, err := mlq.path(ctx)
		if err != nil {
			return err
		}
		mlq.sql = prev
	}
	return nil
}

func (mlq *MJLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MJLog, error) {
	var (
		nodes       = []*MJLog{}
		withFKs     = mlq.withFKs
		_spec       = mlq.querySpec()
		loadedTypes = [2]bool{
			mlq.withMjlogFiles != nil,
			mlq.withGames != nil,
		}
	)
	if mlq.withMjlogFiles != nil || mlq.withGames != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, mjlog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MJLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MJLog{config: mlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mlq.withMjlogFiles; query != nil {
		if err := mlq.loadMjlogFiles(ctx, query, nodes, nil,
			func(n *MJLog, e *MJLogFile) { n.Edges.MjlogFiles = e }); err != nil {
			return nil, err
		}
	}
	if query := mlq.withGames; query != nil {
		if err := mlq.loadGames(ctx, query, nodes, nil,
			func(n *MJLog, e *Game) { n.Edges.Games = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mlq *MJLogQuery) loadMjlogFiles(ctx context.Context, query *MJLogFileQuery, nodes []*MJLog, init func(*MJLog), assign func(*MJLog, *MJLogFile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*MJLog)
	for i := range nodes {
		if nodes[i].mj_log_file_mjlogs == nil {
			continue
		}
		fk := *nodes[i].mj_log_file_mjlogs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(mjlogfile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mj_log_file_mjlogs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mlq *MJLogQuery) loadGames(ctx context.Context, query *GameQuery, nodes []*MJLog, init func(*MJLog), assign func(*MJLog, *Game)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*MJLog)
	for i := range nodes {
		if nodes[i].game_mjlogs == nil {
			continue
		}
		fk := *nodes[i].game_mjlogs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(game.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "game_mjlogs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mlq *MJLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mlq.querySpec()
	_spec.Node.Columns = mlq.fields
	if len(mlq.fields) > 0 {
		_spec.Unique = mlq.unique != nil && *mlq.unique
	}
	return sqlgraph.CountNodes(ctx, mlq.driver, _spec)
}

func (mlq *MJLogQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := mlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (mlq *MJLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mjlog.Table,
			Columns: mjlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mjlog.FieldID,
			},
		},
		From:   mlq.sql,
		Unique: true,
	}
	if unique := mlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mjlog.FieldID)
		for i := range fields {
			if fields[i] != mjlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mlq *MJLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mlq.driver.Dialect())
	t1 := builder.Table(mjlog.Table)
	columns := mlq.fields
	if len(columns) == 0 {
		columns = mjlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mlq.sql != nil {
		selector = mlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mlq.unique != nil && *mlq.unique {
		selector.Distinct()
	}
	for _, p := range mlq.predicates {
		p(selector)
	}
	for _, p := range mlq.order {
		p(selector)
	}
	if offset := mlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MJLogGroupBy is the group-by builder for MJLog entities.
type MJLogGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mlgb *MJLogGroupBy) Aggregate(fns ...AggregateFunc) *MJLogGroupBy {
	mlgb.fns = append(mlgb.fns, fns...)
	return mlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mlgb *MJLogGroupBy) Scan(ctx context.Context, v any) error {
	query, err := mlgb.path(ctx)
	if err != nil {
		return err
	}
	mlgb.sql = query
	return mlgb.sqlScan(ctx, v)
}

func (mlgb *MJLogGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range mlgb.fields {
		if !mjlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mlgb *MJLogGroupBy) sqlQuery() *sql.Selector {
	selector := mlgb.sql.Select()
	aggregation := make([]string, 0, len(mlgb.fns))
	for _, fn := range mlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mlgb.fields)+len(mlgb.fns))
		for _, f := range mlgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mlgb.fields...)...)
}

// MJLogSelect is the builder for selecting fields of MJLog entities.
type MJLogSelect struct {
	*MJLogQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mls *MJLogSelect) Aggregate(fns ...AggregateFunc) *MJLogSelect {
	mls.fns = append(mls.fns, fns...)
	return mls
}

// Scan applies the selector query and scans the result into the given value.
func (mls *MJLogSelect) Scan(ctx context.Context, v any) error {
	if err := mls.prepareQuery(ctx); err != nil {
		return err
	}
	mls.sql = mls.MJLogQuery.sqlQuery(ctx)
	return mls.sqlScan(ctx, v)
}

func (mls *MJLogSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(mls.fns))
	for _, fn := range mls.fns {
		aggregation = append(aggregation, fn(mls.sql))
	}
	switch n := len(*mls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		mls.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		mls.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := mls.sql.Query()
	if err := mls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
