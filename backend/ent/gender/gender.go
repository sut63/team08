// Code generated by entc, DO NOT EDIT.

package gender

const (
	// Label holds the string label denoting the gender type in the database.
	Label = "gender"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGname holds the string denoting the gname field in the database.
	FieldGname = "gname"

	// EdgeFromgender holds the string denoting the fromgender edge name in mutations.
	EdgeFromgender = "fromgender"

	// Table holds the table name of the gender in the database.
	Table = "genders"
	// FromgenderTable is the table the holds the fromgender relation/edge.
	FromgenderTable = "patients"
	// FromgenderInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	FromgenderInverseTable = "patients"
	// FromgenderColumn is the table column denoting the fromgender relation/edge.
	FromgenderColumn = "gender_id"
)

// Columns holds all SQL columns for gender fields.
var Columns = []string{
	FieldID,
	FieldGname,
}

var (
	// GnameValidator is a validator for the "Gname" field. It is called by the builders before save.
	GnameValidator func(string) error
)
