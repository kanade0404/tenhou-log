// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayerhandhai"
)

// GamePlayerHandHaiCreate is the builder for creating a GamePlayerHandHai entity.
type GamePlayerHandHaiCreate struct {
	config
	mutation *GamePlayerHandHaiMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the GamePlayerHandHaiMutation object of the builder.
func (gphhc *GamePlayerHandHaiCreate) Mutation() *GamePlayerHandHaiMutation {
	return gphhc.mutation
}

// Save creates the GamePlayerHandHai in the database.
func (gphhc *GamePlayerHandHaiCreate) Save(ctx context.Context) (*GamePlayerHandHai, error) {
	var (
		err  error
		node *GamePlayerHandHai
	)
	if len(gphhc.hooks) == 0 {
		if err = gphhc.check(); err != nil {
			return nil, err
		}
		node, err = gphhc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GamePlayerHandHaiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gphhc.check(); err != nil {
				return nil, err
			}
			gphhc.mutation = mutation
			if node, err = gphhc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gphhc.hooks) - 1; i >= 0; i-- {
			if gphhc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gphhc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gphhc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GamePlayerHandHai)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GamePlayerHandHaiMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gphhc *GamePlayerHandHaiCreate) SaveX(ctx context.Context) *GamePlayerHandHai {
	v, err := gphhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gphhc *GamePlayerHandHaiCreate) Exec(ctx context.Context) error {
	_, err := gphhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gphhc *GamePlayerHandHaiCreate) ExecX(ctx context.Context) {
	if err := gphhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gphhc *GamePlayerHandHaiCreate) check() error {
	return nil
}

func (gphhc *GamePlayerHandHaiCreate) sqlSave(ctx context.Context) (*GamePlayerHandHai, error) {
	_node, _spec := gphhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gphhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gphhc *GamePlayerHandHaiCreate) createSpec() (*GamePlayerHandHai, *sqlgraph.CreateSpec) {
	var (
		_node = &GamePlayerHandHai{config: gphhc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gameplayerhandhai.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gameplayerhandhai.FieldID,
			},
		}
	)
	_spec.OnConflict = gphhc.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GamePlayerHandHai.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (gphhc *GamePlayerHandHaiCreate) OnConflict(opts ...sql.ConflictOption) *GamePlayerHandHaiUpsertOne {
	gphhc.conflict = opts
	return &GamePlayerHandHaiUpsertOne{
		create: gphhc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GamePlayerHandHai.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gphhc *GamePlayerHandHaiCreate) OnConflictColumns(columns ...string) *GamePlayerHandHaiUpsertOne {
	gphhc.conflict = append(gphhc.conflict, sql.ConflictColumns(columns...))
	return &GamePlayerHandHaiUpsertOne{
		create: gphhc,
	}
}

type (
	// GamePlayerHandHaiUpsertOne is the builder for "upsert"-ing
	//  one GamePlayerHandHai node.
	GamePlayerHandHaiUpsertOne struct {
		create *GamePlayerHandHaiCreate
	}

	// GamePlayerHandHaiUpsert is the "OnConflict" setter.
	GamePlayerHandHaiUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GamePlayerHandHai.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GamePlayerHandHaiUpsertOne) UpdateNewValues() *GamePlayerHandHaiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GamePlayerHandHai.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GamePlayerHandHaiUpsertOne) Ignore() *GamePlayerHandHaiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GamePlayerHandHaiUpsertOne) DoNothing() *GamePlayerHandHaiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GamePlayerHandHaiCreate.OnConflict
// documentation for more info.
func (u *GamePlayerHandHaiUpsertOne) Update(set func(*GamePlayerHandHaiUpsert)) *GamePlayerHandHaiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GamePlayerHandHaiUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *GamePlayerHandHaiUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GamePlayerHandHaiCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GamePlayerHandHaiUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GamePlayerHandHaiUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GamePlayerHandHaiUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GamePlayerHandHaiCreateBulk is the builder for creating many GamePlayerHandHai entities in bulk.
type GamePlayerHandHaiCreateBulk struct {
	config
	builders []*GamePlayerHandHaiCreate
	conflict []sql.ConflictOption
}

// Save creates the GamePlayerHandHai entities in the database.
func (gphhcb *GamePlayerHandHaiCreateBulk) Save(ctx context.Context) ([]*GamePlayerHandHai, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gphhcb.builders))
	nodes := make([]*GamePlayerHandHai, len(gphhcb.builders))
	mutators := make([]Mutator, len(gphhcb.builders))
	for i := range gphhcb.builders {
		func(i int, root context.Context) {
			builder := gphhcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GamePlayerHandHaiMutation)
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
					_, err = mutators[i+1].Mutate(root, gphhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gphhcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gphhcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gphhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gphhcb *GamePlayerHandHaiCreateBulk) SaveX(ctx context.Context) []*GamePlayerHandHai {
	v, err := gphhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gphhcb *GamePlayerHandHaiCreateBulk) Exec(ctx context.Context) error {
	_, err := gphhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gphhcb *GamePlayerHandHaiCreateBulk) ExecX(ctx context.Context) {
	if err := gphhcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GamePlayerHandHai.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (gphhcb *GamePlayerHandHaiCreateBulk) OnConflict(opts ...sql.ConflictOption) *GamePlayerHandHaiUpsertBulk {
	gphhcb.conflict = opts
	return &GamePlayerHandHaiUpsertBulk{
		create: gphhcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GamePlayerHandHai.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gphhcb *GamePlayerHandHaiCreateBulk) OnConflictColumns(columns ...string) *GamePlayerHandHaiUpsertBulk {
	gphhcb.conflict = append(gphhcb.conflict, sql.ConflictColumns(columns...))
	return &GamePlayerHandHaiUpsertBulk{
		create: gphhcb,
	}
}

// GamePlayerHandHaiUpsertBulk is the builder for "upsert"-ing
// a bulk of GamePlayerHandHai nodes.
type GamePlayerHandHaiUpsertBulk struct {
	create *GamePlayerHandHaiCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GamePlayerHandHai.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GamePlayerHandHaiUpsertBulk) UpdateNewValues() *GamePlayerHandHaiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GamePlayerHandHai.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GamePlayerHandHaiUpsertBulk) Ignore() *GamePlayerHandHaiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GamePlayerHandHaiUpsertBulk) DoNothing() *GamePlayerHandHaiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GamePlayerHandHaiCreateBulk.OnConflict
// documentation for more info.
func (u *GamePlayerHandHaiUpsertBulk) Update(set func(*GamePlayerHandHaiUpsert)) *GamePlayerHandHaiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GamePlayerHandHaiUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *GamePlayerHandHaiUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GamePlayerHandHaiCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GamePlayerHandHaiCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GamePlayerHandHaiUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
