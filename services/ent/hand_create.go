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
	"github.com/kanade0404/tenhou-log/services/ent/hand"
	"github.com/kanade0404/tenhou-log/services/ent/round"
	"github.com/kanade0404/tenhou-log/services/ent/turn"
)

// HandCreate is the builder for creating a Hand entity.
type HandCreate struct {
	config
	mutation *HandMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNum sets the "num" field.
func (hc *HandCreate) SetNum(u uint) *HandCreate {
	hc.mutation.SetNum(u)
	return hc
}

// SetContinuePoint sets the "continue_point" field.
func (hc *HandCreate) SetContinuePoint(u uint) *HandCreate {
	hc.mutation.SetContinuePoint(u)
	return hc
}

// SetReachPoint sets the "reach_point" field.
func (hc *HandCreate) SetReachPoint(u uint) *HandCreate {
	hc.mutation.SetReachPoint(u)
	return hc
}

// SetID sets the "id" field.
func (hc *HandCreate) SetID(u uuid.UUID) *HandCreate {
	hc.mutation.SetID(u)
	return hc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hc *HandCreate) SetNillableID(u *uuid.UUID) *HandCreate {
	if u != nil {
		hc.SetID(*u)
	}
	return hc
}

// SetRoundsID sets the "rounds" edge to the Round entity by ID.
func (hc *HandCreate) SetRoundsID(id uuid.UUID) *HandCreate {
	hc.mutation.SetRoundsID(id)
	return hc
}

// SetNillableRoundsID sets the "rounds" edge to the Round entity by ID if the given value is not nil.
func (hc *HandCreate) SetNillableRoundsID(id *uuid.UUID) *HandCreate {
	if id != nil {
		hc = hc.SetRoundsID(*id)
	}
	return hc
}

// SetRounds sets the "rounds" edge to the Round entity.
func (hc *HandCreate) SetRounds(r *Round) *HandCreate {
	return hc.SetRoundsID(r.ID)
}

// AddTurnIDs adds the "turns" edge to the Turn entity by IDs.
func (hc *HandCreate) AddTurnIDs(ids ...uuid.UUID) *HandCreate {
	hc.mutation.AddTurnIDs(ids...)
	return hc
}

// AddTurns adds the "turns" edges to the Turn entity.
func (hc *HandCreate) AddTurns(t ...*Turn) *HandCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return hc.AddTurnIDs(ids...)
}

// Mutation returns the HandMutation object of the builder.
func (hc *HandCreate) Mutation() *HandMutation {
	return hc.mutation
}

// Save creates the Hand in the database.
func (hc *HandCreate) Save(ctx context.Context) (*Hand, error) {
	hc.defaults()
	return withHooks[*Hand, HandMutation](ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HandCreate) SaveX(ctx context.Context) *Hand {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HandCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HandCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HandCreate) defaults() {
	if _, ok := hc.mutation.ID(); !ok {
		v := hand.DefaultID()
		hc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HandCreate) check() error {
	if _, ok := hc.mutation.Num(); !ok {
		return &ValidationError{Name: "num", err: errors.New(`ent: missing required field "Hand.num"`)}
	}
	if _, ok := hc.mutation.ContinuePoint(); !ok {
		return &ValidationError{Name: "continue_point", err: errors.New(`ent: missing required field "Hand.continue_point"`)}
	}
	if v, ok := hc.mutation.ContinuePoint(); ok {
		if err := hand.ContinuePointValidator(v); err != nil {
			return &ValidationError{Name: "continue_point", err: fmt.Errorf(`ent: validator failed for field "Hand.continue_point": %w`, err)}
		}
	}
	if _, ok := hc.mutation.ReachPoint(); !ok {
		return &ValidationError{Name: "reach_point", err: errors.New(`ent: missing required field "Hand.reach_point"`)}
	}
	if v, ok := hc.mutation.ReachPoint(); ok {
		if err := hand.ReachPointValidator(v); err != nil {
			return &ValidationError{Name: "reach_point", err: fmt.Errorf(`ent: validator failed for field "Hand.reach_point": %w`, err)}
		}
	}
	return nil
}

func (hc *HandCreate) sqlSave(ctx context.Context) (*Hand, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
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
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HandCreate) createSpec() (*Hand, *sqlgraph.CreateSpec) {
	var (
		_node = &Hand{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hand.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hand.FieldID,
			},
		}
	)
	_spec.OnConflict = hc.conflict
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hc.mutation.Num(); ok {
		_spec.SetField(hand.FieldNum, field.TypeUint, value)
		_node.Num = value
	}
	if value, ok := hc.mutation.ContinuePoint(); ok {
		_spec.SetField(hand.FieldContinuePoint, field.TypeUint, value)
		_node.ContinuePoint = value
	}
	if value, ok := hc.mutation.ReachPoint(); ok {
		_spec.SetField(hand.FieldReachPoint, field.TypeUint, value)
		_node.ReachPoint = value
	}
	if nodes := hc.mutation.RoundsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hand.RoundsTable,
			Columns: []string{hand.RoundsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: round.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.round_hands = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.TurnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hand.TurnsTable,
			Columns: hand.TurnsPrimaryKey,
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
//	client.Hand.Create().
//		SetNum(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HandUpsert) {
//			SetNum(v+v).
//		}).
//		Exec(ctx)
func (hc *HandCreate) OnConflict(opts ...sql.ConflictOption) *HandUpsertOne {
	hc.conflict = opts
	return &HandUpsertOne{
		create: hc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Hand.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hc *HandCreate) OnConflictColumns(columns ...string) *HandUpsertOne {
	hc.conflict = append(hc.conflict, sql.ConflictColumns(columns...))
	return &HandUpsertOne{
		create: hc,
	}
}

type (
	// HandUpsertOne is the builder for "upsert"-ing
	//  one Hand node.
	HandUpsertOne struct {
		create *HandCreate
	}

	// HandUpsert is the "OnConflict" setter.
	HandUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Hand.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(hand.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *HandUpsertOne) UpdateNewValues() *HandUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(hand.FieldID)
		}
		if _, exists := u.create.mutation.Num(); exists {
			s.SetIgnore(hand.FieldNum)
		}
		if _, exists := u.create.mutation.ContinuePoint(); exists {
			s.SetIgnore(hand.FieldContinuePoint)
		}
		if _, exists := u.create.mutation.ReachPoint(); exists {
			s.SetIgnore(hand.FieldReachPoint)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Hand.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *HandUpsertOne) Ignore() *HandUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HandUpsertOne) DoNothing() *HandUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HandCreate.OnConflict
// documentation for more info.
func (u *HandUpsertOne) Update(set func(*HandUpsert)) *HandUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HandUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *HandUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HandCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HandUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *HandUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: HandUpsertOne.ID is not supported by MySQL driver. Use HandUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *HandUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// HandCreateBulk is the builder for creating many Hand entities in bulk.
type HandCreateBulk struct {
	config
	builders []*HandCreate
	conflict []sql.ConflictOption
}

// Save creates the Hand entities in the database.
func (hcb *HandCreateBulk) Save(ctx context.Context) ([]*Hand, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Hand, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HandMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = hcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HandCreateBulk) SaveX(ctx context.Context) []*Hand {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HandCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HandCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Hand.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.HandUpsert) {
//			SetNum(v+v).
//		}).
//		Exec(ctx)
func (hcb *HandCreateBulk) OnConflict(opts ...sql.ConflictOption) *HandUpsertBulk {
	hcb.conflict = opts
	return &HandUpsertBulk{
		create: hcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Hand.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (hcb *HandCreateBulk) OnConflictColumns(columns ...string) *HandUpsertBulk {
	hcb.conflict = append(hcb.conflict, sql.ConflictColumns(columns...))
	return &HandUpsertBulk{
		create: hcb,
	}
}

// HandUpsertBulk is the builder for "upsert"-ing
// a bulk of Hand nodes.
type HandUpsertBulk struct {
	create *HandCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Hand.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(hand.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *HandUpsertBulk) UpdateNewValues() *HandUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(hand.FieldID)
			}
			if _, exists := b.mutation.Num(); exists {
				s.SetIgnore(hand.FieldNum)
			}
			if _, exists := b.mutation.ContinuePoint(); exists {
				s.SetIgnore(hand.FieldContinuePoint)
			}
			if _, exists := b.mutation.ReachPoint(); exists {
				s.SetIgnore(hand.FieldReachPoint)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Hand.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *HandUpsertBulk) Ignore() *HandUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *HandUpsertBulk) DoNothing() *HandUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the HandCreateBulk.OnConflict
// documentation for more info.
func (u *HandUpsertBulk) Update(set func(*HandUpsert)) *HandUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&HandUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *HandUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the HandCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for HandCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *HandUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
