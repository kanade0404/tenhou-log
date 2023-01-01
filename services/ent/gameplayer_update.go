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
	"github.com/kanade0404/tenhou-log/services/ent/gameplayer"
	"github.com/kanade0404/tenhou-log/services/ent/player"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// GamePlayerUpdate is the builder for updating GamePlayer entities.
type GamePlayerUpdate struct {
	config
	hooks    []Hook
	mutation *GamePlayerMutation
}

// Where appends a list predicates to the GamePlayerUpdate builder.
func (gpu *GamePlayerUpdate) Where(ps ...predicate.GamePlayer) *GamePlayerUpdate {
	gpu.mutation.Where(ps...)
	return gpu
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (gpu *GamePlayerUpdate) AddPlayerIDs(ids ...uuid.UUID) *GamePlayerUpdate {
	gpu.mutation.AddPlayerIDs(ids...)
	return gpu
}

// AddPlayers adds the "players" edges to the Player entity.
func (gpu *GamePlayerUpdate) AddPlayers(p ...*Player) *GamePlayerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gpu.AddPlayerIDs(ids...)
}

// Mutation returns the GamePlayerMutation object of the builder.
func (gpu *GamePlayerUpdate) Mutation() *GamePlayerMutation {
	return gpu.mutation
}

// ClearPlayers clears all "players" edges to the Player entity.
func (gpu *GamePlayerUpdate) ClearPlayers() *GamePlayerUpdate {
	gpu.mutation.ClearPlayers()
	return gpu
}

// RemovePlayerIDs removes the "players" edge to Player entities by IDs.
func (gpu *GamePlayerUpdate) RemovePlayerIDs(ids ...uuid.UUID) *GamePlayerUpdate {
	gpu.mutation.RemovePlayerIDs(ids...)
	return gpu
}

// RemovePlayers removes "players" edges to Player entities.
func (gpu *GamePlayerUpdate) RemovePlayers(p ...*Player) *GamePlayerUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gpu.RemovePlayerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gpu *GamePlayerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gpu.hooks) == 0 {
		affected, err = gpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GamePlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gpu.mutation = mutation
			affected, err = gpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gpu.hooks) - 1; i >= 0; i-- {
			if gpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gpu *GamePlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := gpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gpu *GamePlayerUpdate) Exec(ctx context.Context) error {
	_, err := gpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpu *GamePlayerUpdate) ExecX(ctx context.Context) {
	if err := gpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gpu *GamePlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gameplayer.Table,
			Columns: gameplayer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gameplayer.FieldID,
			},
		},
	}
	if ps := gpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gpu.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: gameplayer.PlayersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: player.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpu.mutation.RemovedPlayersIDs(); len(nodes) > 0 && !gpu.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: gameplayer.PlayersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpu.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: gameplayer.PlayersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameplayer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GamePlayerUpdateOne is the builder for updating a single GamePlayer entity.
type GamePlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GamePlayerMutation
}

// AddPlayerIDs adds the "players" edge to the Player entity by IDs.
func (gpuo *GamePlayerUpdateOne) AddPlayerIDs(ids ...uuid.UUID) *GamePlayerUpdateOne {
	gpuo.mutation.AddPlayerIDs(ids...)
	return gpuo
}

// AddPlayers adds the "players" edges to the Player entity.
func (gpuo *GamePlayerUpdateOne) AddPlayers(p ...*Player) *GamePlayerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gpuo.AddPlayerIDs(ids...)
}

// Mutation returns the GamePlayerMutation object of the builder.
func (gpuo *GamePlayerUpdateOne) Mutation() *GamePlayerMutation {
	return gpuo.mutation
}

// ClearPlayers clears all "players" edges to the Player entity.
func (gpuo *GamePlayerUpdateOne) ClearPlayers() *GamePlayerUpdateOne {
	gpuo.mutation.ClearPlayers()
	return gpuo
}

// RemovePlayerIDs removes the "players" edge to Player entities by IDs.
func (gpuo *GamePlayerUpdateOne) RemovePlayerIDs(ids ...uuid.UUID) *GamePlayerUpdateOne {
	gpuo.mutation.RemovePlayerIDs(ids...)
	return gpuo
}

// RemovePlayers removes "players" edges to Player entities.
func (gpuo *GamePlayerUpdateOne) RemovePlayers(p ...*Player) *GamePlayerUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gpuo.RemovePlayerIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gpuo *GamePlayerUpdateOne) Select(field string, fields ...string) *GamePlayerUpdateOne {
	gpuo.fields = append([]string{field}, fields...)
	return gpuo
}

// Save executes the query and returns the updated GamePlayer entity.
func (gpuo *GamePlayerUpdateOne) Save(ctx context.Context) (*GamePlayer, error) {
	var (
		err  error
		node *GamePlayer
	)
	if len(gpuo.hooks) == 0 {
		node, err = gpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GamePlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gpuo.mutation = mutation
			node, err = gpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gpuo.hooks) - 1; i >= 0; i-- {
			if gpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GamePlayer)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GamePlayerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gpuo *GamePlayerUpdateOne) SaveX(ctx context.Context) *GamePlayer {
	node, err := gpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gpuo *GamePlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := gpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gpuo *GamePlayerUpdateOne) ExecX(ctx context.Context) {
	if err := gpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gpuo *GamePlayerUpdateOne) sqlSave(ctx context.Context) (_node *GamePlayer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gameplayer.Table,
			Columns: gameplayer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: gameplayer.FieldID,
			},
		},
	}
	id, ok := gpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GamePlayer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gameplayer.FieldID)
		for _, f := range fields {
			if !gameplayer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gameplayer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gpuo.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: gameplayer.PlayersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: player.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpuo.mutation.RemovedPlayersIDs(); len(nodes) > 0 && !gpuo.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: gameplayer.PlayersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpuo.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: gameplayer.PlayersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GamePlayer{config: gpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameplayer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
