// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
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
	conflict []sql.ConflictOption
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
	gppc.defaults()
	return withHooks[*GamePlayerPoint, GamePlayerPointMutation](ctx, gppc.sqlSave, gppc.mutation, gppc.hooks)
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
	if err := gppc.check(); err != nil {
		return nil, err
	}
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
	gppc.mutation.id = &_node.ID
	gppc.mutation.done = true
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
	_spec.OnConflict = gppc.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GamePlayerPoint.Create().
//		SetPoint(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GamePlayerPointUpsert) {
//			SetPoint(v+v).
//		}).
//		Exec(ctx)
func (gppc *GamePlayerPointCreate) OnConflict(opts ...sql.ConflictOption) *GamePlayerPointUpsertOne {
	gppc.conflict = opts
	return &GamePlayerPointUpsertOne{
		create: gppc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GamePlayerPoint.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gppc *GamePlayerPointCreate) OnConflictColumns(columns ...string) *GamePlayerPointUpsertOne {
	gppc.conflict = append(gppc.conflict, sql.ConflictColumns(columns...))
	return &GamePlayerPointUpsertOne{
		create: gppc,
	}
}

type (
	// GamePlayerPointUpsertOne is the builder for "upsert"-ing
	//  one GamePlayerPoint node.
	GamePlayerPointUpsertOne struct {
		create *GamePlayerPointCreate
	}

	// GamePlayerPointUpsert is the "OnConflict" setter.
	GamePlayerPointUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GamePlayerPoint.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gameplayerpoint.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GamePlayerPointUpsertOne) UpdateNewValues() *GamePlayerPointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(gameplayerpoint.FieldID)
		}
		if _, exists := u.create.mutation.Point(); exists {
			s.SetIgnore(gameplayerpoint.FieldPoint)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GamePlayerPoint.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GamePlayerPointUpsertOne) Ignore() *GamePlayerPointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GamePlayerPointUpsertOne) DoNothing() *GamePlayerPointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GamePlayerPointCreate.OnConflict
// documentation for more info.
func (u *GamePlayerPointUpsertOne) Update(set func(*GamePlayerPointUpsert)) *GamePlayerPointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GamePlayerPointUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *GamePlayerPointUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GamePlayerPointCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GamePlayerPointUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GamePlayerPointUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GamePlayerPointUpsertOne.ID is not supported by MySQL driver. Use GamePlayerPointUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GamePlayerPointUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GamePlayerPointCreateBulk is the builder for creating many GamePlayerPoint entities in bulk.
type GamePlayerPointCreateBulk struct {
	config
	builders []*GamePlayerPointCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = gppcb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GamePlayerPoint.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GamePlayerPointUpsert) {
//			SetPoint(v+v).
//		}).
//		Exec(ctx)
func (gppcb *GamePlayerPointCreateBulk) OnConflict(opts ...sql.ConflictOption) *GamePlayerPointUpsertBulk {
	gppcb.conflict = opts
	return &GamePlayerPointUpsertBulk{
		create: gppcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GamePlayerPoint.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gppcb *GamePlayerPointCreateBulk) OnConflictColumns(columns ...string) *GamePlayerPointUpsertBulk {
	gppcb.conflict = append(gppcb.conflict, sql.ConflictColumns(columns...))
	return &GamePlayerPointUpsertBulk{
		create: gppcb,
	}
}

// GamePlayerPointUpsertBulk is the builder for "upsert"-ing
// a bulk of GamePlayerPoint nodes.
type GamePlayerPointUpsertBulk struct {
	create *GamePlayerPointCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GamePlayerPoint.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(gameplayerpoint.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GamePlayerPointUpsertBulk) UpdateNewValues() *GamePlayerPointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(gameplayerpoint.FieldID)
			}
			if _, exists := b.mutation.Point(); exists {
				s.SetIgnore(gameplayerpoint.FieldPoint)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GamePlayerPoint.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GamePlayerPointUpsertBulk) Ignore() *GamePlayerPointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GamePlayerPointUpsertBulk) DoNothing() *GamePlayerPointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GamePlayerPointCreateBulk.OnConflict
// documentation for more info.
func (u *GamePlayerPointUpsertBulk) Update(set func(*GamePlayerPointUpsert)) *GamePlayerPointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GamePlayerPointUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *GamePlayerPointUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GamePlayerPointCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GamePlayerPointCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GamePlayerPointUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
