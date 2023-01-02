package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/google/uuid"
)

// GamePlayerPoint holds the schema definition for the GamePlayerPoint entity.
type GamePlayerPoint struct {
	ent.Schema
}

// Fields of the GamePlayerPoint.
func (GamePlayerPoint) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.Uint("point").Immutable().Validate(func(u uint) error {
			if u%100 == 0 {
				return nil
			}
			return fmt.Errorf("point must be multiply by 100. actual: %d", u)
		}),
	}
}

// Edges of the GamePlayerPoint.
func (GamePlayerPoint) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("turns", Turn.Type).Ref("game_players"),
	}
}
