package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/exuan/ruber/pkg/util/format"
)

type Favorite struct {
	ent.Schema
}

func (Favorite) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("user_id"),
		field.Int64("recipe_id"),
		field.Int64("create_time").
			DefaultFunc(format.UnixTimestamp).
			Immutable(),
	}
}

func (Favorite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rb_recipe_favorite"},
	}
}

func (Favorite) Edges() []ent.Edge {
	return nil
}

func (Favorite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "recipe_id").
			Unique(),
	}
}
