// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/discard"
	"github.com/kanade0404/tenhou-log/services/ent/event"
	"github.com/kanade0404/tenhou-log/services/ent/reach"
)

// Reach is the model entity for the Reach schema.
type Reach struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReachQuery when eager-loading is set.
	Edges         ReachEdges `json:"edges"`
	discard_reach *uuid.UUID
	event_reach   *uuid.UUID
}

// ReachEdges holds the relations/edges for other nodes in the graph.
type ReachEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// Discard holds the value of the discard edge.
	Discard *Discard `json:"discard,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReachEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// DiscardOrErr returns the Discard value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReachEdges) DiscardOrErr() (*Discard, error) {
	if e.loadedTypes[1] {
		if e.Discard == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: discard.Label}
		}
		return e.Discard, nil
	}
	return nil, &NotLoadedError{edge: "discard"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Reach) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reach.FieldID:
			values[i] = new(uuid.UUID)
		case reach.ForeignKeys[0]: // discard_reach
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case reach.ForeignKeys[1]: // event_reach
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Reach", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Reach fields.
func (r *Reach) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reach.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case reach.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field discard_reach", values[i])
			} else if value.Valid {
				r.discard_reach = new(uuid.UUID)
				*r.discard_reach = *value.S.(*uuid.UUID)
			}
		case reach.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field event_reach", values[i])
			} else if value.Valid {
				r.event_reach = new(uuid.UUID)
				*r.event_reach = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the Reach entity.
func (r *Reach) QueryEvent() *EventQuery {
	return (&ReachClient{config: r.config}).QueryEvent(r)
}

// QueryDiscard queries the "discard" edge of the Reach entity.
func (r *Reach) QueryDiscard() *DiscardQuery {
	return (&ReachClient{config: r.config}).QueryDiscard(r)
}

// Update returns a builder for updating this Reach.
// Note that you need to call Reach.Unwrap() before calling this method if this Reach
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Reach) Update() *ReachUpdateOne {
	return (&ReachClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Reach entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Reach) Unwrap() *Reach {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Reach is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Reach) String() string {
	var builder strings.Builder
	builder.WriteString("Reach(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Reaches is a parsable slice of Reach.
type Reaches []*Reach

func (r Reaches) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}