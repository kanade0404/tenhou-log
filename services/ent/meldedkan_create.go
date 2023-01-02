// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/meldedkan"
)

// MeldedKanCreate is the builder for creating a MeldedKan entity.
type MeldedKanCreate struct {
	config
	mutation *MeldedKanMutation
	hooks    []Hook
}

// Mutation returns the MeldedKanMutation object of the builder.
func (mkc *MeldedKanCreate) Mutation() *MeldedKanMutation {
	return mkc.mutation
}

// Save creates the MeldedKan in the database.
func (mkc *MeldedKanCreate) Save(ctx context.Context) (*MeldedKan, error) {
	var (
		err  error
		node *MeldedKan
	)
	if len(mkc.hooks) == 0 {
		if err = mkc.check(); err != nil {
			return nil, err
		}
		node, err = mkc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MeldedKanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mkc.check(); err != nil {
				return nil, err
			}
			mkc.mutation = mutation
			if node, err = mkc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mkc.hooks) - 1; i >= 0; i-- {
			if mkc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mkc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mkc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MeldedKan)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MeldedKanMutation", v)
		}
		node = nv
	}
	return node, err
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

// check runs all checks and user-defined validators on the builder.
func (mkc *MeldedKanCreate) check() error {
	return nil
}

func (mkc *MeldedKanCreate) sqlSave(ctx context.Context) (*MeldedKan, error) {
	_node, _spec := mkc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mkc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mkc *MeldedKanCreate) createSpec() (*MeldedKan, *sqlgraph.CreateSpec) {
	var (
		_node = &MeldedKan{config: mkc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: meldedkan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: meldedkan.FieldID,
			},
		}
	)
	return _node, _spec
}

// MeldedKanCreateBulk is the builder for creating many MeldedKan entities in bulk.
type MeldedKanCreateBulk struct {
	config
	builders []*MeldedKanCreate
}

// Save creates the MeldedKan entities in the database.
func (mkcb *MeldedKanCreateBulk) Save(ctx context.Context) ([]*MeldedKan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mkcb.builders))
	nodes := make([]*MeldedKan, len(mkcb.builders))
	mutators := make([]Mutator, len(mkcb.builders))
	for i := range mkcb.builders {
		func(i int, root context.Context) {
			builder := mkcb.builders[i]
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
