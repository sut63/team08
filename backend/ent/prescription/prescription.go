// Code generated by entc, DO NOT EDIT.

package prescription

import (
	"time"
)

const (
	// Label holds the string label denoting the prescription type in the database.
	Label = "prescription"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrescripNote holds the string denoting the prescrip_note field in the database.
	FieldPrescripNote = "prescrip_note"
	// FieldPrescripDateTime holds the string denoting the prescrip_datetime field in the database.
	FieldPrescripDateTime = "prescrip_date_time"

	// EdgeDoctor holds the string denoting the doctor edge name in mutations.
	EdgeDoctor = "doctor"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeNurse holds the string denoting the nurse edge name in mutations.
	EdgeNurse = "nurse"
	// EdgeDrug holds the string denoting the drug edge name in mutations.
	EdgeDrug = "drug"

	// Table holds the table name of the prescription in the database.
	Table = "prescriptions"
	// DoctorTable is the table the holds the doctor relation/edge.
	DoctorTable = "prescriptions"
	// DoctorInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	DoctorInverseTable = "doctors"
	// DoctorColumn is the table column denoting the doctor relation/edge.
	DoctorColumn = "doctor_id"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "prescriptions"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_id"
	// NurseTable is the table the holds the nurse relation/edge.
	NurseTable = "prescriptions"
	// NurseInverseTable is the table name for the Nurse entity.
	// It exists in this package in order to avoid circular dependency with the "nurse" package.
	NurseInverseTable = "nurses"
	// NurseColumn is the table column denoting the nurse relation/edge.
	NurseColumn = "nurse_id"
	// DrugTable is the table the holds the drug relation/edge.
	DrugTable = "prescriptions"
	// DrugInverseTable is the table name for the Drug entity.
	// It exists in this package in order to avoid circular dependency with the "drug" package.
	DrugInverseTable = "drugs"
	// DrugColumn is the table column denoting the drug relation/edge.
	DrugColumn = "drug_id"
)

// Columns holds all SQL columns for prescription fields.
var Columns = []string{
	FieldID,
	FieldPrescripNote,
	FieldPrescripDateTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Prescription type.
var ForeignKeys = []string{
	"doctor_id",
	"drug_id",
	"nurse_id",
	"patient_id",
}

var (
	// DefaultPrescripDateTime holds the default value on creation for the Prescrip_DateTime field.
	DefaultPrescripDateTime func() time.Time
)
