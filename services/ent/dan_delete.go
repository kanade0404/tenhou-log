// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/dan"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// DanDelete is the builder for deleting a Dan entity.
type DanDelete struct {
	config
	hooks    []Hook
	mutation *DanMutation
}

// Where appends a list predicates to the DanDelete builder.
func (dd *DanDelete) Where(ps ...predicate.Dan) *DanDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DanDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DanMutation](ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DanDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DanDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: dan.FieldID,
			},
		},
	}
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DanDeleteOne is the builder for deleting a single Dan entity.
type DanDeleteOne struct {
	dd *DanDelete
}

// Exec executes the deletion query.
func (ddo *DanDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dan.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DanDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}