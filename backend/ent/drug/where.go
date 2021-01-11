// Code generated by entc, DO NOT EDIT.

package drug

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team08/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DrugName applies equality check predicate on the "Drug_Name" field. It's identical to DrugNameEQ.
func DrugName(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugName), v))
	})
}

// DrugNameEQ applies the EQ predicate on the "Drug_Name" field.
func DrugNameEQ(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDrugName), v))
	})
}

// DrugNameNEQ applies the NEQ predicate on the "Drug_Name" field.
func DrugNameNEQ(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDrugName), v))
	})
}

// DrugNameIn applies the In predicate on the "Drug_Name" field.
func DrugNameIn(vs ...string) predicate.Drug {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Drug(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDrugName), v...))
	})
}

// DrugNameNotIn applies the NotIn predicate on the "Drug_Name" field.
func DrugNameNotIn(vs ...string) predicate.Drug {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Drug(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDrugName), v...))
	})
}

// DrugNameGT applies the GT predicate on the "Drug_Name" field.
func DrugNameGT(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDrugName), v))
	})
}

// DrugNameGTE applies the GTE predicate on the "Drug_Name" field.
func DrugNameGTE(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDrugName), v))
	})
}

// DrugNameLT applies the LT predicate on the "Drug_Name" field.
func DrugNameLT(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDrugName), v))
	})
}

// DrugNameLTE applies the LTE predicate on the "Drug_Name" field.
func DrugNameLTE(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDrugName), v))
	})
}

// DrugNameContains applies the Contains predicate on the "Drug_Name" field.
func DrugNameContains(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDrugName), v))
	})
}

// DrugNameHasPrefix applies the HasPrefix predicate on the "Drug_Name" field.
func DrugNameHasPrefix(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDrugName), v))
	})
}

// DrugNameHasSuffix applies the HasSuffix predicate on the "Drug_Name" field.
func DrugNameHasSuffix(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDrugName), v))
	})
}

// DrugNameEqualFold applies the EqualFold predicate on the "Drug_Name" field.
func DrugNameEqualFold(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDrugName), v))
	})
}

// DrugNameContainsFold applies the ContainsFold predicate on the "Drug_Name" field.
func DrugNameContainsFold(v string) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDrugName), v))
	})
}

// HasDrugPrescription applies the HasEdge predicate on the "drug_prescription" edge.
func HasDrugPrescription() predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DrugPrescriptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DrugPrescriptionTable, DrugPrescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDrugPrescriptionWith applies the HasEdge predicate on the "drug_prescription" edge with a given conditions (other predicates).
func HasDrugPrescriptionWith(preds ...predicate.Prescription) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DrugPrescriptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DrugPrescriptionTable, DrugPrescriptionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Drug) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Drug) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
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
func Not(p predicate.Drug) predicate.Drug {
	return predicate.Drug(func(s *sql.Selector) {
		p(s.Not())
	})
}
