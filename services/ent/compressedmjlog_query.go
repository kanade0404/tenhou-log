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
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// CompressedMJLogQuery is the builder for querying CompressedMJLog entities.
type CompressedMJLogQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	inters         []Interceptor
	predicates     []predicate.CompressedMJLog
	withMjlogFiles *MJLogFileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompressedMJLogQuery builder.
func (cmlq *CompressedMJLogQuery) Where(ps ...predicate.CompressedMJLog) *CompressedMJLogQuery {
	cmlq.predicates = append(cmlq.predicates, ps...)
	return cmlq
}

// Limit the number of records to be returned by this query.
func (cmlq *CompressedMJLogQuery) Limit(limit int) *CompressedMJLogQuery {
	cmlq.limit = &limit
	return cmlq
}

// Offset to start from.
func (cmlq *CompressedMJLogQuery) Offset(offset int) *CompressedMJLogQuery {
	cmlq.offset = &offset
	return cmlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cmlq *CompressedMJLogQuery) Unique(unique bool) *CompressedMJLogQuery {
	cmlq.unique = &unique
	return cmlq
}

// Order specifies how the records should be ordered.
func (cmlq *CompressedMJLogQuery) Order(o ...OrderFunc) *CompressedMJLogQuery {
	cmlq.order = append(cmlq.order, o...)
	return cmlq
}

// QueryMjlogFiles chains the current query on the "mjlog_files" edge.
func (cmlq *CompressedMJLogQuery) QueryMjlogFiles() *MJLogFileQuery {
	query := (&MJLogFileClient{config: cmlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cmlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cmlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(compressedmjlog.Table, compressedmjlog.FieldID, selector),
			sqlgraph.To(mjlogfile.Table, mjlogfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, compressedmjlog.MjlogFilesTable, compressedmjlog.MjlogFilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cmlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CompressedMJLog entity from the query.
// Returns a *NotFoundError when no CompressedMJLog was found.
func (cmlq *CompressedMJLogQuery) First(ctx context.Context) (*CompressedMJLog, error) {
	nodes, err := cmlq.Limit(1).All(newQueryContext(ctx, TypeCompressedMJLog, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{compressedmjlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) FirstX(ctx context.Context) *CompressedMJLog {
	node, err := cmlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CompressedMJLog ID from the query.
// Returns a *NotFoundError when no CompressedMJLog ID was found.
func (cmlq *CompressedMJLogQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cmlq.Limit(1).IDs(newQueryContext(ctx, TypeCompressedMJLog, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{compressedmjlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cmlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CompressedMJLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CompressedMJLog entity is found.
// Returns a *NotFoundError when no CompressedMJLog entities are found.
func (cmlq *CompressedMJLogQuery) Only(ctx context.Context) (*CompressedMJLog, error) {
	nodes, err := cmlq.Limit(2).All(newQueryContext(ctx, TypeCompressedMJLog, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{compressedmjlog.Label}
	default:
		return nil, &NotSingularError{compressedmjlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) OnlyX(ctx context.Context) *CompressedMJLog {
	node, err := cmlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CompressedMJLog ID in the query.
// Returns a *NotSingularError when more than one CompressedMJLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (cmlq *CompressedMJLogQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cmlq.Limit(2).IDs(newQueryContext(ctx, TypeCompressedMJLog, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{compressedmjlog.Label}
	default:
		err = &NotSingularError{compressedmjlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cmlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CompressedMJLogs.
func (cmlq *CompressedMJLogQuery) All(ctx context.Context) ([]*CompressedMJLog, error) {
	ctx = newQueryContext(ctx, TypeCompressedMJLog, "All")
	if err := cmlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CompressedMJLog, *CompressedMJLogQuery]()
	return withInterceptors[[]*CompressedMJLog](ctx, cmlq, qr, cmlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) AllX(ctx context.Context) []*CompressedMJLog {
	nodes, err := cmlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CompressedMJLog IDs.
func (cmlq *CompressedMJLogQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeCompressedMJLog, "IDs")
	if err := cmlq.Select(compressedmjlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cmlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cmlq *CompressedMJLogQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeCompressedMJLog, "Count")
	if err := cmlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cmlq, querierCount[*CompressedMJLogQuery](), cmlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) CountX(ctx context.Context) int {
	count, err := cmlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cmlq *CompressedMJLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeCompressedMJLog, "Exist")
	switch _, err := cmlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cmlq *CompressedMJLogQuery) ExistX(ctx context.Context) bool {
	exist, err := cmlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompressedMJLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cmlq *CompressedMJLogQuery) Clone() *CompressedMJLogQuery {
	if cmlq == nil {
		return nil
	}
	return &CompressedMJLogQuery{
		config:         cmlq.config,
		limit:          cmlq.limit,
		offset:         cmlq.offset,
		order:          append([]OrderFunc{}, cmlq.order...),
		inters:         append([]Interceptor{}, cmlq.inters...),
		predicates:     append([]predicate.CompressedMJLog{}, cmlq.predicates...),
		withMjlogFiles: cmlq.withMjlogFiles.Clone(),
		// clone intermediate query.
		sql:    cmlq.sql.Clone(),
		path:   cmlq.path,
		unique: cmlq.unique,
	}
}

// WithMjlogFiles tells the query-builder to eager-load the nodes that are connected to
// the "mjlog_files" edge. The optional arguments are used to configure the query builder of the edge.
func (cmlq *CompressedMJLogQuery) WithMjlogFiles(opts ...func(*MJLogFileQuery)) *CompressedMJLogQuery {
	query := (&MJLogFileClient{config: cmlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cmlq.withMjlogFiles = query
	return cmlq
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
//	client.CompressedMJLog.Query().
//		GroupBy(compressedmjlog.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cmlq *CompressedMJLogQuery) GroupBy(field string, fields ...string) *CompressedMJLogGroupBy {
	cmlq.fields = append([]string{field}, fields...)
	grbuild := &CompressedMJLogGroupBy{build: cmlq}
	grbuild.flds = &cmlq.fields
	grbuild.label = compressedmjlog.Label
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
//	client.CompressedMJLog.Query().
//		Select(compressedmjlog.FieldName).
//		Scan(ctx, &v)
func (cmlq *CompressedMJLogQuery) Select(fields ...string) *CompressedMJLogSelect {
	cmlq.fields = append(cmlq.fields, fields...)
	sbuild := &CompressedMJLogSelect{CompressedMJLogQuery: cmlq}
	sbuild.label = compressedmjlog.Label
	sbuild.flds, sbuild.scan = &cmlq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompressedMJLogSelect configured with the given aggregations.
func (cmlq *CompressedMJLogQuery) Aggregate(fns ...AggregateFunc) *CompressedMJLogSelect {
	return cmlq.Select().Aggregate(fns...)
}

func (cmlq *CompressedMJLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cmlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cmlq); err != nil {
				return err
			}
		}
	}
	for _, f := range cmlq.fields {
		if !compressedmjlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cmlq.path != nil {
		prev, err := cmlq.path(ctx)
		if err != nil {
			return err
		}
		cmlq.sql = prev
	}
	return nil
}

func (cmlq *CompressedMJLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CompressedMJLog, error) {
	var (
		nodes       = []*CompressedMJLog{}
		_spec       = cmlq.querySpec()
		loadedTypes = [1]bool{
			cmlq.withMjlogFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CompressedMJLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CompressedMJLog{config: cmlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cmlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cmlq.withMjlogFiles; query != nil {
		if err := cmlq.loadMjlogFiles(ctx, query, nodes, nil,
			func(n *CompressedMJLog, e *MJLogFile) { n.Edges.MjlogFiles = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cmlq *CompressedMJLogQuery) loadMjlogFiles(ctx context.Context, query *MJLogFileQuery, nodes []*CompressedMJLog, init func(*CompressedMJLog), assign func(*CompressedMJLog, *MJLogFile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*CompressedMJLog)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.MJLogFile(func(s *sql.Selector) {
		s.Where(sql.InValues(compressedmjlog.MjlogFilesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.compressed_mj_log_mjlog_files
		if fk == nil {
			return fmt.Errorf(`foreign-key "compressed_mj_log_mjlog_files" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "compressed_mj_log_mjlog_files" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cmlq *CompressedMJLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cmlq.querySpec()
	_spec.Node.Columns = cmlq.fields
	if len(cmlq.fields) > 0 {
		_spec.Unique = cmlq.unique != nil && *cmlq.unique
	}
	return sqlgraph.CountNodes(ctx, cmlq.driver, _spec)
}

func (cmlq *CompressedMJLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   compressedmjlog.Table,
			Columns: compressedmjlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compressedmjlog.FieldID,
			},
		},
		From:   cmlq.sql,
		Unique: true,
	}
	if unique := cmlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cmlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, compressedmjlog.FieldID)
		for i := range fields {
			if fields[i] != compressedmjlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cmlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cmlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cmlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cmlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cmlq *CompressedMJLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cmlq.driver.Dialect())
	t1 := builder.Table(compressedmjlog.Table)
	columns := cmlq.fields
	if len(columns) == 0 {
		columns = compressedmjlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cmlq.sql != nil {
		selector = cmlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cmlq.unique != nil && *cmlq.unique {
		selector.Distinct()
	}
	for _, p := range cmlq.predicates {
		p(selector)
	}
	for _, p := range cmlq.order {
		p(selector)
	}
	if offset := cmlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cmlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CompressedMJLogGroupBy is the group-by builder for CompressedMJLog entities.
type CompressedMJLogGroupBy struct {
	selector
	build *CompressedMJLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cmlgb *CompressedMJLogGroupBy) Aggregate(fns ...AggregateFunc) *CompressedMJLogGroupBy {
	cmlgb.fns = append(cmlgb.fns, fns...)
	return cmlgb
}

// Scan applies the selector query and scans the result into the given value.
func (cmlgb *CompressedMJLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeCompressedMJLog, "GroupBy")
	if err := cmlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompressedMJLogQuery, *CompressedMJLogGroupBy](ctx, cmlgb.build, cmlgb, cmlgb.build.inters, v)
}

func (cmlgb *CompressedMJLogGroupBy) sqlScan(ctx context.Context, root *CompressedMJLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cmlgb.fns))
	for _, fn := range cmlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cmlgb.flds)+len(cmlgb.fns))
		for _, f := range *cmlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cmlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompressedMJLogSelect is the builder for selecting fields of CompressedMJLog entities.
type CompressedMJLogSelect struct {
	*CompressedMJLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cmls *CompressedMJLogSelect) Aggregate(fns ...AggregateFunc) *CompressedMJLogSelect {
	cmls.fns = append(cmls.fns, fns...)
	return cmls
}

// Scan applies the selector query and scans the result into the given value.
func (cmls *CompressedMJLogSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeCompressedMJLog, "Select")
	if err := cmls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompressedMJLogQuery, *CompressedMJLogSelect](ctx, cmls.CompressedMJLogQuery, cmls, cmls.inters, v)
}

func (cmls *CompressedMJLogSelect) sqlScan(ctx context.Context, root *CompressedMJLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cmls.fns))
	for _, fn := range cmls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cmls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
