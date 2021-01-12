// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team08/ent/gender"
)

// Gender is the model entity for the Gender schema.
type Gender struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Gname holds the value of the "Gname" field.
	Gname string `json:"Gname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GenderQuery when eager-loading is set.
	Edges GenderEdges `json:"edges"`
}

// GenderEdges holds the relations/edges for other nodes in the graph.
type GenderEdges struct {
	// Fromgender holds the value of the fromgender edge.
	Fromgender []*Patient
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FromgenderOrErr returns the Fromgender value or an error if the edge
// was not loaded in eager-loading.
func (e GenderEdges) FromgenderOrErr() ([]*Patient, error) {
	if e.loadedTypes[0] {
		return e.Fromgender, nil
	}
	return nil, &NotLoadedError{edge: "fromgender"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Gender) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Gname
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Gender fields.
func (ge *Gender) assignValues(values ...interface{}) error {
	if m, n := len(values), len(gender.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ge.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Gname", values[0])
	} else if value.Valid {
		ge.Gname = value.String
	}
	return nil
}

// QueryFromgender queries the fromgender edge of the Gender.
func (ge *Gender) QueryFromgender() *PatientQuery {
	return (&GenderClient{config: ge.config}).QueryFromgender(ge)
}

// Update returns a builder for updating this Gender.
// Note that, you need to call Gender.Unwrap() before calling this method, if this Gender
// was returned from a transaction, and the transaction was committed or rolled back.
func (ge *Gender) Update() *GenderUpdateOne {
	return (&GenderClient{config: ge.config}).UpdateOne(ge)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ge *Gender) Unwrap() *Gender {
	tx, ok := ge.config.driver.(*txDriver)
	if !ok {
		panic("ent: Gender is not a transactional entity")
	}
	ge.config.driver = tx.drv
	return ge
}

// String implements the fmt.Stringer.
func (ge *Gender) String() string {
	var builder strings.Builder
	builder.WriteString("Gender(")
	builder.WriteString(fmt.Sprintf("id=%v", ge.ID))
	builder.WriteString(", Gname=")
	builder.WriteString(ge.Gname)
	builder.WriteByte(')')
	return builder.String()
}

// Genders is a parsable slice of Gender.
type Genders []*Gender

func (ge Genders) config(cfg config) {
	for _i := range ge {
		ge[_i].config = cfg
	}
}