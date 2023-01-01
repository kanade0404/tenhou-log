package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Dan holds the schema definition for the Dan entity.
type Dan struct {
	ent.Schema
}

// Fields of the Dan.
func (Dan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("name").Unique().Immutable(),
	}
}

// Edges of the Dan.
func (Dan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game_players", GamePlayer.Type),
	}
}
