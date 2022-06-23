package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/exuan/ruber/pkg/util/format"
)

type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.Int64("parent_id"),
		field.Int64("sort"),
		field.Int("status").
			Default(0),
		field.String("creator").
			Immutable(),
		field.String("modifier"),
		field.Int64("update_time").
			NonNegative().
			DefaultFunc(format.UnixTimestamp),
		field.Int64("create_time").
			NonNegative().
			DefaultFunc(format.UnixTimestamp).
			Immutable(),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rb_recipe_category"},
	}
}

func (Category) Edges() []ent.Edge {
	return nil
}
