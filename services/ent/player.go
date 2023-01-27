// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/player"
)

// Player is the model entity for the Player schema.
type Player struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Sex holds the value of the "sex" field.
	Sex string `json:"sex,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerQuery when eager-loading is set.
	Edges PlayerEdges `json:"edges"`
}

// PlayerEdges holds the relations/edges for other nodes in the graph.
type PlayerEdges struct {
	// GamePlayers holds the value of the game_players edge.
	GamePlayers []*GamePlayer `json:"game_players,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GamePlayersOrErr returns the GamePlayers value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) GamePlayersOrErr() ([]*GamePlayer, error) {
	if e.loadedTypes[0] {
		return e.GamePlayers, nil
	}
	return nil, &NotLoadedError{edge: "game_players"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Player) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case player.FieldName, player.FieldSex:
			values[i] = new(sql.NullString)
		case player.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Player", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Player fields.
func (pl *Player) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case player.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pl.ID = *value
			}
		case player.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case player.FieldSex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sex", values[i])
			} else if value.Valid {
				pl.Sex = value.String
			}
		}
	}
	return nil
}

// QueryGamePlayers queries the "game_players" edge of the Player entity.
func (pl *Player) QueryGamePlayers() *GamePlayerQuery {
	return (&PlayerClient{config: pl.config}).QueryGamePlayers(pl)
}

// Update returns a builder for updating this Player.
// Note that you need to call Player.Unwrap() before calling this method if this Player
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Player) Update() *PlayerUpdateOne {
	return (&PlayerClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Player entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Player) Unwrap() *Player {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Player is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Player) String() string {
	var builder strings.Builder
	builder.WriteString("Player(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("sex=")
	builder.WriteString(pl.Sex)
	builder.WriteByte(')')
	return builder.String()
}

// Players is a parsable slice of Player.
type Players []*Player

func (pl Players) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}
