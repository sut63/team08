package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// covered_person holds the schema definition for the covered_person entity.
type covered_person struct {
	ent.Schema
}

// Fields of the covered_person.
func (covered_person) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time"),
	}
}

// Edges of the covered_person.
func (covered_person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("playlist", Playlist.Type).Ref("covered_person").Unique(),
		edge.From("video", Video.Type).Ref("covered_person").Unique(),
		edge.From("resolution", Resolution.Type).Ref("covered_person").Unique(),
	}
}
