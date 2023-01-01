package schema

import (
	"entgo.io/ent"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		//field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		//field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("game_players", GamePlayer.Type),
	}
}
