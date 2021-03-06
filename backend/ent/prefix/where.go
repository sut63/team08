// Code generated by entc, DO NOT EDIT.

package prefix

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team08/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Pname applies equality check predicate on the "Pname" field. It's identical to PnameEQ.
func Pname(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPname), v))
	})
}

// PnameEQ applies the EQ predicate on the "Pname" field.
func PnameEQ(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPname), v))
	})
}

// PnameNEQ applies the NEQ predicate on the "Pname" field.
func PnameNEQ(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPname), v))
	})
}

// PnameIn applies the In predicate on the "Pname" field.
func PnameIn(vs ...string) predicate.Prefix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPname), v...))
	})
}

// PnameNotIn applies the NotIn predicate on the "Pname" field.
func PnameNotIn(vs ...string) predicate.Prefix {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prefix(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPname), v...))
	})
}

// PnameGT applies the GT predicate on the "Pname" field.
func PnameGT(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPname), v))
	})
}

// PnameGTE applies the GTE predicate on the "Pname" field.
func PnameGTE(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPname), v))
	})
}

// PnameLT applies the LT predicate on the "Pname" field.
func PnameLT(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPname), v))
	})
}

// PnameLTE applies the LTE predicate on the "Pname" field.
func PnameLTE(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPname), v))
	})
}

// PnameContains applies the Contains predicate on the "Pname" field.
func PnameContains(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPname), v))
	})
}

// PnameHasPrefix applies the HasPrefix predicate on the "Pname" field.
func PnameHasPrefix(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPname), v))
	})
}

// PnameHasSuffix applies the HasSuffix predicate on the "Pname" field.
func PnameHasSuffix(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPname), v))
	})
}

// PnameEqualFold applies the EqualFold predicate on the "Pname" field.
func PnameEqualFold(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPname), v))
	})
}

// PnameContainsFold applies the ContainsFold predicate on the "Pname" field.
func PnameContainsFold(v string) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPname), v))
	})
}

// HasFromprefix applies the HasEdge predicate on the "fromprefix" edge.
func HasFromprefix() predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromprefixTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FromprefixTable, FromprefixColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromprefixWith applies the HasEdge predicate on the "fromprefix" edge with a given conditions (other predicates).
func HasFromprefixWith(preds ...predicate.Patient) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromprefixInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FromprefixTable, FromprefixColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
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
func Not(p predicate.Prefix) predicate.Prefix {
	return predicate.Prefix(func(s *sql.Selector) {
		p(s.Not())
	})
}
