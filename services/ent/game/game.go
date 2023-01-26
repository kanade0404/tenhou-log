// Code generated by ent, DO NOT EDIT.

package game

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the game type in the database.
	Label = "game"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// EdgeMjlogs holds the string denoting the mjlogs edge name in mutations.
	EdgeMjlogs = "mjlogs"
	// EdgeGamePlayers holds the string denoting the game_players edge name in mutations.
	EdgeGamePlayers = "game_players"
	// EdgeRooms holds the string denoting the rooms edge name in mutations.
	EdgeRooms = "rooms"
	// EdgeRounds holds the string denoting the rounds edge name in mutations.
	EdgeRounds = "rounds"
	// Table holds the table name of the game in the database.
	Table = "games"
	// MjlogsTable is the table that holds the mjlogs relation/edge.
	MjlogsTable = "mj_logs"
	// MjlogsInverseTable is the table name for the MJLog entity.
	// It exists in this package in order to avoid circular dependency with the "mjlog" package.
	MjlogsInverseTable = "mj_logs"
	// MjlogsColumn is the table column denoting the mjlogs relation/edge.
	MjlogsColumn = "game_mjlogs"
	// GamePlayersTable is the table that holds the game_players relation/edge. The primary key declared below.
	GamePlayersTable = "game_game_players"
	// GamePlayersInverseTable is the table name for the GamePlayer entity.
	// It exists in this package in order to avoid circular dependency with the "gameplayer" package.
	GamePlayersInverseTable = "game_players"
	// RoomsTable is the table that holds the rooms relation/edge.
	RoomsTable = "games"
	// RoomsInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomsInverseTable = "rooms"
	// RoomsColumn is the table column denoting the rooms relation/edge.
	RoomsColumn = "room_games"
	// RoundsTable is the table that holds the rounds relation/edge.
	RoundsTable = "rounds"
	// RoundsInverseTable is the table name for the Round entity.
	// It exists in this package in order to avoid circular dependency with the "round" package.
	RoundsInverseTable = "rounds"
	// RoundsColumn is the table column denoting the rounds relation/edge.
	RoundsColumn = "game_rounds"
)

// Columns holds all SQL columns for game fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStartedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "games"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"room_games",
}

var (
	// GamePlayersPrimaryKey and GamePlayersColumn2 are the table columns denoting the
	// primary key for the game_players relation (M2M).
	GamePlayersPrimaryKey = []string{"game_id", "game_player_id"}
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)