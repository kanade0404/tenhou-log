// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayerhandhai"
)

// GamePlayerHandHai is the model entity for the GamePlayerHandHai schema.
type GamePlayerHandHai struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GamePlayerHandHai) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gameplayerhandhai.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GamePlayerHandHai", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GamePlayerHandHai fields.
func (gphh *GamePlayerHandHai) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gameplayerhandhai.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gphh.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this GamePlayerHandHai.
// Note that you need to call GamePlayerHandHai.Unwrap() before calling this method if this GamePlayerHandHai
// was returned from a transaction, and the transaction was committed or rolled back.
func (gphh *GamePlayerHandHai) Update() *GamePlayerHandHaiUpdateOne {
	return (&GamePlayerHandHaiClient{config: gphh.config}).UpdateOne(gphh)
}

// Unwrap unwraps the GamePlayerHandHai entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gphh *GamePlayerHandHai) Unwrap() *GamePlayerHandHai {
	_tx, ok := gphh.config.driver.(*txDriver)
	if !ok {
		panic("ent: GamePlayerHandHai is not a transactional entity")
	}
	gphh.config.driver = _tx.drv
	return gphh
}

// String implements the fmt.Stringer.
func (gphh *GamePlayerHandHai) String() string {
	var builder strings.Builder
	builder.WriteString("GamePlayerHandHai(")
	builder.WriteString(fmt.Sprintf("id=%v", gphh.ID))
	builder.WriteByte(')')
	return builder.String()
}

// GamePlayerHandHais is a parsable slice of GamePlayerHandHai.
type GamePlayerHandHais []*GamePlayerHandHai

func (gphh GamePlayerHandHais) config(cfg config) {
	for _i := range gphh {
		gphh[_i].config = cfg
	}
}