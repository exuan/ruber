package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/exuan/ruber/pkg/util/format"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("username"),
		field.String("phone"),
		field.String("nickname"),
		field.String("avatar"),
		field.String("password"),
		field.String("salt"),
		field.String("owner"),
		field.String("app").
			Default(""),
		field.JSON("properties", map[string]interface{}{}).
			Optional(),
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

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rb_user"},
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
