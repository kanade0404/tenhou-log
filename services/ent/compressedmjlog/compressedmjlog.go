// Code generated by ent, DO NOT EDIT.

package compressedmjlog

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the compressedmjlog type in the database.
	Label = "compressed_mj_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldInsertedAt holds the string denoting the inserted_at field in the database.
	FieldInsertedAt = "inserted_at"
	// EdgeMjlogFiles holds the string denoting the mjlog_files edge name in mutations.
	EdgeMjlogFiles = "mjlog_files"
	// Table holds the table name of the compressedmjlog in the database.
	Table = "compressed_mj_logs"
	// MjlogFilesTable is the table that holds the mjlog_files relation/edge.
	MjlogFilesTable = "mj_log_files"
	// MjlogFilesInverseTable is the table name for the MJLogFile entity.
	// It exists in this package in order to avoid circular dependency with the "mjlogfile" package.
	MjlogFilesInverseTable = "mj_log_files"
	// MjlogFilesColumn is the table column denoting the mjlog_files relation/edge.
	MjlogFilesColumn = "compressed_mj_log_mjlog_files"
)

// Columns holds all SQL columns for compressedmjlog fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSize,
	FieldInsertedAt,
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
	// DefaultInsertedAt holds the default value on creation for the "inserted_at" field.
	DefaultInsertedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
