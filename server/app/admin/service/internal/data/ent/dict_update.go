// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent/dict"
	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent/predicate"
)

// DictUpdate is the builder for updating Dict entities.
type DictUpdate struct {
	config
	hooks     []Hook
	mutation  *DictMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DictUpdate builder.
func (du *DictUpdate) Where(ps ...predicate.Dict) *DictUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreateBy sets the "create_by" field.
func (du *DictUpdate) SetCreateBy(u uint32) *DictUpdate {
	du.mutation.ResetCreateBy()
	du.mutation.SetCreateBy(u)
	return du
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (du *DictUpdate) SetNillableCreateBy(u *uint32) *DictUpdate {
	if u != nil {
		du.SetCreateBy(*u)
	}
	return du
}

// AddCreateBy adds u to the "create_by" field.
func (du *DictUpdate) AddCreateBy(u int32) *DictUpdate {
	du.mutation.AddCreateBy(u)
	return du
}

// ClearCreateBy clears the value of the "create_by" field.
func (du *DictUpdate) ClearCreateBy() *DictUpdate {
	du.mutation.ClearCreateBy()
	return du
}

// SetName sets the "name" field.
func (du *DictUpdate) SetName(s string) *DictUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DictUpdate) SetNillableName(s *string) *DictUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// ClearName clears the value of the "name" field.
func (du *DictUpdate) ClearName() *DictUpdate {
	du.mutation.ClearName()
	return du
}

// SetDescription sets the "description" field.
func (du *DictUpdate) SetDescription(s string) *DictUpdate {
	du.mutation.SetDescription(s)
	return du
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (du *DictUpdate) SetNillableDescription(s *string) *DictUpdate {
	if s != nil {
		du.SetDescription(*s)
	}
	return du
}

// ClearDescription clears the value of the "description" field.
func (du *DictUpdate) ClearDescription() *DictUpdate {
	du.mutation.ClearDescription()
	return du
}

// Mutation returns the DictMutation object of the builder.
func (du *DictUpdate) Mutation() *DictMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DictUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DictUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DictUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DictUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (du *DictUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DictUpdate {
	du.modifiers = append(du.modifiers, modifiers...)
	return du
}

func (du *DictUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(dict.Table, dict.Columns, sqlgraph.NewFieldSpec(dict.FieldID, field.TypeUint32))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if du.mutation.CreateTimeCleared() {
		_spec.ClearField(dict.FieldCreateTime, field.TypeTime)
	}
	if du.mutation.UpdateTimeCleared() {
		_spec.ClearField(dict.FieldUpdateTime, field.TypeTime)
	}
	if du.mutation.DeleteTimeCleared() {
		_spec.ClearField(dict.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := du.mutation.CreateBy(); ok {
		_spec.SetField(dict.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := du.mutation.AddedCreateBy(); ok {
		_spec.AddField(dict.FieldCreateBy, field.TypeUint32, value)
	}
	if du.mutation.CreateByCleared() {
		_spec.ClearField(dict.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(dict.FieldName, field.TypeString, value)
	}
	if du.mutation.NameCleared() {
		_spec.ClearField(dict.FieldName, field.TypeString)
	}
	if value, ok := du.mutation.Description(); ok {
		_spec.SetField(dict.FieldDescription, field.TypeString, value)
	}
	if du.mutation.DescriptionCleared() {
		_spec.ClearField(dict.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(du.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dict.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DictUpdateOne is the builder for updating a single Dict entity.
type DictUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DictMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreateBy sets the "create_by" field.
func (duo *DictUpdateOne) SetCreateBy(u uint32) *DictUpdateOne {
	duo.mutation.ResetCreateBy()
	duo.mutation.SetCreateBy(u)
	return duo
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (duo *DictUpdateOne) SetNillableCreateBy(u *uint32) *DictUpdateOne {
	if u != nil {
		duo.SetCreateBy(*u)
	}
	return duo
}

// AddCreateBy adds u to the "create_by" field.
func (duo *DictUpdateOne) AddCreateBy(u int32) *DictUpdateOne {
	duo.mutation.AddCreateBy(u)
	return duo
}

// ClearCreateBy clears the value of the "create_by" field.
func (duo *DictUpdateOne) ClearCreateBy() *DictUpdateOne {
	duo.mutation.ClearCreateBy()
	return duo
}

// SetName sets the "name" field.
func (duo *DictUpdateOne) SetName(s string) *DictUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DictUpdateOne) SetNillableName(s *string) *DictUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// ClearName clears the value of the "name" field.
func (duo *DictUpdateOne) ClearName() *DictUpdateOne {
	duo.mutation.ClearName()
	return duo
}

// SetDescription sets the "description" field.
func (duo *DictUpdateOne) SetDescription(s string) *DictUpdateOne {
	duo.mutation.SetDescription(s)
	return duo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (duo *DictUpdateOne) SetNillableDescription(s *string) *DictUpdateOne {
	if s != nil {
		duo.SetDescription(*s)
	}
	return duo
}

// ClearDescription clears the value of the "description" field.
func (duo *DictUpdateOne) ClearDescription() *DictUpdateOne {
	duo.mutation.ClearDescription()
	return duo
}

// Mutation returns the DictMutation object of the builder.
func (duo *DictUpdateOne) Mutation() *DictMutation {
	return duo.mutation
}

// Where appends a list predicates to the DictUpdate builder.
func (duo *DictUpdateOne) Where(ps ...predicate.Dict) *DictUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DictUpdateOne) Select(field string, fields ...string) *DictUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dict entity.
func (duo *DictUpdateOne) Save(ctx context.Context) (*Dict, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DictUpdateOne) SaveX(ctx context.Context) *Dict {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DictUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DictUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (duo *DictUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DictUpdateOne {
	duo.modifiers = append(duo.modifiers, modifiers...)
	return duo
}

func (duo *DictUpdateOne) sqlSave(ctx context.Context) (_node *Dict, err error) {
	_spec := sqlgraph.NewUpdateSpec(dict.Table, dict.Columns, sqlgraph.NewFieldSpec(dict.FieldID, field.TypeUint32))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dict.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dict.FieldID)
		for _, f := range fields {
			if !dict.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dict.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if duo.mutation.CreateTimeCleared() {
		_spec.ClearField(dict.FieldCreateTime, field.TypeTime)
	}
	if duo.mutation.UpdateTimeCleared() {
		_spec.ClearField(dict.FieldUpdateTime, field.TypeTime)
	}
	if duo.mutation.DeleteTimeCleared() {
		_spec.ClearField(dict.FieldDeleteTime, field.TypeTime)
	}
	if value, ok := duo.mutation.CreateBy(); ok {
		_spec.SetField(dict.FieldCreateBy, field.TypeUint32, value)
	}
	if value, ok := duo.mutation.AddedCreateBy(); ok {
		_spec.AddField(dict.FieldCreateBy, field.TypeUint32, value)
	}
	if duo.mutation.CreateByCleared() {
		_spec.ClearField(dict.FieldCreateBy, field.TypeUint32)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(dict.FieldName, field.TypeString, value)
	}
	if duo.mutation.NameCleared() {
		_spec.ClearField(dict.FieldName, field.TypeString)
	}
	if value, ok := duo.mutation.Description(); ok {
		_spec.SetField(dict.FieldDescription, field.TypeString, value)
	}
	if duo.mutation.DescriptionCleared() {
		_spec.ClearField(dict.FieldDescription, field.TypeString)
	}
	_spec.AddModifiers(duo.modifiers...)
	_node = &Dict{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dict.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
