// Code generated by ent, DO NOT EDIT.

package mjlogfile

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the mjlogfile type in the database.
	Label = "mj_log_file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCompressedMjlogFiles holds the string denoting the compressed_mjlog_files edge name in mutations.
	EdgeCompressedMjlogFiles = "compressed_mjlog_files"
	// Table holds the table name of the mjlogfile in the database.
	Table = "mj_log_files"
	// CompressedMjlogFilesTable is the table that holds the compressed_mjlog_files relation/edge.
	CompressedMjlogFilesTable = "mj_log_files"
	// CompressedMjlogFilesInverseTable is the table name for the CompressedMJLog entity.
	// It exists in this package in order to avoid circular dependency with the "compressedmjlog" package.
	CompressedMjlogFilesInverseTable = "compressed_mj_logs"
	// CompressedMjlogFilesColumn is the table column denoting the compressed_mjlog_files relation/edge.
	CompressedMjlogFilesColumn = "compressed_mj_log_mjlog_files"
)

// Columns holds all SQL columns for mjlogfile fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "mj_log_files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"compressed_mj_log_mjlog_files",
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
