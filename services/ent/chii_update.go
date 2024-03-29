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
	"github.com/kanade0404/tenhou-log/services/ent/call"
	"github.com/kanade0404/tenhou-log/services/ent/chii"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ChiiUpdate is the builder for updating Chii entities.
type ChiiUpdate struct {
	config
	hooks    []Hook
	mutation *ChiiMutation
}

// Where appends a list predicates to the ChiiUpdate builder.
func (cu *ChiiUpdate) Where(ps ...predicate.Chii) *ChiiUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (cu *ChiiUpdate) SetCallID(id uuid.UUID) *ChiiUpdate {
	cu.mutation.SetCallID(id)
	return cu
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (cu *ChiiUpdate) SetNillableCallID(id *uuid.UUID) *ChiiUpdate {
	if id != nil {
		cu = cu.SetCallID(*id)
	}
	return cu
}

// SetCall sets the "call" edge to the Call entity.
func (cu *ChiiUpdate) SetCall(c *Call) *ChiiUpdate {
	return cu.SetCallID(c.ID)
}

// Mutation returns the ChiiMutation object of the builder.
func (cu *ChiiUpdate) Mutation() *ChiiMutation {
	return cu.mutation
}

// ClearCall clears the "call" edge to the Call entity.
func (cu *ChiiUpdate) ClearCall() *ChiiUpdate {
	cu.mutation.ClearCall()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChiiUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ChiiMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChiiUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChiiUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChiiUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChiiUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chii.Table,
			Columns: chii.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: chii.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.CallCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   chii.CallTable,
			Columns: []string{chii.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: call.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   chii.CallTable,
			Columns: []string{chii.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: call.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chii.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChiiUpdateOne is the builder for updating a single Chii entity.
type ChiiUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChiiMutation
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (cuo *ChiiUpdateOne) SetCallID(id uuid.UUID) *ChiiUpdateOne {
	cuo.mutation.SetCallID(id)
	return cuo
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (cuo *ChiiUpdateOne) SetNillableCallID(id *uuid.UUID) *ChiiUpdateOne {
	if id != nil {
		cuo = cuo.SetCallID(*id)
	}
	return cuo
}

// SetCall sets the "call" edge to the Call entity.
func (cuo *ChiiUpdateOne) SetCall(c *Call) *ChiiUpdateOne {
	return cuo.SetCallID(c.ID)
}

// Mutation returns the ChiiMutation object of the builder.
func (cuo *ChiiUpdateOne) Mutation() *ChiiMutation {
	return cuo.mutation
}

// ClearCall clears the "call" edge to the Call entity.
func (cuo *ChiiUpdateOne) ClearCall() *ChiiUpdateOne {
	cuo.mutation.ClearCall()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChiiUpdateOne) Select(field string, fields ...string) *ChiiUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chii entity.
func (cuo *ChiiUpdateOne) Save(ctx context.Context) (*Chii, error) {
	return withHooks[*Chii, ChiiMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChiiUpdateOne) SaveX(ctx context.Context) *Chii {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChiiUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChiiUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChiiUpdateOne) sqlSave(ctx context.Context) (_node *Chii, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chii.Table,
			Columns: chii.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: chii.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chii.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chii.FieldID)
		for _, f := range fields {
			if !chii.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chii.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.CallCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   chii.CallTable,
			Columns: []string{chii.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: call.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   chii.CallTable,
			Columns: []string{chii.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: call.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chii{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chii.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
