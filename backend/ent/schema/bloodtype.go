package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bloodtype holds the schema definition for the bloodtype entity.
type Bloodtype struct {
	ent.Schema
}

// Fields of the Bloodtype.
func (Bloodtype) Fields() []ent.Field {
	return []ent.Field{
		field.String("BTname").NotEmpty(),
	}
}

// Edges of the Bloodtype.
func (Bloodtype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("frombloodtype", Patient.Type).StorageKey(edge.Column("bloodtype_id")),
	}
}
