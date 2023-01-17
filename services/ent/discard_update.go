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
	"github.com/kanade0404/tenhou-log/services/ent/discard"
	"github.com/kanade0404/tenhou-log/services/ent/drawn"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
	"github.com/kanade0404/tenhou-log/services/ent/reach"
)

// DiscardUpdate is the builder for updating Discard entities.
type DiscardUpdate struct {
	config
	hooks    []Hook
	mutation *DiscardMutation
}

// Where appends a list predicates to the DiscardUpdate builder.
func (du *DiscardUpdate) Where(ps ...predicate.Discard) *DiscardUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetReachID sets the "reach" edge to the Reach entity by ID.
func (du *DiscardUpdate) SetReachID(id uuid.UUID) *DiscardUpdate {
	du.mutation.SetReachID(id)
	return du
}

// SetNillableReachID sets the "reach" edge to the Reach entity by ID if the given value is not nil.
func (du *DiscardUpdate) SetNillableReachID(id *uuid.UUID) *DiscardUpdate {
	if id != nil {
		du = du.SetReachID(*id)
	}
	return du
}

// SetReach sets the "reach" edge to the Reach entity.
func (du *DiscardUpdate) SetReach(r *Reach) *DiscardUpdate {
	return du.SetReachID(r.ID)
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (du *DiscardUpdate) SetCallID(id uuid.UUID) *DiscardUpdate {
	du.mutation.SetCallID(id)
	return du
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (du *DiscardUpdate) SetNillableCallID(id *uuid.UUID) *DiscardUpdate {
	if id != nil {
		du = du.SetCallID(*id)
	}
	return du
}

// SetCall sets the "call" edge to the Call entity.
func (du *DiscardUpdate) SetCall(c *Call) *DiscardUpdate {
	return du.SetCallID(c.ID)
}

// SetDrawID sets the "draw" edge to the Drawn entity by ID.
func (du *DiscardUpdate) SetDrawID(id uuid.UUID) *DiscardUpdate {
	du.mutation.SetDrawID(id)
	return du
}

// SetNillableDrawID sets the "draw" edge to the Drawn entity by ID if the given value is not nil.
func (du *DiscardUpdate) SetNillableDrawID(id *uuid.UUID) *DiscardUpdate {
	if id != nil {
		du = du.SetDrawID(*id)
	}
	return du
}

// SetDraw sets the "draw" edge to the Drawn entity.
func (du *DiscardUpdate) SetDraw(d *Drawn) *DiscardUpdate {
	return du.SetDrawID(d.ID)
}

// Mutation returns the DiscardMutation object of the builder.
func (du *DiscardUpdate) Mutation() *DiscardMutation {
	return du.mutation
}

// ClearReach clears the "reach" edge to the Reach entity.
func (du *DiscardUpdate) ClearReach() *DiscardUpdate {
	du.mutation.ClearReach()
	return du
}

// ClearCall clears the "call" edge to the Call entity.
func (du *DiscardUpdate) ClearCall() *DiscardUpdate {
	du.mutation.ClearCall()
	return du
}

// ClearDraw clears the "draw" edge to the Drawn entity.
func (du *DiscardUpdate) ClearDraw() *DiscardUpdate {
	du.mutation.ClearDraw()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DiscardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, DiscardMutation](ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DiscardUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DiscardUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DiscardUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DiscardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discard.Table,
			Columns: discard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: discard.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if du.mutation.ReachCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.ReachTable,
			Columns: []string{discard.ReachColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: reach.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ReachIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.ReachTable,
			Columns: []string{discard.ReachColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: reach.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.CallCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.CallTable,
			Columns: []string{discard.CallColumn},
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
	if nodes := du.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.CallTable,
			Columns: []string{discard.CallColumn},
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
	if du.mutation.DrawCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.DrawTable,
			Columns: []string{discard.DrawColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: drawn.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DrawIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.DrawTable,
			Columns: []string{discard.DrawColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: drawn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DiscardUpdateOne is the builder for updating a single Discard entity.
type DiscardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DiscardMutation
}

// SetReachID sets the "reach" edge to the Reach entity by ID.
func (duo *DiscardUpdateOne) SetReachID(id uuid.UUID) *DiscardUpdateOne {
	duo.mutation.SetReachID(id)
	return duo
}

// SetNillableReachID sets the "reach" edge to the Reach entity by ID if the given value is not nil.
func (duo *DiscardUpdateOne) SetNillableReachID(id *uuid.UUID) *DiscardUpdateOne {
	if id != nil {
		duo = duo.SetReachID(*id)
	}
	return duo
}

// SetReach sets the "reach" edge to the Reach entity.
func (duo *DiscardUpdateOne) SetReach(r *Reach) *DiscardUpdateOne {
	return duo.SetReachID(r.ID)
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (duo *DiscardUpdateOne) SetCallID(id uuid.UUID) *DiscardUpdateOne {
	duo.mutation.SetCallID(id)
	return duo
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (duo *DiscardUpdateOne) SetNillableCallID(id *uuid.UUID) *DiscardUpdateOne {
	if id != nil {
		duo = duo.SetCallID(*id)
	}
	return duo
}

// SetCall sets the "call" edge to the Call entity.
func (duo *DiscardUpdateOne) SetCall(c *Call) *DiscardUpdateOne {
	return duo.SetCallID(c.ID)
}

// SetDrawID sets the "draw" edge to the Drawn entity by ID.
func (duo *DiscardUpdateOne) SetDrawID(id uuid.UUID) *DiscardUpdateOne {
	duo.mutation.SetDrawID(id)
	return duo
}

// SetNillableDrawID sets the "draw" edge to the Drawn entity by ID if the given value is not nil.
func (duo *DiscardUpdateOne) SetNillableDrawID(id *uuid.UUID) *DiscardUpdateOne {
	if id != nil {
		duo = duo.SetDrawID(*id)
	}
	return duo
}

// SetDraw sets the "draw" edge to the Drawn entity.
func (duo *DiscardUpdateOne) SetDraw(d *Drawn) *DiscardUpdateOne {
	return duo.SetDrawID(d.ID)
}

// Mutation returns the DiscardMutation object of the builder.
func (duo *DiscardUpdateOne) Mutation() *DiscardMutation {
	return duo.mutation
}

// ClearReach clears the "reach" edge to the Reach entity.
func (duo *DiscardUpdateOne) ClearReach() *DiscardUpdateOne {
	duo.mutation.ClearReach()
	return duo
}

// ClearCall clears the "call" edge to the Call entity.
func (duo *DiscardUpdateOne) ClearCall() *DiscardUpdateOne {
	duo.mutation.ClearCall()
	return duo
}

// ClearDraw clears the "draw" edge to the Drawn entity.
func (duo *DiscardUpdateOne) ClearDraw() *DiscardUpdateOne {
	duo.mutation.ClearDraw()
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DiscardUpdateOne) Select(field string, fields ...string) *DiscardUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Discard entity.
func (duo *DiscardUpdateOne) Save(ctx context.Context) (*Discard, error) {
	return withHooks[*Discard, DiscardMutation](ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DiscardUpdateOne) SaveX(ctx context.Context) *Discard {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DiscardUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DiscardUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DiscardUpdateOne) sqlSave(ctx context.Context) (_node *Discard, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discard.Table,
			Columns: discard.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: discard.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Discard.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discard.FieldID)
		for _, f := range fields {
			if !discard.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != discard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if duo.mutation.ReachCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.ReachTable,
			Columns: []string{discard.ReachColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: reach.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ReachIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.ReachTable,
			Columns: []string{discard.ReachColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: reach.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.CallCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.CallTable,
			Columns: []string{discard.CallColumn},
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
	if nodes := duo.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.CallTable,
			Columns: []string{discard.CallColumn},
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
	if duo.mutation.DrawCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.DrawTable,
			Columns: []string{discard.DrawColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: drawn.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DrawIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   discard.DrawTable,
			Columns: []string{discard.DrawColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: drawn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Discard{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discard.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
