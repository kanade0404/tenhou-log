// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/event"
	"github.com/kanade0404/tenhou-log/services/ent/win"
)

// Win is the model entity for the Win schema.
type Win struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WinQuery when eager-loading is set.
	Edges     WinEdges `json:"edges"`
	event_win *uuid.UUID
}

// WinEdges holds the relations/edges for other nodes in the graph.
type WinEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WinEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Win) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case win.FieldID:
			values[i] = new(uuid.UUID)
		case win.ForeignKeys[0]: // event_win
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Win", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Win fields.
func (w *Win) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case win.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				w.ID = *value
			}
		case win.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field event_win", values[i])
			} else if value.Valid {
				w.event_win = new(uuid.UUID)
				*w.event_win = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryEvent queries the "event" edge of the Win entity.
func (w *Win) QueryEvent() *EventQuery {
	return (&WinClient{config: w.config}).QueryEvent(w)
}

// Update returns a builder for updating this Win.
// Note that you need to call Win.Unwrap() before calling this method if this Win
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Win) Update() *WinUpdateOne {
	return (&WinClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Win entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Win) Unwrap() *Win {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Win is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Win) String() string {
	var builder strings.Builder
	builder.WriteString("Win(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Wins is a parsable slice of Win.
type Wins []*Win

func (w Wins) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
