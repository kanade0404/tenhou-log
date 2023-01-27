// Code generated by ent, DO NOT EDIT.

package hand

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the hand type in the database.
	Label = "hand"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// FieldContinuePoint holds the string denoting the continue_point field in the database.
	FieldContinuePoint = "continue_point"
	// FieldReachPoint holds the string denoting the reach_point field in the database.
	FieldReachPoint = "reach_point"
	// EdgeRounds holds the string denoting the rounds edge name in mutations.
	EdgeRounds = "rounds"
	// EdgeTurns holds the string denoting the turns edge name in mutations.
	EdgeTurns = "turns"
	// Table holds the table name of the hand in the database.
	Table = "hands"
	// RoundsTable is the table that holds the rounds relation/edge.
	RoundsTable = "hands"
	// RoundsInverseTable is the table name for the Round entity.
	// It exists in this package in order to avoid circular dependency with the "round" package.
	RoundsInverseTable = "rounds"
	// RoundsColumn is the table column denoting the rounds relation/edge.
	RoundsColumn = "round_hands"
	// TurnsTable is the table that holds the turns relation/edge. The primary key declared below.
	TurnsTable = "hand_turns"
	// TurnsInverseTable is the table name for the Turn entity.
	// It exists in this package in order to avoid circular dependency with the "turn" package.
	TurnsInverseTable = "turns"
)

// Columns holds all SQL columns for hand fields.
var Columns = []string{
	FieldID,
	FieldNum,
	FieldContinuePoint,
	FieldReachPoint,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "hands"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"round_hands",
}

var (
	// TurnsPrimaryKey and TurnsColumn2 are the table columns denoting the
	// primary key for the turns relation (M2M).
	TurnsPrimaryKey = []string{"hand_id", "turn_id"}
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
	// ContinuePointValidator is a validator for the "continue_point" field. It is called by the builders before save.
	ContinuePointValidator func(uint) error
	// ReachPointValidator is a validator for the "reach_point" field. It is called by the builders before save.
	ReachPointValidator func(uint) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
