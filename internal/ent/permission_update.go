// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-ent-project/internal/ent/permission"
	"go-ent-project/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableName(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetCanRead sets the "can_read" field.
func (pu *PermissionUpdate) SetCanRead(b bool) *PermissionUpdate {
	pu.mutation.SetCanRead(b)
	return pu
}

// SetNillableCanRead sets the "can_read" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCanRead(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetCanRead(*b)
	}
	return pu
}

// SetCanCreate sets the "can_create" field.
func (pu *PermissionUpdate) SetCanCreate(b bool) *PermissionUpdate {
	pu.mutation.SetCanCreate(b)
	return pu
}

// SetNillableCanCreate sets the "can_create" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCanCreate(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetCanCreate(*b)
	}
	return pu
}

// SetCanUpdate sets the "can_update" field.
func (pu *PermissionUpdate) SetCanUpdate(b bool) *PermissionUpdate {
	pu.mutation.SetCanUpdate(b)
	return pu
}

// SetNillableCanUpdate sets the "can_update" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCanUpdate(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetCanUpdate(*b)
	}
	return pu
}

// SetCanDelete sets the "can_delete" field.
func (pu *PermissionUpdate) SetCanDelete(b bool) *PermissionUpdate {
	pu.mutation.SetCanDelete(b)
	return pu
}

// SetNillableCanDelete sets the "can_delete" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCanDelete(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetCanDelete(*b)
	}
	return pu
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Permission.name": %w`, err)}
		}
	}
	return nil
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.CanRead(); ok {
		_spec.SetField(permission.FieldCanRead, field.TypeBool, value)
	}
	if value, ok := pu.mutation.CanCreate(); ok {
		_spec.SetField(permission.FieldCanCreate, field.TypeBool, value)
	}
	if value, ok := pu.mutation.CanUpdate(); ok {
		_spec.SetField(permission.FieldCanUpdate, field.TypeBool, value)
	}
	if value, ok := pu.mutation.CanDelete(); ok {
		_spec.SetField(permission.FieldCanDelete, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableName(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetCanRead sets the "can_read" field.
func (puo *PermissionUpdateOne) SetCanRead(b bool) *PermissionUpdateOne {
	puo.mutation.SetCanRead(b)
	return puo
}

// SetNillableCanRead sets the "can_read" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCanRead(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetCanRead(*b)
	}
	return puo
}

// SetCanCreate sets the "can_create" field.
func (puo *PermissionUpdateOne) SetCanCreate(b bool) *PermissionUpdateOne {
	puo.mutation.SetCanCreate(b)
	return puo
}

// SetNillableCanCreate sets the "can_create" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCanCreate(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetCanCreate(*b)
	}
	return puo
}

// SetCanUpdate sets the "can_update" field.
func (puo *PermissionUpdateOne) SetCanUpdate(b bool) *PermissionUpdateOne {
	puo.mutation.SetCanUpdate(b)
	return puo
}

// SetNillableCanUpdate sets the "can_update" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCanUpdate(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetCanUpdate(*b)
	}
	return puo
}

// SetCanDelete sets the "can_delete" field.
func (puo *PermissionUpdateOne) SetCanDelete(b bool) *PermissionUpdateOne {
	puo.mutation.SetCanDelete(b)
	return puo
}

// SetNillableCanDelete sets the "can_delete" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCanDelete(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetCanDelete(*b)
	}
	return puo
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Permission.name": %w`, err)}
		}
	}
	return nil
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.CanRead(); ok {
		_spec.SetField(permission.FieldCanRead, field.TypeBool, value)
	}
	if value, ok := puo.mutation.CanCreate(); ok {
		_spec.SetField(permission.FieldCanCreate, field.TypeBool, value)
	}
	if value, ok := puo.mutation.CanUpdate(); ok {
		_spec.SetField(permission.FieldCanUpdate, field.TypeBool, value)
	}
	if value, ok := puo.mutation.CanDelete(); ok {
		_spec.SetField(permission.FieldCanDelete, field.TypeBool, value)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
