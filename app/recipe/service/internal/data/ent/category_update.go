// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/category"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/predicate"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CategoryUpdate) SetParentID(i int64) *CategoryUpdate {
	cu.mutation.ResetParentID()
	cu.mutation.SetParentID(i)
	return cu
}

// AddParentID adds i to the "parent_id" field.
func (cu *CategoryUpdate) AddParentID(i int64) *CategoryUpdate {
	cu.mutation.AddParentID(i)
	return cu
}

// SetSort sets the "sort" field.
func (cu *CategoryUpdate) SetSort(i int64) *CategoryUpdate {
	cu.mutation.ResetSort()
	cu.mutation.SetSort(i)
	return cu
}

// AddSort adds i to the "sort" field.
func (cu *CategoryUpdate) AddSort(i int64) *CategoryUpdate {
	cu.mutation.AddSort(i)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CategoryUpdate) SetStatus(i int) *CategoryUpdate {
	cu.mutation.ResetStatus()
	cu.mutation.SetStatus(i)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableStatus(i *int) *CategoryUpdate {
	if i != nil {
		cu.SetStatus(*i)
	}
	return cu
}

// AddStatus adds i to the "status" field.
func (cu *CategoryUpdate) AddStatus(i int) *CategoryUpdate {
	cu.mutation.AddStatus(i)
	return cu
}

// SetModifier sets the "modifier" field.
func (cu *CategoryUpdate) SetModifier(s string) *CategoryUpdate {
	cu.mutation.SetModifier(s)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *CategoryUpdate) SetUpdateTime(i int64) *CategoryUpdate {
	cu.mutation.ResetUpdateTime()
	cu.mutation.SetUpdateTime(i)
	return cu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableUpdateTime(i *int64) *CategoryUpdate {
	if i != nil {
		cu.SetUpdateTime(*i)
	}
	return cu
}

// AddUpdateTime adds i to the "update_time" field.
func (cu *CategoryUpdate) AddUpdateTime(i int64) *CategoryUpdate {
	cu.mutation.AddUpdateTime(i)
	return cu
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CategoryUpdate) check() error {
	if v, ok := cu.mutation.UpdateTime(); ok {
		if err := category.UpdateTimeValidator(v); err != nil {
			return &ValidationError{Name: "update_time", err: fmt.Errorf("ent: validator failed for field \"update_time\": %w", err)}
		}
	}
	return nil
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: category.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if value, ok := cu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cu.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldSort,
		})
	}
	if value, ok := cu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldSort,
		})
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldStatus,
		})
	}
	if value, ok := cu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldStatus,
		})
	}
	if value, ok := cu.mutation.Modifier(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldModifier,
		})
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldUpdateTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CategoryUpdateOne) SetParentID(i int64) *CategoryUpdateOne {
	cuo.mutation.ResetParentID()
	cuo.mutation.SetParentID(i)
	return cuo
}

// AddParentID adds i to the "parent_id" field.
func (cuo *CategoryUpdateOne) AddParentID(i int64) *CategoryUpdateOne {
	cuo.mutation.AddParentID(i)
	return cuo
}

// SetSort sets the "sort" field.
func (cuo *CategoryUpdateOne) SetSort(i int64) *CategoryUpdateOne {
	cuo.mutation.ResetSort()
	cuo.mutation.SetSort(i)
	return cuo
}

// AddSort adds i to the "sort" field.
func (cuo *CategoryUpdateOne) AddSort(i int64) *CategoryUpdateOne {
	cuo.mutation.AddSort(i)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CategoryUpdateOne) SetStatus(i int) *CategoryUpdateOne {
	cuo.mutation.ResetStatus()
	cuo.mutation.SetStatus(i)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableStatus(i *int) *CategoryUpdateOne {
	if i != nil {
		cuo.SetStatus(*i)
	}
	return cuo
}

// AddStatus adds i to the "status" field.
func (cuo *CategoryUpdateOne) AddStatus(i int) *CategoryUpdateOne {
	cuo.mutation.AddStatus(i)
	return cuo
}

// SetModifier sets the "modifier" field.
func (cuo *CategoryUpdateOne) SetModifier(s string) *CategoryUpdateOne {
	cuo.mutation.SetModifier(s)
	return cuo
}

// SetUpdateTime sets the "update_time" field.
func (cuo *CategoryUpdateOne) SetUpdateTime(i int64) *CategoryUpdateOne {
	cuo.mutation.ResetUpdateTime()
	cuo.mutation.SetUpdateTime(i)
	return cuo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableUpdateTime(i *int64) *CategoryUpdateOne {
	if i != nil {
		cuo.SetUpdateTime(*i)
	}
	return cuo
}

// AddUpdateTime adds i to the "update_time" field.
func (cuo *CategoryUpdateOne) AddUpdateTime(i int64) *CategoryUpdateOne {
	cuo.mutation.AddUpdateTime(i)
	return cuo
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CategoryUpdateOne) check() error {
	if v, ok := cuo.mutation.UpdateTime(); ok {
		if err := category.UpdateTimeValidator(v); err != nil {
			return &ValidationError{Name: "update_time", err: fmt.Errorf("ent: validator failed for field \"update_time\": %w", err)}
		}
	}
	return nil
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: category.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Category.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if value, ok := cuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cuo.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldSort,
		})
	}
	if value, ok := cuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldSort,
		})
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldStatus,
		})
	}
	if value, ok := cuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: category.FieldStatus,
		})
	}
	if value, ok := cuo.mutation.Modifier(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldModifier,
		})
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.AddedUpdateTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldUpdateTime,
		})
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
