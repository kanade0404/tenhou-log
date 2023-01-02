// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/hand"
	"github.com/kanade0404/tenhou-log/services/ent/turn"
)

// TurnCreate is the builder for creating a Turn entity.
type TurnCreate struct {
	config
	mutation *TurnMutation
	hooks    []Hook
}

// SetNum sets the "num" field.
func (tc *TurnCreate) SetNum(u uint) *TurnCreate {
	tc.mutation.SetNum(u)
	return tc
}

// SetID sets the "id" field.
func (tc *TurnCreate) SetID(u uuid.UUID) *TurnCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TurnCreate) SetNillableID(u *uuid.UUID) *TurnCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// AddHandIDs adds the "hands" edge to the Hand entity by IDs.
func (tc *TurnCreate) AddHandIDs(ids ...uuid.UUID) *TurnCreate {
	tc.mutation.AddHandIDs(ids...)
	return tc
}

// AddHands adds the "hands" edges to the Hand entity.
func (tc *TurnCreate) AddHands(h ...*Hand) *TurnCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tc.AddHandIDs(ids...)
}

// Mutation returns the TurnMutation object of the builder.
func (tc *TurnCreate) Mutation() *TurnMutation {
	return tc.mutation
}

// Save creates the Turn in the database.
func (tc *TurnCreate) Save(ctx context.Context) (*Turn, error) {
	var (
		err  error
		node *Turn
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TurnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Turn)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TurnMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TurnCreate) SaveX(ctx context.Context) *Turn {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TurnCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TurnCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TurnCreate) defaults() {
	if _, ok := tc.mutation.ID(); !ok {
		v := turn.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TurnCreate) check() error {
	if _, ok := tc.mutation.Num(); !ok {
		return &ValidationError{Name: "num", err: errors.New(`ent: missing required field "Turn.num"`)}
	}
	return nil
}

func (tc *TurnCreate) sqlSave(ctx context.Context) (*Turn, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TurnCreate) createSpec() (*Turn, *sqlgraph.CreateSpec) {
	var (
		_node = &Turn{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: turn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: turn.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Num(); ok {
		_spec.SetField(turn.FieldNum, field.TypeUint, value)
		_node.Num = value
	}
	if nodes := tc.mutation.HandsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   turn.HandsTable,
			Columns: turn.HandsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: hand.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TurnCreateBulk is the builder for creating many Turn entities in bulk.
type TurnCreateBulk struct {
	config
	builders []*TurnCreate
}

// Save creates the Turn entities in the database.
func (tcb *TurnCreateBulk) Save(ctx context.Context) ([]*Turn, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Turn, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TurnMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TurnCreateBulk) SaveX(ctx context.Context) []*Turn {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TurnCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TurnCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
