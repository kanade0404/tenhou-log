// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/game"
	"github.com/kanade0404/tenhou-log/services/ent/mjlog"
	"github.com/kanade0404/tenhou-log/services/ent/mjlogfile"
)

// MJLog is the model entity for the MJLog schema.
type MJLog struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Version holds the value of the "version" field.
	Version float64 `json:"version,omitempty"`
	// Seed holds the value of the "seed" field.
	Seed string `json:"seed,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt time.Time `json:"started_at,omitempty"`
	// InsertedAt holds the value of the "inserted_at" field.
	InsertedAt time.Time `json:"inserted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MJLogQuery when eager-loading is set.
	Edges              MJLogEdges `json:"edges"`
	game_mjlogs        *uuid.UUID
	mj_log_file_mjlogs *uuid.UUID
}

// MJLogEdges holds the relations/edges for other nodes in the graph.
type MJLogEdges struct {
	// MjlogFiles holds the value of the mjlog_files edge.
	MjlogFiles *MJLogFile `json:"mjlog_files,omitempty"`
	// Games holds the value of the games edge.
	Games *Game `json:"games,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MjlogFilesOrErr returns the MjlogFiles value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MJLogEdges) MjlogFilesOrErr() (*MJLogFile, error) {
	if e.loadedTypes[0] {
		if e.MjlogFiles == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mjlogfile.Label}
		}
		return e.MjlogFiles, nil
	}
	return nil, &NotLoadedError{edge: "mjlog_files"}
}

// GamesOrErr returns the Games value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MJLogEdges) GamesOrErr() (*Game, error) {
	if e.loadedTypes[1] {
		if e.Games == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: game.Label}
		}
		return e.Games, nil
	}
	return nil, &NotLoadedError{edge: "games"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MJLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mjlog.FieldVersion:
			values[i] = new(sql.NullFloat64)
		case mjlog.FieldSeed:
			values[i] = new(sql.NullString)
		case mjlog.FieldStartedAt, mjlog.FieldInsertedAt:
			values[i] = new(sql.NullTime)
		case mjlog.FieldID:
			values[i] = new(uuid.UUID)
		case mjlog.ForeignKeys[0]: // game_mjlogs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case mjlog.ForeignKeys[1]: // mj_log_file_mjlogs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type MJLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MJLog fields.
func (ml *MJLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mjlog.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ml.ID = *value
			}
		case mjlog.FieldVersion:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ml.Version = value.Float64
			}
		case mjlog.FieldSeed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seed", values[i])
			} else if value.Valid {
				ml.Seed = value.String
			}
		case mjlog.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				ml.StartedAt = value.Time
			}
		case mjlog.FieldInsertedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field inserted_at", values[i])
			} else if value.Valid {
				ml.InsertedAt = value.Time
			}
		case mjlog.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field game_mjlogs", values[i])
			} else if value.Valid {
				ml.game_mjlogs = new(uuid.UUID)
				*ml.game_mjlogs = *value.S.(*uuid.UUID)
			}
		case mjlog.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field mj_log_file_mjlogs", values[i])
			} else if value.Valid {
				ml.mj_log_file_mjlogs = new(uuid.UUID)
				*ml.mj_log_file_mjlogs = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryMjlogFiles queries the "mjlog_files" edge of the MJLog entity.
func (ml *MJLog) QueryMjlogFiles() *MJLogFileQuery {
	return (&MJLogClient{config: ml.config}).QueryMjlogFiles(ml)
}

// QueryGames queries the "games" edge of the MJLog entity.
func (ml *MJLog) QueryGames() *GameQuery {
	return (&MJLogClient{config: ml.config}).QueryGames(ml)
}

// Update returns a builder for updating this MJLog.
// Note that you need to call MJLog.Unwrap() before calling this method if this MJLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (ml *MJLog) Update() *MJLogUpdateOne {
	return (&MJLogClient{config: ml.config}).UpdateOne(ml)
}

// Unwrap unwraps the MJLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ml *MJLog) Unwrap() *MJLog {
	_tx, ok := ml.config.driver.(*txDriver)
	if !ok {
		panic("ent: MJLog is not a transactional entity")
	}
	ml.config.driver = _tx.drv
	return ml
}

// String implements the fmt.Stringer.
func (ml *MJLog) String() string {
	var builder strings.Builder
	builder.WriteString("MJLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ml.ID))
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", ml.Version))
	builder.WriteString(", ")
	builder.WriteString("seed=")
	builder.WriteString(ml.Seed)
	builder.WriteString(", ")
	builder.WriteString("started_at=")
	builder.WriteString(ml.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("inserted_at=")
	builder.WriteString(ml.InsertedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MJLogs is a parsable slice of MJLog.
type MJLogs []*MJLog

func (ml MJLogs) config(cfg config) {
	for _i := range ml {
		ml[_i].config = cfg
	}
}
