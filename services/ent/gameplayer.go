// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/gameplayer"
)

// GamePlayer is the model entity for the GamePlayer schema.
type GamePlayer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Rate holds the value of the "rate" field.
	Rate float64 `json:"rate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GamePlayerQuery when eager-loading is set.
	Edges GamePlayerEdges `json:"edges"`
}

// GamePlayerEdges holds the relations/edges for other nodes in the graph.
type GamePlayerEdges struct {
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e GamePlayerEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[0] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GamePlayer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gameplayer.FieldRate:
			values[i] = new(sql.NullFloat64)
		case gameplayer.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GamePlayer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GamePlayer fields.
func (gp *GamePlayer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gameplayer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gp.ID = *value
			}
		case gameplayer.FieldRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rate", values[i])
			} else if value.Valid {
				gp.Rate = value.Float64
			}
		}
	}
	return nil
}

// QueryPlayers queries the "players" edge of the GamePlayer entity.
func (gp *GamePlayer) QueryPlayers() *PlayerQuery {
	return (&GamePlayerClient{config: gp.config}).QueryPlayers(gp)
}

// Update returns a builder for updating this GamePlayer.
// Note that you need to call GamePlayer.Unwrap() before calling this method if this GamePlayer
// was returned from a transaction, and the transaction was committed or rolled back.
func (gp *GamePlayer) Update() *GamePlayerUpdateOne {
	return (&GamePlayerClient{config: gp.config}).UpdateOne(gp)
}

// Unwrap unwraps the GamePlayer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gp *GamePlayer) Unwrap() *GamePlayer {
	_tx, ok := gp.config.driver.(*txDriver)
	if !ok {
		panic("ent: GamePlayer is not a transactional entity")
	}
	gp.config.driver = _tx.drv
	return gp
}

// String implements the fmt.Stringer.
func (gp *GamePlayer) String() string {
	var builder strings.Builder
	builder.WriteString("GamePlayer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gp.ID))
	builder.WriteString("rate=")
	builder.WriteString(fmt.Sprintf("%v", gp.Rate))
	builder.WriteByte(')')
	return builder.String()
}

// GamePlayers is a parsable slice of GamePlayer.
type GamePlayers []*GamePlayer

func (gp GamePlayers) config(cfg config) {
	for _i := range gp {
		gp[_i].config = cfg
	}
}
