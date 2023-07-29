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
	"github.com/kanade0404/tenhou-log/services/ent/event"
	"github.com/kanade0404/tenhou-log/services/ent/win"
)

// WinCreate is the builder for creating a Win entity.
type WinCreate struct {
	config
	mutation *WinMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (wc *WinCreate) SetID(u uuid.UUID) *WinCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WinCreate) SetNillableID(u *uuid.UUID) *WinCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (wc *WinCreate) SetEventID(id uuid.UUID) *WinCreate {
	wc.mutation.SetEventID(id)
	return wc
}

// SetEvent sets the "event" edge to the Event entity.
func (wc *WinCreate) SetEvent(e *Event) *WinCreate {
	return wc.SetEventID(e.ID)
}

// Mutation returns the WinMutation object of the builder.
func (wc *WinCreate) Mutation() *WinMutation {
	return wc.mutation
}

// Save creates the Win in the database.
func (wc *WinCreate) Save(ctx context.Context) (*Win, error) {
	wc.defaults()
	return withHooks[*Win, WinMutation](ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WinCreate) SaveX(ctx context.Context) *Win {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WinCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WinCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WinCreate) defaults() {
	if _, ok := wc.mutation.ID(); !ok {
		v := win.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WinCreate) check() error {
	if _, ok := wc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required edge "Win.event"`)}
	}
	return nil
}

func (wc *WinCreate) sqlSave(ctx context.Context) (*Win, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
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
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WinCreate) createSpec() (*Win, *sqlgraph.CreateSpec) {
	var (
		_node = &Win{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(win.Table, sqlgraph.NewFieldSpec(win.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = wc.conflict
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := wc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   win.EventTable,
			Columns: []string{win.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.event_win = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Win.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (wc *WinCreate) OnConflict(opts ...sql.ConflictOption) *WinUpsertOne {
	wc.conflict = opts
	return &WinUpsertOne{
		create: wc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Win.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wc *WinCreate) OnConflictColumns(columns ...string) *WinUpsertOne {
	wc.conflict = append(wc.conflict, sql.ConflictColumns(columns...))
	return &WinUpsertOne{
		create: wc,
	}
}

type (
	// WinUpsertOne is the builder for "upsert"-ing
	//  one Win node.
	WinUpsertOne struct {
		create *WinCreate
	}

	// WinUpsert is the "OnConflict" setter.
	WinUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Win.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(win.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *WinUpsertOne) UpdateNewValues() *WinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(win.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Win.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *WinUpsertOne) Ignore() *WinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WinUpsertOne) DoNothing() *WinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WinCreate.OnConflict
// documentation for more info.
func (u *WinUpsertOne) Update(set func(*WinUpsert)) *WinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WinUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *WinUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WinCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WinUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *WinUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: WinUpsertOne.ID is not supported by MySQL driver. Use WinUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *WinUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// WinCreateBulk is the builder for creating many Win entities in bulk.
type WinCreateBulk struct {
	config
	builders []*WinCreate
	conflict []sql.ConflictOption
}

// Save creates the Win entities in the database.
func (wcb *WinCreateBulk) Save(ctx context.Context) ([]*Win, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Win, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WinMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = wcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WinCreateBulk) SaveX(ctx context.Context) []*Win {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WinCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WinCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Win.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (wcb *WinCreateBulk) OnConflict(opts ...sql.ConflictOption) *WinUpsertBulk {
	wcb.conflict = opts
	return &WinUpsertBulk{
		create: wcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Win.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wcb *WinCreateBulk) OnConflictColumns(columns ...string) *WinUpsertBulk {
	wcb.conflict = append(wcb.conflict, sql.ConflictColumns(columns...))
	return &WinUpsertBulk{
		create: wcb,
	}
}

// WinUpsertBulk is the builder for "upsert"-ing
// a bulk of Win nodes.
type WinUpsertBulk struct {
	create *WinCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Win.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(win.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *WinUpsertBulk) UpdateNewValues() *WinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(win.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Win.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *WinUpsertBulk) Ignore() *WinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WinUpsertBulk) DoNothing() *WinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WinCreateBulk.OnConflict
// documentation for more info.
func (u *WinUpsertBulk) Update(set func(*WinUpsert)) *WinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WinUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *WinUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the WinCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WinCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WinUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
