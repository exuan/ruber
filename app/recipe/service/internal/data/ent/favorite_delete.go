// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/favorite"
	"github.com/exuan/ruber/app/recipe/service/internal/data/ent/predicate"
)

// FavoriteDelete is the builder for deleting a Favorite entity.
type FavoriteDelete struct {
	config
	hooks    []Hook
	mutation *FavoriteMutation
}

// Where appends a list predicates to the FavoriteDelete builder.
func (fd *FavoriteDelete) Where(ps ...predicate.Favorite) *FavoriteDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FavoriteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fd.hooks) == 0 {
		affected, err = fd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavoriteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fd.mutation = mutation
			affected, err = fd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fd.hooks) - 1; i >= 0; i-- {
			if fd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FavoriteDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FavoriteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: favorite.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: favorite.FieldID,
			},
		},
	}
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
}

// FavoriteDeleteOne is the builder for deleting a single Favorite entity.
type FavoriteDeleteOne struct {
	fd *FavoriteDelete
}

// Exec executes the deletion query.
func (fdo *FavoriteDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{favorite.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FavoriteDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}
