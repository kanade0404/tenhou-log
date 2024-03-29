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
	"github.com/kanade0404/tenhou-log/services/ent/chakan"
	"github.com/kanade0404/tenhou-log/services/ent/chii"
	"github.com/kanade0404/tenhou-log/services/ent/concealedkan"
	"github.com/kanade0404/tenhou-log/services/ent/discard"
	"github.com/kanade0404/tenhou-log/services/ent/event"
	"github.com/kanade0404/tenhou-log/services/ent/meldedkan"
	"github.com/kanade0404/tenhou-log/services/ent/pon"
)

// CallCreate is the builder for creating a Call entity.
type CallCreate struct {
	config
	mutation *CallMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (cc *CallCreate) SetID(u uuid.UUID) *CallCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CallCreate) SetNillableID(u *uuid.UUID) *CallCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (cc *CallCreate) SetEventID(id uuid.UUID) *CallCreate {
	cc.mutation.SetEventID(id)
	return cc
}

// SetEvent sets the "event" edge to the Event entity.
func (cc *CallCreate) SetEvent(e *Event) *CallCreate {
	return cc.SetEventID(e.ID)
}

// SetDiscardID sets the "discard" edge to the Discard entity by ID.
func (cc *CallCreate) SetDiscardID(id uuid.UUID) *CallCreate {
	cc.mutation.SetDiscardID(id)
	return cc
}

// SetDiscard sets the "discard" edge to the Discard entity.
func (cc *CallCreate) SetDiscard(d *Discard) *CallCreate {
	return cc.SetDiscardID(d.ID)
}

// SetChiiID sets the "chii" edge to the Chii entity by ID.
func (cc *CallCreate) SetChiiID(id uuid.UUID) *CallCreate {
	cc.mutation.SetChiiID(id)
	return cc
}

// SetChii sets the "chii" edge to the Chii entity.
func (cc *CallCreate) SetChii(c *Chii) *CallCreate {
	return cc.SetChiiID(c.ID)
}

// SetChakanID sets the "chakan" edge to the Chakan entity by ID.
func (cc *CallCreate) SetChakanID(id uuid.UUID) *CallCreate {
	cc.mutation.SetChakanID(id)
	return cc
}

// SetChakan sets the "chakan" edge to the Chakan entity.
func (cc *CallCreate) SetChakan(c *Chakan) *CallCreate {
	return cc.SetChakanID(c.ID)
}

// SetConcealedkanID sets the "concealedkan" edge to the ConcealedKan entity by ID.
func (cc *CallCreate) SetConcealedkanID(id uuid.UUID) *CallCreate {
	cc.mutation.SetConcealedkanID(id)
	return cc
}

// SetConcealedkan sets the "concealedkan" edge to the ConcealedKan entity.
func (cc *CallCreate) SetConcealedkan(c *ConcealedKan) *CallCreate {
	return cc.SetConcealedkanID(c.ID)
}

// SetMeldedkanID sets the "meldedkan" edge to the MeldedKan entity by ID.
func (cc *CallCreate) SetMeldedkanID(id uuid.UUID) *CallCreate {
	cc.mutation.SetMeldedkanID(id)
	return cc
}

// SetMeldedkan sets the "meldedkan" edge to the MeldedKan entity.
func (cc *CallCreate) SetMeldedkan(m *MeldedKan) *CallCreate {
	return cc.SetMeldedkanID(m.ID)
}

// SetPonID sets the "pon" edge to the Pon entity by ID.
func (cc *CallCreate) SetPonID(id uuid.UUID) *CallCreate {
	cc.mutation.SetPonID(id)
	return cc
}

// SetPon sets the "pon" edge to the Pon entity.
func (cc *CallCreate) SetPon(p *Pon) *CallCreate {
	return cc.SetPonID(p.ID)
}

// Mutation returns the CallMutation object of the builder.
func (cc *CallCreate) Mutation() *CallMutation {
	return cc.mutation
}

// Save creates the Call in the database.
func (cc *CallCreate) Save(ctx context.Context) (*Call, error) {
	cc.defaults()
	return withHooks[*Call, CallMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
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

// defaults sets the default values of the builder before save.
func (cc *CallCreate) defaults() {
	if _, ok := cc.mutation.ID(); !ok {
		v := call.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CallCreate) check() error {
	if _, ok := cc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required edge "Call.event"`)}
	}
	if _, ok := cc.mutation.DiscardID(); !ok {
		return &ValidationError{Name: "discard", err: errors.New(`ent: missing required edge "Call.discard"`)}
	}
	if _, ok := cc.mutation.ChiiID(); !ok {
		return &ValidationError{Name: "chii", err: errors.New(`ent: missing required edge "Call.chii"`)}
	}
	if _, ok := cc.mutation.ChakanID(); !ok {
		return &ValidationError{Name: "chakan", err: errors.New(`ent: missing required edge "Call.chakan"`)}
	}
	if _, ok := cc.mutation.ConcealedkanID(); !ok {
		return &ValidationError{Name: "concealedkan", err: errors.New(`ent: missing required edge "Call.concealedkan"`)}
	}
	if _, ok := cc.mutation.MeldedkanID(); !ok {
		return &ValidationError{Name: "meldedkan", err: errors.New(`ent: missing required edge "Call.meldedkan"`)}
	}
	if _, ok := cc.mutation.PonID(); !ok {
		return &ValidationError{Name: "pon", err: errors.New(`ent: missing required edge "Call.pon"`)}
	}
	return nil
}

func (cc *CallCreate) sqlSave(ctx context.Context) (*Call, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CallCreate) createSpec() (*Call, *sqlgraph.CreateSpec) {
	var (
		_node = &Call{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: call.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: call.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := cc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   call.EventTable,
			Columns: []string{call.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.event_call = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.DiscardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.DiscardTable,
			Columns: []string{call.DiscardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: discard.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discard_call = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChiiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChiiTable,
			Columns: []string{call.ChiiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: chii.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.chii_call = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChakanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChakanTable,
			Columns: []string{call.ChakanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: chakan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.chakan_call = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ConcealedkanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ConcealedkanTable,
			Columns: []string{call.ConcealedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: concealedkan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.concealed_kan_call = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.MeldedkanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.MeldedkanTable,
			Columns: []string{call.MeldedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: meldedkan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.melded_kan_call = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.PonTable,
			Columns: []string{call.PonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: pon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pon_call = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
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

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Call.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(call.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CallUpsertOne) UpdateNewValues() *CallUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(call.FieldID)
		}
	}))
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
func (u *CallUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CallUpsertOne.ID is not supported by MySQL driver. Use CallUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CallUpsertOne) IDX(ctx context.Context) uuid.UUID {
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
			builder.defaults()
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
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(call.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CallUpsertBulk) UpdateNewValues() *CallUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(call.FieldID)
			}
		}
	}))
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
