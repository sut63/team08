// Code generated by entc, DO NOT EDIT.

package rent

const (
	// Label holds the string label denoting the rent type in the database.
	Label = "rent"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRentID holds the string denoting the rent_id field in the database.
	FieldRentID = "rent_id"
	// FieldKinTel holds the string denoting the kin_tel field in the database.
	FieldKinTel = "kin_tel"
	// FieldKinName holds the string denoting the kin_name field in the database.
	FieldKinName = "kin_name"
	// FieldAddedTime holds the string denoting the added_time field in the database.
	FieldAddedTime = "added_time"

	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeNurse holds the string denoting the nurse edge name in mutations.
	EdgeNurse = "nurse"

	// Table holds the table name of the rent in the database.
	Table = "rents"
	// RoomTable is the table the holds the room relation/edge.
	RoomTable = "rents"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_id"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "rents"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "Patient_id"
	// NurseTable is the table the holds the nurse relation/edge.
	NurseTable = "rents"
	// NurseInverseTable is the table name for the Nurse entity.
	// It exists in this package in order to avoid circular dependency with the "nurse" package.
	NurseInverseTable = "nurses"
	// NurseColumn is the table column denoting the nurse relation/edge.
	NurseColumn = "nurse_id"
)

// Columns holds all SQL columns for rent fields.
var Columns = []string{
	FieldID,
	FieldRentID,
	FieldKinTel,
	FieldKinName,
	FieldAddedTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Rent type.
var ForeignKeys = []string{
	"nurse_id",
	"Patient_id",
	"room_id",
}

var (
	// RentIDValidator is a validator for the "rent_id" field. It is called by the builders before save.
	RentIDValidator func(string) error
	// KinTelValidator is a validator for the "kin_tel" field. It is called by the builders before save.
	KinTelValidator func(string) error
	// KinNameValidator is a validator for the "kin_name" field. It is called by the builders before save.
	KinNameValidator func(string) error
)
