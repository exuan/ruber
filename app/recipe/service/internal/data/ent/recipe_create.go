// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/recipe"
)

// RecipeCreate is the builder for creating a Recipe entity.
type RecipeCreate struct {
	config
	mutation *RecipeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (rc *RecipeCreate) SetTitle(s string) *RecipeCreate {
	rc.mutation.SetTitle(s)
	return rc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (rc *RecipeCreate) SetNillableTitle(s *string) *RecipeCreate {
	if s != nil {
		rc.SetTitle(*s)
	}
	return rc
}

// SetImage sets the "image" field.
func (rc *RecipeCreate) SetImage(s string) *RecipeCreate {
	rc.mutation.SetImage(s)
	return rc
}

// SetCategoryID sets the "category_id" field.
func (rc *RecipeCreate) SetCategoryID(i int64) *RecipeCreate {
	rc.mutation.SetCategoryID(i)
	return rc
}

// SetWeight sets the "weight" field.
func (rc *RecipeCreate) SetWeight(i int64) *RecipeCreate {
	rc.mutation.SetWeight(i)
	return rc
}

// SetSort sets the "sort" field.
func (rc *RecipeCreate) SetSort(i int64) *RecipeCreate {
	rc.mutation.SetSort(i)
	return rc
}

// SetFavorites sets the "favorites" field.
func (rc *RecipeCreate) SetFavorites(i int64) *RecipeCreate {
	rc.mutation.SetFavorites(i)
	return rc
}

// SetNillableFavorites sets the "favorites" field if the given value is not nil.
func (rc *RecipeCreate) SetNillableFavorites(i *int64) *RecipeCreate {
	if i != nil {
		rc.SetFavorites(*i)
	}
	return rc
}

// SetContent sets the "content" field.
func (rc *RecipeCreate) SetContent(s string) *RecipeCreate {
	rc.mutation.SetContent(s)
	return rc
}

// SetStatus sets the "status" field.
func (rc *RecipeCreate) SetStatus(i int) *RecipeCreate {
	rc.mutation.SetStatus(i)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RecipeCreate) SetNillableStatus(i *int) *RecipeCreate {
	if i != nil {
		rc.SetStatus(*i)
	}
	return rc
}

// SetCreator sets the "creator" field.
func (rc *RecipeCreate) SetCreator(s string) *RecipeCreate {
	rc.mutation.SetCreator(s)
	return rc
}

// SetModifier sets the "modifier" field.
func (rc *RecipeCreate) SetModifier(s string) *RecipeCreate {
	rc.mutation.SetModifier(s)
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *RecipeCreate) SetUpdateTime(i int64) *RecipeCreate {
	rc.mutation.SetUpdateTime(i)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *RecipeCreate) SetNillableUpdateTime(i *int64) *RecipeCreate {
	if i != nil {
		rc.SetUpdateTime(*i)
	}
	return rc
}

// SetCreateTime sets the "create_time" field.
func (rc *RecipeCreate) SetCreateTime(i int64) *RecipeCreate {
	rc.mutation.SetCreateTime(i)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *RecipeCreate) SetNillableCreateTime(i *int64) *RecipeCreate {
	if i != nil {
		rc.SetCreateTime(*i)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RecipeCreate) SetID(i int64) *RecipeCreate {
	rc.mutation.SetID(i)
	return rc
}

// Mutation returns the RecipeMutation object of the builder.
func (rc *RecipeCreate) Mutation() *RecipeMutation {
	return rc.mutation
}

// Save creates the Recipe in the database.
func (rc *RecipeCreate) Save(ctx context.Context) (*Recipe, error) {
	var (
		err  error
		node *Recipe
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecipeCreate) SaveX(ctx context.Context) *Recipe {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecipeCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecipeCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RecipeCreate) defaults() {
	if _, ok := rc.mutation.Title(); !ok {
		v := recipe.DefaultTitle
		rc.mutation.SetTitle(v)
	}
	if _, ok := rc.mutation.Status(); !ok {
		v := recipe.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		v := recipe.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.CreateTime(); !ok {
		v := recipe.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecipeCreate) check() error {
	if _, ok := rc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if _, ok := rc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "image"`)}
	}
	if _, ok := rc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "category_id"`)}
	}
	if _, ok := rc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "weight"`)}
	}
	if _, ok := rc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "sort"`)}
	}
	if _, ok := rc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "content"`)}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if _, ok := rc.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required field "creator"`)}
	}
	if _, ok := rc.mutation.Modifier(); !ok {
		return &ValidationError{Name: "modifier", err: errors.New(`ent: missing required field "modifier"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if v, ok := rc.mutation.UpdateTime(); ok {
		if err := recipe.UpdateTimeValidator(v); err != nil {
			return &ValidationError{Name: "update_time", err: fmt.Errorf(`ent: validator failed for field "update_time": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if v, ok := rc.mutation.CreateTime(); ok {
		if err := recipe.CreateTimeValidator(v); err != nil {
			return &ValidationError{Name: "create_time", err: fmt.Errorf(`ent: validator failed for field "create_time": %w`, err)}
		}
	}
	return nil
}

func (rc *RecipeCreate) sqlSave(ctx context.Context) (*Recipe, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (rc *RecipeCreate) createSpec() (*Recipe, *sqlgraph.CreateSpec) {
	var (
		_node = &Recipe{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recipe.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: recipe.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := rc.mutation.Image(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldImage,
		})
		_node.Image = value
	}
	if value, ok := rc.mutation.CategoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recipe.FieldCategoryID,
		})
		_node.CategoryID = value
	}
	if value, ok := rc.mutation.Weight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recipe.FieldWeight,
		})
		_node.Weight = value
	}
	if value, ok := rc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recipe.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := rc.mutation.Favorites(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recipe.FieldFavorites,
		})
		_node.Favorites = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recipe.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := rc.mutation.Creator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldCreator,
		})
		_node.Creator = value
	}
	if value, ok := rc.mutation.Modifier(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldModifier,
		})
		_node.Modifier = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recipe.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recipe.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Recipe.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecipeUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
//
func (rc *RecipeCreate) OnConflict(opts ...sql.ConflictOption) *RecipeUpsertOne {
	rc.conflict = opts
	return &RecipeUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rc *RecipeCreate) OnConflictColumns(columns ...string) *RecipeUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RecipeUpsertOne{
		create: rc,
	}
}

type (
	// RecipeUpsertOne is the builder for "upsert"-ing
	//  one Recipe node.
	RecipeUpsertOne struct {
		create *RecipeCreate
	}

	// RecipeUpsert is the "OnConflict" setter.
	RecipeUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *RecipeUpsert) SetTitle(v string) *RecipeUpsert {
	u.Set(recipe.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateTitle() *RecipeUpsert {
	u.SetExcluded(recipe.FieldTitle)
	return u
}

// SetImage sets the "image" field.
func (u *RecipeUpsert) SetImage(v string) *RecipeUpsert {
	u.Set(recipe.FieldImage, v)
	return u
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateImage() *RecipeUpsert {
	u.SetExcluded(recipe.FieldImage)
	return u
}

// SetCategoryID sets the "category_id" field.
func (u *RecipeUpsert) SetCategoryID(v int64) *RecipeUpsert {
	u.Set(recipe.FieldCategoryID, v)
	return u
}

// UpdateCategoryID sets the "category_id" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateCategoryID() *RecipeUpsert {
	u.SetExcluded(recipe.FieldCategoryID)
	return u
}

// SetWeight sets the "weight" field.
func (u *RecipeUpsert) SetWeight(v int64) *RecipeUpsert {
	u.Set(recipe.FieldWeight, v)
	return u
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateWeight() *RecipeUpsert {
	u.SetExcluded(recipe.FieldWeight)
	return u
}

// SetSort sets the "sort" field.
func (u *RecipeUpsert) SetSort(v int64) *RecipeUpsert {
	u.Set(recipe.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateSort() *RecipeUpsert {
	u.SetExcluded(recipe.FieldSort)
	return u
}

// SetFavorites sets the "favorites" field.
func (u *RecipeUpsert) SetFavorites(v int64) *RecipeUpsert {
	u.Set(recipe.FieldFavorites, v)
	return u
}

// UpdateFavorites sets the "favorites" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateFavorites() *RecipeUpsert {
	u.SetExcluded(recipe.FieldFavorites)
	return u
}

// ClearFavorites clears the value of the "favorites" field.
func (u *RecipeUpsert) ClearFavorites() *RecipeUpsert {
	u.SetNull(recipe.FieldFavorites)
	return u
}

// SetContent sets the "content" field.
func (u *RecipeUpsert) SetContent(v string) *RecipeUpsert {
	u.Set(recipe.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateContent() *RecipeUpsert {
	u.SetExcluded(recipe.FieldContent)
	return u
}

// SetStatus sets the "status" field.
func (u *RecipeUpsert) SetStatus(v int) *RecipeUpsert {
	u.Set(recipe.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateStatus() *RecipeUpsert {
	u.SetExcluded(recipe.FieldStatus)
	return u
}

// SetCreator sets the "creator" field.
func (u *RecipeUpsert) SetCreator(v string) *RecipeUpsert {
	u.Set(recipe.FieldCreator, v)
	return u
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateCreator() *RecipeUpsert {
	u.SetExcluded(recipe.FieldCreator)
	return u
}

// SetModifier sets the "modifier" field.
func (u *RecipeUpsert) SetModifier(v string) *RecipeUpsert {
	u.Set(recipe.FieldModifier, v)
	return u
}

// UpdateModifier sets the "modifier" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateModifier() *RecipeUpsert {
	u.SetExcluded(recipe.FieldModifier)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *RecipeUpsert) SetUpdateTime(v int64) *RecipeUpsert {
	u.Set(recipe.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateUpdateTime() *RecipeUpsert {
	u.SetExcluded(recipe.FieldUpdateTime)
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *RecipeUpsert) SetCreateTime(v int64) *RecipeUpsert {
	u.Set(recipe.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *RecipeUpsert) UpdateCreateTime() *RecipeUpsert {
	u.SetExcluded(recipe.FieldCreateTime)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(recipe.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *RecipeUpsertOne) UpdateNewValues() *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(recipe.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Recipe.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *RecipeUpsertOne) Ignore() *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecipeUpsertOne) DoNothing() *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecipeCreate.OnConflict
// documentation for more info.
func (u *RecipeUpsertOne) Update(set func(*RecipeUpsert)) *RecipeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecipeUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *RecipeUpsertOne) SetTitle(v string) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateTitle() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateTitle()
	})
}

// SetImage sets the "image" field.
func (u *RecipeUpsertOne) SetImage(v string) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateImage() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateImage()
	})
}

// SetCategoryID sets the "category_id" field.
func (u *RecipeUpsertOne) SetCategoryID(v int64) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetCategoryID(v)
	})
}

// UpdateCategoryID sets the "category_id" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateCategoryID() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateCategoryID()
	})
}

// SetWeight sets the "weight" field.
func (u *RecipeUpsertOne) SetWeight(v int64) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateWeight() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateWeight()
	})
}

// SetSort sets the "sort" field.
func (u *RecipeUpsertOne) SetSort(v int64) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateSort() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateSort()
	})
}

// SetFavorites sets the "favorites" field.
func (u *RecipeUpsertOne) SetFavorites(v int64) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetFavorites(v)
	})
}

// UpdateFavorites sets the "favorites" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateFavorites() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateFavorites()
	})
}

// ClearFavorites clears the value of the "favorites" field.
func (u *RecipeUpsertOne) ClearFavorites() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.ClearFavorites()
	})
}

// SetContent sets the "content" field.
func (u *RecipeUpsertOne) SetContent(v string) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateContent() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateContent()
	})
}

// SetStatus sets the "status" field.
func (u *RecipeUpsertOne) SetStatus(v int) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateStatus() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateStatus()
	})
}

// SetCreator sets the "creator" field.
func (u *RecipeUpsertOne) SetCreator(v string) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetCreator(v)
	})
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateCreator() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateCreator()
	})
}

// SetModifier sets the "modifier" field.
func (u *RecipeUpsertOne) SetModifier(v string) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetModifier(v)
	})
}

// UpdateModifier sets the "modifier" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateModifier() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateModifier()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *RecipeUpsertOne) SetUpdateTime(v int64) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateUpdateTime() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *RecipeUpsertOne) SetCreateTime(v int64) *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *RecipeUpsertOne) UpdateCreateTime() *RecipeUpsertOne {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateCreateTime()
	})
}

// Exec executes the query.
func (u *RecipeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecipeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecipeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RecipeUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RecipeUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RecipeCreateBulk is the builder for creating many Recipe entities in bulk.
type RecipeCreateBulk struct {
	config
	builders []*RecipeCreate
	conflict []sql.ConflictOption
}

// Save creates the Recipe entities in the database.
func (rcb *RecipeCreateBulk) Save(ctx context.Context) ([]*Recipe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Recipe, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecipeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecipeCreateBulk) SaveX(ctx context.Context) []*Recipe {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecipeCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecipeCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Recipe.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecipeUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
//
func (rcb *RecipeCreateBulk) OnConflict(opts ...sql.ConflictOption) *RecipeUpsertBulk {
	rcb.conflict = opts
	return &RecipeUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rcb *RecipeCreateBulk) OnConflictColumns(columns ...string) *RecipeUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RecipeUpsertBulk{
		create: rcb,
	}
}

// RecipeUpsertBulk is the builder for "upsert"-ing
// a bulk of Recipe nodes.
type RecipeUpsertBulk struct {
	create *RecipeCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(recipe.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *RecipeUpsertBulk) UpdateNewValues() *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(recipe.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Recipe.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *RecipeUpsertBulk) Ignore() *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecipeUpsertBulk) DoNothing() *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecipeCreateBulk.OnConflict
// documentation for more info.
func (u *RecipeUpsertBulk) Update(set func(*RecipeUpsert)) *RecipeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecipeUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *RecipeUpsertBulk) SetTitle(v string) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateTitle() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateTitle()
	})
}

// SetImage sets the "image" field.
func (u *RecipeUpsertBulk) SetImage(v string) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateImage() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateImage()
	})
}

// SetCategoryID sets the "category_id" field.
func (u *RecipeUpsertBulk) SetCategoryID(v int64) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetCategoryID(v)
	})
}

// UpdateCategoryID sets the "category_id" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateCategoryID() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateCategoryID()
	})
}

// SetWeight sets the "weight" field.
func (u *RecipeUpsertBulk) SetWeight(v int64) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateWeight() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateWeight()
	})
}

// SetSort sets the "sort" field.
func (u *RecipeUpsertBulk) SetSort(v int64) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateSort() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateSort()
	})
}

// SetFavorites sets the "favorites" field.
func (u *RecipeUpsertBulk) SetFavorites(v int64) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetFavorites(v)
	})
}

// UpdateFavorites sets the "favorites" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateFavorites() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateFavorites()
	})
}

// ClearFavorites clears the value of the "favorites" field.
func (u *RecipeUpsertBulk) ClearFavorites() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.ClearFavorites()
	})
}

// SetContent sets the "content" field.
func (u *RecipeUpsertBulk) SetContent(v string) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateContent() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateContent()
	})
}

// SetStatus sets the "status" field.
func (u *RecipeUpsertBulk) SetStatus(v int) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateStatus() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateStatus()
	})
}

// SetCreator sets the "creator" field.
func (u *RecipeUpsertBulk) SetCreator(v string) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetCreator(v)
	})
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateCreator() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateCreator()
	})
}

// SetModifier sets the "modifier" field.
func (u *RecipeUpsertBulk) SetModifier(v string) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetModifier(v)
	})
}

// UpdateModifier sets the "modifier" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateModifier() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateModifier()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *RecipeUpsertBulk) SetUpdateTime(v int64) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateUpdateTime() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *RecipeUpsertBulk) SetCreateTime(v int64) *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *RecipeUpsertBulk) UpdateCreateTime() *RecipeUpsertBulk {
	return u.Update(func(s *RecipeUpsert) {
		s.UpdateCreateTime()
	})
}

// Exec executes the query.
func (u *RecipeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RecipeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecipeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecipeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
