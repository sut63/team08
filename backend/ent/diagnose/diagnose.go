// Code generated by entc, DO NOT EDIT.

package diagnose

const (
	// Label holds the string label denoting the diagnose type in the database.
	Label = "diagnose"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDiagnoseID holds the string denoting the diagnose_id field in the database.
	FieldDiagnoseID = "diagnose_id"
	// FieldDiagnoseSymptoms holds the string denoting the diagnose_symptoms field in the database.
	FieldDiagnoseSymptoms = "diagnose_symptoms"
	// FieldDiagnoseNote holds the string denoting the diagnose_note field in the database.
	FieldDiagnoseNote = "diagnose_note"

	// EdgeDisease holds the string denoting the disease edge name in mutations.
	EdgeDisease = "disease"
	// EdgeDepartment holds the string denoting the department edge name in mutations.
	EdgeDepartment = "department"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeDoctor holds the string denoting the doctor edge name in mutations.
	EdgeDoctor = "doctor"

	// Table holds the table name of the diagnose in the database.
	Table = "diagnoses"
	// DiseaseTable is the table the holds the disease relation/edge.
	DiseaseTable = "diagnoses"
	// DiseaseInverseTable is the table name for the Disease entity.
	// It exists in this package in order to avoid circular dependency with the "disease" package.
	DiseaseInverseTable = "diseases"
	// DiseaseColumn is the table column denoting the disease relation/edge.
	DiseaseColumn = "disease_id"
	// DepartmentTable is the table the holds the department relation/edge.
	DepartmentTable = "diagnoses"
	// DepartmentInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentInverseTable = "departments"
	// DepartmentColumn is the table column denoting the department relation/edge.
	DepartmentColumn = "department_id"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "diagnoses"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_id"
	// DoctorTable is the table the holds the doctor relation/edge.
	DoctorTable = "diagnoses"
	// DoctorInverseTable is the table name for the Doctor entity.
	// It exists in this package in order to avoid circular dependency with the "doctor" package.
	DoctorInverseTable = "doctors"
	// DoctorColumn is the table column denoting the doctor relation/edge.
	DoctorColumn = "doctor_id"
)

// Columns holds all SQL columns for diagnose fields.
var Columns = []string{
	FieldID,
	FieldDiagnoseID,
	FieldDiagnoseSymptoms,
	FieldDiagnoseNote,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Diagnose type.
var ForeignKeys = []string{
	"department_id",
	"disease_id",
	"doctor_id",
	"patient_id",
}

var (
	// DiagnoseIDValidator is a validator for the "Diagnose_ID" field. It is called by the builders before save.
	DiagnoseIDValidator func(string) error
	// DiagnoseSymptomsValidator is a validator for the "Diagnose_Symptoms" field. It is called by the builders before save.
	DiagnoseSymptomsValidator func(string) error
	// DiagnoseNoteValidator is a validator for the "Diagnose_Note" field. It is called by the builders before save.
	DiagnoseNoteValidator func(string) error
)