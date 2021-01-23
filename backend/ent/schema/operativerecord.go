package schema

import (
	"errors"
	"regexp"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Operativerecord holds the schema definition for the Operativerecord entity.
type Operativerecord struct {
	ent.Schema
}

// Fields of the Operativerecord.
func (Operativerecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("Nurse_Number").Validate(func(s string) error {
			match, _ := regexp.MatchString("[N]\\d{6}", s)
			if !match {
				return errors.New("รูปแบบรหัสพยาบาลไม่ถูกต้อง")
			}
			return nil
		}),

		field.Time("OperativeTime").Default(time.Now),
	}
}

// Edges of the Operativerecord.
func (Operativerecord) Edges() []ent.Edge {
	return []ent.Edge{edge.From("Examinationroom", Examinationroom.Type).Ref("Examinationroom_Operativerecord").Unique(),
		edge.From("Nurse", Nurse.Type).Ref("Nurse_Operativerecord").Unique(),
		edge.From("Operative", Operative.Type).Ref("Operative_Operativerecord").Unique(),
		edge.From("Tool", Tool.Type).Ref("Tool_Operativerecord").Unique(),
	}
}
