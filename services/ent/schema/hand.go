package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/google/uuid"
)

// Hand holds the schema definition for the Hand entity.
type Hand struct {
	ent.Schema
}

// Fields of the Hand.
func (Hand) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Uint("num").Immutable(),
		field.Uint("continue_point").Immutable().Validate(func(u uint) error {
			if u%100 == 0 {
				return nil
			}
			return fmt.Errorf("continue_point must be multiply by 100. actual: %d", u)
		}),
		field.Uint("reach_point").Immutable().Validate(func(u uint) error {
			if u%1000 == 0 {
				return nil
			}
			return fmt.Errorf("reach_point must be multiply by 1000. actual: %d", u)
		}),
	}
}

// Edges of the Hand.
func (Hand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rounds", Round.Type).Ref("hands").Unique(),
		edge.To("turns", Turn.Type),
	}
}
