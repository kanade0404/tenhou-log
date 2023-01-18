package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// CompressedMJLog holds the schema definition for the CompressedMJLog entity.
type CompressedMJLog struct {
	ent.Schema
}

// Fields of the CompressedMJLog.
func (CompressedMJLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Immutable(),
		field.Uint("size"),
		field.Time("inserted_at").Immutable().Default(time.Now),
	}
}

// Edges of the CompressedMJLog.
func (CompressedMJLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mjlog_files", MJLogFile.Type).Unique(),
	}
}
