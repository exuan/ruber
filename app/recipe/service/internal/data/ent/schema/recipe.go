package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/exuan/ruber/pkg/util/format"
)

type Recipe struct {
	ent.Schema
}

func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title").
			Default(""),
		field.String("image"),
		field.Int64("category_id"),
		field.Int64("weight"),
		field.Int64("sort"),
		field.Int64("favorites").
			Optional(),
		field.Text("content"),
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

func (Recipe) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rb_recipe"},
	}
}

func (Recipe) Edges() []ent.Edge {
	return nil
}
