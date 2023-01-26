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
	"github.com/kanade0404/tenhou-log/services/ent/concealedkan"
)

// ConcealedKanCreate is the builder for creating a ConcealedKan entity.
type ConcealedKanCreate struct {
	config
	mutation *ConcealedKanMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (ckc *ConcealedKanCreate) SetID(u uuid.UUID) *ConcealedKanCreate {
	ckc.mutation.SetID(u)
	return ckc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ckc *ConcealedKanCreate) SetNillableID(u *uuid.UUID) *ConcealedKanCreate {
	if u != nil {
		ckc.SetID(*u)
	}
	return ckc
}

// SetCallID sets the "call" edge to the Call entity by ID.
func (ckc *ConcealedKanCreate) SetCallID(id uuid.UUID) *ConcealedKanCreate {
	ckc.mutation.SetCallID(id)
	return ckc
}

// SetNillableCallID sets the "call" edge to the Call entity by ID if the given value is not nil.
func (ckc *ConcealedKanCreate) SetNillableCallID(id *uuid.UUID) *ConcealedKanCreate {
	if id != nil {
		ckc = ckc.SetCallID(*id)
	}
	return ckc
}

// SetCall sets the "call" edge to the Call entity.
func (ckc *ConcealedKanCreate) SetCall(c *Call) *ConcealedKanCreate {
	return ckc.SetCallID(c.ID)
}

// Mutation returns the ConcealedKanMutation object of the builder.
func (ckc *ConcealedKanCreate) Mutation() *ConcealedKanMutation {
	return ckc.mutation
}

// Save creates the ConcealedKan in the database.
func (ckc *ConcealedKanCreate) Save(ctx context.Context) (*ConcealedKan, error) {
	ckc.defaults()
	return withHooks[*ConcealedKan, ConcealedKanMutation](ctx, ckc.sqlSave, ckc.mutation, ckc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ckc *ConcealedKanCreate) SaveX(ctx context.Context) *ConcealedKan {
	v, err := ckc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ckc *ConcealedKanCreate) Exec(ctx context.Context) error {
	_, err := ckc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ckc *ConcealedKanCreate) ExecX(ctx context.Context) {
	if err := ckc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ckc *ConcealedKanCreate) defaults() {
	if _, ok := ckc.mutation.ID(); !ok {
		v := concealedkan.DefaultID()
		ckc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ckc *ConcealedKanCreate) check() error {
	return nil
}

func (ckc *ConcealedKanCreate) sqlSave(ctx context.Context) (*ConcealedKan, error) {
	if err := ckc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ckc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ckc.driver, _spec); err != nil {
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
	ckc.mutation.id = &_node.ID
	ckc.mutation.done = true
	return _node, nil
}

func (ckc *ConcealedKanCreate) createSpec() (*ConcealedKan, *sqlgraph.CreateSpec) {
	var (
		_node = &ConcealedKan{config: ckc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: concealedkan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: concealedkan.FieldID,
			},
		}
	)
	_spec.OnConflict = ckc.conflict
	if id, ok := ckc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := ckc.mutation.CallIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   concealedkan.CallTable,
			Columns: []string{concealedkan.CallColumn},
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
//	client.ConcealedKan.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (ckc *ConcealedKanCreate) OnConflict(opts ...sql.ConflictOption) *ConcealedKanUpsertOne {
	ckc.conflict = opts
	return &ConcealedKanUpsertOne{
		create: ckc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ConcealedKan.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ckc *ConcealedKanCreate) OnConflictColumns(columns ...string) *ConcealedKanUpsertOne {
	ckc.conflict = append(ckc.conflict, sql.ConflictColumns(columns...))
	return &ConcealedKanUpsertOne{
		create: ckc,
	}
}

type (
	// ConcealedKanUpsertOne is the builder for "upsert"-ing
	//  one ConcealedKan node.
	ConcealedKanUpsertOne struct {
		create *ConcealedKanCreate
	}

	// ConcealedKanUpsert is the "OnConflict" setter.
	ConcealedKanUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ConcealedKan.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(concealedkan.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ConcealedKanUpsertOne) UpdateNewValues() *ConcealedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(concealedkan.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ConcealedKan.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ConcealedKanUpsertOne) Ignore() *ConcealedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ConcealedKanUpsertOne) DoNothing() *ConcealedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ConcealedKanCreate.OnConflict
// documentation for more info.
func (u *ConcealedKanUpsertOne) Update(set func(*ConcealedKanUpsert)) *ConcealedKanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ConcealedKanUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ConcealedKanUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ConcealedKanCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ConcealedKanUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ConcealedKanUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ConcealedKanUpsertOne.ID is not supported by MySQL driver. Use ConcealedKanUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ConcealedKanUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ConcealedKanCreateBulk is the builder for creating many ConcealedKan entities in bulk.
type ConcealedKanCreateBulk struct {
	config
	builders []*ConcealedKanCreate
	conflict []sql.ConflictOption
}

// Save creates the ConcealedKan entities in the database.
func (ckcb *ConcealedKanCreateBulk) Save(ctx context.Context) ([]*ConcealedKan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ckcb.builders))
	nodes := make([]*ConcealedKan, len(ckcb.builders))
	mutators := make([]Mutator, len(ckcb.builders))
	for i := range ckcb.builders {
		func(i int, root context.Context) {
			builder := ckcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConcealedKanMutation)
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
					_, err = mutators[i+1].Mutate(root, ckcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ckcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ckcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ckcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ckcb *ConcealedKanCreateBulk) SaveX(ctx context.Context) []*ConcealedKan {
	v, err := ckcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ckcb *ConcealedKanCreateBulk) Exec(ctx context.Context) error {
	_, err := ckcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ckcb *ConcealedKanCreateBulk) ExecX(ctx context.Context) {
	if err := ckcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ConcealedKan.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (ckcb *ConcealedKanCreateBulk) OnConflict(opts ...sql.ConflictOption) *ConcealedKanUpsertBulk {
	ckcb.conflict = opts
	return &ConcealedKanUpsertBulk{
		create: ckcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ConcealedKan.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ckcb *ConcealedKanCreateBulk) OnConflictColumns(columns ...string) *ConcealedKanUpsertBulk {
	ckcb.conflict = append(ckcb.conflict, sql.ConflictColumns(columns...))
	return &ConcealedKanUpsertBulk{
		create: ckcb,
	}
}

// ConcealedKanUpsertBulk is the builder for "upsert"-ing
// a bulk of ConcealedKan nodes.
type ConcealedKanUpsertBulk struct {
	create *ConcealedKanCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ConcealedKan.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(concealedkan.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ConcealedKanUpsertBulk) UpdateNewValues() *ConcealedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(concealedkan.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ConcealedKan.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ConcealedKanUpsertBulk) Ignore() *ConcealedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ConcealedKanUpsertBulk) DoNothing() *ConcealedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ConcealedKanCreateBulk.OnConflict
// documentation for more info.
func (u *ConcealedKanUpsertBulk) Update(set func(*ConcealedKanUpsert)) *ConcealedKanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ConcealedKanUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ConcealedKanUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ConcealedKanCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ConcealedKanCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ConcealedKanUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}