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
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// MJLogFileQuery is the builder for querying MJLogFile entities.
type MJLogFileQuery struct {
	config
	limit                    *int
	offset                   *int
	unique                   *bool
	order                    []OrderFunc
	fields                   []string
	inters                   []Interceptor
	predicates               []predicate.MJLogFile
	withCompressedMjlogFiles *CompressedMJLogQuery
	withMjlogs               *MJLogQuery
	withFKs                  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MJLogFileQuery builder.
func (mlfq *MJLogFileQuery) Where(ps ...predicate.MJLogFile) *MJLogFileQuery {
	mlfq.predicates = append(mlfq.predicates, ps...)
	return mlfq
}

// Limit the number of records to be returned by this query.
func (mlfq *MJLogFileQuery) Limit(limit int) *MJLogFileQuery {
	mlfq.limit = &limit
	return mlfq
}

// Offset to start from.
func (mlfq *MJLogFileQuery) Offset(offset int) *MJLogFileQuery {
	mlfq.offset = &offset
	return mlfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mlfq *MJLogFileQuery) Unique(unique bool) *MJLogFileQuery {
	mlfq.unique = &unique
	return mlfq
}

// Order specifies how the records should be ordered.
func (mlfq *MJLogFileQuery) Order(o ...OrderFunc) *MJLogFileQuery {
	mlfq.order = append(mlfq.order, o...)
	return mlfq
}

// QueryCompressedMjlogFiles chains the current query on the "compressed_mjlog_files" edge.
func (mlfq *MJLogFileQuery) QueryCompressedMjlogFiles() *CompressedMJLogQuery {
	query := (&CompressedMJLogClient{config: mlfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mlfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mlfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mjlogfile.Table, mjlogfile.FieldID, selector),
			sqlgraph.To(compressedmjlog.Table, compressedmjlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, mjlogfile.CompressedMjlogFilesTable, mjlogfile.CompressedMjlogFilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mlfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMjlogs chains the current query on the "mjlogs" edge.
func (mlfq *MJLogFileQuery) QueryMjlogs() *MJLogQuery {
	query := (&MJLogClient{config: mlfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mlfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mlfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mjlogfile.Table, mjlogfile.FieldID, selector),
			sqlgraph.To(mjlog.Table, mjlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, mjlogfile.MjlogsTable, mjlogfile.MjlogsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mlfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MJLogFile entity from the query.
// Returns a *NotFoundError when no MJLogFile was found.
func (mlfq *MJLogFileQuery) First(ctx context.Context) (*MJLogFile, error) {
	nodes, err := mlfq.Limit(1).All(newQueryContext(ctx, TypeMJLogFile, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mjlogfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mlfq *MJLogFileQuery) FirstX(ctx context.Context) *MJLogFile {
	node, err := mlfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MJLogFile ID from the query.
// Returns a *NotFoundError when no MJLogFile ID was found.
func (mlfq *MJLogFileQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mlfq.Limit(1).IDs(newQueryContext(ctx, TypeMJLogFile, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mjlogfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mlfq *MJLogFileQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mlfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MJLogFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MJLogFile entity is found.
// Returns a *NotFoundError when no MJLogFile entities are found.
func (mlfq *MJLogFileQuery) Only(ctx context.Context) (*MJLogFile, error) {
	nodes, err := mlfq.Limit(2).All(newQueryContext(ctx, TypeMJLogFile, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mjlogfile.Label}
	default:
		return nil, &NotSingularError{mjlogfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mlfq *MJLogFileQuery) OnlyX(ctx context.Context) *MJLogFile {
	node, err := mlfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MJLogFile ID in the query.
// Returns a *NotSingularError when more than one MJLogFile ID is found.
// Returns a *NotFoundError when no entities are found.
func (mlfq *MJLogFileQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mlfq.Limit(2).IDs(newQueryContext(ctx, TypeMJLogFile, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mjlogfile.Label}
	default:
		err = &NotSingularError{mjlogfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mlfq *MJLogFileQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mlfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MJLogFiles.
func (mlfq *MJLogFileQuery) All(ctx context.Context) ([]*MJLogFile, error) {
	ctx = newQueryContext(ctx, TypeMJLogFile, "All")
	if err := mlfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MJLogFile, *MJLogFileQuery]()
	return withInterceptors[[]*MJLogFile](ctx, mlfq, qr, mlfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mlfq *MJLogFileQuery) AllX(ctx context.Context) []*MJLogFile {
	nodes, err := mlfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MJLogFile IDs.
func (mlfq *MJLogFileQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeMJLogFile, "IDs")
	if err := mlfq.Select(mjlogfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mlfq *MJLogFileQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mlfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mlfq *MJLogFileQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeMJLogFile, "Count")
	if err := mlfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mlfq, querierCount[*MJLogFileQuery](), mlfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mlfq *MJLogFileQuery) CountX(ctx context.Context) int {
	count, err := mlfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mlfq *MJLogFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeMJLogFile, "Exist")
	switch _, err := mlfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mlfq *MJLogFileQuery) ExistX(ctx context.Context) bool {
	exist, err := mlfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MJLogFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mlfq *MJLogFileQuery) Clone() *MJLogFileQuery {
	if mlfq == nil {
		return nil
	}
	return &MJLogFileQuery{
		config:                   mlfq.config,
		limit:                    mlfq.limit,
		offset:                   mlfq.offset,
		order:                    append([]OrderFunc{}, mlfq.order...),
		inters:                   append([]Interceptor{}, mlfq.inters...),
		predicates:               append([]predicate.MJLogFile{}, mlfq.predicates...),
		withCompressedMjlogFiles: mlfq.withCompressedMjlogFiles.Clone(),
		withMjlogs:               mlfq.withMjlogs.Clone(),
		// clone intermediate query.
		sql:    mlfq.sql.Clone(),
		path:   mlfq.path,
		unique: mlfq.unique,
	}
}

// WithCompressedMjlogFiles tells the query-builder to eager-load the nodes that are connected to
// the "compressed_mjlog_files" edge. The optional arguments are used to configure the query builder of the edge.
func (mlfq *MJLogFileQuery) WithCompressedMjlogFiles(opts ...func(*CompressedMJLogQuery)) *MJLogFileQuery {
	query := (&CompressedMJLogClient{config: mlfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mlfq.withCompressedMjlogFiles = query
	return mlfq
}

// WithMjlogs tells the query-builder to eager-load the nodes that are connected to
// the "mjlogs" edge. The optional arguments are used to configure the query builder of the edge.
func (mlfq *MJLogFileQuery) WithMjlogs(opts ...func(*MJLogQuery)) *MJLogFileQuery {
	query := (&MJLogClient{config: mlfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mlfq.withMjlogs = query
	return mlfq
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
//	client.MJLogFile.Query().
//		GroupBy(mjlogfile.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mlfq *MJLogFileQuery) GroupBy(field string, fields ...string) *MJLogFileGroupBy {
	mlfq.fields = append([]string{field}, fields...)
	grbuild := &MJLogFileGroupBy{build: mlfq}
	grbuild.flds = &mlfq.fields
	grbuild.label = mjlogfile.Label
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
//	client.MJLogFile.Query().
//		Select(mjlogfile.FieldName).
//		Scan(ctx, &v)
func (mlfq *MJLogFileQuery) Select(fields ...string) *MJLogFileSelect {
	mlfq.fields = append(mlfq.fields, fields...)
	sbuild := &MJLogFileSelect{MJLogFileQuery: mlfq}
	sbuild.label = mjlogfile.Label
	sbuild.flds, sbuild.scan = &mlfq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MJLogFileSelect configured with the given aggregations.
func (mlfq *MJLogFileQuery) Aggregate(fns ...AggregateFunc) *MJLogFileSelect {
	return mlfq.Select().Aggregate(fns...)
}

func (mlfq *MJLogFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mlfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mlfq); err != nil {
				return err
			}
		}
	}
	for _, f := range mlfq.fields {
		if !mjlogfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mlfq.path != nil {
		prev, err := mlfq.path(ctx)
		if err != nil {
			return err
		}
		mlfq.sql = prev
	}
	return nil
}

func (mlfq *MJLogFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MJLogFile, error) {
	var (
		nodes       = []*MJLogFile{}
		withFKs     = mlfq.withFKs
		_spec       = mlfq.querySpec()
		loadedTypes = [2]bool{
			mlfq.withCompressedMjlogFiles != nil,
			mlfq.withMjlogs != nil,
		}
	)
	if mlfq.withCompressedMjlogFiles != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, mjlogfile.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MJLogFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MJLogFile{config: mlfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mlfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mlfq.withCompressedMjlogFiles; query != nil {
		if err := mlfq.loadCompressedMjlogFiles(ctx, query, nodes, nil,
			func(n *MJLogFile, e *CompressedMJLog) { n.Edges.CompressedMjlogFiles = e }); err != nil {
			return nil, err
		}
	}
	if query := mlfq.withMjlogs; query != nil {
		if err := mlfq.loadMjlogs(ctx, query, nodes, nil,
			func(n *MJLogFile, e *MJLog) { n.Edges.Mjlogs = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mlfq *MJLogFileQuery) loadCompressedMjlogFiles(ctx context.Context, query *CompressedMJLogQuery, nodes []*MJLogFile, init func(*MJLogFile), assign func(*MJLogFile, *CompressedMJLog)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*MJLogFile)
	for i := range nodes {
		if nodes[i].compressed_mj_log_mjlog_files == nil {
			continue
		}
		fk := *nodes[i].compressed_mj_log_mjlog_files
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(compressedmjlog.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "compressed_mj_log_mjlog_files" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mlfq *MJLogFileQuery) loadMjlogs(ctx context.Context, query *MJLogQuery, nodes []*MJLogFile, init func(*MJLogFile), assign func(*MJLogFile, *MJLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*MJLogFile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.InValues(mjlogfile.MjlogsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.mj_log_file_mjlogs
		if fk == nil {
			return fmt.Errorf(`foreign-key "mj_log_file_mjlogs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mj_log_file_mjlogs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mlfq *MJLogFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mlfq.querySpec()
	_spec.Node.Columns = mlfq.fields
	if len(mlfq.fields) > 0 {
		_spec.Unique = mlfq.unique != nil && *mlfq.unique
	}
	return sqlgraph.CountNodes(ctx, mlfq.driver, _spec)
}

func (mlfq *MJLogFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mjlogfile.Table,
			Columns: mjlogfile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mjlogfile.FieldID,
			},
		},
		From:   mlfq.sql,
		Unique: true,
	}
	if unique := mlfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mlfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mjlogfile.FieldID)
		for i := range fields {
			if fields[i] != mjlogfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mlfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mlfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mlfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mlfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mlfq *MJLogFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mlfq.driver.Dialect())
	t1 := builder.Table(mjlogfile.Table)
	columns := mlfq.fields
	if len(columns) == 0 {
		columns = mjlogfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mlfq.sql != nil {
		selector = mlfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mlfq.unique != nil && *mlfq.unique {
		selector.Distinct()
	}
	for _, p := range mlfq.predicates {
		p(selector)
	}
	for _, p := range mlfq.order {
		p(selector)
	}
	if offset := mlfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mlfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MJLogFileGroupBy is the group-by builder for MJLogFile entities.
type MJLogFileGroupBy struct {
	selector
	build *MJLogFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mlfgb *MJLogFileGroupBy) Aggregate(fns ...AggregateFunc) *MJLogFileGroupBy {
	mlfgb.fns = append(mlfgb.fns, fns...)
	return mlfgb
}

// Scan applies the selector query and scans the result into the given value.
func (mlfgb *MJLogFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeMJLogFile, "GroupBy")
	if err := mlfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MJLogFileQuery, *MJLogFileGroupBy](ctx, mlfgb.build, mlfgb, mlfgb.build.inters, v)
}

func (mlfgb *MJLogFileGroupBy) sqlScan(ctx context.Context, root *MJLogFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mlfgb.fns))
	for _, fn := range mlfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mlfgb.flds)+len(mlfgb.fns))
		for _, f := range *mlfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mlfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mlfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MJLogFileSelect is the builder for selecting fields of MJLogFile entities.
type MJLogFileSelect struct {
	*MJLogFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mlfs *MJLogFileSelect) Aggregate(fns ...AggregateFunc) *MJLogFileSelect {
	mlfs.fns = append(mlfs.fns, fns...)
	return mlfs
}

// Scan applies the selector query and scans the result into the given value.
func (mlfs *MJLogFileSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeMJLogFile, "Select")
	if err := mlfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MJLogFileQuery, *MJLogFileSelect](ctx, mlfs.MJLogFileQuery, mlfs, mlfs.inters, v)
}

func (mlfs *MJLogFileSelect) sqlScan(ctx context.Context, root *MJLogFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mlfs.fns))
	for _, fn := range mlfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mlfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mlfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
