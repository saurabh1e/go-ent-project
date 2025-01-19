// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"go-ent-project/ent/schema"
	"go-ent-project/internal/ent/camera"
	"go-ent-project/internal/ent/car"
	"go-ent-project/internal/ent/permission"
	"go-ent-project/internal/ent/policestation"
	"go-ent-project/internal/ent/role"
	"go-ent-project/internal/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cameraMixin := schema.Camera{}.Mixin()
	cameraMixinFields0 := cameraMixin[0].Fields()
	_ = cameraMixinFields0
	cameraFields := schema.Camera{}.Fields()
	_ = cameraFields
	// cameraDescCreatedAt is the schema descriptor for created_at field.
	cameraDescCreatedAt := cameraMixinFields0[1].Descriptor()
	// camera.DefaultCreatedAt holds the default value on creation for the created_at field.
	camera.DefaultCreatedAt = cameraDescCreatedAt.Default.(func() time.Time)
	// cameraDescUpdatedAt is the schema descriptor for updated_at field.
	cameraDescUpdatedAt := cameraMixinFields0[2].Descriptor()
	// camera.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	camera.DefaultUpdatedAt = cameraDescUpdatedAt.Default.(func() time.Time)
	// camera.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	camera.UpdateDefaultUpdatedAt = cameraDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cameraDescName is the schema descriptor for name field.
	cameraDescName := cameraFields[0].Descriptor()
	// camera.NameValidator is a validator for the "name" field. It is called by the builders before save.
	camera.NameValidator = cameraDescName.Validators[0].(func(string) error)
	// cameraDescModel is the schema descriptor for model field.
	cameraDescModel := cameraFields[1].Descriptor()
	// camera.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	camera.ModelValidator = cameraDescModel.Validators[0].(func(string) error)
	// cameraDescImei is the schema descriptor for imei field.
	cameraDescImei := cameraFields[2].Descriptor()
	// camera.ImeiValidator is a validator for the "imei" field. It is called by the builders before save.
	camera.ImeiValidator = cameraDescImei.Validators[0].(func(string) error)
	// cameraDescActive is the schema descriptor for active field.
	cameraDescActive := cameraFields[4].Descriptor()
	// camera.DefaultActive holds the default value on creation for the active field.
	camera.DefaultActive = cameraDescActive.Default.(bool)
	carMixin := schema.Car{}.Mixin()
	carMixinFields0 := carMixin[0].Fields()
	_ = carMixinFields0
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescCreatedAt is the schema descriptor for created_at field.
	carDescCreatedAt := carMixinFields0[1].Descriptor()
	// car.DefaultCreatedAt holds the default value on creation for the created_at field.
	car.DefaultCreatedAt = carDescCreatedAt.Default.(func() time.Time)
	// carDescUpdatedAt is the schema descriptor for updated_at field.
	carDescUpdatedAt := carMixinFields0[2].Descriptor()
	// car.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	car.DefaultUpdatedAt = carDescUpdatedAt.Default.(func() time.Time)
	// car.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	car.UpdateDefaultUpdatedAt = carDescUpdatedAt.UpdateDefault.(func() time.Time)
	// carDescMake is the schema descriptor for make field.
	carDescMake := carFields[0].Descriptor()
	// car.MakeValidator is a validator for the "make" field. It is called by the builders before save.
	car.MakeValidator = carDescMake.Validators[0].(func(string) error)
	// carDescModel is the schema descriptor for model field.
	carDescModel := carFields[1].Descriptor()
	// car.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	car.ModelValidator = carDescModel.Validators[0].(func(string) error)
	// carDescYear is the schema descriptor for year field.
	carDescYear := carFields[2].Descriptor()
	// car.YearValidator is a validator for the "year" field. It is called by the builders before save.
	car.YearValidator = carDescYear.Validators[0].(func(int) error)
	// carDescRegistration is the schema descriptor for registration field.
	carDescRegistration := carFields[3].Descriptor()
	// car.RegistrationValidator is a validator for the "registration" field. It is called by the builders before save.
	car.RegistrationValidator = carDescRegistration.Validators[0].(func(string) error)
	// carDescColor is the schema descriptor for color field.
	carDescColor := carFields[4].Descriptor()
	// car.ColorValidator is a validator for the "color" field. It is called by the builders before save.
	car.ColorValidator = carDescColor.Validators[0].(func(string) error)
	permissionMixin := schema.Permission{}.Mixin()
	permissionMixinFields0 := permissionMixin[0].Fields()
	_ = permissionMixinFields0
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionMixinFields0[1].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(func() time.Time)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionMixinFields0[2].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(func() time.Time)
	// permission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	permission.UpdateDefaultUpdatedAt = permissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// permissionDescName is the schema descriptor for name field.
	permissionDescName := permissionFields[0].Descriptor()
	// permission.NameValidator is a validator for the "name" field. It is called by the builders before save.
	permission.NameValidator = permissionDescName.Validators[0].(func(string) error)
	// permissionDescCanRead is the schema descriptor for can_read field.
	permissionDescCanRead := permissionFields[1].Descriptor()
	// permission.DefaultCanRead holds the default value on creation for the can_read field.
	permission.DefaultCanRead = permissionDescCanRead.Default.(bool)
	// permissionDescCanCreate is the schema descriptor for can_create field.
	permissionDescCanCreate := permissionFields[2].Descriptor()
	// permission.DefaultCanCreate holds the default value on creation for the can_create field.
	permission.DefaultCanCreate = permissionDescCanCreate.Default.(bool)
	// permissionDescCanUpdate is the schema descriptor for can_update field.
	permissionDescCanUpdate := permissionFields[3].Descriptor()
	// permission.DefaultCanUpdate holds the default value on creation for the can_update field.
	permission.DefaultCanUpdate = permissionDescCanUpdate.Default.(bool)
	// permissionDescCanDelete is the schema descriptor for can_delete field.
	permissionDescCanDelete := permissionFields[4].Descriptor()
	// permission.DefaultCanDelete holds the default value on creation for the can_delete field.
	permission.DefaultCanDelete = permissionDescCanDelete.Default.(bool)
	policestationMixin := schema.PoliceStation{}.Mixin()
	policestationMixinFields0 := policestationMixin[0].Fields()
	_ = policestationMixinFields0
	policestationFields := schema.PoliceStation{}.Fields()
	_ = policestationFields
	// policestationDescCreatedAt is the schema descriptor for created_at field.
	policestationDescCreatedAt := policestationMixinFields0[1].Descriptor()
	// policestation.DefaultCreatedAt holds the default value on creation for the created_at field.
	policestation.DefaultCreatedAt = policestationDescCreatedAt.Default.(func() time.Time)
	// policestationDescUpdatedAt is the schema descriptor for updated_at field.
	policestationDescUpdatedAt := policestationMixinFields0[2].Descriptor()
	// policestation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	policestation.DefaultUpdatedAt = policestationDescUpdatedAt.Default.(func() time.Time)
	// policestation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	policestation.UpdateDefaultUpdatedAt = policestationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// policestationDescName is the schema descriptor for name field.
	policestationDescName := policestationFields[0].Descriptor()
	// policestation.NameValidator is a validator for the "name" field. It is called by the builders before save.
	policestation.NameValidator = policestationDescName.Validators[0].(func(string) error)
	// policestationDescCode is the schema descriptor for code field.
	policestationDescCode := policestationFields[2].Descriptor()
	// policestation.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	policestation.CodeValidator = policestationDescCode.Validators[0].(func(string) error)
	// policestationDescIdentifier is the schema descriptor for identifier field.
	policestationDescIdentifier := policestationFields[3].Descriptor()
	// policestation.IdentifierValidator is a validator for the "identifier" field. It is called by the builders before save.
	policestation.IdentifierValidator = policestationDescIdentifier.Validators[0].(func(string) error)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[1].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[2].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[4].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
}

const (
	Version = "v0.14.1"                                         // Version of ent codegen.
	Sum     = "h1:fUERL506Pqr92EPHJqr8EYxbPioflJo6PudkrEA8a/s=" // Sum of ent codegen.
)
