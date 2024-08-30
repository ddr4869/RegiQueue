// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ddr4869/RegiQueue/ent/course"
)

// CourseCreate is the builder for creating a Course entity.
type CourseCreate struct {
	config
	mutation *CourseMutation
	hooks    []Hook
}

// SetCourseName sets the "course_name" field.
func (cc *CourseCreate) SetCourseName(i int) *CourseCreate {
	cc.mutation.SetCourseName(i)
	return cc
}

// SetUserID sets the "user_id" field.
func (cc *CourseCreate) SetUserID(s string) *CourseCreate {
	cc.mutation.SetUserID(s)
	return cc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cc *CourseCreate) SetNillableUserID(s *string) *CourseCreate {
	if s != nil {
		cc.SetUserID(*s)
	}
	return cc
}

// Mutation returns the CourseMutation object of the builder.
func (cc *CourseCreate) Mutation() *CourseMutation {
	return cc.mutation
}

// Save creates the Course in the database.
func (cc *CourseCreate) Save(ctx context.Context) (*Course, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CourseCreate) SaveX(ctx context.Context) *Course {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CourseCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CourseCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CourseCreate) defaults() {
	if _, ok := cc.mutation.UserID(); !ok {
		v := course.DefaultUserID
		cc.mutation.SetUserID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CourseCreate) check() error {
	if _, ok := cc.mutation.CourseName(); !ok {
		return &ValidationError{Name: "course_name", err: errors.New(`ent: missing required field "Course.course_name"`)}
	}
	if v, ok := cc.mutation.CourseName(); ok {
		if err := course.CourseNameValidator(v); err != nil {
			return &ValidationError{Name: "course_name", err: fmt.Errorf(`ent: validator failed for field "Course.course_name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Course.user_id"`)}
	}
	return nil
}

func (cc *CourseCreate) sqlSave(ctx context.Context) (*Course, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CourseCreate) createSpec() (*Course, *sqlgraph.CreateSpec) {
	var (
		_node = &Course{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(course.Table, sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CourseName(); ok {
		_spec.SetField(course.FieldCourseName, field.TypeInt, value)
		_node.CourseName = value
	}
	if value, ok := cc.mutation.UserID(); ok {
		_spec.SetField(course.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	return _node, _spec
}

// CourseCreateBulk is the builder for creating many Course entities in bulk.
type CourseCreateBulk struct {
	config
	err      error
	builders []*CourseCreate
}

// Save creates the Course entities in the database.
func (ccb *CourseCreateBulk) Save(ctx context.Context) ([]*Course, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Course, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CourseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
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
func (ccb *CourseCreateBulk) SaveX(ctx context.Context) []*Course {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CourseCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CourseCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
