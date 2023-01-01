package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("name").Unique().Immutable(),
		field.Time("started_at").Immutable(),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mjlogs", MJLog.Type).Unique(),
		edge.To("game_players", GamePlayer.Type),
		edge.From("rooms", Room.Type).Ref("games").Unique(),
		//edge.To("rounds", Round.Type),
	}
}
