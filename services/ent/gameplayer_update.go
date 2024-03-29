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
	"github.com/kanade0404/tenhou-log/services/ent/dan"
	"github.com/kanade0404/tenhou-log/services/ent/game"
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

// SetStartPosition sets the "start_position" field.
func (gpu *GamePlayerUpdate) SetStartPosition(s string) *GamePlayerUpdate {
	gpu.mutation.SetStartPosition(s)
	return gpu
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (gpu *GamePlayerUpdate) AddGameIDs(ids ...uuid.UUID) *GamePlayerUpdate {
	gpu.mutation.AddGameIDs(ids...)
	return gpu
}

// AddGames adds the "games" edges to the Game entity.
func (gpu *GamePlayerUpdate) AddGames(g ...*Game) *GamePlayerUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gpu.AddGameIDs(ids...)
}

// SetPlayersID sets the "players" edge to the Player entity by ID.
func (gpu *GamePlayerUpdate) SetPlayersID(id uuid.UUID) *GamePlayerUpdate {
	gpu.mutation.SetPlayersID(id)
	return gpu
}

// SetNillablePlayersID sets the "players" edge to the Player entity by ID if the given value is not nil.
func (gpu *GamePlayerUpdate) SetNillablePlayersID(id *uuid.UUID) *GamePlayerUpdate {
	if id != nil {
		gpu = gpu.SetPlayersID(*id)
	}
	return gpu
}

// SetPlayers sets the "players" edge to the Player entity.
func (gpu *GamePlayerUpdate) SetPlayers(p *Player) *GamePlayerUpdate {
	return gpu.SetPlayersID(p.ID)
}

// SetDansID sets the "dans" edge to the Dan entity by ID.
func (gpu *GamePlayerUpdate) SetDansID(id uuid.UUID) *GamePlayerUpdate {
	gpu.mutation.SetDansID(id)
	return gpu
}

// SetNillableDansID sets the "dans" edge to the Dan entity by ID if the given value is not nil.
func (gpu *GamePlayerUpdate) SetNillableDansID(id *uuid.UUID) *GamePlayerUpdate {
	if id != nil {
		gpu = gpu.SetDansID(*id)
	}
	return gpu
}

// SetDans sets the "dans" edge to the Dan entity.
func (gpu *GamePlayerUpdate) SetDans(d *Dan) *GamePlayerUpdate {
	return gpu.SetDansID(d.ID)
}

// Mutation returns the GamePlayerMutation object of the builder.
func (gpu *GamePlayerUpdate) Mutation() *GamePlayerMutation {
	return gpu.mutation
}

// ClearGames clears all "games" edges to the Game entity.
func (gpu *GamePlayerUpdate) ClearGames() *GamePlayerUpdate {
	gpu.mutation.ClearGames()
	return gpu
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (gpu *GamePlayerUpdate) RemoveGameIDs(ids ...uuid.UUID) *GamePlayerUpdate {
	gpu.mutation.RemoveGameIDs(ids...)
	return gpu
}

// RemoveGames removes "games" edges to Game entities.
func (gpu *GamePlayerUpdate) RemoveGames(g ...*Game) *GamePlayerUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gpu.RemoveGameIDs(ids...)
}

// ClearPlayers clears the "players" edge to the Player entity.
func (gpu *GamePlayerUpdate) ClearPlayers() *GamePlayerUpdate {
	gpu.mutation.ClearPlayers()
	return gpu
}

// ClearDans clears the "dans" edge to the Dan entity.
func (gpu *GamePlayerUpdate) ClearDans() *GamePlayerUpdate {
	gpu.mutation.ClearDans()
	return gpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gpu *GamePlayerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GamePlayerMutation](ctx, gpu.sqlSave, gpu.mutation, gpu.hooks)
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

// check runs all checks and user-defined validators on the builder.
func (gpu *GamePlayerUpdate) check() error {
	if v, ok := gpu.mutation.StartPosition(); ok {
		if err := gameplayer.StartPositionValidator(v); err != nil {
			return &ValidationError{Name: "start_position", err: fmt.Errorf(`ent: validator failed for field "GamePlayer.start_position": %w`, err)}
		}
	}
	return nil
}

func (gpu *GamePlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gpu.check(); err != nil {
		return n, err
	}
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
	if value, ok := gpu.mutation.StartPosition(); ok {
		_spec.SetField(gameplayer.FieldStartPosition, field.TypeString, value)
	}
	if gpu.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.GamesTable,
			Columns: gameplayer.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpu.mutation.RemovedGamesIDs(); len(nodes) > 0 && !gpu.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.GamesTable,
			Columns: gameplayer.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpu.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.GamesTable,
			Columns: gameplayer.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gpu.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: []string{gameplayer.PlayersColumn},
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
	if nodes := gpu.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: []string{gameplayer.PlayersColumn},
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
	if gpu.mutation.DansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.DansTable,
			Columns: []string{gameplayer.DansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpu.mutation.DansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.DansTable,
			Columns: []string{gameplayer.DansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dan.FieldID,
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
	gpu.mutation.done = true
	return n, nil
}

// GamePlayerUpdateOne is the builder for updating a single GamePlayer entity.
type GamePlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GamePlayerMutation
}

// SetStartPosition sets the "start_position" field.
func (gpuo *GamePlayerUpdateOne) SetStartPosition(s string) *GamePlayerUpdateOne {
	gpuo.mutation.SetStartPosition(s)
	return gpuo
}

// AddGameIDs adds the "games" edge to the Game entity by IDs.
func (gpuo *GamePlayerUpdateOne) AddGameIDs(ids ...uuid.UUID) *GamePlayerUpdateOne {
	gpuo.mutation.AddGameIDs(ids...)
	return gpuo
}

// AddGames adds the "games" edges to the Game entity.
func (gpuo *GamePlayerUpdateOne) AddGames(g ...*Game) *GamePlayerUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gpuo.AddGameIDs(ids...)
}

// SetPlayersID sets the "players" edge to the Player entity by ID.
func (gpuo *GamePlayerUpdateOne) SetPlayersID(id uuid.UUID) *GamePlayerUpdateOne {
	gpuo.mutation.SetPlayersID(id)
	return gpuo
}

// SetNillablePlayersID sets the "players" edge to the Player entity by ID if the given value is not nil.
func (gpuo *GamePlayerUpdateOne) SetNillablePlayersID(id *uuid.UUID) *GamePlayerUpdateOne {
	if id != nil {
		gpuo = gpuo.SetPlayersID(*id)
	}
	return gpuo
}

// SetPlayers sets the "players" edge to the Player entity.
func (gpuo *GamePlayerUpdateOne) SetPlayers(p *Player) *GamePlayerUpdateOne {
	return gpuo.SetPlayersID(p.ID)
}

// SetDansID sets the "dans" edge to the Dan entity by ID.
func (gpuo *GamePlayerUpdateOne) SetDansID(id uuid.UUID) *GamePlayerUpdateOne {
	gpuo.mutation.SetDansID(id)
	return gpuo
}

// SetNillableDansID sets the "dans" edge to the Dan entity by ID if the given value is not nil.
func (gpuo *GamePlayerUpdateOne) SetNillableDansID(id *uuid.UUID) *GamePlayerUpdateOne {
	if id != nil {
		gpuo = gpuo.SetDansID(*id)
	}
	return gpuo
}

// SetDans sets the "dans" edge to the Dan entity.
func (gpuo *GamePlayerUpdateOne) SetDans(d *Dan) *GamePlayerUpdateOne {
	return gpuo.SetDansID(d.ID)
}

// Mutation returns the GamePlayerMutation object of the builder.
func (gpuo *GamePlayerUpdateOne) Mutation() *GamePlayerMutation {
	return gpuo.mutation
}

// ClearGames clears all "games" edges to the Game entity.
func (gpuo *GamePlayerUpdateOne) ClearGames() *GamePlayerUpdateOne {
	gpuo.mutation.ClearGames()
	return gpuo
}

// RemoveGameIDs removes the "games" edge to Game entities by IDs.
func (gpuo *GamePlayerUpdateOne) RemoveGameIDs(ids ...uuid.UUID) *GamePlayerUpdateOne {
	gpuo.mutation.RemoveGameIDs(ids...)
	return gpuo
}

// RemoveGames removes "games" edges to Game entities.
func (gpuo *GamePlayerUpdateOne) RemoveGames(g ...*Game) *GamePlayerUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gpuo.RemoveGameIDs(ids...)
}

// ClearPlayers clears the "players" edge to the Player entity.
func (gpuo *GamePlayerUpdateOne) ClearPlayers() *GamePlayerUpdateOne {
	gpuo.mutation.ClearPlayers()
	return gpuo
}

// ClearDans clears the "dans" edge to the Dan entity.
func (gpuo *GamePlayerUpdateOne) ClearDans() *GamePlayerUpdateOne {
	gpuo.mutation.ClearDans()
	return gpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gpuo *GamePlayerUpdateOne) Select(field string, fields ...string) *GamePlayerUpdateOne {
	gpuo.fields = append([]string{field}, fields...)
	return gpuo
}

// Save executes the query and returns the updated GamePlayer entity.
func (gpuo *GamePlayerUpdateOne) Save(ctx context.Context) (*GamePlayer, error) {
	return withHooks[*GamePlayer, GamePlayerMutation](ctx, gpuo.sqlSave, gpuo.mutation, gpuo.hooks)
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

// check runs all checks and user-defined validators on the builder.
func (gpuo *GamePlayerUpdateOne) check() error {
	if v, ok := gpuo.mutation.StartPosition(); ok {
		if err := gameplayer.StartPositionValidator(v); err != nil {
			return &ValidationError{Name: "start_position", err: fmt.Errorf(`ent: validator failed for field "GamePlayer.start_position": %w`, err)}
		}
	}
	return nil
}

func (gpuo *GamePlayerUpdateOne) sqlSave(ctx context.Context) (_node *GamePlayer, err error) {
	if err := gpuo.check(); err != nil {
		return _node, err
	}
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
	if value, ok := gpuo.mutation.StartPosition(); ok {
		_spec.SetField(gameplayer.FieldStartPosition, field.TypeString, value)
	}
	if gpuo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.GamesTable,
			Columns: gameplayer.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: game.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpuo.mutation.RemovedGamesIDs(); len(nodes) > 0 && !gpuo.mutation.GamesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.GamesTable,
			Columns: gameplayer.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpuo.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameplayer.GamesTable,
			Columns: gameplayer.GamesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gpuo.mutation.PlayersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: []string{gameplayer.PlayersColumn},
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
	if nodes := gpuo.mutation.PlayersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.PlayersTable,
			Columns: []string{gameplayer.PlayersColumn},
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
	if gpuo.mutation.DansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.DansTable,
			Columns: []string{gameplayer.DansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gpuo.mutation.DansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gameplayer.DansTable,
			Columns: []string{gameplayer.DansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dan.FieldID,
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
	gpuo.mutation.done = true
	return _node, nil
}
