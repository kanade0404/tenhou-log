// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/game"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: game.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	var (
		err  error
		node *Game
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GameMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Game)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GameMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   game.Table,
			Columns: game.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: game.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Game.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
