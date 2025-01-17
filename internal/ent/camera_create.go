// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-ent-project/internal/ent/camera"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CameraCreate is the builder for creating a Camera entity.
type CameraCreate struct {
	config
	mutation *CameraMutation
	hooks    []Hook
	conflict []sql.ConflictOption
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
	_spec.OnConflict = cc.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Camera.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CameraUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (cc *CameraCreate) OnConflict(opts ...sql.ConflictOption) *CameraUpsertOne {
	cc.conflict = opts
	return &CameraUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Camera.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CameraCreate) OnConflictColumns(columns ...string) *CameraUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CameraUpsertOne{
		create: cc,
	}
}

type (
	// CameraUpsertOne is the builder for "upsert"-ing
	//  one Camera node.
	CameraUpsertOne struct {
		create *CameraCreate
	}

	// CameraUpsert is the "OnConflict" setter.
	CameraUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *CameraUpsert) SetName(v string) *CameraUpsert {
	u.Set(camera.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CameraUpsert) UpdateName() *CameraUpsert {
	u.SetExcluded(camera.FieldName)
	return u
}

// SetModel sets the "model" field.
func (u *CameraUpsert) SetModel(v string) *CameraUpsert {
	u.Set(camera.FieldModel, v)
	return u
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *CameraUpsert) UpdateModel() *CameraUpsert {
	u.SetExcluded(camera.FieldModel)
	return u
}

// SetImei sets the "imei" field.
func (u *CameraUpsert) SetImei(v string) *CameraUpsert {
	u.Set(camera.FieldImei, v)
	return u
}

// UpdateImei sets the "imei" field to the value that was provided on create.
func (u *CameraUpsert) UpdateImei() *CameraUpsert {
	u.SetExcluded(camera.FieldImei)
	return u
}

// SetLocation sets the "location" field.
func (u *CameraUpsert) SetLocation(v map[string]interface{}) *CameraUpsert {
	u.Set(camera.FieldLocation, v)
	return u
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *CameraUpsert) UpdateLocation() *CameraUpsert {
	u.SetExcluded(camera.FieldLocation)
	return u
}

// SetActive sets the "active" field.
func (u *CameraUpsert) SetActive(v bool) *CameraUpsert {
	u.Set(camera.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *CameraUpsert) UpdateActive() *CameraUpsert {
	u.SetExcluded(camera.FieldActive)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Camera.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CameraUpsertOne) UpdateNewValues() *CameraUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Camera.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CameraUpsertOne) Ignore() *CameraUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CameraUpsertOne) DoNothing() *CameraUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CameraCreate.OnConflict
// documentation for more info.
func (u *CameraUpsertOne) Update(set func(*CameraUpsert)) *CameraUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CameraUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CameraUpsertOne) SetName(v string) *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CameraUpsertOne) UpdateName() *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateName()
	})
}

// SetModel sets the "model" field.
func (u *CameraUpsertOne) SetModel(v string) *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *CameraUpsertOne) UpdateModel() *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateModel()
	})
}

// SetImei sets the "imei" field.
func (u *CameraUpsertOne) SetImei(v string) *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.SetImei(v)
	})
}

// UpdateImei sets the "imei" field to the value that was provided on create.
func (u *CameraUpsertOne) UpdateImei() *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateImei()
	})
}

// SetLocation sets the "location" field.
func (u *CameraUpsertOne) SetLocation(v map[string]interface{}) *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *CameraUpsertOne) UpdateLocation() *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateLocation()
	})
}

// SetActive sets the "active" field.
func (u *CameraUpsertOne) SetActive(v bool) *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *CameraUpsertOne) UpdateActive() *CameraUpsertOne {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateActive()
	})
}

// Exec executes the query.
func (u *CameraUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CameraCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CameraUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CameraUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CameraUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CameraCreateBulk is the builder for creating many Camera entities in bulk.
type CameraCreateBulk struct {
	config
	err      error
	builders []*CameraCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = ccb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Camera.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CameraUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ccb *CameraCreateBulk) OnConflict(opts ...sql.ConflictOption) *CameraUpsertBulk {
	ccb.conflict = opts
	return &CameraUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Camera.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CameraCreateBulk) OnConflictColumns(columns ...string) *CameraUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CameraUpsertBulk{
		create: ccb,
	}
}

// CameraUpsertBulk is the builder for "upsert"-ing
// a bulk of Camera nodes.
type CameraUpsertBulk struct {
	create *CameraCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Camera.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CameraUpsertBulk) UpdateNewValues() *CameraUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Camera.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CameraUpsertBulk) Ignore() *CameraUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CameraUpsertBulk) DoNothing() *CameraUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CameraCreateBulk.OnConflict
// documentation for more info.
func (u *CameraUpsertBulk) Update(set func(*CameraUpsert)) *CameraUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CameraUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *CameraUpsertBulk) SetName(v string) *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CameraUpsertBulk) UpdateName() *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateName()
	})
}

// SetModel sets the "model" field.
func (u *CameraUpsertBulk) SetModel(v string) *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *CameraUpsertBulk) UpdateModel() *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateModel()
	})
}

// SetImei sets the "imei" field.
func (u *CameraUpsertBulk) SetImei(v string) *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.SetImei(v)
	})
}

// UpdateImei sets the "imei" field to the value that was provided on create.
func (u *CameraUpsertBulk) UpdateImei() *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateImei()
	})
}

// SetLocation sets the "location" field.
func (u *CameraUpsertBulk) SetLocation(v map[string]interface{}) *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.SetLocation(v)
	})
}

// UpdateLocation sets the "location" field to the value that was provided on create.
func (u *CameraUpsertBulk) UpdateLocation() *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateLocation()
	})
}

// SetActive sets the "active" field.
func (u *CameraUpsertBulk) SetActive(v bool) *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *CameraUpsertBulk) UpdateActive() *CameraUpsertBulk {
	return u.Update(func(s *CameraUpsert) {
		s.UpdateActive()
	})
}

// Exec executes the query.
func (u *CameraUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CameraCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CameraCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CameraUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
