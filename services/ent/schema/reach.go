package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Reach holds the schema definition for the Reach entity.
type Reach struct {
	ent.Schema
}

// Fields of the Reach.
func (Reach) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Reach.
func (Reach) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("reach").Unique().Required(),
		edge.From("discard", Discard.Type).Ref("reach").Unique().Required(),
	}
}
