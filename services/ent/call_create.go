// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/call"
)

// CallCreate is the builder for creating a Call entity.
type CallCreate struct {
	config
	mutation *CallMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the CallMutation object of the builder.
func (cc *CallCreate) Mutation() *CallMutation {
	return cc.mutation
}

// Save creates the Call in the database.
func (cc *CallCreate) Save(ctx context.Context) (*Call, error) {
	var (
		err  error
		node *Call
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CallMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Call)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CallMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CallCreate) SaveX(ctx context.Context) *Call {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CallCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CallCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CallCreate) check() error {
	return nil
}

func (cc *CallCreate) sqlSave(ctx context.Context) (*Call, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CallCreate) createSpec() (*Call, *sqlgraph.CreateSpec) {
	var (
		_node = &Call{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: call.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: call.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Call.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (cc *CallCreate) OnConflict(opts ...sql.ConflictOption) *CallUpsertOne {
	cc.conflict = opts
	return &CallUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Call.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CallCreate) OnConflictColumns(columns ...string) *CallUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CallUpsertOne{
		create: cc,
	}
}

type (
	// CallUpsertOne is the builder for "upsert"-ing
	//  one Call node.
	CallUpsertOne struct {
		create *CallCreate
	}

	// CallUpsert is the "OnConflict" setter.
	CallUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Call.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CallUpsertOne) UpdateNewValues() *CallUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Call.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CallUpsertOne) Ignore() *CallUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CallUpsertOne) DoNothing() *CallUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CallCreate.OnConflict
// documentation for more info.
func (u *CallUpsertOne) Update(set func(*CallUpsert)) *CallUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CallUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *CallUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CallCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CallUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CallUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CallUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CallCreateBulk is the builder for creating many Call entities in bulk.
type CallCreateBulk struct {
	config
	builders []*CallCreate
	conflict []sql.ConflictOption
}

// Save creates the Call entities in the database.
func (ccb *CallCreateBulk) Save(ctx context.Context) ([]*Call, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Call, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CallMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CallCreateBulk) SaveX(ctx context.Context) []*Call {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CallCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CallCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Call.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (ccb *CallCreateBulk) OnConflict(opts ...sql.ConflictOption) *CallUpsertBulk {
	ccb.conflict = opts
	return &CallUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Call.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CallCreateBulk) OnConflictColumns(columns ...string) *CallUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CallUpsertBulk{
		create: ccb,
	}
}

// CallUpsertBulk is the builder for "upsert"-ing
// a bulk of Call nodes.
type CallUpsertBulk struct {
	create *CallCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Call.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CallUpsertBulk) UpdateNewValues() *CallUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Call.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CallUpsertBulk) Ignore() *CallUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CallUpsertBulk) DoNothing() *CallUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CallCreateBulk.OnConflict
// documentation for more info.
func (u *CallUpsertBulk) Update(set func(*CallUpsert)) *CallUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CallUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *CallUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CallCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CallCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CallUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
