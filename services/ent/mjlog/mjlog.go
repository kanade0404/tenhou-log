// Code generated by ent, DO NOT EDIT.

package mjlog

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the mjlog type in the database.
	Label = "mj_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldSeed holds the string denoting the seed field in the database.
	FieldSeed = "seed"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldInsertedAt holds the string denoting the inserted_at field in the database.
	FieldInsertedAt = "inserted_at"
	// EdgeMjlogFiles holds the string denoting the mjlog_files edge name in mutations.
	EdgeMjlogFiles = "mjlog_files"
	// EdgeGames holds the string denoting the games edge name in mutations.
	EdgeGames = "games"
	// Table holds the table name of the mjlog in the database.
	Table = "mj_logs"
	// MjlogFilesTable is the table that holds the mjlog_files relation/edge.
	MjlogFilesTable = "mj_logs"
	// MjlogFilesInverseTable is the table name for the MJLogFile entity.
	// It exists in this package in order to avoid circular dependency with the "mjlogfile" package.
	MjlogFilesInverseTable = "mj_log_files"
	// MjlogFilesColumn is the table column denoting the mjlog_files relation/edge.
	MjlogFilesColumn = "mj_log_file_mjlogs"
	// GamesTable is the table that holds the games relation/edge.
	GamesTable = "mj_logs"
	// GamesInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GamesInverseTable = "games"
	// GamesColumn is the table column denoting the games relation/edge.
	GamesColumn = "game_mjlogs"
)

// Columns holds all SQL columns for mjlog fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldSeed,
	FieldStartedAt,
	FieldInsertedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "mj_logs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_mjlogs",
	"mj_log_file_mjlogs",
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
	// DefaultInsertedAt holds the default value on creation for the "inserted_at" field.
	DefaultInsertedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
