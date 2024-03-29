// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
)

// CompressedMJLogCreate is the builder for creating a CompressedMJLog entity.
type CompressedMJLogCreate struct {
	config
	mutation *CompressedMJLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (cmlc *CompressedMJLogCreate) SetName(s string) *CompressedMJLogCreate {
	cmlc.mutation.SetName(s)
	return cmlc
}

// SetSize sets the "size" field.
func (cmlc *CompressedMJLogCreate) SetSize(u uint) *CompressedMJLogCreate {
	cmlc.mutation.SetSize(u)
	return cmlc
}

// SetInsertedAt sets the "inserted_at" field.
func (cmlc *CompressedMJLogCreate) SetInsertedAt(t time.Time) *CompressedMJLogCreate {
	cmlc.mutation.SetInsertedAt(t)
	return cmlc
}

// SetNillableInsertedAt sets the "inserted_at" field if the given value is not nil.
func (cmlc *CompressedMJLogCreate) SetNillableInsertedAt(t *time.Time) *CompressedMJLogCreate {
	if t != nil {
		cmlc.SetInsertedAt(*t)
	}
	return cmlc
}

// SetID sets the "id" field.
func (cmlc *CompressedMJLogCreate) SetID(u uuid.UUID) *CompressedMJLogCreate {
	cmlc.mutation.SetID(u)
	return cmlc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cmlc *CompressedMJLogCreate) SetNillableID(u *uuid.UUID) *CompressedMJLogCreate {
	if u != nil {
		cmlc.SetID(*u)
	}
	return cmlc
}

// SetMjlogFilesID sets the "mjlog_files" edge to the MJLogFile entity by ID.
func (cmlc *CompressedMJLogCreate) SetMjlogFilesID(id uuid.UUID) *CompressedMJLogCreate {
	cmlc.mutation.SetMjlogFilesID(id)
	return cmlc
}

// SetNillableMjlogFilesID sets the "mjlog_files" edge to the MJLogFile entity by ID if the given value is not nil.
func (cmlc *CompressedMJLogCreate) SetNillableMjlogFilesID(id *uuid.UUID) *CompressedMJLogCreate {
	if id != nil {
		cmlc = cmlc.SetMjlogFilesID(*id)
	}
	return cmlc
}

// SetMjlogFiles sets the "mjlog_files" edge to the MJLogFile entity.
func (cmlc *CompressedMJLogCreate) SetMjlogFiles(m *MJLogFile) *CompressedMJLogCreate {
	return cmlc.SetMjlogFilesID(m.ID)
}

// Mutation returns the CompressedMJLogMutation object of the builder.
func (cmlc *CompressedMJLogCreate) Mutation() *CompressedMJLogMutation {
	return cmlc.mutation
}

// Save creates the CompressedMJLog in the database.
func (cmlc *CompressedMJLogCreate) Save(ctx context.Context) (*CompressedMJLog, error) {
	cmlc.defaults()
	return withHooks[*CompressedMJLog, CompressedMJLogMutation](ctx, cmlc.sqlSave, cmlc.mutation, cmlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cmlc *CompressedMJLogCreate) SaveX(ctx context.Context) *CompressedMJLog {
	v, err := cmlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cmlc *CompressedMJLogCreate) Exec(ctx context.Context) error {
	_, err := cmlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmlc *CompressedMJLogCreate) ExecX(ctx context.Context) {
	if err := cmlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cmlc *CompressedMJLogCreate) defaults() {
	if _, ok := cmlc.mutation.InsertedAt(); !ok {
		v := compressedmjlog.DefaultInsertedAt()
		cmlc.mutation.SetInsertedAt(v)
	}
	if _, ok := cmlc.mutation.ID(); !ok {
		v := compressedmjlog.DefaultID()
		cmlc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmlc *CompressedMJLogCreate) check() error {
	if _, ok := cmlc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CompressedMJLog.name"`)}
	}
	if _, ok := cmlc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "CompressedMJLog.size"`)}
	}
	if _, ok := cmlc.mutation.InsertedAt(); !ok {
		return &ValidationError{Name: "inserted_at", err: errors.New(`ent: missing required field "CompressedMJLog.inserted_at"`)}
	}
	return nil
}

func (cmlc *CompressedMJLogCreate) sqlSave(ctx context.Context) (*CompressedMJLog, error) {
	if err := cmlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cmlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cmlc.driver, _spec); err != nil {
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
	cmlc.mutation.id = &_node.ID
	cmlc.mutation.done = true
	return _node, nil
}

func (cmlc *CompressedMJLogCreate) createSpec() (*CompressedMJLog, *sqlgraph.CreateSpec) {
	var (
		_node = &CompressedMJLog{config: cmlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: compressedmjlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compressedmjlog.FieldID,
			},
		}
	)
	_spec.OnConflict = cmlc.conflict
	if id, ok := cmlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cmlc.mutation.Name(); ok {
		_spec.SetField(compressedmjlog.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cmlc.mutation.Size(); ok {
		_spec.SetField(compressedmjlog.FieldSize, field.TypeUint, value)
		_node.Size = value
	}
	if value, ok := cmlc.mutation.InsertedAt(); ok {
		_spec.SetField(compressedmjlog.FieldInsertedAt, field.TypeTime, value)
		_node.InsertedAt = value
	}
	if nodes := cmlc.mutation.MjlogFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   compressedmjlog.MjlogFilesTable,
			Columns: []string{compressedmjlog.MjlogFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mjlogfile.FieldID,
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
//	client.CompressedMJLog.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CompressedMJLogUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (cmlc *CompressedMJLogCreate) OnConflict(opts ...sql.ConflictOption) *CompressedMJLogUpsertOne {
	cmlc.conflict = opts
	return &CompressedMJLogUpsertOne{
		create: cmlc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CompressedMJLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cmlc *CompressedMJLogCreate) OnConflictColumns(columns ...string) *CompressedMJLogUpsertOne {
	cmlc.conflict = append(cmlc.conflict, sql.ConflictColumns(columns...))
	return &CompressedMJLogUpsertOne{
		create: cmlc,
	}
}

type (
	// CompressedMJLogUpsertOne is the builder for "upsert"-ing
	//  one CompressedMJLog node.
	CompressedMJLogUpsertOne struct {
		create *CompressedMJLogCreate
	}

	// CompressedMJLogUpsert is the "OnConflict" setter.
	CompressedMJLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetSize sets the "size" field.
func (u *CompressedMJLogUpsert) SetSize(v uint) *CompressedMJLogUpsert {
	u.Set(compressedmjlog.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *CompressedMJLogUpsert) UpdateSize() *CompressedMJLogUpsert {
	u.SetExcluded(compressedmjlog.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *CompressedMJLogUpsert) AddSize(v uint) *CompressedMJLogUpsert {
	u.Add(compressedmjlog.FieldSize, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CompressedMJLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(compressedmjlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CompressedMJLogUpsertOne) UpdateNewValues() *CompressedMJLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(compressedmjlog.FieldID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(compressedmjlog.FieldName)
		}
		if _, exists := u.create.mutation.InsertedAt(); exists {
			s.SetIgnore(compressedmjlog.FieldInsertedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CompressedMJLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CompressedMJLogUpsertOne) Ignore() *CompressedMJLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CompressedMJLogUpsertOne) DoNothing() *CompressedMJLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CompressedMJLogCreate.OnConflict
// documentation for more info.
func (u *CompressedMJLogUpsertOne) Update(set func(*CompressedMJLogUpsert)) *CompressedMJLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CompressedMJLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetSize sets the "size" field.
func (u *CompressedMJLogUpsertOne) SetSize(v uint) *CompressedMJLogUpsertOne {
	return u.Update(func(s *CompressedMJLogUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *CompressedMJLogUpsertOne) AddSize(v uint) *CompressedMJLogUpsertOne {
	return u.Update(func(s *CompressedMJLogUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *CompressedMJLogUpsertOne) UpdateSize() *CompressedMJLogUpsertOne {
	return u.Update(func(s *CompressedMJLogUpsert) {
		s.UpdateSize()
	})
}

// Exec executes the query.
func (u *CompressedMJLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CompressedMJLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CompressedMJLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CompressedMJLogUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CompressedMJLogUpsertOne.ID is not supported by MySQL driver. Use CompressedMJLogUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CompressedMJLogUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CompressedMJLogCreateBulk is the builder for creating many CompressedMJLog entities in bulk.
type CompressedMJLogCreateBulk struct {
	config
	builders []*CompressedMJLogCreate
	conflict []sql.ConflictOption
}

// Save creates the CompressedMJLog entities in the database.
func (cmlcb *CompressedMJLogCreateBulk) Save(ctx context.Context) ([]*CompressedMJLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cmlcb.builders))
	nodes := make([]*CompressedMJLog, len(cmlcb.builders))
	mutators := make([]Mutator, len(cmlcb.builders))
	for i := range cmlcb.builders {
		func(i int, root context.Context) {
			builder := cmlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompressedMJLogMutation)
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
					_, err = mutators[i+1].Mutate(root, cmlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cmlcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cmlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cmlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cmlcb *CompressedMJLogCreateBulk) SaveX(ctx context.Context) []*CompressedMJLog {
	v, err := cmlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cmlcb *CompressedMJLogCreateBulk) Exec(ctx context.Context) error {
	_, err := cmlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmlcb *CompressedMJLogCreateBulk) ExecX(ctx context.Context) {
	if err := cmlcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CompressedMJLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CompressedMJLogUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (cmlcb *CompressedMJLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *CompressedMJLogUpsertBulk {
	cmlcb.conflict = opts
	return &CompressedMJLogUpsertBulk{
		create: cmlcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CompressedMJLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cmlcb *CompressedMJLogCreateBulk) OnConflictColumns(columns ...string) *CompressedMJLogUpsertBulk {
	cmlcb.conflict = append(cmlcb.conflict, sql.ConflictColumns(columns...))
	return &CompressedMJLogUpsertBulk{
		create: cmlcb,
	}
}

// CompressedMJLogUpsertBulk is the builder for "upsert"-ing
// a bulk of CompressedMJLog nodes.
type CompressedMJLogUpsertBulk struct {
	create *CompressedMJLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CompressedMJLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(compressedmjlog.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CompressedMJLogUpsertBulk) UpdateNewValues() *CompressedMJLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(compressedmjlog.FieldID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(compressedmjlog.FieldName)
			}
			if _, exists := b.mutation.InsertedAt(); exists {
				s.SetIgnore(compressedmjlog.FieldInsertedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CompressedMJLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CompressedMJLogUpsertBulk) Ignore() *CompressedMJLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CompressedMJLogUpsertBulk) DoNothing() *CompressedMJLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CompressedMJLogCreateBulk.OnConflict
// documentation for more info.
func (u *CompressedMJLogUpsertBulk) Update(set func(*CompressedMJLogUpsert)) *CompressedMJLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CompressedMJLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetSize sets the "size" field.
func (u *CompressedMJLogUpsertBulk) SetSize(v uint) *CompressedMJLogUpsertBulk {
	return u.Update(func(s *CompressedMJLogUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *CompressedMJLogUpsertBulk) AddSize(v uint) *CompressedMJLogUpsertBulk {
	return u.Update(func(s *CompressedMJLogUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *CompressedMJLogUpsertBulk) UpdateSize() *CompressedMJLogUpsertBulk {
	return u.Update(func(s *CompressedMJLogUpsert) {
		s.UpdateSize()
	})
}

// Exec executes the query.
func (u *CompressedMJLogUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CompressedMJLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CompressedMJLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CompressedMJLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
