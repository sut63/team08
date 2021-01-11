package schema

import (
	"github.com/facebookincubator/ent"

	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Doctor struct {
	ent.Schema
}

func (Doctor) Fields() []ent.Field {

	return []ent.Field{

		field.String("Doctor_Name").NotEmpty(),
		field.String("Doctor_Password").NotEmpty(),
		field.String("Doctor_Email").NotEmpty(),
		field.String("Doctor_tel").NotEmpty(),
	}
}
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doctor_diagnose", Diagnose.Type).StorageKey(edge.Column("doctor_id")),
		edge.To("doctor_prescription", Prescription.Type).StorageKey(edge.Column("doctor_id")),
	}
}
