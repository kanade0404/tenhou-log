// Code generated by ent, DO NOT EDIT.

package win

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the win type in the database.
	Label = "win"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"
	// Table holds the table name of the win in the database.
	Table = "wins"
	// EventTable is the table that holds the event relation/edge.
	EventTable = "wins"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_win"
)

// Columns holds all SQL columns for win fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "wins"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_win",
}

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
