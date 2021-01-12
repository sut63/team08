package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Gender holds the schema definition for the gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the gender.
func (Gender) Fields() []ent.Field {
	return []ent.Field{
		field.String("Gname").NotEmpty(),
	}
}

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fromgender", Patient.Type).StorageKey(edge.Column("gender_id")),
	}
}
