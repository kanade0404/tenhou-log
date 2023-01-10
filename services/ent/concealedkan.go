// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/call"
	"github.com/kanade0404/tenhou-log/services/ent/concealedkan"
)

// ConcealedKan is the model entity for the ConcealedKan schema.
type ConcealedKan struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConcealedKanQuery when eager-loading is set.
	Edges ConcealedKanEdges `json:"edges"`
}

// ConcealedKanEdges holds the relations/edges for other nodes in the graph.
type ConcealedKanEdges struct {
	// Call holds the value of the call edge.
	Call *Call `json:"call,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CallOrErr returns the Call value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConcealedKanEdges) CallOrErr() (*Call, error) {
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
func (*ConcealedKan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case concealedkan.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ConcealedKan", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ConcealedKan fields.
func (ck *ConcealedKan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case concealedkan.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ck.ID = *value
			}
		}
	}
	return nil
}

// QueryCall queries the "call" edge of the ConcealedKan entity.
func (ck *ConcealedKan) QueryCall() *CallQuery {
	return (&ConcealedKanClient{config: ck.config}).QueryCall(ck)
}

// Update returns a builder for updating this ConcealedKan.
// Note that you need to call ConcealedKan.Unwrap() before calling this method if this ConcealedKan
// was returned from a transaction, and the transaction was committed or rolled back.
func (ck *ConcealedKan) Update() *ConcealedKanUpdateOne {
	return (&ConcealedKanClient{config: ck.config}).UpdateOne(ck)
}

// Unwrap unwraps the ConcealedKan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ck *ConcealedKan) Unwrap() *ConcealedKan {
	_tx, ok := ck.config.driver.(*txDriver)
	if !ok {
		panic("ent: ConcealedKan is not a transactional entity")
	}
	ck.config.driver = _tx.drv
	return ck
}

// String implements the fmt.Stringer.
func (ck *ConcealedKan) String() string {
	var builder strings.Builder
	builder.WriteString("ConcealedKan(")
	builder.WriteString(fmt.Sprintf("id=%v", ck.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ConcealedKans is a parsable slice of ConcealedKan.
type ConcealedKans []*ConcealedKan

func (ck ConcealedKans) config(cfg config) {
	for _i := range ck {
		ck[_i].config = cfg
	}
}
