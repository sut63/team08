package schema

import (
	"time"

	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Prescription holds the schema definition for the Prescription entity.
type Prescription struct {
	ent.Schema
}

// Fields of the Prescription.
func (Prescription) Fields() []ent.Field {

	return []ent.Field{
		field.String("Prescrip_Number").Validate(func(s string) error {
			match, _ := regexp.MatchString("[P]\\d{3}", s)
			if !match {
				return errors.New("รูปแบบรหัสยาไม่ถูกต้อง")
			}
			return nil
		}),

		field.String("Prescrip_Issue").NotEmpty(),

		field.String("Prescrip_Note").NotEmpty(),

		field.Time("Prescrip_DateTime").Default(time.Now),
	}
}

// Edges of the Prescription.
func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor", Doctor.Type).Ref("doctor_prescription").Unique(),
		edge.From("patient", Patient.Type).Ref("patient_prescription").Unique(),
		edge.From("nurse", Nurse.Type).Ref("nurse_prescription").Unique(),
		edge.From("drug", Drug.Type).Ref("drug_prescription").Unique(),
	}
}
