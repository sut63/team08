package schema

import (
	"github.com/facebookincubator/ent"

	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type Department struct {
	ent.Schema
}

func (Department) Fields() []ent.Field {

	return []ent.Field{

		field.String("Department_Name").NotEmpty(),
	}

}
func (Department) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("department_diagnose", Diagnose.Type).StorageKey(edge.Column("department_id")),
	}
}