package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Disease struct {
	ent.Schema
}

func (Disease) Fields() []ent.Field {

	return []ent.Field{

		field.String("Disease_Name").NotEmpty(),
		
	}
}
func (Disease) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("disease_diagnose", Diagnose.Type).StorageKey(edge.Column("disease_id")),
	}
}