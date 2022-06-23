// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/favorite"
)

// Favorite is the model entity for the Favorite schema.
type Favorite struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// RecipeID holds the value of the "recipe_id" field.
	RecipeID int64 `json:"recipe_id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime int64 `json:"create_time,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Favorite) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case favorite.FieldID, favorite.FieldRecipeID, favorite.FieldCreateTime:
			values[i] = new(sql.NullInt64)
		case favorite.FieldUserID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Favorite", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Favorite fields.
func (f *Favorite) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case favorite.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int64(value.Int64)
		case favorite.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				f.UserID = value.String
			}
		case favorite.FieldRecipeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field recipe_id", values[i])
			} else if value.Valid {
				f.RecipeID = value.Int64
			}
		case favorite.FieldCreateTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				f.CreateTime = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Favorite.
// Note that you need to call Favorite.Unwrap() before calling this method if this Favorite
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Favorite) Update() *FavoriteUpdateOne {
	return (&FavoriteClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Favorite entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Favorite) Unwrap() *Favorite {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Favorite is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Favorite) String() string {
	var builder strings.Builder
	builder.WriteString("Favorite(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(f.UserID)
	builder.WriteString(", recipe_id=")
	builder.WriteString(fmt.Sprintf("%v", f.RecipeID))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", f.CreateTime))
	builder.WriteByte(')')
	return builder.String()
}

// Favorites is a parsable slice of Favorite.
type Favorites []*Favorite

func (f Favorites) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}