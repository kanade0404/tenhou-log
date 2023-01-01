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
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
	"github.com/kanade0404/tenhou-log/services/ent/room"
)

// RoomQuery is the builder for querying Room entities.
type RoomQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Room
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoomQuery builder.
func (rq *RoomQuery) Where(ps ...predicate.Room) *RoomQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RoomQuery) Limit(limit int) *RoomQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RoomQuery) Offset(offset int) *RoomQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RoomQuery) Unique(unique bool) *RoomQuery {
	rq.unique = &unique
	return rq
}

// Order adds an order step to the query.
func (rq *RoomQuery) Order(o ...OrderFunc) *RoomQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// First returns the first Room entity from the query.
// Returns a *NotFoundError when no Room was found.
func (rq *RoomQuery) First(ctx context.Context) (*Room, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{room.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoomQuery) FirstX(ctx context.Context) *Room {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Room ID from the query.
// Returns a *NotFoundError when no Room ID was found.
func (rq *RoomQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{room.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RoomQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Room entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Room entity is found.
// Returns a *NotFoundError when no Room entities are found.
func (rq *RoomQuery) Only(ctx context.Context) (*Room, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{room.Label}
	default:
		return nil, &NotSingularError{room.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoomQuery) OnlyX(ctx context.Context) *Room {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Room ID in the query.
// Returns a *NotSingularError when more than one Room ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RoomQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{room.Label}
	default:
		err = &NotSingularError{room.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoomQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Rooms.
func (rq *RoomQuery) All(ctx context.Context) ([]*Room, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoomQuery) AllX(ctx context.Context) []*Room {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Room IDs.
func (rq *RoomQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rq.Select(room.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoomQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoomQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoomQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoomQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoomQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoomQuery) Clone() *RoomQuery {
	if rq == nil {
		return nil
	}
	return &RoomQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		predicates: append([]predicate.Room{}, rq.predicates...),
		// clone intermediate query.
		sql:    rq.sql.Clone(),
		path:   rq.path,
		unique: rq.unique,
	}
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
//	client.Room.Query().
//		GroupBy(room.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RoomQuery) GroupBy(field string, fields ...string) *RoomGroupBy {
	grbuild := &RoomGroupBy{config: rq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(ctx), nil
	}
	grbuild.label = room.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.Room.Query().
//		Select(room.FieldName).
//		Scan(ctx, &v)
//
func (rq *RoomQuery) Select(fields ...string) *RoomSelect {
	rq.fields = append(rq.fields, fields...)
	selbuild := &RoomSelect{RoomQuery: rq}
	selbuild.label = room.Label
	selbuild.flds, selbuild.scan = &rq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a RoomSelect configured with the given aggregations.
func (rq *RoomQuery) Aggregate(fns ...AggregateFunc) *RoomSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RoomQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rq.fields {
		if !room.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RoomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Room, error) {
	var (
		nodes = []*Room{}
		_spec = rq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Room).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Room{config: rq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rq *RoomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.fields
	if len(rq.fields) > 0 {
		_spec.Unique = rq.unique != nil && *rq.unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoomQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (rq *RoomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: room.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, room.FieldID)
		for i := range fields {
			if fields[i] != room.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RoomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(room.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = room.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.unique != nil && *rq.unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoomGroupBy is the group-by builder for Room entities.
type RoomGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoomGroupBy) Aggregate(fns ...AggregateFunc) *RoomGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rgb *RoomGroupBy) Scan(ctx context.Context, v any) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

func (rgb *RoomGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range rgb.fields {
		if !room.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RoomGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql.Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
		for _, f := range rgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rgb.fields...)...)
}

// RoomSelect is the builder for selecting fields of Room entities.
type RoomSelect struct {
	*RoomQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RoomSelect) Aggregate(fns ...AggregateFunc) *RoomSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RoomSelect) Scan(ctx context.Context, v any) error {
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	rs.sql = rs.RoomQuery.sqlQuery(ctx)
	return rs.sqlScan(ctx, v)
}

func (rs *RoomSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(rs.sql))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		rs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		rs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := rs.sql.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
