package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MJLogFileCompressed holds the schema definition for the MJLogFileCompressed entity.
type MJLogFileCompressed struct {
	ent.Schema
}

// Fields of the MJLogFileCompressed.
func (MJLogFileCompressed) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
	}
}

// Edges of the MJLogFileCompressed.
func (MJLogFileCompressed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("compressed_mjlog_files", CompressedMJLog.Type).Unique(),
	}
}
