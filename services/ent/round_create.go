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
	"github.com/kanade0404/tenhou-log/services/ent/game"
	"github.com/kanade0404/tenhou-log/services/ent/hand"
	"github.com/kanade0404/tenhou-log/services/ent/round"
)

// RoundCreate is the builder for creating a Round entity.
type RoundCreate struct {
	config
	mutation *RoundMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWind sets the "wind" field.
func (rc *RoundCreate) SetWind(s string) *RoundCreate {
	rc.mutation.SetWind(s)
	return rc
}

// SetID sets the "id" field.
func (rc *RoundCreate) SetID(u uuid.UUID) *RoundCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RoundCreate) SetNillableID(u *uuid.UUID) *RoundCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetGamesID sets the "games" edge to the Game entity by ID.
func (rc *RoundCreate) SetGamesID(id uuid.UUID) *RoundCreate {
	rc.mutation.SetGamesID(id)
	return rc
}

// SetNillableGamesID sets the "games" edge to the Game entity by ID if the given value is not nil.
func (rc *RoundCreate) SetNillableGamesID(id *uuid.UUID) *RoundCreate {
	if id != nil {
		rc = rc.SetGamesID(*id)
	}
	return rc
}

// SetGames sets the "games" edge to the Game entity.
func (rc *RoundCreate) SetGames(g *Game) *RoundCreate {
	return rc.SetGamesID(g.ID)
}

// AddHandIDs adds the "hands" edge to the Hand entity by IDs.
func (rc *RoundCreate) AddHandIDs(ids ...uuid.UUID) *RoundCreate {
	rc.mutation.AddHandIDs(ids...)
	return rc
}

// AddHands adds the "hands" edges to the Hand entity.
func (rc *RoundCreate) AddHands(h ...*Hand) *RoundCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return rc.AddHandIDs(ids...)
}

// Mutation returns the RoundMutation object of the builder.
func (rc *RoundCreate) Mutation() *RoundMutation {
	return rc.mutation
}

// Save creates the Round in the database.
func (rc *RoundCreate) Save(ctx context.Context) (*Round, error) {
	rc.defaults()
	return withHooks[*Round, RoundMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoundCreate) SaveX(ctx context.Context) *Round {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoundCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoundCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoundCreate) defaults() {
	if _, ok := rc.mutation.ID(); !ok {
		v := round.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoundCreate) check() error {
	if _, ok := rc.mutation.Wind(); !ok {
		return &ValidationError{Name: "wind", err: errors.New(`ent: missing required field "Round.wind"`)}
	}
	if v, ok := rc.mutation.Wind(); ok {
		if err := round.WindValidator(v); err != nil {
			return &ValidationError{Name: "wind", err: fmt.Errorf(`ent: validator failed for field "Round.wind": %w`, err)}
		}
	}
	return nil
}

func (rc *RoundCreate) sqlSave(ctx context.Context) (*Round, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
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
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoundCreate) createSpec() (*Round, *sqlgraph.CreateSpec) {
	var (
		_node = &Round{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: round.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: round.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.Wind(); ok {
		_spec.SetField(round.FieldWind, field.TypeString, value)
		_node.Wind = value
	}
	if nodes := rc.mutation.GamesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   round.GamesTable,
			Columns: []string{round.GamesColumn},
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
		_node.game_rounds = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.HandsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   round.HandsTable,
			Columns: []string{round.HandsColumn},
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Round.Create().
//		SetWind(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoundUpsert) {
//			SetWind(v+v).
//		}).
//		Exec(ctx)
func (rc *RoundCreate) OnConflict(opts ...sql.ConflictOption) *RoundUpsertOne {
	rc.conflict = opts
	return &RoundUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Round.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RoundCreate) OnConflictColumns(columns ...string) *RoundUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RoundUpsertOne{
		create: rc,
	}
}

type (
	// RoundUpsertOne is the builder for "upsert"-ing
	//  one Round node.
	RoundUpsertOne struct {
		create *RoundCreate
	}

	// RoundUpsert is the "OnConflict" setter.
	RoundUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Round.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(round.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoundUpsertOne) UpdateNewValues() *RoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(round.FieldID)
		}
		if _, exists := u.create.mutation.Wind(); exists {
			s.SetIgnore(round.FieldWind)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Round.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoundUpsertOne) Ignore() *RoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoundUpsertOne) DoNothing() *RoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoundCreate.OnConflict
// documentation for more info.
func (u *RoundUpsertOne) Update(set func(*RoundUpsert)) *RoundUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoundUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *RoundUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoundCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoundUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RoundUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: RoundUpsertOne.ID is not supported by MySQL driver. Use RoundUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RoundUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RoundCreateBulk is the builder for creating many Round entities in bulk.
type RoundCreateBulk struct {
	config
	builders []*RoundCreate
	conflict []sql.ConflictOption
}

// Save creates the Round entities in the database.
func (rcb *RoundCreateBulk) Save(ctx context.Context) ([]*Round, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Round, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoundMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoundCreateBulk) SaveX(ctx context.Context) []*Round {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoundCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoundCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Round.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoundUpsert) {
//			SetWind(v+v).
//		}).
//		Exec(ctx)
func (rcb *RoundCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoundUpsertBulk {
	rcb.conflict = opts
	return &RoundUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Round.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RoundCreateBulk) OnConflictColumns(columns ...string) *RoundUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RoundUpsertBulk{
		create: rcb,
	}
}

// RoundUpsertBulk is the builder for "upsert"-ing
// a bulk of Round nodes.
type RoundUpsertBulk struct {
	create *RoundCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Round.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(round.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoundUpsertBulk) UpdateNewValues() *RoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(round.FieldID)
			}
			if _, exists := b.mutation.Wind(); exists {
				s.SetIgnore(round.FieldWind)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Round.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoundUpsertBulk) Ignore() *RoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoundUpsertBulk) DoNothing() *RoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoundCreateBulk.OnConflict
// documentation for more info.
func (u *RoundUpsertBulk) Update(set func(*RoundUpsert)) *RoundUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoundUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *RoundUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RoundCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RoundCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoundUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}