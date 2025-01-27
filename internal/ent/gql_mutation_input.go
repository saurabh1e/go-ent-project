// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateCameraInput represents a mutation input for creating cameras.
type CreateCameraInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	Model     string
	Imei      string
	Location  map[string]interface{}
	Active    *bool
}

// Mutate applies the CreateCameraInput on the CameraMutation builder.
func (i *CreateCameraInput) Mutate(m *CameraMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	m.SetModel(i.Model)
	m.SetImei(i.Imei)
	if v := i.Location; v != nil {
		m.SetLocation(v)
	}
	if v := i.Active; v != nil {
		m.SetActive(*v)
	}
}

// SetInput applies the change-set in the CreateCameraInput on the CameraCreate builder.
func (c *CameraCreate) SetInput(i CreateCameraInput) *CameraCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCameraInput represents a mutation input for updating cameras.
type UpdateCameraInput struct {
	UpdatedAt *time.Time
	Name      *string
	Model     *string
	Imei      *string
	Location  map[string]interface{}
	Active    *bool
}

// Mutate applies the UpdateCameraInput on the CameraMutation builder.
func (i *UpdateCameraInput) Mutate(m *CameraMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Imei; v != nil {
		m.SetImei(*v)
	}
	if v := i.Location; v != nil {
		m.SetLocation(v)
	}
	if v := i.Active; v != nil {
		m.SetActive(*v)
	}
}

// SetInput applies the change-set in the UpdateCameraInput on the CameraUpdate builder.
func (c *CameraUpdate) SetInput(i UpdateCameraInput) *CameraUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCameraInput on the CameraUpdateOne builder.
func (c *CameraUpdateOne) SetInput(i UpdateCameraInput) *CameraUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateCarInput represents a mutation input for creating cars.
type CreateCarInput struct {
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Make         string
	Model        string
	Year         int
	Registration string
	Color        string
}

// Mutate applies the CreateCarInput on the CarMutation builder.
func (i *CreateCarInput) Mutate(m *CarMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetMake(i.Make)
	m.SetModel(i.Model)
	m.SetYear(i.Year)
	m.SetRegistration(i.Registration)
	m.SetColor(i.Color)
}

// SetInput applies the change-set in the CreateCarInput on the CarCreate builder.
func (c *CarCreate) SetInput(i CreateCarInput) *CarCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateCarInput represents a mutation input for updating cars.
type UpdateCarInput struct {
	UpdatedAt    *time.Time
	Make         *string
	Model        *string
	Year         *int
	Registration *string
	Color        *string
}

// Mutate applies the UpdateCarInput on the CarMutation builder.
func (i *UpdateCarInput) Mutate(m *CarMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Make; v != nil {
		m.SetMake(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Year; v != nil {
		m.SetYear(*v)
	}
	if v := i.Registration; v != nil {
		m.SetRegistration(*v)
	}
	if v := i.Color; v != nil {
		m.SetColor(*v)
	}
}

// SetInput applies the change-set in the UpdateCarInput on the CarUpdate builder.
func (c *CarUpdate) SetInput(i UpdateCarInput) *CarUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateCarInput on the CarUpdateOne builder.
func (c *CarUpdateOne) SetInput(i UpdateCarInput) *CarUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePermissionInput represents a mutation input for creating permissions.
type CreatePermissionInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	CanRead   *bool
	CanCreate *bool
	CanUpdate *bool
	CanDelete *bool
}

// Mutate applies the CreatePermissionInput on the PermissionMutation builder.
func (i *CreatePermissionInput) Mutate(m *PermissionMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	if v := i.CanRead; v != nil {
		m.SetCanRead(*v)
	}
	if v := i.CanCreate; v != nil {
		m.SetCanCreate(*v)
	}
	if v := i.CanUpdate; v != nil {
		m.SetCanUpdate(*v)
	}
	if v := i.CanDelete; v != nil {
		m.SetCanDelete(*v)
	}
}

// SetInput applies the change-set in the CreatePermissionInput on the PermissionCreate builder.
func (c *PermissionCreate) SetInput(i CreatePermissionInput) *PermissionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePermissionInput represents a mutation input for updating permissions.
type UpdatePermissionInput struct {
	UpdatedAt *time.Time
	Name      *string
	CanRead   *bool
	CanCreate *bool
	CanUpdate *bool
	CanDelete *bool
}

// Mutate applies the UpdatePermissionInput on the PermissionMutation builder.
func (i *UpdatePermissionInput) Mutate(m *PermissionMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.CanRead; v != nil {
		m.SetCanRead(*v)
	}
	if v := i.CanCreate; v != nil {
		m.SetCanCreate(*v)
	}
	if v := i.CanUpdate; v != nil {
		m.SetCanUpdate(*v)
	}
	if v := i.CanDelete; v != nil {
		m.SetCanDelete(*v)
	}
}

// SetInput applies the change-set in the UpdatePermissionInput on the PermissionUpdate builder.
func (c *PermissionUpdate) SetInput(i UpdatePermissionInput) *PermissionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePermissionInput on the PermissionUpdateOne builder.
func (c *PermissionUpdateOne) SetInput(i UpdatePermissionInput) *PermissionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePoliceStationInput represents a mutation input for creating policestations.
type CreatePoliceStationInput struct {
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	Name            string
	Location        map[string]interface{}
	Code            string
	Identifier      string
	UserIDs         []int
	ParentStationID *int
	ChildStationIDs []int
}

// Mutate applies the CreatePoliceStationInput on the PoliceStationMutation builder.
func (i *CreatePoliceStationInput) Mutate(m *PoliceStationMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	if v := i.Location; v != nil {
		m.SetLocation(v)
	}
	m.SetCode(i.Code)
	m.SetIdentifier(i.Identifier)
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.ParentStationID; v != nil {
		m.SetParentStationID(*v)
	}
	if v := i.ChildStationIDs; len(v) > 0 {
		m.AddChildStationIDs(v...)
	}
}

// SetInput applies the change-set in the CreatePoliceStationInput on the PoliceStationCreate builder.
func (c *PoliceStationCreate) SetInput(i CreatePoliceStationInput) *PoliceStationCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePoliceStationInput represents a mutation input for updating policestations.
type UpdatePoliceStationInput struct {
	UpdatedAt             *time.Time
	Name                  *string
	ClearLocation         bool
	Location              map[string]interface{}
	Code                  *string
	Identifier            *string
	ClearUsers            bool
	AddUserIDs            []int
	RemoveUserIDs         []int
	ClearParentStation    bool
	ParentStationID       *int
	ClearChildStations    bool
	AddChildStationIDs    []int
	RemoveChildStationIDs []int
}

// Mutate applies the UpdatePoliceStationInput on the PoliceStationMutation builder.
func (i *UpdatePoliceStationInput) Mutate(m *PoliceStationMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearLocation {
		m.ClearLocation()
	}
	if v := i.Location; v != nil {
		m.SetLocation(v)
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if v := i.Identifier; v != nil {
		m.SetIdentifier(*v)
	}
	if i.ClearUsers {
		m.ClearUsers()
	}
	if v := i.AddUserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.RemoveUserIDs; len(v) > 0 {
		m.RemoveUserIDs(v...)
	}
	if i.ClearParentStation {
		m.ClearParentStation()
	}
	if v := i.ParentStationID; v != nil {
		m.SetParentStationID(*v)
	}
	if i.ClearChildStations {
		m.ClearChildStations()
	}
	if v := i.AddChildStationIDs; len(v) > 0 {
		m.AddChildStationIDs(v...)
	}
	if v := i.RemoveChildStationIDs; len(v) > 0 {
		m.RemoveChildStationIDs(v...)
	}
}

// SetInput applies the change-set in the UpdatePoliceStationInput on the PoliceStationUpdate builder.
func (c *PoliceStationUpdate) SetInput(i UpdatePoliceStationInput) *PoliceStationUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePoliceStationInput on the PoliceStationUpdateOne builder.
func (c *PoliceStationUpdateOne) SetInput(i UpdatePoliceStationInput) *PoliceStationUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateRoleInput represents a mutation input for creating roles.
type CreateRoleInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Name          string
	PermissionIDs []int
	UserIDs       []int
}

// Mutate applies the CreateRoleInput on the RoleMutation builder.
func (i *CreateRoleInput) Mutate(m *RoleMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	if v := i.PermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreateRoleInput on the RoleCreate builder.
func (c *RoleCreate) SetInput(i CreateRoleInput) *RoleCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRoleInput represents a mutation input for updating roles.
type UpdateRoleInput struct {
	UpdatedAt           *time.Time
	Name                *string
	ClearPermissions    bool
	AddPermissionIDs    []int
	RemovePermissionIDs []int
	ClearUsers          bool
	AddUserIDs          []int
	RemoveUserIDs       []int
}

// Mutate applies the UpdateRoleInput on the RoleMutation builder.
func (i *UpdateRoleInput) Mutate(m *RoleMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearPermissions {
		m.ClearPermissions()
	}
	if v := i.AddPermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
	if v := i.RemovePermissionIDs; len(v) > 0 {
		m.RemovePermissionIDs(v...)
	}
	if i.ClearUsers {
		m.ClearUsers()
	}
	if v := i.AddUserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.RemoveUserIDs; len(v) > 0 {
		m.RemoveUserIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateRoleInput on the RoleUpdate builder.
func (c *RoleUpdate) SetInput(i UpdateRoleInput) *RoleUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRoleInput on the RoleUpdateOne builder.
func (c *RoleUpdateOne) SetInput(i UpdateRoleInput) *RoleUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string
	Email     string
	Password  string
	Phone     *string
	Active    *bool
	RoleID    *int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetName(i.Name)
	m.SetEmail(i.Email)
	m.SetPassword(i.Password)
	if v := i.Phone; v != nil {
		m.SetPhone(*v)
	}
	if v := i.Active; v != nil {
		m.SetActive(*v)
	}
	if v := i.RoleID; v != nil {
		m.SetRoleID(*v)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	UpdatedAt  *time.Time
	Name       *string
	Email      *string
	Password   *string
	ClearPhone bool
	Phone      *string
	Active     *bool
	ClearRole  bool
	RoleID     *int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if i.ClearPhone {
		m.ClearPhone()
	}
	if v := i.Phone; v != nil {
		m.SetPhone(*v)
	}
	if v := i.Active; v != nil {
		m.SetActive(*v)
	}
	if i.ClearRole {
		m.ClearRole()
	}
	if v := i.RoleID; v != nil {
		m.SetRoleID(*v)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateVehicleDataInput represents a mutation input for creating vehicledataslice.
type CreateVehicleDataInput struct {
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
	PlateBoundingBox     []int
	PlateChannel         *int
	PlateIsExist         *bool
	PlateColor           *string
	PlateNumber          *string
	PlateType            *string
	PlateRegion          *string
	PlateUploadNum       *int
	SnapAllowUser        *bool
	SnapAllowUserEndTime *string
	SnapDefenceCode      *string
	SnapDeviceID         *string
	SnapInCarPeopleNum   *int
	SnapLanNo            *int
	SnapOpenStrobe       *bool
	VehicleBoundingBox   []int
	VehicleSeries        *string
}

// Mutate applies the CreateVehicleDataInput on the VehicleDataMutation builder.
func (i *CreateVehicleDataInput) Mutate(m *VehicleDataMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.PlateBoundingBox; v != nil {
		m.SetPlateBoundingBox(v)
	}
	if v := i.PlateChannel; v != nil {
		m.SetPlateChannel(*v)
	}
	if v := i.PlateIsExist; v != nil {
		m.SetPlateIsExist(*v)
	}
	if v := i.PlateColor; v != nil {
		m.SetPlateColor(*v)
	}
	if v := i.PlateNumber; v != nil {
		m.SetPlateNumber(*v)
	}
	if v := i.PlateType; v != nil {
		m.SetPlateType(*v)
	}
	if v := i.PlateRegion; v != nil {
		m.SetPlateRegion(*v)
	}
	if v := i.PlateUploadNum; v != nil {
		m.SetPlateUploadNum(*v)
	}
	if v := i.SnapAllowUser; v != nil {
		m.SetSnapAllowUser(*v)
	}
	if v := i.SnapAllowUserEndTime; v != nil {
		m.SetSnapAllowUserEndTime(*v)
	}
	if v := i.SnapDefenceCode; v != nil {
		m.SetSnapDefenceCode(*v)
	}
	if v := i.SnapDeviceID; v != nil {
		m.SetSnapDeviceID(*v)
	}
	if v := i.SnapInCarPeopleNum; v != nil {
		m.SetSnapInCarPeopleNum(*v)
	}
	if v := i.SnapLanNo; v != nil {
		m.SetSnapLanNo(*v)
	}
	if v := i.SnapOpenStrobe; v != nil {
		m.SetSnapOpenStrobe(*v)
	}
	if v := i.VehicleBoundingBox; v != nil {
		m.SetVehicleBoundingBox(v)
	}
	if v := i.VehicleSeries; v != nil {
		m.SetVehicleSeries(*v)
	}
}

// SetInput applies the change-set in the CreateVehicleDataInput on the VehicleDataCreate builder.
func (c *VehicleDataCreate) SetInput(i CreateVehicleDataInput) *VehicleDataCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateVehicleDataInput represents a mutation input for updating vehicledataslice.
type UpdateVehicleDataInput struct {
	UpdatedAt                 *time.Time
	ClearPlateBoundingBox     bool
	PlateBoundingBox          []int
	AppendPlateBoundingBox    []int
	ClearPlateChannel         bool
	PlateChannel              *int
	ClearPlateIsExist         bool
	PlateIsExist              *bool
	ClearPlateColor           bool
	PlateColor                *string
	ClearPlateNumber          bool
	PlateNumber               *string
	ClearPlateType            bool
	PlateType                 *string
	ClearPlateRegion          bool
	PlateRegion               *string
	ClearPlateUploadNum       bool
	PlateUploadNum            *int
	ClearSnapAllowUser        bool
	SnapAllowUser             *bool
	ClearSnapAllowUserEndTime bool
	SnapAllowUserEndTime      *string
	ClearSnapDefenceCode      bool
	SnapDefenceCode           *string
	ClearSnapDeviceID         bool
	SnapDeviceID              *string
	ClearSnapInCarPeopleNum   bool
	SnapInCarPeopleNum        *int
	ClearSnapLanNo            bool
	SnapLanNo                 *int
	ClearSnapOpenStrobe       bool
	SnapOpenStrobe            *bool
	ClearVehicleBoundingBox   bool
	VehicleBoundingBox        []int
	AppendVehicleBoundingBox  []int
	ClearVehicleSeries        bool
	VehicleSeries             *string
}

// Mutate applies the UpdateVehicleDataInput on the VehicleDataMutation builder.
func (i *UpdateVehicleDataInput) Mutate(m *VehicleDataMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearPlateBoundingBox {
		m.ClearPlateBoundingBox()
	}
	if v := i.PlateBoundingBox; v != nil {
		m.SetPlateBoundingBox(v)
	}
	if i.AppendPlateBoundingBox != nil {
		m.AppendPlateBoundingBox(i.PlateBoundingBox)
	}
	if i.ClearPlateChannel {
		m.ClearPlateChannel()
	}
	if v := i.PlateChannel; v != nil {
		m.SetPlateChannel(*v)
	}
	if i.ClearPlateIsExist {
		m.ClearPlateIsExist()
	}
	if v := i.PlateIsExist; v != nil {
		m.SetPlateIsExist(*v)
	}
	if i.ClearPlateColor {
		m.ClearPlateColor()
	}
	if v := i.PlateColor; v != nil {
		m.SetPlateColor(*v)
	}
	if i.ClearPlateNumber {
		m.ClearPlateNumber()
	}
	if v := i.PlateNumber; v != nil {
		m.SetPlateNumber(*v)
	}
	if i.ClearPlateType {
		m.ClearPlateType()
	}
	if v := i.PlateType; v != nil {
		m.SetPlateType(*v)
	}
	if i.ClearPlateRegion {
		m.ClearPlateRegion()
	}
	if v := i.PlateRegion; v != nil {
		m.SetPlateRegion(*v)
	}
	if i.ClearPlateUploadNum {
		m.ClearPlateUploadNum()
	}
	if v := i.PlateUploadNum; v != nil {
		m.SetPlateUploadNum(*v)
	}
	if i.ClearSnapAllowUser {
		m.ClearSnapAllowUser()
	}
	if v := i.SnapAllowUser; v != nil {
		m.SetSnapAllowUser(*v)
	}
	if i.ClearSnapAllowUserEndTime {
		m.ClearSnapAllowUserEndTime()
	}
	if v := i.SnapAllowUserEndTime; v != nil {
		m.SetSnapAllowUserEndTime(*v)
	}
	if i.ClearSnapDefenceCode {
		m.ClearSnapDefenceCode()
	}
	if v := i.SnapDefenceCode; v != nil {
		m.SetSnapDefenceCode(*v)
	}
	if i.ClearSnapDeviceID {
		m.ClearSnapDeviceID()
	}
	if v := i.SnapDeviceID; v != nil {
		m.SetSnapDeviceID(*v)
	}
	if i.ClearSnapInCarPeopleNum {
		m.ClearSnapInCarPeopleNum()
	}
	if v := i.SnapInCarPeopleNum; v != nil {
		m.SetSnapInCarPeopleNum(*v)
	}
	if i.ClearSnapLanNo {
		m.ClearSnapLanNo()
	}
	if v := i.SnapLanNo; v != nil {
		m.SetSnapLanNo(*v)
	}
	if i.ClearSnapOpenStrobe {
		m.ClearSnapOpenStrobe()
	}
	if v := i.SnapOpenStrobe; v != nil {
		m.SetSnapOpenStrobe(*v)
	}
	if i.ClearVehicleBoundingBox {
		m.ClearVehicleBoundingBox()
	}
	if v := i.VehicleBoundingBox; v != nil {
		m.SetVehicleBoundingBox(v)
	}
	if i.AppendVehicleBoundingBox != nil {
		m.AppendVehicleBoundingBox(i.VehicleBoundingBox)
	}
	if i.ClearVehicleSeries {
		m.ClearVehicleSeries()
	}
	if v := i.VehicleSeries; v != nil {
		m.SetVehicleSeries(*v)
	}
}

// SetInput applies the change-set in the UpdateVehicleDataInput on the VehicleDataUpdate builder.
func (c *VehicleDataUpdate) SetInput(i UpdateVehicleDataInput) *VehicleDataUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateVehicleDataInput on the VehicleDataUpdateOne builder.
func (c *VehicleDataUpdateOne) SetInput(i UpdateVehicleDataInput) *VehicleDataUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
