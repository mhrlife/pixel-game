package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Hype struct {
	ent.Schema
}

func (Hype) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount_remaining"),
		field.Int("max_hype"),
		field.Time("last_updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Int("hype_per_minute").Default(2),
	}
}

func (Hype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("hype").Unique().Required(),
	}
}
