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
	"github.com/kanade0404/tenhou-log/services/ent/call"
	"github.com/kanade0404/tenhou-log/services/ent/meldedkan"
)

// MeldedKanCreate is the builder for creating a MeldedKan entity.
type MeldedKanCreate struct {
	config
	mutation *MeldedKanMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (mkc *MeldedKanCreate) SetID(u uuid.UUID) *MeldedKanCreate {
	mkc.mutation.SetID(u)
	return mkc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mkc *MeldedKanCreate) SetNillableID(u *uuid.UUID) *MeldedKanCreate {
	if u != nil {
		mkc.SetID(*u)
	}
	return mkc
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (mkc *MeldedKanCreate) SetCallID(id uuid.UUID) *MeldedKanCreate {
	mkc.mutation.SetCallID(id)
	return mkc
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (mkc *MeldedKanCreate) SetNillableCallID(id *uuid.UUID) *MeldedKanCreate {
	if id != nil {
		mkc = mkc.SetCallID(*id)
	}
	return mkc
}

// SetCall sets the "call" edge to the Call entity.
func (mkc *MeldedKanCreate) SetCall(c *Call) *MeldedKanCreate {
	return mkc.SetCallID(c.ID)
}

// Mutation returns the MeldedKanMutation object of the builder.
func (mkc *MeldedKanCreate) Mutation() *MeldedKanMutation {
	return mkc.mutation
}

// Save creates the MeldedKan in the database.
func (mkc *MeldedKanCreate) Save(ctx context.Context) (*MeldedKan, error) {
	mkc.defaults()
	return withHooks[*MeldedKan, MeldedKanMutation](ctx, mkc.sqlSave, mkc.mutation, mkc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mkc *MeldedKanCreate) SaveX(ctx context.Context) *MeldedKan {
	v, err := mkc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mkc *MeldedKanCreate) Exec(ctx context.Context) error {
	_, err := mkc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mkc *MeldedKanCreate) ExecX(ctx context.Context) {
	if err := mkc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mkc *MeldedKanCreate) defaults() {
	if _, ok := mkc.mutation.ID(); !ok {
		v := meldedkan.DefaultID()
		mkc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mkc *MeldedKanCreate) check() error {
	return nil
}

func (mkc *MeldedKanCreate) sqlSave(ctx context.Context) (*MeldedKan, error) {
	if err := mkc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mkc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mkc.driver, _spec); err != nil {
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
	mkc.mutation.id = &_node.ID
	mkc.mutation.done = true
	return _node, nil
}

func (mkc *MeldedKanCreate) createSpec() (*MeldedKan, *sqlgraph.CreateSpec) {
	var (
		_node = &MeldedKan{config: mkc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: meldedkan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: meldedkan.FieldID,
			},
		}
	)
	_spec.OnConflict = mkc.conflict
	if id, ok := mkc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := mkc.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   meldedkan.CallTable,
			Columns: []string{meldedkan.CallColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: call.FieldID,
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
//	client.MeldedKan.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (mkc *MeldedKanCreate) OnConflict(opts ...sql.ConflictOption) *MeldedKanUpsertOne {
	mkc.conflict = opts
	return &MeldedKanUpsertOne{
		create: mkc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MeldedKan.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mkc *MeldedKanCreate) OnConflictColumns(columns ...string) *MeldedKanUpsertOne {
	mkc.conflict = append(mkc.conflict, sql.ConflictColumns(columns...))
	return &MeldedKanUpsertOne{
		create: mkc,
	}
}

type (
	// MeldedKanUpsertOne is the builder for "upsert"-ing
	//  one MeldedKan node.
	MeldedKanUpsertOne struct {
		create *MeldedKanCreate
	}

	// MeldedKanUpsert is the "OnConflict" setter.
	MeldedKanUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MeldedKan.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(meldedkan.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MeldedKanUpsertOne) UpdateNewValues() *MeldedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(meldedkan.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MeldedKan.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MeldedKanUpsertOne) Ignore() *MeldedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MeldedKanUpsertOne) DoNothing() *MeldedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MeldedKanCreate.OnConflict
// documentation for more info.
func (u *MeldedKanUpsertOne) Update(set func(*MeldedKanUpsert)) *MeldedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MeldedKanUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *MeldedKanUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MeldedKanCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MeldedKanUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MeldedKanUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MeldedKanUpsertOne.ID is not supported by MySQL driver. Use MeldedKanUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MeldedKanUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MeldedKanCreateBulk is the builder for creating many MeldedKan entities in bulk.
type MeldedKanCreateBulk struct {
	config
	builders []*MeldedKanCreate
	conflict []sql.ConflictOption
}

// Save creates the MeldedKan entities in the database.
func (mkcb *MeldedKanCreateBulk) Save(ctx context.Context) ([]*MeldedKan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mkcb.builders))
	nodes := make([]*MeldedKan, len(mkcb.builders))
	mutators := make([]Mutator, len(mkcb.builders))
	for i := range mkcb.builders {
		func(i int, root context.Context) {
			builder := mkcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MeldedKanMutation)
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
					_, err = mutators[i+1].Mutate(root, mkcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mkcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mkcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mkcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mkcb *MeldedKanCreateBulk) SaveX(ctx context.Context) []*MeldedKan {
	v, err := mkcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mkcb *MeldedKanCreateBulk) Exec(ctx context.Context) error {
	_, err := mkcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mkcb *MeldedKanCreateBulk) ExecX(ctx context.Context) {
	if err := mkcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MeldedKan.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (mkcb *MeldedKanCreateBulk) OnConflict(opts ...sql.ConflictOption) *MeldedKanUpsertBulk {
	mkcb.conflict = opts
	return &MeldedKanUpsertBulk{
		create: mkcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MeldedKan.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mkcb *MeldedKanCreateBulk) OnConflictColumns(columns ...string) *MeldedKanUpsertBulk {
	mkcb.conflict = append(mkcb.conflict, sql.ConflictColumns(columns...))
	return &MeldedKanUpsertBulk{
		create: mkcb,
	}
}

// MeldedKanUpsertBulk is the builder for "upsert"-ing
// a bulk of MeldedKan nodes.
type MeldedKanUpsertBulk struct {
	create *MeldedKanCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MeldedKan.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(meldedkan.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MeldedKanUpsertBulk) UpdateNewValues() *MeldedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(meldedkan.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MeldedKan.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MeldedKanUpsertBulk) Ignore() *MeldedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MeldedKanUpsertBulk) DoNothing() *MeldedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MeldedKanCreateBulk.OnConflict
// documentation for more info.
func (u *MeldedKanUpsertBulk) Update(set func(*MeldedKanUpsert)) *MeldedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MeldedKanUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *MeldedKanUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MeldedKanCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MeldedKanCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MeldedKanUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}