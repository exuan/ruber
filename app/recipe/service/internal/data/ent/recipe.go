// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/recipe"
)

// Recipe is the model entity for the Recipe schema.
type Recipe struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID int64 `json:"category_id,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight int64 `json:"weight,omitempty"`
	// Sort holds the value of the "sort" field.
	Sort int64 `json:"sort,omitempty"`
	// Favorites holds the value of the "favorites" field.
	Favorites int64 `json:"favorites,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Creator holds the value of the "creator" field.
	Creator string `json:"creator,omitempty"`
	// Modifier holds the value of the "modifier" field.
	Modifier string `json:"modifier,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime int64 `json:"update_time,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime int64 `json:"create_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Recipe) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case recipe.FieldID, recipe.FieldCategoryID, recipe.FieldWeight, recipe.FieldSort, recipe.FieldFavorites, recipe.FieldStatus, recipe.FieldUpdateTime, recipe.FieldCreateTime:
			values[i] = new(sql.NullInt64)
		case recipe.FieldTitle, recipe.FieldImage, recipe.FieldContent, recipe.FieldCreator, recipe.FieldModifier:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Recipe", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Recipe fields.
func (r *Recipe) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recipe.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case recipe.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				r.Title = value.String
			}
		case recipe.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				r.Image = value.String
			}
		case recipe.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				r.CategoryID = value.Int64
			}
		case recipe.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				r.Weight = value.Int64
			}
		case recipe.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				r.Sort = value.Int64
			}
		case recipe.FieldFavorites:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field favorites", values[i])
			} else if value.Valid {
				r.Favorites = value.Int64
			}
		case recipe.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				r.Content = value.String
			}
		case recipe.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = int(value.Int64)
			}
		case recipe.FieldCreator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creator", values[i])
			} else if value.Valid {
				r.Creator = value.String
			}
		case recipe.FieldModifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field modifier", values[i])
			} else if value.Valid {
				r.Modifier = value.String
			}
		case recipe.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = value.Int64
			}
		case recipe.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Recipe.
// Note that you need to call Recipe.Unwrap() before calling this method if this Recipe
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Recipe) Update() *RecipeUpdateOne {
	return (&RecipeClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Recipe entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Recipe) Unwrap() *Recipe {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Recipe is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Recipe) String() string {
	var builder strings.Builder
	builder.WriteString("Recipe(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", title=")
	builder.WriteString(r.Title)
	builder.WriteString(", image=")
	builder.WriteString(r.Image)
	builder.WriteString(", category_id=")
	builder.WriteString(fmt.Sprintf("%v", r.CategoryID))
	builder.WriteString(", weight=")
	builder.WriteString(fmt.Sprintf("%v", r.Weight))
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", r.Sort))
	builder.WriteString(", favorites=")
	builder.WriteString(fmt.Sprintf("%v", r.Favorites))
	builder.WriteString(", content=")
	builder.WriteString(r.Content)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", creator=")
	builder.WriteString(r.Creator)
	builder.WriteString(", modifier=")
	builder.WriteString(r.Modifier)
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdateTime))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", r.CreateTime))
	builder.WriteByte(')')
	return builder.String()
}

// Recipes is a parsable slice of Recipe.
type Recipes []*Recipe

func (r Recipes) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}