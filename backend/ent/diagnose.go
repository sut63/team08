// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team08/ent/department"
	"github.com/sut63/team08/ent/diagnose"
	"github.com/sut63/team08/ent/disease"
	"github.com/sut63/team08/ent/doctor"
	"github.com/sut63/team08/ent/patient"
)

// Diagnose is the model entity for the Diagnose schema.
type Diagnose struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DiagnoseID holds the value of the "Diagnose_ID" field.
	DiagnoseID string `json:"Diagnose_ID,omitempty"`
	// DiagnoseSymptoms holds the value of the "Diagnose_Symptoms" field.
	DiagnoseSymptoms string `json:"Diagnose_Symptoms,omitempty"`
	// DiagnoseNote holds the value of the "Diagnose_Note" field.
	DiagnoseNote string `json:"Diagnose_Note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiagnoseQuery when eager-loading is set.
	Edges         DiagnoseEdges `json:"edges"`
	department_id *int
	disease_id    *int
	doctor_id     *int
	patient_id    *int
}

// DiagnoseEdges holds the relations/edges for other nodes in the graph.
type DiagnoseEdges struct {
	// Disease holds the value of the disease edge.
	Disease *Disease
	// Department holds the value of the department edge.
	Department *Department
	// Patient holds the value of the patient edge.
	Patient *Patient
	// Doctor holds the value of the doctor edge.
	Doctor *Doctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// DiseaseOrErr returns the Disease value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiagnoseEdges) DiseaseOrErr() (*Disease, error) {
	if e.loadedTypes[0] {
		if e.Disease == nil {
			// The edge disease was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: disease.Label}
		}
		return e.Disease, nil
	}
	return nil, &NotLoadedError{edge: "disease"}
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiagnoseEdges) DepartmentOrErr() (*Department, error) {
	if e.loadedTypes[1] {
		if e.Department == nil {
			// The edge department was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.Department, nil
	}
	return nil, &NotLoadedError{edge: "department"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiagnoseEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[2] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// DoctorOrErr returns the Doctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiagnoseEdges) DoctorOrErr() (*Doctor, error) {
	if e.loadedTypes[3] {
		if e.Doctor == nil {
			// The edge doctor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.Doctor, nil
	}
	return nil, &NotLoadedError{edge: "doctor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Diagnose) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Diagnose_ID
		&sql.NullString{}, // Diagnose_Symptoms
		&sql.NullString{}, // Diagnose_Note
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Diagnose) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // department_id
		&sql.NullInt64{}, // disease_id
		&sql.NullInt64{}, // doctor_id
		&sql.NullInt64{}, // patient_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Diagnose fields.
func (d *Diagnose) assignValues(values ...interface{}) error {
	if m, n := len(values), len(diagnose.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	d.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Diagnose_ID", values[0])
	} else if value.Valid {
		d.DiagnoseID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Diagnose_Symptoms", values[1])
	} else if value.Valid {
		d.DiagnoseSymptoms = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Diagnose_Note", values[2])
	} else if value.Valid {
		d.DiagnoseNote = value.String
	}
	values = values[3:]
	if len(values) == len(diagnose.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_id", value)
		} else if value.Valid {
			d.department_id = new(int)
			*d.department_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field disease_id", value)
		} else if value.Valid {
			d.disease_id = new(int)
			*d.disease_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field doctor_id", value)
		} else if value.Valid {
			d.doctor_id = new(int)
			*d.doctor_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field patient_id", value)
		} else if value.Valid {
			d.patient_id = new(int)
			*d.patient_id = int(value.Int64)
		}
	}
	return nil
}

// QueryDisease queries the disease edge of the Diagnose.
func (d *Diagnose) QueryDisease() *DiseaseQuery {
	return (&DiagnoseClient{config: d.config}).QueryDisease(d)
}

// QueryDepartment queries the department edge of the Diagnose.
func (d *Diagnose) QueryDepartment() *DepartmentQuery {
	return (&DiagnoseClient{config: d.config}).QueryDepartment(d)
}

// QueryPatient queries the patient edge of the Diagnose.
func (d *Diagnose) QueryPatient() *PatientQuery {
	return (&DiagnoseClient{config: d.config}).QueryPatient(d)
}

// QueryDoctor queries the doctor edge of the Diagnose.
func (d *Diagnose) QueryDoctor() *DoctorQuery {
	return (&DiagnoseClient{config: d.config}).QueryDoctor(d)
}

// Update returns a builder for updating this Diagnose.
// Note that, you need to call Diagnose.Unwrap() before calling this method, if this Diagnose
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Diagnose) Update() *DiagnoseUpdateOne {
	return (&DiagnoseClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Diagnose) Unwrap() *Diagnose {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Diagnose is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Diagnose) String() string {
	var builder strings.Builder
	builder.WriteString("Diagnose(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", Diagnose_ID=")
	builder.WriteString(d.DiagnoseID)
	builder.WriteString(", Diagnose_Symptoms=")
	builder.WriteString(d.DiagnoseSymptoms)
	builder.WriteString(", Diagnose_Note=")
	builder.WriteString(d.DiagnoseNote)
	builder.WriteByte(')')
	return builder.String()
}

// Diagnoses is a parsable slice of Diagnose.
type Diagnoses []*Diagnose

func (d Diagnoses) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
