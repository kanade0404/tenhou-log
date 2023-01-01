// Code generated by ent, DO NOT EDIT.

package wind

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the wind type in the database.
	Label = "wind"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the wind in the database.
	Table = "winds"
)

// Columns holds all SQL columns for wind fields.
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
