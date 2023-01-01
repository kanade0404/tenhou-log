// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayerpoint"
)

// GamePlayerPointCreate is the builder for creating a GamePlayerPoint entity.
type GamePlayerPointCreate struct {
	config
	mutation *GamePlayerPointMutation
	hooks    []Hook
}

// Mutation returns the GamePlayerPointMutation object of the builder.
func (gppc *GamePlayerPointCreate) Mutation() *GamePlayerPointMutation {
	return gppc.mutation
}

// Save creates the GamePlayerPoint in the database.
func (gppc *GamePlayerPointCreate) Save(ctx context.Context) (*GamePlayerPoint, error) {
	var (
		err  error
		node *GamePlayerPoint
	)
	if len(gppc.hooks) == 0 {
		if err = gppc.check(); err != nil {
			return nil, err
		}
		node, err = gppc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GamePlayerPointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gppc.check(); err != nil {
				return nil, err
			}
			gppc.mutation = mutation
			if node, err = gppc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gppc.hooks) - 1; i >= 0; i-- {
			if gppc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gppc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gppc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GamePlayerPoint)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GamePlayerPointMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gppc *GamePlayerPointCreate) SaveX(ctx context.Context) *GamePlayerPoint {
	v, err := gppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gppc *GamePlayerPointCreate) Exec(ctx context.Context) error {
	_, err := gppc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gppc *GamePlayerPointCreate) ExecX(ctx context.Context) {
	if err := gppc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gppc *GamePlayerPointCreate) check() error {
	return nil
}

func (gppc *GamePlayerPointCreate) sqlSave(ctx context.Context) (*GamePlayerPoint, error) {
	_node, _spec := gppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gppc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gppc *GamePlayerPointCreate) createSpec() (*GamePlayerPoint, *sqlgraph.CreateSpec) {
	var (
		_node = &GamePlayerPoint{config: gppc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gameplayerpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gameplayerpoint.FieldID,
			},
		}
	)
	return _node, _spec
}

// GamePlayerPointCreateBulk is the builder for creating many GamePlayerPoint entities in bulk.
type GamePlayerPointCreateBulk struct {
	config
	builders []*GamePlayerPointCreate
}

// Save creates the GamePlayerPoint entities in the database.
func (gppcb *GamePlayerPointCreateBulk) Save(ctx context.Context) ([]*GamePlayerPoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gppcb.builders))
	nodes := make([]*GamePlayerPoint, len(gppcb.builders))
	mutators := make([]Mutator, len(gppcb.builders))
	for i := range gppcb.builders {
		func(i int, root context.Context) {
			builder := gppcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GamePlayerPointMutation)
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
					_, err = mutators[i+1].Mutate(root, gppcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gppcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, gppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gppcb *GamePlayerPointCreateBulk) SaveX(ctx context.Context) []*GamePlayerPoint {
	v, err := gppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gppcb *GamePlayerPointCreateBulk) Exec(ctx context.Context) error {
	_, err := gppcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gppcb *GamePlayerPointCreateBulk) ExecX(ctx context.Context) {
	if err := gppcb.Exec(ctx); err != nil {
		panic(err)
	}
}
