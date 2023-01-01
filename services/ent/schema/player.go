package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"regexp"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("name").Unique().Immutable(),
		field.String("sex").Match(regexp.MustCompile("^[F|M]$")).Immutable(),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game_players", GamePlayer.Type),
	}
}
