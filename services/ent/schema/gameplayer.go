package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GamePlayer holds the schema definition for the GamePlayer entity.
type GamePlayer struct {
	ent.Schema
}

// Fields of the GamePlayer.
func (GamePlayer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.Float("rate").Immutable(),
	}
}

// Edges of the GamePlayer.
func (GamePlayer) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("games", Game.Type).Ref("game_players"),
		edge.From("players", Player.Type).Ref("game_players"),
		//edge.From("dans", Dan.Type).Ref("game_players"),
		//edge.To("go_arounds", GoAround.Type),
	}
}
