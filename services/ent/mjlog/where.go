// Code generated by ent, DO NOT EDIT.

package mjlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v float64) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// Seed applies equality check predicate on the "seed" field. It's identical to SeedEQ.
func Seed(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeed), v))
	})
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// InsertedAt applies equality check predicate on the "inserted_at" field. It's identical to InsertedAtEQ.
func InsertedAt(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsertedAt), v))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v float64) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v float64) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...float64) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...float64) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v float64) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v float64) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v float64) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v float64) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// SeedEQ applies the EQ predicate on the "seed" field.
func SeedEQ(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeed), v))
	})
}

// SeedNEQ applies the NEQ predicate on the "seed" field.
func SeedNEQ(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeed), v))
	})
}

// SeedIn applies the In predicate on the "seed" field.
func SeedIn(vs ...string) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSeed), v...))
	})
}

// SeedNotIn applies the NotIn predicate on the "seed" field.
func SeedNotIn(vs ...string) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSeed), v...))
	})
}

// SeedGT applies the GT predicate on the "seed" field.
func SeedGT(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeed), v))
	})
}

// SeedGTE applies the GTE predicate on the "seed" field.
func SeedGTE(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeed), v))
	})
}

// SeedLT applies the LT predicate on the "seed" field.
func SeedLT(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeed), v))
	})
}

// SeedLTE applies the LTE predicate on the "seed" field.
func SeedLTE(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeed), v))
	})
}

// SeedContains applies the Contains predicate on the "seed" field.
func SeedContains(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeed), v))
	})
}

// SeedHasPrefix applies the HasPrefix predicate on the "seed" field.
func SeedHasPrefix(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeed), v))
	})
}

// SeedHasSuffix applies the HasSuffix predicate on the "seed" field.
func SeedHasSuffix(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeed), v))
	})
}

// SeedEqualFold applies the EqualFold predicate on the "seed" field.
func SeedEqualFold(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeed), v))
	})
}

// SeedContainsFold applies the ContainsFold predicate on the "seed" field.
func SeedContainsFold(v string) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeed), v))
	})
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// InsertedAtEQ applies the EQ predicate on the "inserted_at" field.
func InsertedAtEQ(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsertedAt), v))
	})
}

// InsertedAtNEQ applies the NEQ predicate on the "inserted_at" field.
func InsertedAtNEQ(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInsertedAt), v))
	})
}

// InsertedAtIn applies the In predicate on the "inserted_at" field.
func InsertedAtIn(vs ...time.Time) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInsertedAt), v...))
	})
}

// InsertedAtNotIn applies the NotIn predicate on the "inserted_at" field.
func InsertedAtNotIn(vs ...time.Time) predicate.MJLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInsertedAt), v...))
	})
}

// InsertedAtGT applies the GT predicate on the "inserted_at" field.
func InsertedAtGT(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInsertedAt), v))
	})
}

// InsertedAtGTE applies the GTE predicate on the "inserted_at" field.
func InsertedAtGTE(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInsertedAt), v))
	})
}

// InsertedAtLT applies the LT predicate on the "inserted_at" field.
func InsertedAtLT(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInsertedAt), v))
	})
}

// InsertedAtLTE applies the LTE predicate on the "inserted_at" field.
func InsertedAtLTE(v time.Time) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInsertedAt), v))
	})
}

// HasMjlogFiles applies the HasEdge predicate on the "mjlog_files" edge.
func HasMjlogFiles() predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MjlogFilesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MjlogFilesTable, MjlogFilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMjlogFilesWith applies the HasEdge predicate on the "mjlog_files" edge with a given conditions (other predicates).
func HasMjlogFilesWith(preds ...predicate.MJLogFile) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MjlogFilesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MjlogFilesTable, MjlogFilesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGames applies the HasEdge predicate on the "games" edge.
func HasGames() predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GamesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, GamesTable, GamesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGamesWith applies the HasEdge predicate on the "games" edge with a given conditions (other predicates).
func HasGamesWith(preds ...predicate.Game) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GamesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, GamesTable, GamesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MJLog) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MJLog) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
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
func Not(p predicate.MJLog) predicate.MJLog {
	return predicate.MJLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
