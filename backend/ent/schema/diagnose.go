package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Diagnose struct {
	ent.Schema
}

func (Diagnose) Fields() []ent.Field {
	return []ent.Field{
		field.String("Diagnose_ID").Validate(func(s string) error {
			match, _ := regexp.MatchString("[D]\\d{4}", s)
			if !match {
				return errors.New("รหัสต้องขึ้นต้นด้วย D ตามด้วยตัวเลข 4 ตัว")
			}
			return nil
		}),
		field.String("Diagnose_Symptoms").NotEmpty(),

		field.String("Diagnose_Note").NotEmpty(),
	}
}
//Edges of the Diagnose.
func (Diagnose) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("disease", Disease.Type).Ref("disease_diagnose").Unique(),
		edge.From("department", Department.Type).Ref("department_diagnose").Unique(),
		edge.From("patient", Patient.Type).Ref("patient_diagnose").Unique(),
		edge.From("doctor", Doctor.Type).Ref("doctor_diagnose").Unique(),
	}

}
