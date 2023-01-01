// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
	"github.com/kanade0404/tenhou-log/services/ent/round"
	"github.com/kanade0404/tenhou-log/services/ent/wind"
)

// WindUpdate is the builder for updating Wind entities.
type WindUpdate struct {
	config
	hooks    []Hook
	mutation *WindMutation
}

// Where appends a list predicates to the WindUpdate builder.
func (wu *WindUpdate) Where(ps ...predicate.Wind) *WindUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// AddRoundIDs adds the "rounds" edge to the Round entity by IDs.
func (wu *WindUpdate) AddRoundIDs(ids ...uuid.UUID) *WindUpdate {
	wu.mutation.AddRoundIDs(ids...)
	return wu
}

// AddRounds adds the "rounds" edges to the Round entity.
func (wu *WindUpdate) AddRounds(r ...*Round) *WindUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wu.AddRoundIDs(ids...)
}

// Mutation returns the WindMutation object of the builder.
func (wu *WindUpdate) Mutation() *WindMutation {
	return wu.mutation
}

// ClearRounds clears all "rounds" edges to the Round entity.
func (wu *WindUpdate) ClearRounds() *WindUpdate {
	wu.mutation.ClearRounds()
	return wu
}

// RemoveRoundIDs removes the "rounds" edge to Round entities by IDs.
func (wu *WindUpdate) RemoveRoundIDs(ids ...uuid.UUID) *WindUpdate {
	wu.mutation.RemoveRoundIDs(ids...)
	return wu
}

// RemoveRounds removes "rounds" edges to Round entities.
func (wu *WindUpdate) RemoveRounds(r ...*Round) *WindUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wu.RemoveRoundIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WindUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WindMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WindUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WindUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WindUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wu *WindUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wind.Table,
			Columns: wind.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wind.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wu.mutation.RoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wind.RoundsTable,
			Columns: []string{wind.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: round.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedRoundsIDs(); len(nodes) > 0 && !wu.mutation.RoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wind.RoundsTable,
			Columns: []string{wind.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: round.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wind.RoundsTable,
			Columns: []string{wind.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: round.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wind.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WindUpdateOne is the builder for updating a single Wind entity.
type WindUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WindMutation
}

// AddRoundIDs adds the "rounds" edge to the Round entity by IDs.
func (wuo *WindUpdateOne) AddRoundIDs(ids ...uuid.UUID) *WindUpdateOne {
	wuo.mutation.AddRoundIDs(ids...)
	return wuo
}

// AddRounds adds the "rounds" edges to the Round entity.
func (wuo *WindUpdateOne) AddRounds(r ...*Round) *WindUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wuo.AddRoundIDs(ids...)
}

// Mutation returns the WindMutation object of the builder.
func (wuo *WindUpdateOne) Mutation() *WindMutation {
	return wuo.mutation
}

// ClearRounds clears all "rounds" edges to the Round entity.
func (wuo *WindUpdateOne) ClearRounds() *WindUpdateOne {
	wuo.mutation.ClearRounds()
	return wuo
}

// RemoveRoundIDs removes the "rounds" edge to Round entities by IDs.
func (wuo *WindUpdateOne) RemoveRoundIDs(ids ...uuid.UUID) *WindUpdateOne {
	wuo.mutation.RemoveRoundIDs(ids...)
	return wuo
}

// RemoveRounds removes "rounds" edges to Round entities.
func (wuo *WindUpdateOne) RemoveRounds(r ...*Round) *WindUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wuo.RemoveRoundIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WindUpdateOne) Select(field string, fields ...string) *WindUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wind entity.
func (wuo *WindUpdateOne) Save(ctx context.Context) (*Wind, error) {
	var (
		err  error
		node *Wind
	)
	if len(wuo.hooks) == 0 {
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WindMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Wind)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WindMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WindUpdateOne) SaveX(ctx context.Context) *Wind {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WindUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WindUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wuo *WindUpdateOne) sqlSave(ctx context.Context) (_node *Wind, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wind.Table,
			Columns: wind.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wind.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wind.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wind.FieldID)
		for _, f := range fields {
			if !wind.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wind.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if wuo.mutation.RoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wind.RoundsTable,
			Columns: []string{wind.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: round.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedRoundsIDs(); len(nodes) > 0 && !wuo.mutation.RoundsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wind.RoundsTable,
			Columns: []string{wind.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: round.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wind.RoundsTable,
			Columns: []string{wind.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: round.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wind{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wind.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
