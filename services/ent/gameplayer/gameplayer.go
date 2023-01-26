// Code generated by ent, DO NOT EDIT.

package gameplayer

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the gameplayer type in the database.
	Label = "game_player"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRate holds the string denoting the rate field in the database.
	FieldRate = "rate"
	// FieldStartPosition holds the string denoting the start_position field in the database.
	FieldStartPosition = "start_position"
	// EdgeGames holds the string denoting the games edge name in mutations.
	EdgeGames = "games"
	// EdgePlayers holds the string denoting the players edge name in mutations.
	EdgePlayers = "players"
	// EdgeDans holds the string denoting the dans edge name in mutations.
	EdgeDans = "dans"
	// Table holds the table name of the gameplayer in the database.
	Table = "game_players"
	// GamesTable is the table that holds the games relation/edge. The primary key declared below.
	GamesTable = "game_game_players"
	// GamesInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GamesInverseTable = "games"
	// PlayersTable is the table that holds the players relation/edge.
	PlayersTable = "game_players"
	// PlayersInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayersInverseTable = "players"
	// PlayersColumn is the table column denoting the players relation/edge.
	PlayersColumn = "player_game_players"
	// DansTable is the table that holds the dans relation/edge.
	DansTable = "game_players"
	// DansInverseTable is the table name for the Dan entity.
	// It exists in this package in order to avoid circular dependency with the "dan" package.
	DansInverseTable = "dans"
	// DansColumn is the table column denoting the dans relation/edge.
	DansColumn = "dan_game_players"
)

// Columns holds all SQL columns for gameplayer fields.
var Columns = []string{
	FieldID,
	FieldRate,
	FieldStartPosition,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "game_players"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"dan_game_players",
	"player_game_players",
}

var (
	// GamesPrimaryKey and GamesColumn2 are the table columns denoting the
	// primary key for the games relation (M2M).
	GamesPrimaryKey = []string{"game_id", "game_player_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// StartPositionValidator is a validator for the "start_position" field. It is called by the builders before save.
	StartPositionValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)