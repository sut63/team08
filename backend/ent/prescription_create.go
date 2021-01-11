// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/doctor"
	"github.com/sut63/team08/ent/drug"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/prescription"
)

// PrescriptionCreate is the builder for creating a Prescription entity.
type PrescriptionCreate struct {
	config
	mutation *PrescriptionMutation
	hooks    []Hook
}

// SetPrescripNote sets the Prescrip_Note field.
func (pc *PrescriptionCreate) SetPrescripNote(s string) *PrescriptionCreate {
	pc.mutation.SetPrescripNote(s)
	return pc
}

// SetPrescripDateTime sets the Prescrip_DateTime field.
func (pc *PrescriptionCreate) SetPrescripDateTime(t time.Time) *PrescriptionCreate {
	pc.mutation.SetPrescripDateTime(t)
	return pc
}

// SetNillablePrescripDateTime sets the Prescrip_DateTime field if the given value is not nil.
func (pc *PrescriptionCreate) SetNillablePrescripDateTime(t *time.Time) *PrescriptionCreate {
	if t != nil {
		pc.SetPrescripDateTime(*t)
	}
	return pc
}

// SetDoctorID sets the doctor edge to Doctor by id.
func (pc *PrescriptionCreate) SetDoctorID(id int) *PrescriptionCreate {
	pc.mutation.SetDoctorID(id)
	return pc
}

// SetNillableDoctorID sets the doctor edge to Doctor by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableDoctorID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetDoctorID(*id)
	}
	return pc
}

// SetDoctor sets the doctor edge to Doctor.
func (pc *PrescriptionCreate) SetDoctor(d *Doctor) *PrescriptionCreate {
	return pc.SetDoctorID(d.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (pc *PrescriptionCreate) SetPatientID(id int) *PrescriptionCreate {
	pc.mutation.SetPatientID(id)
	return pc
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillablePatientID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetPatientID(*id)
	}
	return pc
}

// SetPatient sets the patient edge to Patient.
func (pc *PrescriptionCreate) SetPatient(p *Patient) *PrescriptionCreate {
	return pc.SetPatientID(p.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (pc *PrescriptionCreate) SetNurseID(id int) *PrescriptionCreate {
	pc.mutation.SetNurseID(id)
	return pc
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableNurseID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetNurseID(*id)
	}
	return pc
}

// SetNurse sets the nurse edge to Nurse.
func (pc *PrescriptionCreate) SetNurse(n *Nurse) *PrescriptionCreate {
	return pc.SetNurseID(n.ID)
}

// SetDrugID sets the drug edge to Drug by id.
func (pc *PrescriptionCreate) SetDrugID(id int) *PrescriptionCreate {
	pc.mutation.SetDrugID(id)
	return pc
}

// SetNillableDrugID sets the drug edge to Drug by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableDrugID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetDrugID(*id)
	}
	return pc
}

// SetDrug sets the drug edge to Drug.
func (pc *PrescriptionCreate) SetDrug(d *Drug) *PrescriptionCreate {
	return pc.SetDrugID(d.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (pc *PrescriptionCreate) Mutation() *PrescriptionMutation {
	return pc.mutation
}

// Save creates the Prescription in the database.
func (pc *PrescriptionCreate) Save(ctx context.Context) (*Prescription, error) {
	if _, ok := pc.mutation.PrescripNote(); !ok {
		return nil, &ValidationError{Name: "Prescrip_Note", err: errors.New("ent: missing required field \"Prescrip_Note\"")}
	}
	if _, ok := pc.mutation.PrescripDateTime(); !ok {
		v := prescription.DefaultPrescripDateTime()
		pc.mutation.SetPrescripDateTime(v)
	}
	var (
		err  error
		node *Prescription
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PrescriptionCreate) SaveX(ctx context.Context) *Prescription {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PrescriptionCreate) sqlSave(ctx context.Context) (*Prescription, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *PrescriptionCreate) createSpec() (*Prescription, *sqlgraph.CreateSpec) {
	var (
		pr    = &Prescription{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: prescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PrescripNote(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prescription.FieldPrescripNote,
		})
		pr.PrescripNote = value
	}
	if value, ok := pc.mutation.PrescripDateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: prescription.FieldPrescripDateTime,
		})
		pr.PrescripDateTime = value
	}
	if nodes := pc.mutation.DoctorIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.NurseIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DrugIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}
