// Code generated by ent, DO NOT EDIT.

package turn

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the turn type in the database.
	Label = "turn"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// EdgeHands holds the string denoting the hands edge name in mutations.
	EdgeHands = "hands"
	// Table holds the table name of the turn in the database.
	Table = "turns"
	// HandsTable is the table that holds the hands relation/edge. The primary key declared below.
	HandsTable = "hand_turns"
	// HandsInverseTable is the table name for the Hand entity.
	// It exists in this package in order to avoid circular dependency with the "hand" package.
	HandsInverseTable = "hands"
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
