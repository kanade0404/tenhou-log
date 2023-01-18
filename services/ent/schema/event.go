package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("turn", Turn.Type).Ref("event").Unique().Required(),
		edge.To("win", Win.Type),
		edge.To("call", Call.Type),
		edge.To("draw", Drawn.Type),
		edge.To("reach", Reach.Type),
	}
}
