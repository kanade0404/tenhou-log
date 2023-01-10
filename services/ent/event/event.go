// Code generated by ent, DO NOT EDIT.

package event

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeTurn holds the string denoting the turn edge name in mutations.
	EdgeTurn = "turn"
	// EdgeWin holds the string denoting the win edge name in mutations.
	EdgeWin = "win"
	// EdgeCall holds the string denoting the call edge name in mutations.
	EdgeCall = "call"
	// EdgeDraw holds the string denoting the draw edge name in mutations.
	EdgeDraw = "draw"
	// EdgeReach holds the string denoting the reach edge name in mutations.
	EdgeReach = "reach"
	// Table holds the table name of the event in the database.
	Table = "events"
	// TurnTable is the table that holds the turn relation/edge.
	TurnTable = "events"
	// TurnInverseTable is the table name for the Turn entity.
	// It exists in this package in order to avoid circular dependency with the "turn" package.
	TurnInverseTable = "turns"
	// TurnColumn is the table column denoting the turn relation/edge.
	TurnColumn = "turn_event"
	// WinTable is the table that holds the win relation/edge.
	WinTable = "wins"
	// WinInverseTable is the table name for the Win entity.
	// It exists in this package in order to avoid circular dependency with the "win" package.
	WinInverseTable = "wins"
	// WinColumn is the table column denoting the win relation/edge.
	WinColumn = "event_win"
	// CallTable is the table that holds the call relation/edge.
	CallTable = "calls"
	// CallInverseTable is the table name for the Call entity.
	// It exists in this package in order to avoid circular dependency with the "call" package.
	CallInverseTable = "calls"
	// CallColumn is the table column denoting the call relation/edge.
	CallColumn = "event_call"
	// DrawTable is the table that holds the draw relation/edge.
	DrawTable = "drawns"
	// DrawInverseTable is the table name for the Drawn entity.
	// It exists in this package in order to avoid circular dependency with the "drawn" package.
	DrawInverseTable = "drawns"
	// DrawColumn is the table column denoting the draw relation/edge.
	DrawColumn = "event_draw"
	// ReachTable is the table that holds the reach relation/edge.
	ReachTable = "reaches"
	// ReachInverseTable is the table name for the Reach entity.
	// It exists in this package in order to avoid circular dependency with the "reach" package.
	ReachInverseTable = "reaches"
	// ReachColumn is the table column denoting the reach relation/edge.
	ReachColumn = "event_reach"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"turn_event",
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
