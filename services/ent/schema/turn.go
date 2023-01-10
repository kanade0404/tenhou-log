package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Turn holds the schema definition for the Turn entity.
type Turn struct {
	ent.Schema
}

// Fields of the Turn.
func (Turn) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Uint("num").Immutable(),
	}
}

// Edges of the Turn.
func (Turn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hands", Hand.Type).Ref("turns"),
		edge.To("game_player_points", GamePlayerPoint.Type),
		edge.To("event", Event.Type),
	}
}
