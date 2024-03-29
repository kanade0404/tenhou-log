// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/compressedmjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
)

// CompressedMJLog is the model entity for the CompressedMJLog schema.
type CompressedMJLog struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Size holds the value of the "size" field.
	Size uint `json:"size,omitempty"`
	// InsertedAt holds the value of the "inserted_at" field.
	InsertedAt time.Time `json:"inserted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompressedMJLogQuery when eager-loading is set.
	Edges CompressedMJLogEdges `json:"edges"`
}

// CompressedMJLogEdges holds the relations/edges for other nodes in the graph.
type CompressedMJLogEdges struct {
	// MjlogFiles holds the value of the mjlog_files edge.
	MjlogFiles *MJLogFile `json:"mjlog_files,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MjlogFilesOrErr returns the MjlogFiles value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompressedMJLogEdges) MjlogFilesOrErr() (*MJLogFile, error) {
	if e.loadedTypes[0] {
		if e.MjlogFiles == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mjlogfile.Label}
		}
		return e.MjlogFiles, nil
	}
	return nil, &NotLoadedError{edge: "mjlog_files"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CompressedMJLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case compressedmjlog.FieldSize:
			values[i] = new(sql.NullInt64)
		case compressedmjlog.FieldName:
			values[i] = new(sql.NullString)
		case compressedmjlog.FieldInsertedAt:
			values[i] = new(sql.NullTime)
		case compressedmjlog.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CompressedMJLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CompressedMJLog fields.
func (cml *CompressedMJLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case compressedmjlog.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cml.ID = *value
			}
		case compressedmjlog.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cml.Name = value.String
			}
		case compressedmjlog.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				cml.Size = uint(value.Int64)
			}
		case compressedmjlog.FieldInsertedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field inserted_at", values[i])
			} else if value.Valid {
				cml.InsertedAt = value.Time
			}
		}
	}
	return nil
}

// QueryMjlogFiles queries the "mjlog_files" edge of the CompressedMJLog entity.
func (cml *CompressedMJLog) QueryMjlogFiles() *MJLogFileQuery {
	return (&CompressedMJLogClient{config: cml.config}).QueryMjlogFiles(cml)
}

// Update returns a builder for updating this CompressedMJLog.
// Note that you need to call CompressedMJLog.Unwrap() before calling this method if this CompressedMJLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (cml *CompressedMJLog) Update() *CompressedMJLogUpdateOne {
	return (&CompressedMJLogClient{config: cml.config}).UpdateOne(cml)
}

// Unwrap unwraps the CompressedMJLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cml *CompressedMJLog) Unwrap() *CompressedMJLog {
	_tx, ok := cml.config.driver.(*txDriver)
	if !ok {
		panic("ent: CompressedMJLog is not a transactional entity")
	}
	cml.config.driver = _tx.drv
	return cml
}

// String implements the fmt.Stringer.
func (cml *CompressedMJLog) String() string {
	var builder strings.Builder
	builder.WriteString("CompressedMJLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cml.ID))
	builder.WriteString("name=")
	builder.WriteString(cml.Name)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", cml.Size))
	builder.WriteString(", ")
	builder.WriteString("inserted_at=")
	builder.WriteString(cml.InsertedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CompressedMJLogs is a parsable slice of CompressedMJLog.
type CompressedMJLogs []*CompressedMJLog

func (cml CompressedMJLogs) config(cfg config) {
	for _i := range cml {
		cml[_i].config = cfg
	}
}
