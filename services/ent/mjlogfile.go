// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
)

// MJLogFile is the model entity for the MJLogFile schema.
type MJLogFile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MJLogFileQuery when eager-loading is set.
	Edges                         MJLogFileEdges `json:"edges"`
	compressed_mj_log_mjlog_files *uuid.UUID
}

// MJLogFileEdges holds the relations/edges for other nodes in the graph.
type MJLogFileEdges struct {
	// CompressedMjlogFiles holds the value of the compressed_mjlog_files edge.
	CompressedMjlogFiles *CompressedMJLog `json:"compressed_mjlog_files,omitempty"`
	// Mjlogs holds the value of the mjlogs edge.
	Mjlogs *MJLog `json:"mjlogs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CompressedMjlogFilesOrErr returns the CompressedMjlogFiles value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MJLogFileEdges) CompressedMjlogFilesOrErr() (*CompressedMJLog, error) {
	if e.loadedTypes[0] {
		if e.CompressedMjlogFiles == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: compressedmjlog.Label}
		}
		return e.CompressedMjlogFiles, nil
	}
	return nil, &NotLoadedError{edge: "compressed_mjlog_files"}
}

// MjlogsOrErr returns the Mjlogs value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MJLogFileEdges) MjlogsOrErr() (*MJLog, error) {
	if e.loadedTypes[1] {
		if e.Mjlogs == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mjlog.Label}
		}
		return e.Mjlogs, nil
	}
	return nil, &NotLoadedError{edge: "mjlogs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MJLogFile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mjlogfile.FieldName:
			values[i] = new(sql.NullString)
		case mjlogfile.FieldID:
			values[i] = new(uuid.UUID)
		case mjlogfile.ForeignKeys[0]: // compressed_mj_log_mjlog_files
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type MJLogFile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MJLogFile fields.
func (mlf *MJLogFile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mjlogfile.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				mlf.ID = *value
			}
		case mjlogfile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mlf.Name = value.String
			}
		case mjlogfile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field compressed_mj_log_mjlog_files", values[i])
			} else if value.Valid {
				mlf.compressed_mj_log_mjlog_files = new(uuid.UUID)
				*mlf.compressed_mj_log_mjlog_files = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCompressedMjlogFiles queries the "compressed_mjlog_files" edge of the MJLogFile entity.
func (mlf *MJLogFile) QueryCompressedMjlogFiles() *CompressedMJLogQuery {
	return (&MJLogFileClient{config: mlf.config}).QueryCompressedMjlogFiles(mlf)
}

// QueryMjlogs queries the "mjlogs" edge of the MJLogFile entity.
func (mlf *MJLogFile) QueryMjlogs() *MJLogQuery {
	return (&MJLogFileClient{config: mlf.config}).QueryMjlogs(mlf)
}

// Update returns a builder for updating this MJLogFile.
// Note that you need to call MJLogFile.Unwrap() before calling this method if this MJLogFile
// was returned from a transaction, and the transaction was committed or rolled back.
func (mlf *MJLogFile) Update() *MJLogFileUpdateOne {
	return (&MJLogFileClient{config: mlf.config}).UpdateOne(mlf)
}

// Unwrap unwraps the MJLogFile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mlf *MJLogFile) Unwrap() *MJLogFile {
	_tx, ok := mlf.config.driver.(*txDriver)
	if !ok {
		panic("ent: MJLogFile is not a transactional entity")
	}
	mlf.config.driver = _tx.drv
	return mlf
}

// String implements the fmt.Stringer.
func (mlf *MJLogFile) String() string {
	var builder strings.Builder
	builder.WriteString("MJLogFile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mlf.ID))
	builder.WriteString("name=")
	builder.WriteString(mlf.Name)
	builder.WriteByte(')')
	return builder.String()
}

// MJLogFiles is a parsable slice of MJLogFile.
type MJLogFiles []*MJLogFile

func (mlf MJLogFiles) config(cfg config) {
	for _i := range mlf {
		mlf[_i].config = cfg
	}
}
