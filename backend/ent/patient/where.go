// Code generated by entc, DO NOT EDIT.

package patient

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team08/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PatientName applies equality check predicate on the "Patient_name" field. It's identical to PatientNameEQ.
func PatientName(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientName), v))
	})
}

// PatientAge applies equality check predicate on the "Patient_age" field. It's identical to PatientAgeEQ.
func PatientAge(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientAge), v))
	})
}

// PatientWeight applies equality check predicate on the "Patient_weight" field. It's identical to PatientWeightEQ.
func PatientWeight(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientWeight), v))
	})
}

// PatientHeight applies equality check predicate on the "Patient_height" field. It's identical to PatientHeightEQ.
func PatientHeight(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientHeight), v))
	})
}

// PatientNameEQ applies the EQ predicate on the "Patient_name" field.
func PatientNameEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientName), v))
	})
}

// PatientNameNEQ applies the NEQ predicate on the "Patient_name" field.
func PatientNameNEQ(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientName), v))
	})
}

// PatientNameIn applies the In predicate on the "Patient_name" field.
func PatientNameIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientName), v...))
	})
}

// PatientNameNotIn applies the NotIn predicate on the "Patient_name" field.
func PatientNameNotIn(vs ...string) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientName), v...))
	})
}

// PatientNameGT applies the GT predicate on the "Patient_name" field.
func PatientNameGT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientName), v))
	})
}

// PatientNameGTE applies the GTE predicate on the "Patient_name" field.
func PatientNameGTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientName), v))
	})
}

// PatientNameLT applies the LT predicate on the "Patient_name" field.
func PatientNameLT(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientName), v))
	})
}

// PatientNameLTE applies the LTE predicate on the "Patient_name" field.
func PatientNameLTE(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientName), v))
	})
}

// PatientNameContains applies the Contains predicate on the "Patient_name" field.
func PatientNameContains(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPatientName), v))
	})
}

// PatientNameHasPrefix applies the HasPrefix predicate on the "Patient_name" field.
func PatientNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPatientName), v))
	})
}

// PatientNameHasSuffix applies the HasSuffix predicate on the "Patient_name" field.
func PatientNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPatientName), v))
	})
}

// PatientNameEqualFold applies the EqualFold predicate on the "Patient_name" field.
func PatientNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPatientName), v))
	})
}

// PatientNameContainsFold applies the ContainsFold predicate on the "Patient_name" field.
func PatientNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPatientName), v))
	})
}

// PatientAgeEQ applies the EQ predicate on the "Patient_age" field.
func PatientAgeEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientAge), v))
	})
}

// PatientAgeNEQ applies the NEQ predicate on the "Patient_age" field.
func PatientAgeNEQ(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientAge), v))
	})
}

// PatientAgeIn applies the In predicate on the "Patient_age" field.
func PatientAgeIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientAge), v...))
	})
}

// PatientAgeNotIn applies the NotIn predicate on the "Patient_age" field.
func PatientAgeNotIn(vs ...int) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientAge), v...))
	})
}

// PatientAgeGT applies the GT predicate on the "Patient_age" field.
func PatientAgeGT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientAge), v))
	})
}

// PatientAgeGTE applies the GTE predicate on the "Patient_age" field.
func PatientAgeGTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientAge), v))
	})
}

// PatientAgeLT applies the LT predicate on the "Patient_age" field.
func PatientAgeLT(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientAge), v))
	})
}

// PatientAgeLTE applies the LTE predicate on the "Patient_age" field.
func PatientAgeLTE(v int) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientAge), v))
	})
}

// PatientWeightEQ applies the EQ predicate on the "Patient_weight" field.
func PatientWeightEQ(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientWeight), v))
	})
}

// PatientWeightNEQ applies the NEQ predicate on the "Patient_weight" field.
func PatientWeightNEQ(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientWeight), v))
	})
}

// PatientWeightIn applies the In predicate on the "Patient_weight" field.
func PatientWeightIn(vs ...float64) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientWeight), v...))
	})
}

// PatientWeightNotIn applies the NotIn predicate on the "Patient_weight" field.
func PatientWeightNotIn(vs ...float64) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientWeight), v...))
	})
}

// PatientWeightGT applies the GT predicate on the "Patient_weight" field.
func PatientWeightGT(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientWeight), v))
	})
}

// PatientWeightGTE applies the GTE predicate on the "Patient_weight" field.
func PatientWeightGTE(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientWeight), v))
	})
}

// PatientWeightLT applies the LT predicate on the "Patient_weight" field.
func PatientWeightLT(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientWeight), v))
	})
}

// PatientWeightLTE applies the LTE predicate on the "Patient_weight" field.
func PatientWeightLTE(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientWeight), v))
	})
}

// PatientHeightEQ applies the EQ predicate on the "Patient_height" field.
func PatientHeightEQ(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientHeight), v))
	})
}

// PatientHeightNEQ applies the NEQ predicate on the "Patient_height" field.
func PatientHeightNEQ(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientHeight), v))
	})
}

// PatientHeightIn applies the In predicate on the "Patient_height" field.
func PatientHeightIn(vs ...float64) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientHeight), v...))
	})
}

// PatientHeightNotIn applies the NotIn predicate on the "Patient_height" field.
func PatientHeightNotIn(vs ...float64) predicate.Patient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientHeight), v...))
	})
}

// PatientHeightGT applies the GT predicate on the "Patient_height" field.
func PatientHeightGT(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientHeight), v))
	})
}

// PatientHeightGTE applies the GTE predicate on the "Patient_height" field.
func PatientHeightGTE(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientHeight), v))
	})
}

// PatientHeightLT applies the LT predicate on the "Patient_height" field.
func PatientHeightLT(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientHeight), v))
	})
}

// PatientHeightLTE applies the LTE predicate on the "Patient_height" field.
func PatientHeightLTE(v float64) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientHeight), v))
	})
}

// HasGender applies the HasEdge predicate on the "gender" edge.
func HasGender() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenderWith applies the HasEdge predicate on the "gender" edge with a given conditions (other predicates).
func HasGenderWith(preds ...predicate.Gender) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenderTable, GenderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrefix applies the HasEdge predicate on the "prefix" edge.
func HasPrefix() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrefixTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrefixTable, PrefixColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrefixWith applies the HasEdge predicate on the "prefix" edge with a given conditions (other predicates).
func HasPrefixWith(preds ...predicate.Prefix) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrefixInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrefixTable, PrefixColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBloodtype applies the HasEdge predicate on the "bloodtype" edge.
func HasBloodtype() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BloodtypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BloodtypeTable, BloodtypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBloodtypeWith applies the HasEdge predicate on the "bloodtype" edge with a given conditions (other predicates).
func HasBloodtypeWith(preds ...predicate.Bloodtype) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BloodtypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BloodtypeTable, BloodtypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFrompatient applies the HasEdge predicate on the "frompatient" edge.
func HasFrompatient() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FrompatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FrompatientTable, FrompatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFrompatientWith applies the HasEdge predicate on the "frompatient" edge with a given conditions (other predicates).
func HasFrompatientWith(preds ...predicate.Rent) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FrompatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, FrompatientTable, FrompatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientCoveredPerson applies the HasEdge predicate on the "Patient_CoveredPerson" edge.
func HasPatientCoveredPerson() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientCoveredPersonTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientCoveredPersonTable, PatientCoveredPersonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientCoveredPersonWith applies the HasEdge predicate on the "Patient_CoveredPerson" edge with a given conditions (other predicates).
func HasPatientCoveredPersonWith(preds ...predicate.CoveredPerson) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientCoveredPersonInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientCoveredPersonTable, PatientCoveredPersonColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientDiagnose applies the HasEdge predicate on the "patient_diagnose" edge.
func HasPatientDiagnose() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientDiagnoseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientDiagnoseTable, PatientDiagnoseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientDiagnoseWith applies the HasEdge predicate on the "patient_diagnose" edge with a given conditions (other predicates).
func HasPatientDiagnoseWith(preds ...predicate.Diagnose) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientDiagnoseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientDiagnoseTable, PatientDiagnoseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientPrescription applies the HasEdge predicate on the "patient_prescription" edge.
func HasPatientPrescription() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientPrescriptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientPrescriptionTable, PatientPrescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientPrescriptionWith applies the HasEdge predicate on the "patient_prescription" edge with a given conditions (other predicates).
func HasPatientPrescriptionWith(preds ...predicate.Prescription) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientPrescriptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientPrescriptionTable, PatientPrescriptionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
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
func Not(p predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		p(s.Not())
	})
}
