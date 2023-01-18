package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// MJLog holds the schema definition for the MJLog entity.
type MJLog struct {
	ent.Schema
}

// Fields of the MJLog.
func (MJLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Float("version").Immutable(),
		field.String("seed").Immutable(),
		field.Time("started_at").Immutable(),
		field.Time("inserted_at").Default(time.Now).Immutable(),
	}
}

// Edges of the MJLog.
func (MJLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mjlog_files", MJLogFile.Type).Ref("mjlogs").Unique(),
		edge.From("games", Game.Type).Ref("mjlogs").Unique().Required(),
	}
}
