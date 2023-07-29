// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

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
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// CallUpdate is the builder for updating Call entities.
type CallUpdate struct {
	config
	hooks    []Hook
	mutation *CallMutation
}

// Where appends a list predicates to the CallUpdate builder.
func (cu *CallUpdate) Where(ps ...predicate.Call) *CallUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (cu *CallUpdate) SetEventID(id uuid.UUID) *CallUpdate {
	cu.mutation.SetEventID(id)
	return cu
}

// SetEvent sets the "event" edge to the Event entity.
func (cu *CallUpdate) SetEvent(e *Event) *CallUpdate {
	return cu.SetEventID(e.ID)
}

// SetDiscardID sets the "discard" edge to the Discard entity by ID.
func (cu *CallUpdate) SetDiscardID(id uuid.UUID) *CallUpdate {
	cu.mutation.SetDiscardID(id)
	return cu
}

// SetDiscard sets the "discard" edge to the Discard entity.
func (cu *CallUpdate) SetDiscard(d *Discard) *CallUpdate {
	return cu.SetDiscardID(d.ID)
}

// SetChiiID sets the "chii" edge to the Chii entity by ID.
func (cu *CallUpdate) SetChiiID(id uuid.UUID) *CallUpdate {
	cu.mutation.SetChiiID(id)
	return cu
}

// SetChii sets the "chii" edge to the Chii entity.
func (cu *CallUpdate) SetChii(c *Chii) *CallUpdate {
	return cu.SetChiiID(c.ID)
}

// SetChakanID sets the "chakan" edge to the Chakan entity by ID.
func (cu *CallUpdate) SetChakanID(id uuid.UUID) *CallUpdate {
	cu.mutation.SetChakanID(id)
	return cu
}

// SetChakan sets the "chakan" edge to the Chakan entity.
func (cu *CallUpdate) SetChakan(c *Chakan) *CallUpdate {
	return cu.SetChakanID(c.ID)
}

// SetConcealedkanID sets the "concealedkan" edge to the ConcealedKan entity by ID.
func (cu *CallUpdate) SetConcealedkanID(id uuid.UUID) *CallUpdate {
	cu.mutation.SetConcealedkanID(id)
	return cu
}

// SetConcealedkan sets the "concealedkan" edge to the ConcealedKan entity.
func (cu *CallUpdate) SetConcealedkan(c *ConcealedKan) *CallUpdate {
	return cu.SetConcealedkanID(c.ID)
}

// SetMeldedkanID sets the "meldedkan" edge to the MeldedKan entity by ID.
func (cu *CallUpdate) SetMeldedkanID(id uuid.UUID) *CallUpdate {
	cu.mutation.SetMeldedkanID(id)
	return cu
}

// SetMeldedkan sets the "meldedkan" edge to the MeldedKan entity.
func (cu *CallUpdate) SetMeldedkan(m *MeldedKan) *CallUpdate {
	return cu.SetMeldedkanID(m.ID)
}

// SetPonID sets the "pon" edge to the Pon entity by ID.
func (cu *CallUpdate) SetPonID(id uuid.UUID) *CallUpdate {
	cu.mutation.SetPonID(id)
	return cu
}

// SetPon sets the "pon" edge to the Pon entity.
func (cu *CallUpdate) SetPon(p *Pon) *CallUpdate {
	return cu.SetPonID(p.ID)
}

// Mutation returns the CallMutation object of the builder.
func (cu *CallUpdate) Mutation() *CallMutation {
	return cu.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (cu *CallUpdate) ClearEvent() *CallUpdate {
	cu.mutation.ClearEvent()
	return cu
}

// ClearDiscard clears the "discard" edge to the Discard entity.
func (cu *CallUpdate) ClearDiscard() *CallUpdate {
	cu.mutation.ClearDiscard()
	return cu
}

// ClearChii clears the "chii" edge to the Chii entity.
func (cu *CallUpdate) ClearChii() *CallUpdate {
	cu.mutation.ClearChii()
	return cu
}

// ClearChakan clears the "chakan" edge to the Chakan entity.
func (cu *CallUpdate) ClearChakan() *CallUpdate {
	cu.mutation.ClearChakan()
	return cu
}

// ClearConcealedkan clears the "concealedkan" edge to the ConcealedKan entity.
func (cu *CallUpdate) ClearConcealedkan() *CallUpdate {
	cu.mutation.ClearConcealedkan()
	return cu
}

// ClearMeldedkan clears the "meldedkan" edge to the MeldedKan entity.
func (cu *CallUpdate) ClearMeldedkan() *CallUpdate {
	cu.mutation.ClearMeldedkan()
	return cu
}

// ClearPon clears the "pon" edge to the Pon entity.
func (cu *CallUpdate) ClearPon() *CallUpdate {
	cu.mutation.ClearPon()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CallUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CallMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CallUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CallUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CallUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CallUpdate) check() error {
	if _, ok := cu.mutation.EventID(); cu.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.event"`)
	}
	if _, ok := cu.mutation.DiscardID(); cu.mutation.DiscardCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.discard"`)
	}
	if _, ok := cu.mutation.ChiiID(); cu.mutation.ChiiCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.chii"`)
	}
	if _, ok := cu.mutation.ChakanID(); cu.mutation.ChakanCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.chakan"`)
	}
	if _, ok := cu.mutation.ConcealedkanID(); cu.mutation.ConcealedkanCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.concealedkan"`)
	}
	if _, ok := cu.mutation.MeldedkanID(); cu.mutation.MeldedkanCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.meldedkan"`)
	}
	if _, ok := cu.mutation.PonID(); cu.mutation.PonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.pon"`)
	}
	return nil
}

func (cu *CallUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(call.Table, call.Columns, sqlgraph.NewFieldSpec(call.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   call.EventTable,
			Columns: []string{call.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   call.EventTable,
			Columns: []string{call.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.DiscardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.DiscardTable,
			Columns: []string{call.DiscardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discard.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.DiscardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.DiscardTable,
			Columns: []string{call.DiscardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discard.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ChiiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChiiTable,
			Columns: []string{call.ChiiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chii.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChiiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChiiTable,
			Columns: []string{call.ChiiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chii.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ChakanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChakanTable,
			Columns: []string{call.ChakanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chakan.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChakanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChakanTable,
			Columns: []string{call.ChakanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chakan.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ConcealedkanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ConcealedkanTable,
			Columns: []string{call.ConcealedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(concealedkan.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ConcealedkanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ConcealedkanTable,
			Columns: []string{call.ConcealedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(concealedkan.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.MeldedkanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.MeldedkanTable,
			Columns: []string{call.MeldedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meldedkan.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MeldedkanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.MeldedkanTable,
			Columns: []string{call.MeldedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meldedkan.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.PonTable,
			Columns: []string{call.PonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pon.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.PonTable,
			Columns: []string{call.PonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pon.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{call.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CallUpdateOne is the builder for updating a single Call entity.
type CallUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CallMutation
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (cuo *CallUpdateOne) SetEventID(id uuid.UUID) *CallUpdateOne {
	cuo.mutation.SetEventID(id)
	return cuo
}

// SetEvent sets the "event" edge to the Event entity.
func (cuo *CallUpdateOne) SetEvent(e *Event) *CallUpdateOne {
	return cuo.SetEventID(e.ID)
}

// SetDiscardID sets the "discard" edge to the Discard entity by ID.
func (cuo *CallUpdateOne) SetDiscardID(id uuid.UUID) *CallUpdateOne {
	cuo.mutation.SetDiscardID(id)
	return cuo
}

// SetDiscard sets the "discard" edge to the Discard entity.
func (cuo *CallUpdateOne) SetDiscard(d *Discard) *CallUpdateOne {
	return cuo.SetDiscardID(d.ID)
}

// SetChiiID sets the "chii" edge to the Chii entity by ID.
func (cuo *CallUpdateOne) SetChiiID(id uuid.UUID) *CallUpdateOne {
	cuo.mutation.SetChiiID(id)
	return cuo
}

// SetChii sets the "chii" edge to the Chii entity.
func (cuo *CallUpdateOne) SetChii(c *Chii) *CallUpdateOne {
	return cuo.SetChiiID(c.ID)
}

// SetChakanID sets the "chakan" edge to the Chakan entity by ID.
func (cuo *CallUpdateOne) SetChakanID(id uuid.UUID) *CallUpdateOne {
	cuo.mutation.SetChakanID(id)
	return cuo
}

// SetChakan sets the "chakan" edge to the Chakan entity.
func (cuo *CallUpdateOne) SetChakan(c *Chakan) *CallUpdateOne {
	return cuo.SetChakanID(c.ID)
}

// SetConcealedkanID sets the "concealedkan" edge to the ConcealedKan entity by ID.
func (cuo *CallUpdateOne) SetConcealedkanID(id uuid.UUID) *CallUpdateOne {
	cuo.mutation.SetConcealedkanID(id)
	return cuo
}

// SetConcealedkan sets the "concealedkan" edge to the ConcealedKan entity.
func (cuo *CallUpdateOne) SetConcealedkan(c *ConcealedKan) *CallUpdateOne {
	return cuo.SetConcealedkanID(c.ID)
}

// SetMeldedkanID sets the "meldedkan" edge to the MeldedKan entity by ID.
func (cuo *CallUpdateOne) SetMeldedkanID(id uuid.UUID) *CallUpdateOne {
	cuo.mutation.SetMeldedkanID(id)
	return cuo
}

// SetMeldedkan sets the "meldedkan" edge to the MeldedKan entity.
func (cuo *CallUpdateOne) SetMeldedkan(m *MeldedKan) *CallUpdateOne {
	return cuo.SetMeldedkanID(m.ID)
}

// SetPonID sets the "pon" edge to the Pon entity by ID.
func (cuo *CallUpdateOne) SetPonID(id uuid.UUID) *CallUpdateOne {
	cuo.mutation.SetPonID(id)
	return cuo
}

// SetPon sets the "pon" edge to the Pon entity.
func (cuo *CallUpdateOne) SetPon(p *Pon) *CallUpdateOne {
	return cuo.SetPonID(p.ID)
}

// Mutation returns the CallMutation object of the builder.
func (cuo *CallUpdateOne) Mutation() *CallMutation {
	return cuo.mutation
}

// ClearEvent clears the "event" edge to the Event entity.
func (cuo *CallUpdateOne) ClearEvent() *CallUpdateOne {
	cuo.mutation.ClearEvent()
	return cuo
}

// ClearDiscard clears the "discard" edge to the Discard entity.
func (cuo *CallUpdateOne) ClearDiscard() *CallUpdateOne {
	cuo.mutation.ClearDiscard()
	return cuo
}

// ClearChii clears the "chii" edge to the Chii entity.
func (cuo *CallUpdateOne) ClearChii() *CallUpdateOne {
	cuo.mutation.ClearChii()
	return cuo
}

// ClearChakan clears the "chakan" edge to the Chakan entity.
func (cuo *CallUpdateOne) ClearChakan() *CallUpdateOne {
	cuo.mutation.ClearChakan()
	return cuo
}

// ClearConcealedkan clears the "concealedkan" edge to the ConcealedKan entity.
func (cuo *CallUpdateOne) ClearConcealedkan() *CallUpdateOne {
	cuo.mutation.ClearConcealedkan()
	return cuo
}

// ClearMeldedkan clears the "meldedkan" edge to the MeldedKan entity.
func (cuo *CallUpdateOne) ClearMeldedkan() *CallUpdateOne {
	cuo.mutation.ClearMeldedkan()
	return cuo
}

// ClearPon clears the "pon" edge to the Pon entity.
func (cuo *CallUpdateOne) ClearPon() *CallUpdateOne {
	cuo.mutation.ClearPon()
	return cuo
}

// Where appends a list predicates to the CallUpdate builder.
func (cuo *CallUpdateOne) Where(ps ...predicate.Call) *CallUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CallUpdateOne) Select(field string, fields ...string) *CallUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Call entity.
func (cuo *CallUpdateOne) Save(ctx context.Context) (*Call, error) {
	return withHooks[*Call, CallMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CallUpdateOne) SaveX(ctx context.Context) *Call {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CallUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CallUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CallUpdateOne) check() error {
	if _, ok := cuo.mutation.EventID(); cuo.mutation.EventCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.event"`)
	}
	if _, ok := cuo.mutation.DiscardID(); cuo.mutation.DiscardCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.discard"`)
	}
	if _, ok := cuo.mutation.ChiiID(); cuo.mutation.ChiiCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.chii"`)
	}
	if _, ok := cuo.mutation.ChakanID(); cuo.mutation.ChakanCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.chakan"`)
	}
	if _, ok := cuo.mutation.ConcealedkanID(); cuo.mutation.ConcealedkanCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.concealedkan"`)
	}
	if _, ok := cuo.mutation.MeldedkanID(); cuo.mutation.MeldedkanCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.meldedkan"`)
	}
	if _, ok := cuo.mutation.PonID(); cuo.mutation.PonCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Call.pon"`)
	}
	return nil
}

func (cuo *CallUpdateOne) sqlSave(ctx context.Context) (_node *Call, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(call.Table, call.Columns, sqlgraph.NewFieldSpec(call.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Call.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, call.FieldID)
		for _, f := range fields {
			if !call.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != call.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   call.EventTable,
			Columns: []string{call.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   call.EventTable,
			Columns: []string{call.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.DiscardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.DiscardTable,
			Columns: []string{call.DiscardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discard.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.DiscardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.DiscardTable,
			Columns: []string{call.DiscardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discard.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ChiiCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChiiTable,
			Columns: []string{call.ChiiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chii.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChiiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChiiTable,
			Columns: []string{call.ChiiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chii.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ChakanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChakanTable,
			Columns: []string{call.ChakanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chakan.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChakanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ChakanTable,
			Columns: []string{call.ChakanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chakan.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ConcealedkanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ConcealedkanTable,
			Columns: []string{call.ConcealedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(concealedkan.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ConcealedkanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.ConcealedkanTable,
			Columns: []string{call.ConcealedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(concealedkan.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.MeldedkanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.MeldedkanTable,
			Columns: []string{call.MeldedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meldedkan.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MeldedkanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.MeldedkanTable,
			Columns: []string{call.MeldedkanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meldedkan.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.PonTable,
			Columns: []string{call.PonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pon.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   call.PonTable,
			Columns: []string{call.PonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pon.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Call{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{call.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
