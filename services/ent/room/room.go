// Code generated by ent, DO NOT EDIT.

package room

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeGames holds the string denoting the games edge name in mutations.
	EdgeGames = "games"
	// Table holds the table name of the room in the database.
	Table = "rooms"
	// GamesTable is the table that holds the games relation/edge.
	GamesTable = "games"
	// GamesInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GamesInverseTable = "games"
	// GamesColumn is the table column denoting the games relation/edge.
	GamesColumn = "room_games"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldName,
}

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
