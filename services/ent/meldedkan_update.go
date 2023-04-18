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
	"github.com/kanade0404/tenhou-log/services/ent/meldedkan"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// MeldedKanUpdate is the builder for updating MeldedKan entities.
type MeldedKanUpdate struct {
	config
	hooks    []Hook
	mutation *MeldedKanMutation
}

// Where appends a list predicates to the MeldedKanUpdate builder.
func (mku *MeldedKanUpdate) Where(ps ...predicate.MeldedKan) *MeldedKanUpdate {
	mku.mutation.Where(ps...)
	return mku
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (mku *MeldedKanUpdate) SetCallID(id uuid.UUID) *MeldedKanUpdate {
	mku.mutation.SetCallID(id)
	return mku
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (mku *MeldedKanUpdate) SetNillableCallID(id *uuid.UUID) *MeldedKanUpdate {
	if id != nil {
		mku = mku.SetCallID(*id)
	}
	return mku
}

// SetCall sets the "call" edge to the Call entity.
func (mku *MeldedKanUpdate) SetCall(c *Call) *MeldedKanUpdate {
	return mku.SetCallID(c.ID)
}

// Mutation returns the MeldedKanMutation object of the builder.
func (mku *MeldedKanUpdate) Mutation() *MeldedKanMutation {
	return mku.mutation
}

// ClearCall clears the "call" edge to the Call entity.
func (mku *MeldedKanUpdate) ClearCall() *MeldedKanUpdate {
	mku.mutation.ClearCall()
	return mku
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mku *MeldedKanUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MeldedKanMutation](ctx, mku.sqlSave, mku.mutation, mku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mku *MeldedKanUpdate) SaveX(ctx context.Context) int {
	affected, err := mku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mku *MeldedKanUpdate) Exec(ctx context.Context) error {
	_, err := mku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mku *MeldedKanUpdate) ExecX(ctx context.Context) {
	if err := mku.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mku *MeldedKanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(meldedkan.Table, meldedkan.Columns, sqlgraph.NewFieldSpec(meldedkan.FieldID, field.TypeUUID))
	if ps := mku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mku.mutation.CallCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   meldedkan.CallTable,
			Columns: []string{meldedkan.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(call.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mku.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   meldedkan.CallTable,
			Columns: []string{meldedkan.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(call.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meldedkan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mku.mutation.done = true
	return n, nil
}

// MeldedKanUpdateOne is the builder for updating a single MeldedKan entity.
type MeldedKanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MeldedKanMutation
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (mkuo *MeldedKanUpdateOne) SetCallID(id uuid.UUID) *MeldedKanUpdateOne {
	mkuo.mutation.SetCallID(id)
	return mkuo
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (mkuo *MeldedKanUpdateOne) SetNillableCallID(id *uuid.UUID) *MeldedKanUpdateOne {
	if id != nil {
		mkuo = mkuo.SetCallID(*id)
	}
	return mkuo
}

// SetCall sets the "call" edge to the Call entity.
func (mkuo *MeldedKanUpdateOne) SetCall(c *Call) *MeldedKanUpdateOne {
	return mkuo.SetCallID(c.ID)
}

// Mutation returns the MeldedKanMutation object of the builder.
func (mkuo *MeldedKanUpdateOne) Mutation() *MeldedKanMutation {
	return mkuo.mutation
}

// ClearCall clears the "call" edge to the Call entity.
func (mkuo *MeldedKanUpdateOne) ClearCall() *MeldedKanUpdateOne {
	mkuo.mutation.ClearCall()
	return mkuo
}

// Where appends a list predicates to the MeldedKanUpdate builder.
func (mkuo *MeldedKanUpdateOne) Where(ps ...predicate.MeldedKan) *MeldedKanUpdateOne {
	mkuo.mutation.Where(ps...)
	return mkuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mkuo *MeldedKanUpdateOne) Select(field string, fields ...string) *MeldedKanUpdateOne {
	mkuo.fields = append([]string{field}, fields...)
	return mkuo
}

// Save executes the query and returns the updated MeldedKan entity.
func (mkuo *MeldedKanUpdateOne) Save(ctx context.Context) (*MeldedKan, error) {
	return withHooks[*MeldedKan, MeldedKanMutation](ctx, mkuo.sqlSave, mkuo.mutation, mkuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mkuo *MeldedKanUpdateOne) SaveX(ctx context.Context) *MeldedKan {
	node, err := mkuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mkuo *MeldedKanUpdateOne) Exec(ctx context.Context) error {
	_, err := mkuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mkuo *MeldedKanUpdateOne) ExecX(ctx context.Context) {
	if err := mkuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mkuo *MeldedKanUpdateOne) sqlSave(ctx context.Context) (_node *MeldedKan, err error) {
	_spec := sqlgraph.NewUpdateSpec(meldedkan.Table, meldedkan.Columns, sqlgraph.NewFieldSpec(meldedkan.FieldID, field.TypeUUID))
	id, ok := mkuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MeldedKan.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mkuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meldedkan.FieldID)
		for _, f := range fields {
			if !meldedkan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != meldedkan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mkuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mkuo.mutation.CallCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   meldedkan.CallTable,
			Columns: []string{meldedkan.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(call.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mkuo.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   meldedkan.CallTable,
			Columns: []string{meldedkan.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(call.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MeldedKan{config: mkuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mkuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meldedkan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mkuo.mutation.done = true
	return _node, nil
}
