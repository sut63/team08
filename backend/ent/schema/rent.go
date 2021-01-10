package schema

import (
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
