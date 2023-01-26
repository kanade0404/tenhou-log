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
	"github.com/kanade0404/tenhou-log/services/ent/event"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
	"github.com/kanade0404/tenhou-log/services/ent/win"
)

// WinUpdate is the builder for updating Win entities.
type WinUpdate struct {
	config
	hooks    []Hook
	mutation *WinMutation
}

// Where appends a list predicates to the WinUpdate builder.
func (wu *WinUpdate) Where(ps ...predicate.Win) *WinUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (wu *WinUpdate) SetEventID(id uuid.UUID) *WinUpdate {
	wu.mutation.SetEventID(id)
	return wu
}

// SetEvent sets the "event" edge to the Event entity.
func (wu *WinUpdate) SetEvent(e *Event) *WinUpdate {
	return wu.SetEventID(e.ID)
}

// Mutation returns the WinMutation object of the builder.
func (wu *WinUpdate) Mutation() *WinMutation {
	return wu.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (wu *WinUpdate) ClearEvent() *WinUpdate {
	wu.mutation.ClearEvent()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WinUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, WinMutation](ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WinUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WinUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WinUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WinUpdate) check() error {
	if _, ok := wu.mutation.EventID(); wu.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Win.event"`)
	}
	return nil
}

func (wu *WinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   win.Table,
			Columns: win.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: win.FieldID,
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
	if wu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   win.EventTable,
			Columns: []string{win.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   win.EventTable,
			Columns: []string{win.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
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
			err = &NotFoundError{win.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WinUpdateOne is the builder for updating a single Win entity.
type WinUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WinMutation
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (wuo *WinUpdateOne) SetEventID(id uuid.UUID) *WinUpdateOne {
	wuo.mutation.SetEventID(id)
	return wuo
}

// SetEvent sets the "event" edge to the Event entity.
func (wuo *WinUpdateOne) SetEvent(e *Event) *WinUpdateOne {
	return wuo.SetEventID(e.ID)
}

// Mutation returns the WinMutation object of the builder.
func (wuo *WinUpdateOne) Mutation() *WinMutation {
	return wuo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (wuo *WinUpdateOne) ClearEvent() *WinUpdateOne {
	wuo.mutation.ClearEvent()
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WinUpdateOne) Select(field string, fields ...string) *WinUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Win entity.
func (wuo *WinUpdateOne) Save(ctx context.Context) (*Win, error) {
	return withHooks[*Win, WinMutation](ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WinUpdateOne) SaveX(ctx context.Context) *Win {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WinUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WinUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WinUpdateOne) check() error {
	if _, ok := wuo.mutation.EventID(); wuo.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Win.event"`)
	}
	return nil
}

func (wuo *WinUpdateOne) sqlSave(ctx context.Context) (_node *Win, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   win.Table,
			Columns: win.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: win.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Win.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, win.FieldID)
		for _, f := range fields {
			if !win.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != win.FieldID {
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
	if wuo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   win.EventTable,
			Columns: []string{win.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   win.EventTable,
			Columns: []string{win.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Win{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{win.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}