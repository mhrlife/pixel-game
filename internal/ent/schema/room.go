package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/teris-io/shortid"
	"time"
)

type Room struct {
	ent.Schema
}

func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("short_id").DefaultFunc(shortid.MustGenerate),
		field.Enum("status").Values("active", "inactive").Default("active"),
		field.Int8("max_participants").Default(5),

		field.Enum("permission_video").
			Values("all_participants", "only_admins").Default("all_participants"),
		field.Enum("permission_audio").
			Values("all_participants", "only_admins").Default("all_participants"),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
	}
}

func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		// admins
		edge.To("admins", User.Type),
	}
}

func (Room) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("short_id"),
	}
}
