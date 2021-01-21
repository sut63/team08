// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/doctor"
	"github.com/sut63/team08/ent/drug"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/predicate"
	"github.com/sut63/team08/ent/prescription"
)

// PrescriptionUpdate is the builder for updating Prescription entities.
type PrescriptionUpdate struct {
	config
	hooks      []Hook
	mutation   *PrescriptionMutation
	predicates []predicate.Prescription
}

// Where adds a new predicate for the builder.
func (pu *PrescriptionUpdate) Where(ps ...predicate.Prescription) *PrescriptionUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPrescripNumber sets the Prescrip_Number field.
func (pu *PrescriptionUpdate) SetPrescripNumber(s string) *PrescriptionUpdate {
	pu.mutation.SetPrescripNumber(s)
	return pu
}

// SetPrescripIssue sets the Prescrip_Issue field.
func (pu *PrescriptionUpdate) SetPrescripIssue(s string) *PrescriptionUpdate {
	pu.mutation.SetPrescripIssue(s)
	return pu
}

// SetPrescripNote sets the Prescrip_Note field.
func (pu *PrescriptionUpdate) SetPrescripNote(s string) *PrescriptionUpdate {
	pu.mutation.SetPrescripNote(s)
	return pu
}

// SetPrescripDateTime sets the Prescrip_DateTime field.
func (pu *PrescriptionUpdate) SetPrescripDateTime(t time.Time) *PrescriptionUpdate {
	pu.mutation.SetPrescripDateTime(t)
	return pu
}

// SetNillablePrescripDateTime sets the Prescrip_DateTime field if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillablePrescripDateTime(t *time.Time) *PrescriptionUpdate {
	if t != nil {
		pu.SetPrescripDateTime(*t)
	}
	return pu
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (pu *PrescriptionUpdate) SetDoctorID(id int) *PrescriptionUpdate {
	pu.mutation.SetDoctorID(id)
	return pu
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillableDoctorID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetDoctorID(*id)
	}
	return pu
}

// SetDoctor sets the doctor edge to Doctor.
func (pu *PrescriptionUpdate) SetDoctor(d *Doctor) *PrescriptionUpdate {
	return pu.SetDoctorID(d.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (pu *PrescriptionUpdate) SetPatientID(id int) *PrescriptionUpdate {
	pu.mutation.SetPatientID(id)
	return pu
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillablePatientID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetPatientID(*id)
	}
	return pu
}

// SetPatient sets the patient edge to Patient.
func (pu *PrescriptionUpdate) SetPatient(p *Patient) *PrescriptionUpdate {
	return pu.SetPatientID(p.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (pu *PrescriptionUpdate) SetNurseID(id int) *PrescriptionUpdate {
	pu.mutation.SetNurseID(id)
	return pu
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillableNurseID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetNurseID(*id)
	}
	return pu
}

// SetNurse sets the nurse edge to Nurse.
func (pu *PrescriptionUpdate) SetNurse(n *Nurse) *PrescriptionUpdate {
	return pu.SetNurseID(n.ID)
}

// SetDrugID sets the drug edge to Drug by id.
func (pu *PrescriptionUpdate) SetDrugID(id int) *PrescriptionUpdate {
	pu.mutation.SetDrugID(id)
	return pu
}

// SetNillableDrugID sets the drug edge to Drug by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillableDrugID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetDrugID(*id)
	}
	return pu
}

// SetDrug sets the drug edge to Drug.
func (pu *PrescriptionUpdate) SetDrug(d *Drug) *PrescriptionUpdate {
	return pu.SetDrugID(d.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (pu *PrescriptionUpdate) Mutation() *PrescriptionMutation {
	return pu.mutation
}

// ClearDoctor clears the doctor edge to Doctor.
func (pu *PrescriptionUpdate) ClearDoctor() *PrescriptionUpdate {
	pu.mutation.ClearDoctor()
	return pu
}

// ClearPatient clears the patient edge to Patient.
func (pu *PrescriptionUpdate) ClearPatient() *PrescriptionUpdate {
	pu.mutation.ClearPatient()
	return pu
}

// ClearNurse clears the nurse edge to Nurse.
func (pu *PrescriptionUpdate) ClearNurse() *PrescriptionUpdate {
	pu.mutation.ClearNurse()
	return pu
}

// ClearDrug clears the drug edge to Drug.
func (pu *PrescriptionUpdate) ClearDrug() *PrescriptionUpdate {
	pu.mutation.ClearDrug()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PrescriptionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.PrescripNumber(); ok {
		if err := prescription.PrescripNumberValidator(v); err != nil {
			return 0, &ValidationError{Name: "Prescrip_Number", err: fmt.Errorf("ent: validator failed for field \"Prescrip_Number\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PrescripIssue(); ok {
		if err := prescription.PrescripIssueValidator(v); err != nil {
			return 0, &ValidationError{Name: "Prescrip_Issue", err: fmt.Errorf("ent: validator failed for field \"Prescrip_Issue\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PrescripNote(); ok {
		if err := prescription.PrescripNoteValidator(v); err != nil {
			return 0, &ValidationError{Name: "Prescrip_Note", err: fmt.Errorf("ent: validator failed for field \"Prescrip_Note\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrescriptionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrescriptionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prescription.Table,
			Columns: prescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.PrescripNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prescription.FieldPrescripNumber,
		})
	}
	if value, ok := pu.mutation.PrescripIssue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prescription.FieldPrescripIssue,
		})
	}
	if value, ok := pu.mutation.PrescripNote(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prescription.FieldPrescripNote,
		})
	}
	if value, ok := pu.mutation.PrescripDateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: prescription.FieldPrescripDateTime,
		})
	}
	if pu.mutation.DoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.NurseTable,
			Columns: []string{prescription.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.NurseTable,
			Columns: []string{prescription.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DrugCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DrugTable,
			Columns: []string{prescription.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DrugIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DrugTable,
			Columns: []string{prescription.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescription.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PrescriptionUpdateOne is the builder for updating a single Prescription entity.
type PrescriptionUpdateOne struct {
	config
	hooks    []Hook
	mutation *PrescriptionMutation
}

// SetPrescripNumber sets the Prescrip_Number field.
func (puo *PrescriptionUpdateOne) SetPrescripNumber(s string) *PrescriptionUpdateOne {
	puo.mutation.SetPrescripNumber(s)
	return puo
}

// SetPrescripIssue sets the Prescrip_Issue field.
func (puo *PrescriptionUpdateOne) SetPrescripIssue(s string) *PrescriptionUpdateOne {
	puo.mutation.SetPrescripIssue(s)
	return puo
}

// SetPrescripNote sets the Prescrip_Note field.
func (puo *PrescriptionUpdateOne) SetPrescripNote(s string) *PrescriptionUpdateOne {
	puo.mutation.SetPrescripNote(s)
	return puo
}

// SetPrescripDateTime sets the Prescrip_DateTime field.
func (puo *PrescriptionUpdateOne) SetPrescripDateTime(t time.Time) *PrescriptionUpdateOne {
	puo.mutation.SetPrescripDateTime(t)
	return puo
}

// SetNillablePrescripDateTime sets the Prescrip_DateTime field if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillablePrescripDateTime(t *time.Time) *PrescriptionUpdateOne {
	if t != nil {
		puo.SetPrescripDateTime(*t)
	}
	return puo
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (puo *PrescriptionUpdateOne) SetDoctorID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetDoctorID(id)
	return puo
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillableDoctorID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetDoctorID(*id)
	}
	return puo
}

// SetDoctor sets the doctor edge to Doctor.
func (puo *PrescriptionUpdateOne) SetDoctor(d *Doctor) *PrescriptionUpdateOne {
	return puo.SetDoctorID(d.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (puo *PrescriptionUpdateOne) SetPatientID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetPatientID(id)
	return puo
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillablePatientID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetPatientID(*id)
	}
	return puo
}

// SetPatient sets the patient edge to Patient.
func (puo *PrescriptionUpdateOne) SetPatient(p *Patient) *PrescriptionUpdateOne {
	return puo.SetPatientID(p.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (puo *PrescriptionUpdateOne) SetNurseID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetNurseID(id)
	return puo
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillableNurseID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetNurseID(*id)
	}
	return puo
}

// SetNurse sets the nurse edge to Nurse.
func (puo *PrescriptionUpdateOne) SetNurse(n *Nurse) *PrescriptionUpdateOne {
	return puo.SetNurseID(n.ID)
}

// SetDrugID sets the drug edge to Drug by id.
func (puo *PrescriptionUpdateOne) SetDrugID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetDrugID(id)
	return puo
}

// SetNillableDrugID sets the drug edge to Drug by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillableDrugID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetDrugID(*id)
	}
	return puo
}

// SetDrug sets the drug edge to Drug.
func (puo *PrescriptionUpdateOne) SetDrug(d *Drug) *PrescriptionUpdateOne {
	return puo.SetDrugID(d.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (puo *PrescriptionUpdateOne) Mutation() *PrescriptionMutation {
	return puo.mutation
}

// ClearDoctor clears the doctor edge to Doctor.
func (puo *PrescriptionUpdateOne) ClearDoctor() *PrescriptionUpdateOne {
	puo.mutation.ClearDoctor()
	return puo
}

// ClearPatient clears the patient edge to Patient.
func (puo *PrescriptionUpdateOne) ClearPatient() *PrescriptionUpdateOne {
	puo.mutation.ClearPatient()
	return puo
}

// ClearNurse clears the nurse edge to Nurse.
func (puo *PrescriptionUpdateOne) ClearNurse() *PrescriptionUpdateOne {
	puo.mutation.ClearNurse()
	return puo
}

// ClearDrug clears the drug edge to Drug.
func (puo *PrescriptionUpdateOne) ClearDrug() *PrescriptionUpdateOne {
	puo.mutation.ClearDrug()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PrescriptionUpdateOne) Save(ctx context.Context) (*Prescription, error) {
	if v, ok := puo.mutation.PrescripNumber(); ok {
		if err := prescription.PrescripNumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "Prescrip_Number", err: fmt.Errorf("ent: validator failed for field \"Prescrip_Number\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PrescripIssue(); ok {
		if err := prescription.PrescripIssueValidator(v); err != nil {
			return nil, &ValidationError{Name: "Prescrip_Issue", err: fmt.Errorf("ent: validator failed for field \"Prescrip_Issue\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PrescripNote(); ok {
		if err := prescription.PrescripNoteValidator(v); err != nil {
			return nil, &ValidationError{Name: "Prescrip_Note", err: fmt.Errorf("ent: validator failed for field \"Prescrip_Note\": %w", err)}
		}
	}

	var (
		err  error
		node *Prescription
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrescriptionUpdateOne) SaveX(ctx context.Context) *Prescription {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PrescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrescriptionUpdateOne) sqlSave(ctx context.Context) (pr *Prescription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prescription.Table,
			Columns: prescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Prescription.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.PrescripNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prescription.FieldPrescripNumber,
		})
	}
	if value, ok := puo.mutation.PrescripIssue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prescription.FieldPrescripIssue,
		})
	}
	if value, ok := puo.mutation.PrescripNote(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prescription.FieldPrescripNote,
		})
	}
	if value, ok := puo.mutation.PrescripDateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: prescription.FieldPrescripDateTime,
		})
	}
	if puo.mutation.DoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DoctorTable,
			Columns: []string{prescription.DoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PatientTable,
			Columns: []string{prescription.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.NurseTable,
			Columns: []string{prescription.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.NurseTable,
			Columns: []string{prescription.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DrugCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DrugTable,
			Columns: []string{prescription.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DrugIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.DrugTable,
			Columns: []string{prescription.DrugColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drug.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Prescription{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescription.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
