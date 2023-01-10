package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Drawn holds the schema definition for the Drawn entity.
type Drawn struct {
	ent.Schema
}

// Fields of the Drawn.
func (Drawn) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Drawn.
func (Drawn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("draw").Unique().Required(),
		edge.From("discard", Discard.Type).Ref("draw").Unique().Required(),
	}
}
