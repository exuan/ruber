// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/category"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetParentID sets the "parent_id" field.
func (cc *CategoryCreate) SetParentID(i int64) *CategoryCreate {
	cc.mutation.SetParentID(i)
	return cc
}

// SetSort sets the "sort" field.
func (cc *CategoryCreate) SetSort(i int64) *CategoryCreate {
	cc.mutation.SetSort(i)
	return cc
}

// SetStatus sets the "status" field.
func (cc *CategoryCreate) SetStatus(i int) *CategoryCreate {
	cc.mutation.SetStatus(i)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableStatus(i *int) *CategoryCreate {
	if i != nil {
		cc.SetStatus(*i)
	}
	return cc
}

// SetCreator sets the "creator" field.
func (cc *CategoryCreate) SetCreator(s string) *CategoryCreate {
	cc.mutation.SetCreator(s)
	return cc
}

// SetModifier sets the "modifier" field.
func (cc *CategoryCreate) SetModifier(s string) *CategoryCreate {
	cc.mutation.SetModifier(s)
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CategoryCreate) SetUpdateTime(i int64) *CategoryCreate {
	cc.mutation.SetUpdateTime(i)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableUpdateTime(i *int64) *CategoryCreate {
	if i != nil {
		cc.SetUpdateTime(*i)
	}
	return cc
}

// SetCreateTime sets the "create_time" field.
func (cc *CategoryCreate) SetCreateTime(i int64) *CategoryCreate {
	cc.mutation.SetCreateTime(i)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableCreateTime(i *int64) *CategoryCreate {
	if i != nil {
		cc.SetCreateTime(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CategoryCreate) SetID(i int64) *CategoryCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CategoryCreate) defaults() {
	if _, ok := cc.mutation.Status(); !ok {
		v := category.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := category.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := category.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := cc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "parent_id"`)}
	}
	if _, ok := cc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "sort"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if _, ok := cc.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required field "creator"`)}
	}
	if _, ok := cc.mutation.Modifier(); !ok {
		return &ValidationError{Name: "modifier", err: errors.New(`ent: missing required field "modifier"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if v, ok := cc.mutation.UpdateTime(); ok {
		if err := category.UpdateTimeValidator(v); err != nil {
			return &ValidationError{Name: "update_time", err: fmt.Errorf(`ent: validator failed for field "update_time": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if v, ok := cc.mutation.CreateTime(); ok {
		if err := category.CreateTimeValidator(v); err != nil {
			return &ValidationError{Name: "create_time", err: fmt.Errorf(`ent: validator failed for field "create_time": %w`, err)}
		}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: category.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: category.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.ParentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldParentID,
		})
		_node.ParentID = value
	}
	if value, ok := cc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := cc.mutation.Creator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldCreator,
		})
		_node.Creator = value
	}
	if value, ok := cc.mutation.Modifier(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldModifier,
		})
		_node.Modifier = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Category.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CategoryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CategoryCreate) OnConflict(opts ...sql.ConflictOption) *CategoryUpsertOne {
	cc.conflict = opts
	return &CategoryUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CategoryCreate) OnConflictColumns(columns ...string) *CategoryUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CategoryUpsertOne{
		create: cc,
	}
}

type (
	// CategoryUpsertOne is the builder for "upsert"-ing
	//  one Category node.
	CategoryUpsertOne struct {
		create *CategoryCreate
	}

	// CategoryUpsert is the "OnConflict" setter.
	CategoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *CategoryUpsert) SetName(v string) *CategoryUpsert {
	u.Set(category.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateName() *CategoryUpsert {
	u.SetExcluded(category.FieldName)
	return u
}

// SetParentID sets the "parent_id" field.
func (u *CategoryUpsert) SetParentID(v int64) *CategoryUpsert {
	u.Set(category.FieldParentID, v)
	return u
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateParentID() *CategoryUpsert {
	u.SetExcluded(category.FieldParentID)
	return u
}

// SetSort sets the "sort" field.
func (u *CategoryUpsert) SetSort(v int64) *CategoryUpsert {
	u.Set(category.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateSort() *CategoryUpsert {
	u.SetExcluded(category.FieldSort)
	return u
}

// SetStatus sets the "status" field.
func (u *CategoryUpsert) SetStatus(v int) *CategoryUpsert {
	u.Set(category.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateStatus() *CategoryUpsert {
	u.SetExcluded(category.FieldStatus)
	return u
}

// SetCreator sets the "creator" field.
func (u *CategoryUpsert) SetCreator(v string) *CategoryUpsert {
	u.Set(category.FieldCreator, v)
	return u
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateCreator() *CategoryUpsert {
	u.SetExcluded(category.FieldCreator)
	return u
}

// SetModifier sets the "modifier" field.
func (u *CategoryUpsert) SetModifier(v string) *CategoryUpsert {
	u.Set(category.FieldModifier, v)
	return u
}

// UpdateModifier sets the "modifier" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateModifier() *CategoryUpsert {
	u.SetExcluded(category.FieldModifier)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CategoryUpsert) SetUpdateTime(v int64) *CategoryUpsert {
	u.Set(category.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateUpdateTime() *CategoryUpsert {
	u.SetExcluded(category.FieldUpdateTime)
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CategoryUpsert) SetCreateTime(v int64) *CategoryUpsert {
	u.Set(category.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateCreateTime() *CategoryUpsert {
	u.SetExcluded(category.FieldCreateTime)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(category.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CategoryUpsertOne) UpdateNewValues() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(category.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Category.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CategoryUpsertOne) Ignore() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CategoryUpsertOne) DoNothing() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CategoryCreate.OnConflict
// documentation for more info.
func (u *CategoryUpsertOne) Update(set func(*CategoryUpsert)) *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CategoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CategoryUpsertOne) SetName(v string) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateName() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *CategoryUpsertOne) SetParentID(v int64) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateParentID() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateParentID()
	})
}

// SetSort sets the "sort" field.
func (u *CategoryUpsertOne) SetSort(v int64) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateSort() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateSort()
	})
}

// SetStatus sets the "status" field.
func (u *CategoryUpsertOne) SetStatus(v int) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateStatus() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateStatus()
	})
}

// SetCreator sets the "creator" field.
func (u *CategoryUpsertOne) SetCreator(v string) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetCreator(v)
	})
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateCreator() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateCreator()
	})
}

// SetModifier sets the "modifier" field.
func (u *CategoryUpsertOne) SetModifier(v string) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetModifier(v)
	})
}

// UpdateModifier sets the "modifier" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateModifier() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateModifier()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CategoryUpsertOne) SetUpdateTime(v int64) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateUpdateTime() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *CategoryUpsertOne) SetCreateTime(v int64) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateCreateTime() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateCreateTime()
	})
}

// Exec executes the query.
func (u *CategoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CategoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CategoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CategoryUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CategoryUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	builders []*CategoryCreate
	conflict []sql.ConflictOption
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Category.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CategoryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CategoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CategoryUpsertBulk {
	ccb.conflict = opts
	return &CategoryUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CategoryCreateBulk) OnConflictColumns(columns ...string) *CategoryUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CategoryUpsertBulk{
		create: ccb,
	}
}

// CategoryUpsertBulk is the builder for "upsert"-ing
// a bulk of Category nodes.
type CategoryUpsertBulk struct {
	create *CategoryCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(category.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CategoryUpsertBulk) UpdateNewValues() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(category.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CategoryUpsertBulk) Ignore() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CategoryUpsertBulk) DoNothing() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CategoryCreateBulk.OnConflict
// documentation for more info.
func (u *CategoryUpsertBulk) Update(set func(*CategoryUpsert)) *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CategoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CategoryUpsertBulk) SetName(v string) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateName() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *CategoryUpsertBulk) SetParentID(v int64) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateParentID() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateParentID()
	})
}

// SetSort sets the "sort" field.
func (u *CategoryUpsertBulk) SetSort(v int64) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateSort() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateSort()
	})
}

// SetStatus sets the "status" field.
func (u *CategoryUpsertBulk) SetStatus(v int) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateStatus() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateStatus()
	})
}

// SetCreator sets the "creator" field.
func (u *CategoryUpsertBulk) SetCreator(v string) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetCreator(v)
	})
}

// UpdateCreator sets the "creator" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateCreator() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateCreator()
	})
}

// SetModifier sets the "modifier" field.
func (u *CategoryUpsertBulk) SetModifier(v string) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetModifier(v)
	})
}

// UpdateModifier sets the "modifier" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateModifier() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateModifier()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CategoryUpsertBulk) SetUpdateTime(v int64) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateUpdateTime() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *CategoryUpsertBulk) SetCreateTime(v int64) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateCreateTime() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateCreateTime()
	})
}

// Exec executes the query.
func (u *CategoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CategoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CategoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CategoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
