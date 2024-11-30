package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Pixel holds the schema definition for the Pixel entity.
type Pixel struct {
	ent.Schema
}

// Fields of the Pixel.
func (Pixel) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StructTag(`json:"oid,omitempty"`),
		field.String("color"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}

}

// Edges of the Pixel.
func (Pixel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("pixels").
			Unique(),
	}
}
