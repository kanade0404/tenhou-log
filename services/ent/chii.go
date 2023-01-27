// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/call"
	"github.com/kanade0404/tenhou-log/services/ent/chii"
)

// Chii is the model entity for the Chii schema.
type Chii struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChiiQuery when eager-loading is set.
	Edges ChiiEdges `json:"edges"`
}

// ChiiEdges holds the relations/edges for other nodes in the graph.
type ChiiEdges struct {
	// Call holds the value of the call edge.
	Call *Call `json:"call,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CallOrErr returns the Call value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChiiEdges) CallOrErr() (*Call, error) {
	if e.loadedTypes[0] {
		if e.Call == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: call.Label}
		}
		return e.Call, nil
	}
	return nil, &NotLoadedError{edge: "call"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chii) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chii.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Chii", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chii fields.
func (c *Chii) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chii.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		}
	}
	return nil
}

// QueryCall queries the "call" edge of the Chii entity.
func (c *Chii) QueryCall() *CallQuery {
	return (&ChiiClient{config: c.config}).QueryCall(c)
}

// Update returns a builder for updating this Chii.
// Note that you need to call Chii.Unwrap() before calling this method if this Chii
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chii) Update() *ChiiUpdateOne {
	return (&ChiiClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Chii entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chii) Unwrap() *Chii {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chii is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chii) String() string {
	var builder strings.Builder
	builder.WriteString("Chii(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Chiis is a parsable slice of Chii.
type Chiis []*Chii

func (c Chiis) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
