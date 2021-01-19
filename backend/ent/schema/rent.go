package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Rent holds the schema definition for the Rent entity.
type Rent struct {
	ent.Schema
}

// Fields of the Rent.
func (Rent) Fields() []ent.Field {
	return []ent.Field{
		field.String("rent_id").Validate(func(s string)error {
			match, _ := regexp.MatchString("[R]\\d{4}", s)
			if !match{
				return errors.New("รูปแบบรหัสการจองไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("kin_tel").Validate(func(s string)error {
			match, _ := regexp.MatchString("[0]\\d{9}", s)
			if !match{
				return errors.New("รูปแบบของเบอร์โทรศัพท์ไม่ถูกต้อง")
			}
			return nil
		}),
		field.String("kin_name").NotEmpty(),
		field.Time("added_time"),

	}
}

// Edges of the Rent.
func (Rent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Ref("rents").Unique(),
		edge.From("patient", Patient.Type).Ref("frompatient").Unique(),
		edge.From("nurse", Nurse.Type).Ref("fromnurse").Unique(),
	}
}
