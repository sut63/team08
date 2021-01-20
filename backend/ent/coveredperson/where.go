// Code generated by entc, DO NOT EDIT.

package coveredperson

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team08/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
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
func IDGT(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CoveredPersonNumber applies equality check predicate on the "CoveredPerson_Number" field. It's identical to CoveredPersonNumberEQ.
func CoveredPersonNumber(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNote applies equality check predicate on the "CoveredPerson_Note" field. It's identical to CoveredPersonNoteEQ.
func CoveredPersonNote(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoveredPersonNote), v))
	})
}

// FundTitle applies equality check predicate on the "Fund_Title" field. It's identical to FundTitleEQ.
func FundTitle(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFundTitle), v))
	})
}

// CoveredPersonNumberEQ applies the EQ predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberEQ(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberNEQ applies the NEQ predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberNEQ(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberIn applies the In predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberIn(vs ...string) predicate.CoveredPerson {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoveredPerson(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoveredPersonNumber), v...))
	})
}

// CoveredPersonNumberNotIn applies the NotIn predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberNotIn(vs ...string) predicate.CoveredPerson {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoveredPerson(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoveredPersonNumber), v...))
	})
}

// CoveredPersonNumberGT applies the GT predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberGT(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberGTE applies the GTE predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberGTE(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberLT applies the LT predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberLT(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberLTE applies the LTE predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberLTE(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberContains applies the Contains predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberContains(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberHasPrefix applies the HasPrefix predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberHasPrefix(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberHasSuffix applies the HasSuffix predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberHasSuffix(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberEqualFold applies the EqualFold predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberEqualFold(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNumberContainsFold applies the ContainsFold predicate on the "CoveredPerson_Number" field.
func CoveredPersonNumberContainsFold(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCoveredPersonNumber), v))
	})
}

// CoveredPersonNoteEQ applies the EQ predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteEQ(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteNEQ applies the NEQ predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteNEQ(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteIn applies the In predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteIn(vs ...string) predicate.CoveredPerson {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoveredPerson(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoveredPersonNote), v...))
	})
}

// CoveredPersonNoteNotIn applies the NotIn predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteNotIn(vs ...string) predicate.CoveredPerson {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoveredPerson(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoveredPersonNote), v...))
	})
}

// CoveredPersonNoteGT applies the GT predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteGT(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteGTE applies the GTE predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteGTE(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteLT applies the LT predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteLT(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteLTE applies the LTE predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteLTE(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteContains applies the Contains predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteContains(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteHasPrefix applies the HasPrefix predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteHasPrefix(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteHasSuffix applies the HasSuffix predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteHasSuffix(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteEqualFold applies the EqualFold predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteEqualFold(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCoveredPersonNote), v))
	})
}

// CoveredPersonNoteContainsFold applies the ContainsFold predicate on the "CoveredPerson_Note" field.
func CoveredPersonNoteContainsFold(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCoveredPersonNote), v))
	})
}

// FundTitleEQ applies the EQ predicate on the "Fund_Title" field.
func FundTitleEQ(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFundTitle), v))
	})
}

// FundTitleNEQ applies the NEQ predicate on the "Fund_Title" field.
func FundTitleNEQ(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFundTitle), v))
	})
}

// FundTitleIn applies the In predicate on the "Fund_Title" field.
func FundTitleIn(vs ...string) predicate.CoveredPerson {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoveredPerson(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFundTitle), v...))
	})
}

// FundTitleNotIn applies the NotIn predicate on the "Fund_Title" field.
func FundTitleNotIn(vs ...string) predicate.CoveredPerson {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoveredPerson(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFundTitle), v...))
	})
}

// FundTitleGT applies the GT predicate on the "Fund_Title" field.
func FundTitleGT(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFundTitle), v))
	})
}

// FundTitleGTE applies the GTE predicate on the "Fund_Title" field.
func FundTitleGTE(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFundTitle), v))
	})
}

// FundTitleLT applies the LT predicate on the "Fund_Title" field.
func FundTitleLT(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFundTitle), v))
	})
}

// FundTitleLTE applies the LTE predicate on the "Fund_Title" field.
func FundTitleLTE(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFundTitle), v))
	})
}

// FundTitleContains applies the Contains predicate on the "Fund_Title" field.
func FundTitleContains(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFundTitle), v))
	})
}

// FundTitleHasPrefix applies the HasPrefix predicate on the "Fund_Title" field.
func FundTitleHasPrefix(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFundTitle), v))
	})
}

// FundTitleHasSuffix applies the HasSuffix predicate on the "Fund_Title" field.
func FundTitleHasSuffix(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFundTitle), v))
	})
}

// FundTitleEqualFold applies the EqualFold predicate on the "Fund_Title" field.
func FundTitleEqualFold(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFundTitle), v))
	})
}

// FundTitleContainsFold applies the ContainsFold predicate on the "Fund_Title" field.
func FundTitleContainsFold(v string) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFundTitle), v))
	})
}

// HasPatient applies the HasEdge predicate on the "Patient" edge.
func HasPatient() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "Patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSchemeType applies the HasEdge predicate on the "SchemeType" edge.
func HasSchemeType() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SchemeTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SchemeTypeTable, SchemeTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSchemeTypeWith applies the HasEdge predicate on the "SchemeType" edge with a given conditions (other predicates).
func HasSchemeTypeWith(preds ...predicate.SchemeType) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SchemeTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SchemeTypeTable, SchemeTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFund applies the HasEdge predicate on the "Fund" edge.
func HasFund() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FundTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FundTable, FundColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFundWith applies the HasEdge predicate on the "Fund" edge with a given conditions (other predicates).
func HasFundWith(preds ...predicate.Fund) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FundInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FundTable, FundColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCertificate applies the HasEdge predicate on the "Certificate" edge.
func HasCertificate() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertificateTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CertificateTable, CertificateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCertificateWith applies the HasEdge predicate on the "Certificate" edge with a given conditions (other predicates).
func HasCertificateWith(preds ...predicate.Certificate) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertificateInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CertificateTable, CertificateColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMedical applies the HasEdge predicate on the "Medical" edge.
func HasMedical() predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MedicalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MedicalTable, MedicalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMedicalWith applies the HasEdge predicate on the "Medical" edge with a given conditions (other predicates).
func HasMedicalWith(preds ...predicate.Medical) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MedicalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MedicalTable, MedicalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.CoveredPerson) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.CoveredPerson) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
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
func Not(p predicate.CoveredPerson) predicate.CoveredPerson {
	return predicate.CoveredPerson(func(s *sql.Selector) {
		p(s.Not())
	})
}
