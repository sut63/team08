package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// RoomType holds the schema definition for the RoomType entity.
type RoomType struct {
	ent.Schema
}

// Fields of the RoomType.
func (RoomType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("bathroom").Positive(),
		field.Float("areasize").Positive(),
		field.String("etc").NotEmpty(),
	}
}

// Edges of the RoomType.
func (RoomType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("room", Room.Type).StorageKey(edge.Column("roomtype_id")),
	}
}
