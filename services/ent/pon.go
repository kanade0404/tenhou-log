// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/call"
	"github.com/kanade0404/tenhou-log/services/ent/pon"
)

// Pon is the model entity for the Pon schema.
type Pon struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PonQuery when eager-loading is set.
	Edges PonEdges `json:"edges"`
}

// PonEdges holds the relations/edges for other nodes in the graph.
type PonEdges struct {
	// Call holds the value of the call edge.
	Call *Call `json:"call,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CallOrErr returns the Call value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PonEdges) CallOrErr() (*Call, error) {
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
func (*Pon) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pon.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pon", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pon fields.
func (po *Pon) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pon.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				po.ID = *value
			}
		}
	}
	return nil
}

// QueryCall queries the "call" edge of the Pon entity.
func (po *Pon) QueryCall() *CallQuery {
	return (&PonClient{config: po.config}).QueryCall(po)
}

// Update returns a builder for updating this Pon.
// Note that you need to call Pon.Unwrap() before calling this method if this Pon
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pon) Update() *PonUpdateOne {
	return (&PonClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Pon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pon) Unwrap() *Pon {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pon is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pon) String() string {
	var builder strings.Builder
	builder.WriteString("Pon(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Pons is a parsable slice of Pon.
type Pons []*Pon

func (po Pons) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
