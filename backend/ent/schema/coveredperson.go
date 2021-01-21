package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// CoveredPerson holds the schema definition for the CoveredPerson entity.
type CoveredPerson struct {
	ent.Schema
}

// Fields of the CoveredPerson.
func (CoveredPerson) Fields() []ent.Field {
	return []ent.Field{

		field.String("CoveredPerson_Number").Validate(func(s string) error {
			match, _ := regexp.MatchString("[CP]\\d{3}", s)
			if !match {
				return errors.New("รูปแบบรหัสสิทธิการรักษาไม่ถูกต้อง")
			}
			return nil
		}),

		field.String("CoveredPerson_Note").NotEmpty(),
		field.String("Fund_Title").NotEmpty(),
	}
}

//Edges of the CoveredPerson.
func (CoveredPerson) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Patient", Patient.Type).Ref("Patient_CoveredPerson").Unique(),
		edge.From("SchemeType", SchemeType.Type).Ref("SchemeType_CoveredPerson").Unique(),
		edge.From("Fund", Fund.Type).Ref("Fund_CoveredPerson").Unique(),
		edge.From("Certificate", Certificate.Type).Ref("Certificate_CoveredPerson").Unique(),
	}
}
