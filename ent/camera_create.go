// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-ent-project/ent/camera"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CameraCreate is the builder for creating a Camera entity.
type CameraCreate struct {
	config
	mutation *CameraMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CameraCreate) SetName(s string) *CameraCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetModel sets the "model" field.
func (cc *CameraCreate) SetModel(s string) *CameraCreate {
	cc.mutation.SetModel(s)
	return cc
}

// SetImei sets the "imei" field.
func (cc *CameraCreate) SetImei(s string) *CameraCreate {
	cc.mutation.SetImei(s)
	return cc
}

// SetLocation sets the "location" field.
func (cc *CameraCreate) SetLocation(m map[string]interface{}) *CameraCreate {
	cc.mutation.SetLocation(m)
	return cc
}

// SetActive sets the "active" field.
func (cc *CameraCreate) SetActive(b bool) *CameraCreate {
	cc.mutation.SetActive(b)
	return cc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (cc *CameraCreate) SetNillableActive(b *bool) *CameraCreate {
	if b != nil {
		cc.SetActive(*b)
	}
	return cc
}

// Mutation returns the CameraMutation object of the builder.
func (cc *CameraCreate) Mutation() *CameraMutation {
	return cc.mutation
}

// Save creates the Camera in the database.
func (cc *CameraCreate) Save(ctx context.Context) (*Camera, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CameraCreate) SaveX(ctx context.Context) *Camera {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CameraCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CameraCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CameraCreate) defaults() {
	if _, ok := cc.mutation.Active(); !ok {
		v := camera.DefaultActive
		cc.mutation.SetActive(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CameraCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Camera.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := camera.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Camera.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "Camera.model"`)}
	}
	if v, ok := cc.mutation.Model(); ok {
		if err := camera.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`ent: validator failed for field "Camera.model": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Imei(); !ok {
		return &ValidationError{Name: "imei", err: errors.New(`ent: missing required field "Camera.imei"`)}
	}
	if v, ok := cc.mutation.Imei(); ok {
		if err := camera.ImeiValidator(v); err != nil {
			return &ValidationError{Name: "imei", err: fmt.Errorf(`ent: validator failed for field "Camera.imei": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "Camera.location"`)}
	}
	if _, ok := cc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Camera.active"`)}
	}
	return nil
}

func (cc *CameraCreate) sqlSave(ctx context.Context) (*Camera, error) {
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

func (cc *CameraCreate) createSpec() (*Camera, *sqlgraph.CreateSpec) {
	var (
		_node = &Camera{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(camera.Table, sqlgraph.NewFieldSpec(camera.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(camera.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Model(); ok {
		_spec.SetField(camera.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if value, ok := cc.mutation.Imei(); ok {
		_spec.SetField(camera.FieldImei, field.TypeString, value)
		_node.Imei = value
	}
	if value, ok := cc.mutation.Location(); ok {
		_spec.SetField(camera.FieldLocation, field.TypeJSON, value)
		_node.Location = value
	}
	if value, ok := cc.mutation.Active(); ok {
		_spec.SetField(camera.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	return _node, _spec
}

// CameraCreateBulk is the builder for creating many Camera entities in bulk.
type CameraCreateBulk struct {
	config
	err      error
	builders []*CameraCreate
}

// Save creates the Camera entities in the database.
func (ccb *CameraCreateBulk) Save(ctx context.Context) ([]*Camera, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Camera, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CameraMutation)
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
func (ccb *CameraCreateBulk) SaveX(ctx context.Context) []*Camera {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CameraCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CameraCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
