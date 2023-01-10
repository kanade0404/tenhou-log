// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/turn"
)

// Turn is the model entity for the Turn schema.
type Turn struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Num holds the value of the "num" field.
	Num uint `json:"num,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TurnQuery when eager-loading is set.
	Edges TurnEdges `json:"edges"`
}

// TurnEdges holds the relations/edges for other nodes in the graph.
type TurnEdges struct {
	// Hands holds the value of the hands edge.
	Hands []*Hand `json:"hands,omitempty"`
	// GamePlayerPoints holds the value of the game_player_points edge.
	GamePlayerPoints []*GamePlayerPoint `json:"game_player_points,omitempty"`
	// Event holds the value of the event edge.
	Event []*Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// HandsOrErr returns the Hands value or an error if the edge
// was not loaded in eager-loading.
func (e TurnEdges) HandsOrErr() ([]*Hand, error) {
	if e.loadedTypes[0] {
		return e.Hands, nil
	}
	return nil, &NotLoadedError{edge: "hands"}
}

// GamePlayerPointsOrErr returns the GamePlayerPoints value or an error if the edge
// was not loaded in eager-loading.
func (e TurnEdges) GamePlayerPointsOrErr() ([]*GamePlayerPoint, error) {
	if e.loadedTypes[1] {
		return e.GamePlayerPoints, nil
	}
	return nil, &NotLoadedError{edge: "game_player_points"}
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading.
func (e TurnEdges) EventOrErr() ([]*Event, error) {
	if e.loadedTypes[2] {
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Turn) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case turn.FieldNum:
			values[i] = new(sql.NullInt64)
		case turn.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Turn", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Turn fields.
func (t *Turn) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case turn.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case turn.FieldNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num", values[i])
			} else if value.Valid {
				t.Num = uint(value.Int64)
			}
		}
	}
	return nil
}

// QueryHands queries the "hands" edge of the Turn entity.
func (t *Turn) QueryHands() *HandQuery {
	return (&TurnClient{config: t.config}).QueryHands(t)
}

// QueryGamePlayerPoints queries the "game_player_points" edge of the Turn entity.
func (t *Turn) QueryGamePlayerPoints() *GamePlayerPointQuery {
	return (&TurnClient{config: t.config}).QueryGamePlayerPoints(t)
}

// QueryEvent queries the "event" edge of the Turn entity.
func (t *Turn) QueryEvent() *EventQuery {
	return (&TurnClient{config: t.config}).QueryEvent(t)
}

// Update returns a builder for updating this Turn.
// Note that you need to call Turn.Unwrap() before calling this method if this Turn
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Turn) Update() *TurnUpdateOne {
	return (&TurnClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Turn entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Turn) Unwrap() *Turn {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Turn is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Turn) String() string {
	var builder strings.Builder
	builder.WriteString("Turn(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("num=")
	builder.WriteString(fmt.Sprintf("%v", t.Num))
	builder.WriteByte(')')
	return builder.String()
}

// Turns is a parsable slice of Turn.
type Turns []*Turn

func (t Turns) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
