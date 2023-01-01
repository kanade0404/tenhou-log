// Code generated by ent, DO NOT EDIT.

package hand

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the hand type in the database.
	Label = "hand"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// Table holds the table name of the hand in the database.
	Table = "hands"
)

// Columns holds all SQL columns for hand fields.
var Columns = []string{
	FieldID,
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
