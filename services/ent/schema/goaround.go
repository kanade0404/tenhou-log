package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GoAround holds the schema definition for the GoAround entity.
type GoAround struct {
	ent.Schema
}

// Fields of the GoAround.
func (GoAround) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
	}
}

// Edges of the GoAround.
func (GoAround) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("game_players", GamePlayer.Type).Ref("go_arounds"),
	}
}
