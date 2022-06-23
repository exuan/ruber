// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RbRecipeCategoryColumns holds the columns for the "rb_recipe_category" table.
	RbRecipeCategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "parent_id", Type: field.TypeInt64},
		{Name: "sort", Type: field.TypeInt64},
		{Name: "status", Type: field.TypeInt, Default: 0},
		{Name: "creator", Type: field.TypeString},
		{Name: "modifier", Type: field.TypeString},
		{Name: "update_time", Type: field.TypeInt64},
		{Name: "create_time", Type: field.TypeInt64},
	}
	// RbRecipeCategoryTable holds the schema information for the "rb_recipe_category" table.
	RbRecipeCategoryTable = &schema.Table{
		Name:       "rb_recipe_category",
		Columns:    RbRecipeCategoryColumns,
		PrimaryKey: []*schema.Column{RbRecipeCategoryColumns[0]},
	}
	// RbRecipeFavoriteColumns holds the columns for the "rb_recipe_favorite" table.
	RbRecipeFavoriteColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "recipe_id", Type: field.TypeInt64},
		{Name: "create_time", Type: field.TypeInt64},
	}
	// RbRecipeFavoriteTable holds the schema information for the "rb_recipe_favorite" table.
	RbRecipeFavoriteTable = &schema.Table{
		Name:       "rb_recipe_favorite",
		Columns:    RbRecipeFavoriteColumns,
		PrimaryKey: []*schema.Column{RbRecipeFavoriteColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "favorite_user_id_recipe_id",
				Unique:  true,
				Columns: []*schema.Column{RbRecipeFavoriteColumns[1], RbRecipeFavoriteColumns[2]},
			},
		},
	}
	// RbRecipeColumns holds the columns for the "rb_recipe" table.
	RbRecipeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString, Default: ""},
		{Name: "image", Type: field.TypeString},
		{Name: "category_id", Type: field.TypeInt64},
		{Name: "weight", Type: field.TypeInt64},
		{Name: "sort", Type: field.TypeInt64},
		{Name: "favorites", Type: field.TypeInt64, Nullable: true},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeInt, Default: 0},
		{Name: "creator", Type: field.TypeString},
		{Name: "modifier", Type: field.TypeString},
		{Name: "update_time", Type: field.TypeInt64},
		{Name: "create_time", Type: field.TypeInt64},
	}
	// RbRecipeTable holds the schema information for the "rb_recipe" table.
	RbRecipeTable = &schema.Table{
		Name:       "rb_recipe",
		Columns:    RbRecipeColumns,
		PrimaryKey: []*schema.Column{RbRecipeColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RbRecipeCategoryTable,
		RbRecipeFavoriteTable,
		RbRecipeTable,
	}
)

func init() {
	RbRecipeCategoryTable.Annotation = &entsql.Annotation{
		Table: "rb_recipe_category",
	}
	RbRecipeFavoriteTable.Annotation = &entsql.Annotation{
		Table: "rb_recipe_favorite",
	}
	RbRecipeTable.Annotation = &entsql.Annotation{
		Table: "rb_recipe",
	}
}
