// Code generated by ent, DO NOT EDIT.

package turn

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the turn type in the database.
	Label = "turn"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// EdgeHands holds the string denoting the hands edge name in mutations.
	EdgeHands = "hands"
	// EdgeGamePlayerPoints holds the string denoting the game_player_points edge name in mutations.
	EdgeGamePlayerPoints = "game_player_points"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// EdgeGameplayerhandhai holds the string denoting the gameplayerhandhai edge name in mutations.
	EdgeGameplayerhandhai = "gameplayerhandhai"
	// Table holds the table name of the turn in the database.
	Table = "turns"
	// HandsTable is the table that holds the hands relation/edge. The primary key declared below.
	HandsTable = "hand_turns"
	// HandsInverseTable is the table name for the Hand entity.
	// It exists in this package in order to avoid circular dependency with the "hand" package.
	HandsInverseTable = "hands"
	// GamePlayerPointsTable is the table that holds the game_player_points relation/edge. The primary key declared below.
	GamePlayerPointsTable = "turn_game_player_points"
	// GamePlayerPointsInverseTable is the table name for the GamePlayerPoint entity.
	// It exists in this package in order to avoid circular dependency with the "gameplayerpoint" package.
	GamePlayerPointsInverseTable = "game_player_points"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "events"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "turn_event"
	// GameplayerhandhaiTable is the table that holds the gameplayerhandhai relation/edge.
	GameplayerhandhaiTable = "game_player_hand_hais"
	// GameplayerhandhaiInverseTable is the table name for the GamePlayerHandHai entity.
	// It exists in this package in order to avoid circular dependency with the "gameplayerhandhai" package.
	GameplayerhandhaiInverseTable = "game_player_hand_hais"
	// GameplayerhandhaiColumn is the table column denoting the gameplayerhandhai relation/edge.
	GameplayerhandhaiColumn = "turn_gameplayerhandhai"
)

// Columns holds all SQL columns for turn fields.
var Columns = []string{
	FieldID,
	FieldNum,
}

var (
	// HandsPrimaryKey and HandsColumn2 are the table columns denoting the
	// primary key for the hands relation (M2M).
	HandsPrimaryKey = []string{"hand_id", "turn_id"}
	// GamePlayerPointsPrimaryKey and GamePlayerPointsColumn2 are the table columns denoting the
	// primary key for the game_player_points relation (M2M).
	GamePlayerPointsPrimaryKey = []string{"turn_id", "game_player_point_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)