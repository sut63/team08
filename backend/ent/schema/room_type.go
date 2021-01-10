package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Roomtype holds the schema definition for the Roomtype entity.
type Roomtype struct {
	ent.Schema
}

// Fields of the Roomtype.
func (Roomtype) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("bathroom"),
		field.Int("toilets"),
		field.Float("areasize").Positive(),
		field.String("etc").NotEmpty(),
	}
}

// Edges of the Roomtype.
func (Roomtype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rooms", Room.Type).StorageKey(edge.Column("roomtype_id")),
	}
}
