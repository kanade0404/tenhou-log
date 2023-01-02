// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayerpoint"
	"github.com/kanade0404/tenhou-log/services/ent/turn"
)

// GamePlayerPointCreate is the builder for creating a GamePlayerPoint entity.
type GamePlayerPointCreate struct {
	config
	mutation *GamePlayerPointMutation
	hooks    []Hook
}

// SetPoint sets the "point" field.
func (gppc *GamePlayerPointCreate) SetPoint(u uint) *GamePlayerPointCreate {
	gppc.mutation.SetPoint(u)
	return gppc
}

// SetID sets the "id" field.
func (gppc *GamePlayerPointCreate) SetID(u uuid.UUID) *GamePlayerPointCreate {
	gppc.mutation.SetID(u)
	return gppc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gppc *GamePlayerPointCreate) SetNillableID(u *uuid.UUID) *GamePlayerPointCreate {
	if u != nil {
		gppc.SetID(*u)
	}
	return gppc
}

// AddTurnIDs adds the "turns" edge to the Turn entity by IDs.
func (gppc *GamePlayerPointCreate) AddTurnIDs(ids ...uuid.UUID) *GamePlayerPointCreate {
	gppc.mutation.AddTurnIDs(ids...)
	return gppc
}

// AddTurns adds the "turns" edges to the Turn entity.
func (gppc *GamePlayerPointCreate) AddTurns(t ...*Turn) *GamePlayerPointCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return gppc.AddTurnIDs(ids...)
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
	gppc.defaults()
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

// defaults sets the default values of the builder before save.
func (gppc *GamePlayerPointCreate) defaults() {
	if _, ok := gppc.mutation.ID(); !ok {
		v := gameplayerpoint.DefaultID()
		gppc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gppc *GamePlayerPointCreate) check() error {
	if _, ok := gppc.mutation.Point(); !ok {
		return &ValidationError{Name: "point", err: errors.New(`ent: missing required field "GamePlayerPoint.point"`)}
	}
	if v, ok := gppc.mutation.Point(); ok {
		if err := gameplayerpoint.PointValidator(v); err != nil {
			return &ValidationError{Name: "point", err: fmt.Errorf(`ent: validator failed for field "GamePlayerPoint.point": %w`, err)}
		}
	}
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
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (gppc *GamePlayerPointCreate) createSpec() (*GamePlayerPoint, *sqlgraph.CreateSpec) {
	var (
		_node = &GamePlayerPoint{config: gppc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gameplayerpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gameplayerpoint.FieldID,
			},
		}
	)
	if id, ok := gppc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gppc.mutation.Point(); ok {
		_spec.SetField(gameplayerpoint.FieldPoint, field.TypeUint, value)
		_node.Point = value
	}
	if nodes := gppc.mutation.TurnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayerpoint.TurnsTable,
			Columns: gameplayerpoint.TurnsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: turn.FieldID,
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
			builder.defaults()
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
