package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MJLogFile holds the schema definition for the MJLogFile entity.
type MJLogFile struct {
	ent.Schema
}

// Fields of the MJLogFile.
func (MJLogFile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("name").Unique().Immutable(),
	}
}

// Edges of the MJLogFile.
func (MJLogFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("compressed_mjlog_files", CompressedMJLog.Type).Ref("mjlog_files").Unique().Required(),
		//edge.To("mjlogs", MJLog.Type).Unique(),
	}
}
