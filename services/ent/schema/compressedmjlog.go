package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CompressedMJLog holds the schema definition for the CompressedMJLog entity.
type CompressedMJLog struct {
	ent.Schema
}

// Fields of the CompressedMJLog.
func (CompressedMJLog) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("name").Unique().Immutable(),
		field.Uint("size"),
	}
}

// Edges of the CompressedMJLog.
func (CompressedMJLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mjlog_files", MJLogFile.Type).Unique(),
	}
}
