// Code generated by ent, DO NOT EDIT.

package hand

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Num applies equality check predicate on the "num" field. It's identical to NumEQ.
func Num(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNum), v))
	})
}

// ContinuePoint applies equality check predicate on the "continue_point" field. It's identical to ContinuePointEQ.
func ContinuePoint(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContinuePoint), v))
	})
}

// ReachPoint applies equality check predicate on the "reach_point" field. It's identical to ReachPointEQ.
func ReachPoint(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReachPoint), v))
	})
}

// NumEQ applies the EQ predicate on the "num" field.
func NumEQ(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNum), v))
	})
}

// NumNEQ applies the NEQ predicate on the "num" field.
func NumNEQ(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNum), v))
	})
}

// NumIn applies the In predicate on the "num" field.
func NumIn(vs ...uint) predicate.Hand {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNum), v...))
	})
}

// NumNotIn applies the NotIn predicate on the "num" field.
func NumNotIn(vs ...uint) predicate.Hand {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNum), v...))
	})
}

// NumGT applies the GT predicate on the "num" field.
func NumGT(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNum), v))
	})
}

// NumGTE applies the GTE predicate on the "num" field.
func NumGTE(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNum), v))
	})
}

// NumLT applies the LT predicate on the "num" field.
func NumLT(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNum), v))
	})
}

// NumLTE applies the LTE predicate on the "num" field.
func NumLTE(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNum), v))
	})
}

// ContinuePointEQ applies the EQ predicate on the "continue_point" field.
func ContinuePointEQ(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContinuePoint), v))
	})
}

// ContinuePointNEQ applies the NEQ predicate on the "continue_point" field.
func ContinuePointNEQ(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContinuePoint), v))
	})
}

// ContinuePointIn applies the In predicate on the "continue_point" field.
func ContinuePointIn(vs ...uint) predicate.Hand {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContinuePoint), v...))
	})
}

// ContinuePointNotIn applies the NotIn predicate on the "continue_point" field.
func ContinuePointNotIn(vs ...uint) predicate.Hand {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContinuePoint), v...))
	})
}

// ContinuePointGT applies the GT predicate on the "continue_point" field.
func ContinuePointGT(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContinuePoint), v))
	})
}

// ContinuePointGTE applies the GTE predicate on the "continue_point" field.
func ContinuePointGTE(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContinuePoint), v))
	})
}

// ContinuePointLT applies the LT predicate on the "continue_point" field.
func ContinuePointLT(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContinuePoint), v))
	})
}

// ContinuePointLTE applies the LTE predicate on the "continue_point" field.
func ContinuePointLTE(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContinuePoint), v))
	})
}

// ReachPointEQ applies the EQ predicate on the "reach_point" field.
func ReachPointEQ(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReachPoint), v))
	})
}

// ReachPointNEQ applies the NEQ predicate on the "reach_point" field.
func ReachPointNEQ(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReachPoint), v))
	})
}

// ReachPointIn applies the In predicate on the "reach_point" field.
func ReachPointIn(vs ...uint) predicate.Hand {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReachPoint), v...))
	})
}

// ReachPointNotIn applies the NotIn predicate on the "reach_point" field.
func ReachPointNotIn(vs ...uint) predicate.Hand {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReachPoint), v...))
	})
}

// ReachPointGT applies the GT predicate on the "reach_point" field.
func ReachPointGT(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReachPoint), v))
	})
}

// ReachPointGTE applies the GTE predicate on the "reach_point" field.
func ReachPointGTE(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReachPoint), v))
	})
}

// ReachPointLT applies the LT predicate on the "reach_point" field.
func ReachPointLT(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReachPoint), v))
	})
}

// ReachPointLTE applies the LTE predicate on the "reach_point" field.
func ReachPointLTE(v uint) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReachPoint), v))
	})
}

// HasRounds applies the HasEdge predicate on the "rounds" edge.
func HasRounds() predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoundsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoundsTable, RoundsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoundsWith applies the HasEdge predicate on the "rounds" edge with a given conditions (other predicates).
func HasRoundsWith(preds ...predicate.Round) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoundsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoundsTable, RoundsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Hand) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Hand) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Hand) predicate.Hand {
	return predicate.Hand(func(s *sql.Selector) {
		p(s.Not())
	})
}
