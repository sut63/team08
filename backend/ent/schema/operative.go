package schema
 
import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)
// Operative holds the schema definition for the Operative entity.
type Operative struct {
    ent.Schema
}
// Fields of the Operative.
func (Operative) Fields() []ent.Field {
    return []ent.Field{
		field.String("operative_Name").NotEmpty().Unique(),
    }
 }
 //Edges of the Operative.
 func (Operative) Edges() []ent.Edge {
    return []ent.Edge{
		edge. To("Operative_Operativerecord",Operativerecord.Type).StorageKey(edge.Column("Operative_id")),
 }
 }