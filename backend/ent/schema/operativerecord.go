package schema

import (
	"time"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)

// Operativerecord holds the schema definition for the Operativerecord entity.
type Operativerecord struct {
	ent.Schema
}

// Fields of the Operativerecord.
func (Operativerecord) Fields() []ent.Field {
	return []ent.Field{

		field.Time("OperativeTime").Default(time.Now),
	}
}

// Edges of the Operativerecord.
func (Operativerecord) Edges() []ent.Edge {
	return []ent.Edge{ edge.From("Examinationroom", Examinationroom.Type).Ref("Examinationroom_Operativerecord").Unique(),
	edge.From("Nurse", Nurse.Type).Ref("Nurse_Operativerecord").Unique(),
	edge.From("Operative", Operative.Type).Ref("Operative_Operativerecord").Unique(),
	edge.From("Tool", Tool.Type).Ref("Tool_Operativerecord").Unique(),
}
}