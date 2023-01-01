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
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfilecompressed"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// MJLogFileCompressedQuery is the builder for querying MJLogFileCompressed entities.
type MJLogFileCompressedQuery struct {
	config
	limit                    *int
	offset                   *int
	unique                   *bool
	order                    []OrderFunc
	fields                   []string
	predicates               []predicate.MJLogFileCompressed
	withCompressedMjlogFiles *CompressedMJLogQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MJLogFileCompressedQuery builder.
func (mlfcq *MJLogFileCompressedQuery) Where(ps ...predicate.MJLogFileCompressed) *MJLogFileCompressedQuery {
	mlfcq.predicates = append(mlfcq.predicates, ps...)
	return mlfcq
}

// Limit adds a limit step to the query.
func (mlfcq *MJLogFileCompressedQuery) Limit(limit int) *MJLogFileCompressedQuery {
	mlfcq.limit = &limit
	return mlfcq
}

// Offset adds an offset step to the query.
func (mlfcq *MJLogFileCompressedQuery) Offset(offset int) *MJLogFileCompressedQuery {
	mlfcq.offset = &offset
	return mlfcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mlfcq *MJLogFileCompressedQuery) Unique(unique bool) *MJLogFileCompressedQuery {
	mlfcq.unique = &unique
	return mlfcq
}

// Order adds an order step to the query.
func (mlfcq *MJLogFileCompressedQuery) Order(o ...OrderFunc) *MJLogFileCompressedQuery {
	mlfcq.order = append(mlfcq.order, o...)
	return mlfcq
}

// QueryCompressedMjlogFiles chains the current query on the "compressed_mjlog_files" edge.
func (mlfcq *MJLogFileCompressedQuery) QueryCompressedMjlogFiles() *CompressedMJLogQuery {
	query := &CompressedMJLogQuery{config: mlfcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mlfcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mlfcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mjlogfilecompressed.Table, mjlogfilecompressed.FieldID, selector),
			sqlgraph.To(compressedmjlog.Table, compressedmjlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, mjlogfilecompressed.CompressedMjlogFilesTable, mjlogfilecompressed.CompressedMjlogFilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mlfcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MJLogFileCompressed entity from the query.
// Returns a *NotFoundError when no MJLogFileCompressed was found.
func (mlfcq *MJLogFileCompressedQuery) First(ctx context.Context) (*MJLogFileCompressed, error) {
	nodes, err := mlfcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mjlogfilecompressed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) FirstX(ctx context.Context) *MJLogFileCompressed {
	node, err := mlfcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MJLogFileCompressed ID from the query.
// Returns a *NotFoundError when no MJLogFileCompressed ID was found.
func (mlfcq *MJLogFileCompressedQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mlfcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mjlogfilecompressed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mlfcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MJLogFileCompressed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MJLogFileCompressed entity is found.
// Returns a *NotFoundError when no MJLogFileCompressed entities are found.
func (mlfcq *MJLogFileCompressedQuery) Only(ctx context.Context) (*MJLogFileCompressed, error) {
	nodes, err := mlfcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mjlogfilecompressed.Label}
	default:
		return nil, &NotSingularError{mjlogfilecompressed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) OnlyX(ctx context.Context) *MJLogFileCompressed {
	node, err := mlfcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MJLogFileCompressed ID in the query.
// Returns a *NotSingularError when more than one MJLogFileCompressed ID is found.
// Returns a *NotFoundError when no entities are found.
func (mlfcq *MJLogFileCompressedQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mlfcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mjlogfilecompressed.Label}
	default:
		err = &NotSingularError{mjlogfilecompressed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mlfcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MJLogFileCompresseds.
func (mlfcq *MJLogFileCompressedQuery) All(ctx context.Context) ([]*MJLogFileCompressed, error) {
	if err := mlfcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mlfcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) AllX(ctx context.Context) []*MJLogFileCompressed {
	nodes, err := mlfcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MJLogFileCompressed IDs.
func (mlfcq *MJLogFileCompressedQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := mlfcq.Select(mjlogfilecompressed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mlfcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mlfcq *MJLogFileCompressedQuery) Count(ctx context.Context) (int, error) {
	if err := mlfcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mlfcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) CountX(ctx context.Context) int {
	count, err := mlfcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mlfcq *MJLogFileCompressedQuery) Exist(ctx context.Context) (bool, error) {
	if err := mlfcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mlfcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mlfcq *MJLogFileCompressedQuery) ExistX(ctx context.Context) bool {
	exist, err := mlfcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MJLogFileCompressedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mlfcq *MJLogFileCompressedQuery) Clone() *MJLogFileCompressedQuery {
	if mlfcq == nil {
		return nil
	}
	return &MJLogFileCompressedQuery{
		config:                   mlfcq.config,
		limit:                    mlfcq.limit,
		offset:                   mlfcq.offset,
		order:                    append([]OrderFunc{}, mlfcq.order...),
		predicates:               append([]predicate.MJLogFileCompressed{}, mlfcq.predicates...),
		withCompressedMjlogFiles: mlfcq.withCompressedMjlogFiles.Clone(),
		// clone intermediate query.
		sql:    mlfcq.sql.Clone(),
		path:   mlfcq.path,
		unique: mlfcq.unique,
	}
}

// WithCompressedMjlogFiles tells the query-builder to eager-load the nodes that are connected to
// the "compressed_mjlog_files" edge. The optional arguments are used to configure the query builder of the edge.
func (mlfcq *MJLogFileCompressedQuery) WithCompressedMjlogFiles(opts ...func(*CompressedMJLogQuery)) *MJLogFileCompressedQuery {
	query := &CompressedMJLogQuery{config: mlfcq.config}
	for _, opt := range opts {
		opt(query)
	}
	mlfcq.withCompressedMjlogFiles = query
	return mlfcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (mlfcq *MJLogFileCompressedQuery) GroupBy(field string, fields ...string) *MJLogFileCompressedGroupBy {
	grbuild := &MJLogFileCompressedGroupBy{config: mlfcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mlfcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mlfcq.sqlQuery(ctx), nil
	}
	grbuild.label = mjlogfilecompressed.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (mlfcq *MJLogFileCompressedQuery) Select(fields ...string) *MJLogFileCompressedSelect {
	mlfcq.fields = append(mlfcq.fields, fields...)
	selbuild := &MJLogFileCompressedSelect{MJLogFileCompressedQuery: mlfcq}
	selbuild.label = mjlogfilecompressed.Label
	selbuild.flds, selbuild.scan = &mlfcq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a MJLogFileCompressedSelect configured with the given aggregations.
func (mlfcq *MJLogFileCompressedQuery) Aggregate(fns ...AggregateFunc) *MJLogFileCompressedSelect {
	return mlfcq.Select().Aggregate(fns...)
}

func (mlfcq *MJLogFileCompressedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mlfcq.fields {
		if !mjlogfilecompressed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mlfcq.path != nil {
		prev, err := mlfcq.path(ctx)
		if err != nil {
			return err
		}
		mlfcq.sql = prev
	}
	return nil
}

func (mlfcq *MJLogFileCompressedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MJLogFileCompressed, error) {
	var (
		nodes       = []*MJLogFileCompressed{}
		_spec       = mlfcq.querySpec()
		loadedTypes = [1]bool{
			mlfcq.withCompressedMjlogFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MJLogFileCompressed).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MJLogFileCompressed{config: mlfcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mlfcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mlfcq.withCompressedMjlogFiles; query != nil {
		if err := mlfcq.loadCompressedMjlogFiles(ctx, query, nodes, nil,
			func(n *MJLogFileCompressed, e *CompressedMJLog) { n.Edges.CompressedMjlogFiles = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mlfcq *MJLogFileCompressedQuery) loadCompressedMjlogFiles(ctx context.Context, query *CompressedMJLogQuery, nodes []*MJLogFileCompressed, init func(*MJLogFileCompressed), assign func(*MJLogFileCompressed, *CompressedMJLog)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*MJLogFileCompressed)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.CompressedMJLog(func(s *sql.Selector) {
		s.Where(sql.InValues(mjlogfilecompressed.CompressedMjlogFilesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.mj_log_file_compressed_compressed_mjlog_files
		if fk == nil {
			return fmt.Errorf(`foreign-key "mj_log_file_compressed_compressed_mjlog_files" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "mj_log_file_compressed_compressed_mjlog_files" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mlfcq *MJLogFileCompressedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mlfcq.querySpec()
	_spec.Node.Columns = mlfcq.fields
	if len(mlfcq.fields) > 0 {
		_spec.Unique = mlfcq.unique != nil && *mlfcq.unique
	}
	return sqlgraph.CountNodes(ctx, mlfcq.driver, _spec)
}

func (mlfcq *MJLogFileCompressedQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := mlfcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (mlfcq *MJLogFileCompressedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mjlogfilecompressed.Table,
			Columns: mjlogfilecompressed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mjlogfilecompressed.FieldID,
			},
		},
		From:   mlfcq.sql,
		Unique: true,
	}
	if unique := mlfcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mlfcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mjlogfilecompressed.FieldID)
		for i := range fields {
			if fields[i] != mjlogfilecompressed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mlfcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mlfcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mlfcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mlfcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mlfcq *MJLogFileCompressedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mlfcq.driver.Dialect())
	t1 := builder.Table(mjlogfilecompressed.Table)
	columns := mlfcq.fields
	if len(columns) == 0 {
		columns = mjlogfilecompressed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mlfcq.sql != nil {
		selector = mlfcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mlfcq.unique != nil && *mlfcq.unique {
		selector.Distinct()
	}
	for _, p := range mlfcq.predicates {
		p(selector)
	}
	for _, p := range mlfcq.order {
		p(selector)
	}
	if offset := mlfcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mlfcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MJLogFileCompressedGroupBy is the group-by builder for MJLogFileCompressed entities.
type MJLogFileCompressedGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mlfcgb *MJLogFileCompressedGroupBy) Aggregate(fns ...AggregateFunc) *MJLogFileCompressedGroupBy {
	mlfcgb.fns = append(mlfcgb.fns, fns...)
	return mlfcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mlfcgb *MJLogFileCompressedGroupBy) Scan(ctx context.Context, v any) error {
	query, err := mlfcgb.path(ctx)
	if err != nil {
		return err
	}
	mlfcgb.sql = query
	return mlfcgb.sqlScan(ctx, v)
}

func (mlfcgb *MJLogFileCompressedGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range mlfcgb.fields {
		if !mjlogfilecompressed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mlfcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mlfcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mlfcgb *MJLogFileCompressedGroupBy) sqlQuery() *sql.Selector {
	selector := mlfcgb.sql.Select()
	aggregation := make([]string, 0, len(mlfcgb.fns))
	for _, fn := range mlfcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mlfcgb.fields)+len(mlfcgb.fns))
		for _, f := range mlfcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mlfcgb.fields...)...)
}

// MJLogFileCompressedSelect is the builder for selecting fields of MJLogFileCompressed entities.
type MJLogFileCompressedSelect struct {
	*MJLogFileCompressedQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mlfcs *MJLogFileCompressedSelect) Aggregate(fns ...AggregateFunc) *MJLogFileCompressedSelect {
	mlfcs.fns = append(mlfcs.fns, fns...)
	return mlfcs
}

// Scan applies the selector query and scans the result into the given value.
func (mlfcs *MJLogFileCompressedSelect) Scan(ctx context.Context, v any) error {
	if err := mlfcs.prepareQuery(ctx); err != nil {
		return err
	}
	mlfcs.sql = mlfcs.MJLogFileCompressedQuery.sqlQuery(ctx)
	return mlfcs.sqlScan(ctx, v)
}

func (mlfcs *MJLogFileCompressedSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(mlfcs.fns))
	for _, fn := range mlfcs.fns {
		aggregation = append(aggregation, fn(mlfcs.sql))
	}
	switch n := len(*mlfcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		mlfcs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		mlfcs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := mlfcs.sql.Query()
	if err := mlfcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
