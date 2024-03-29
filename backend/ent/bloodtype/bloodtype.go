// Code generated by entc, DO NOT EDIT.

package bloodtype

const (
	// Label holds the string label denoting the bloodtype type in the database.
	Label = "bloodtype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBTname holds the string denoting the btname field in the database.
	FieldBTname = "btname"

	// EdgeFrombloodtype holds the string denoting the frombloodtype edge name in mutations.
	EdgeFrombloodtype = "frombloodtype"

	// Table holds the table name of the bloodtype in the database.
	Table = "bloodtypes"
	// FrombloodtypeTable is the table the holds the frombloodtype relation/edge.
	FrombloodtypeTable = "patients"
	// FrombloodtypeInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	FrombloodtypeInverseTable = "patients"
	// FrombloodtypeColumn is the table column denoting the frombloodtype relation/edge.
	FrombloodtypeColumn = "bloodtype_id"
)

// Columns holds all SQL columns for bloodtype fields.
var Columns = []string{
	FieldID,
	FieldBTname,
}

var (
	// BTnameValidator is a validator for the "BTname" field. It is called by the builders before save.
	BTnameValidator func(string) error
)
