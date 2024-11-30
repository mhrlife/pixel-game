package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("display_name"),
		field.String("game_id"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pixels", Pixel.Type),
		edge.To("hype", Hype.Type).Unique(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("game_id"),
	}
}
