// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/panupong/app/ent/predicate"
	"github.com/panupong/app/ent/title"
)

// TitleDelete is the builder for deleting a Title entity.
type TitleDelete struct {
	config
	hooks      []Hook
	mutation   *TitleMutation
	predicates []predicate.Title
}

// Where adds a new predicate to the delete builder.
func (td *TitleDelete) Where(ps ...predicate.Title) *TitleDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TitleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TitleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TitleDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TitleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: title.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: title.FieldID,
			},
		},
	}
	if ps := td.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TitleDeleteOne is the builder for deleting a single Title entity.
type TitleDeleteOne struct {
	td *TitleDelete
}

// Exec executes the deletion query.
func (tdo *TitleDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{title.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TitleDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}