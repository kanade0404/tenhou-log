// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// CompressedMJLogDelete is the builder for deleting a CompressedMJLog entity.
type CompressedMJLogDelete struct {
	config
	hooks    []Hook
	mutation *CompressedMJLogMutation
}

// Where appends a list predicates to the CompressedMJLogDelete builder.
func (cmld *CompressedMJLogDelete) Where(ps ...predicate.CompressedMJLog) *CompressedMJLogDelete {
	cmld.mutation.Where(ps...)
	return cmld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmld *CompressedMJLogDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cmld.hooks) == 0 {
		affected, err = cmld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompressedMJLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cmld.mutation = mutation
			affected, err = cmld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmld.hooks) - 1; i >= 0; i-- {
			if cmld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cmld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmld *CompressedMJLogDelete) ExecX(ctx context.Context) int {
	n, err := cmld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmld *CompressedMJLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: compressedmjlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compressedmjlog.FieldID,
			},
		},
	}
	if ps := cmld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cmld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CompressedMJLogDeleteOne is the builder for deleting a single CompressedMJLog entity.
type CompressedMJLogDeleteOne struct {
	cmld *CompressedMJLogDelete
}

// Exec executes the deletion query.
func (cmldo *CompressedMJLogDeleteOne) Exec(ctx context.Context) error {
	n, err := cmldo.cmld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{compressedmjlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmldo *CompressedMJLogDeleteOne) ExecX(ctx context.Context) {
	cmldo.cmld.ExecX(ctx)
}
