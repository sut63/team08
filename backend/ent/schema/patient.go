package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("Patient_name").NotEmpty(),
		field.Int("Patient_age").Positive(),
		field.Float("Patient_weight").Positive(),
		field.Float("Patient_height").Positive(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).Ref("fromgender").Unique(),
		edge.From("prefix", Prefix.Type).Ref("fromprefix").Unique(),
		edge.From("bloodtype", Bloodtype.Type).Ref("frombloodtype").Unique(),

		edge.To("frompatient", Rent.Type).StorageKey(edge.Column("Patient_id")).Unique(),
		edge.To("Patient_CoveredPerson", CoveredPerson.Type).StorageKey(edge.Column("Patient_id")),
		edge.To("patient_diagnose", Diagnose.Type).StorageKey(edge.Column("patient_id")),
		edge.To("patient_prescription", Prescription.Type).StorageKey(edge.Column("patient_id")),
	}
}
