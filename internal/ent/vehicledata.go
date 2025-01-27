// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"go-ent-project/internal/ent/vehicledata"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VehicleData is the model entity for the VehicleData schema.
type VehicleData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Bounding box coordinates of the plate
	PlateBoundingBox []int `json:"plate_bounding_box,omitempty"`
	// Channel of the plate
	PlateChannel int `json:"plate_channel,omitempty"`
	// Indicates whether the plate exists
	PlateIsExist bool `json:"plate_is_exist,omitempty"`
	// Color of the plate
	PlateColor string `json:"plate_color,omitempty"`
	// Number on the plate
	PlateNumber string `json:"plate_number,omitempty"`
	// Type of the plate
	PlateType string `json:"plate_type,omitempty"`
	// Region of the plate
	PlateRegion string `json:"plate_region,omitempty"`
	// Upload number of the plate
	PlateUploadNum int `json:"plate_upload_num,omitempty"`
	// Indicates if user interaction is allowed
	SnapAllowUser bool `json:"snap_allow_user,omitempty"`
	// End time for user interaction allowance
	SnapAllowUserEndTime string `json:"snap_allow_user_end_time,omitempty"`
	// Defence code
	SnapDefenceCode string `json:"snap_defence_code,omitempty"`
	// Device ID
	SnapDeviceID string `json:"snap_device_id,omitempty"`
	// Number of people in the car
	SnapInCarPeopleNum int `json:"snap_in_car_people_num,omitempty"`
	// LAN number
	SnapLanNo int `json:"snap_lan_no,omitempty"`
	// Indicates if strobe is open
	SnapOpenStrobe bool `json:"snap_open_strobe,omitempty"`
	// Bounding box coordinates of the vehicle
	VehicleBoundingBox []int `json:"vehicle_bounding_box,omitempty"`
	// Vehicle series
	VehicleSeries string `json:"vehicle_series,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VehicleData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehicledata.FieldPlateBoundingBox, vehicledata.FieldVehicleBoundingBox:
			values[i] = new([]byte)
		case vehicledata.FieldPlateIsExist, vehicledata.FieldSnapAllowUser, vehicledata.FieldSnapOpenStrobe:
			values[i] = new(sql.NullBool)
		case vehicledata.FieldID, vehicledata.FieldPlateChannel, vehicledata.FieldPlateUploadNum, vehicledata.FieldSnapInCarPeopleNum, vehicledata.FieldSnapLanNo:
			values[i] = new(sql.NullInt64)
		case vehicledata.FieldPlateColor, vehicledata.FieldPlateNumber, vehicledata.FieldPlateType, vehicledata.FieldPlateRegion, vehicledata.FieldSnapAllowUserEndTime, vehicledata.FieldSnapDefenceCode, vehicledata.FieldSnapDeviceID, vehicledata.FieldVehicleSeries:
			values[i] = new(sql.NullString)
		case vehicledata.FieldCreatedAt, vehicledata.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VehicleData fields.
func (vd *VehicleData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehicledata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vd.ID = int(value.Int64)
		case vehicledata.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vd.CreatedAt = value.Time
			}
		case vehicledata.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vd.UpdatedAt = value.Time
			}
		case vehicledata.FieldPlateBoundingBox:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field plate_bounding_box", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &vd.PlateBoundingBox); err != nil {
					return fmt.Errorf("unmarshal field plate_bounding_box: %w", err)
				}
			}
		case vehicledata.FieldPlateChannel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field plate_channel", values[i])
			} else if value.Valid {
				vd.PlateChannel = int(value.Int64)
			}
		case vehicledata.FieldPlateIsExist:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field plate_is_exist", values[i])
			} else if value.Valid {
				vd.PlateIsExist = value.Bool
			}
		case vehicledata.FieldPlateColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plate_color", values[i])
			} else if value.Valid {
				vd.PlateColor = value.String
			}
		case vehicledata.FieldPlateNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plate_number", values[i])
			} else if value.Valid {
				vd.PlateNumber = value.String
			}
		case vehicledata.FieldPlateType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plate_type", values[i])
			} else if value.Valid {
				vd.PlateType = value.String
			}
		case vehicledata.FieldPlateRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plate_region", values[i])
			} else if value.Valid {
				vd.PlateRegion = value.String
			}
		case vehicledata.FieldPlateUploadNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field plate_upload_num", values[i])
			} else if value.Valid {
				vd.PlateUploadNum = int(value.Int64)
			}
		case vehicledata.FieldSnapAllowUser:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field snap_allow_user", values[i])
			} else if value.Valid {
				vd.SnapAllowUser = value.Bool
			}
		case vehicledata.FieldSnapAllowUserEndTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field snap_allow_user_end_time", values[i])
			} else if value.Valid {
				vd.SnapAllowUserEndTime = value.String
			}
		case vehicledata.FieldSnapDefenceCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field snap_defence_code", values[i])
			} else if value.Valid {
				vd.SnapDefenceCode = value.String
			}
		case vehicledata.FieldSnapDeviceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field snap_device_id", values[i])
			} else if value.Valid {
				vd.SnapDeviceID = value.String
			}
		case vehicledata.FieldSnapInCarPeopleNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field snap_in_car_people_num", values[i])
			} else if value.Valid {
				vd.SnapInCarPeopleNum = int(value.Int64)
			}
		case vehicledata.FieldSnapLanNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field snap_lan_no", values[i])
			} else if value.Valid {
				vd.SnapLanNo = int(value.Int64)
			}
		case vehicledata.FieldSnapOpenStrobe:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field snap_open_strobe", values[i])
			} else if value.Valid {
				vd.SnapOpenStrobe = value.Bool
			}
		case vehicledata.FieldVehicleBoundingBox:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field vehicle_bounding_box", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &vd.VehicleBoundingBox); err != nil {
					return fmt.Errorf("unmarshal field vehicle_bounding_box: %w", err)
				}
			}
		case vehicledata.FieldVehicleSeries:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vehicle_series", values[i])
			} else if value.Valid {
				vd.VehicleSeries = value.String
			}
		default:
			vd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VehicleData.
// This includes values selected through modifiers, order, etc.
func (vd *VehicleData) Value(name string) (ent.Value, error) {
	return vd.selectValues.Get(name)
}

// Update returns a builder for updating this VehicleData.
// Note that you need to call VehicleData.Unwrap() before calling this method if this VehicleData
// was returned from a transaction, and the transaction was committed or rolled back.
func (vd *VehicleData) Update() *VehicleDataUpdateOne {
	return NewVehicleDataClient(vd.config).UpdateOne(vd)
}

// Unwrap unwraps the VehicleData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vd *VehicleData) Unwrap() *VehicleData {
	_tx, ok := vd.config.driver.(*txDriver)
	if !ok {
		panic("ent: VehicleData is not a transactional entity")
	}
	vd.config.driver = _tx.drv
	return vd
}

// String implements the fmt.Stringer.
func (vd *VehicleData) String() string {
	var builder strings.Builder
	builder.WriteString("VehicleData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(vd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("plate_bounding_box=")
	builder.WriteString(fmt.Sprintf("%v", vd.PlateBoundingBox))
	builder.WriteString(", ")
	builder.WriteString("plate_channel=")
	builder.WriteString(fmt.Sprintf("%v", vd.PlateChannel))
	builder.WriteString(", ")
	builder.WriteString("plate_is_exist=")
	builder.WriteString(fmt.Sprintf("%v", vd.PlateIsExist))
	builder.WriteString(", ")
	builder.WriteString("plate_color=")
	builder.WriteString(vd.PlateColor)
	builder.WriteString(", ")
	builder.WriteString("plate_number=")
	builder.WriteString(vd.PlateNumber)
	builder.WriteString(", ")
	builder.WriteString("plate_type=")
	builder.WriteString(vd.PlateType)
	builder.WriteString(", ")
	builder.WriteString("plate_region=")
	builder.WriteString(vd.PlateRegion)
	builder.WriteString(", ")
	builder.WriteString("plate_upload_num=")
	builder.WriteString(fmt.Sprintf("%v", vd.PlateUploadNum))
	builder.WriteString(", ")
	builder.WriteString("snap_allow_user=")
	builder.WriteString(fmt.Sprintf("%v", vd.SnapAllowUser))
	builder.WriteString(", ")
	builder.WriteString("snap_allow_user_end_time=")
	builder.WriteString(vd.SnapAllowUserEndTime)
	builder.WriteString(", ")
	builder.WriteString("snap_defence_code=")
	builder.WriteString(vd.SnapDefenceCode)
	builder.WriteString(", ")
	builder.WriteString("snap_device_id=")
	builder.WriteString(vd.SnapDeviceID)
	builder.WriteString(", ")
	builder.WriteString("snap_in_car_people_num=")
	builder.WriteString(fmt.Sprintf("%v", vd.SnapInCarPeopleNum))
	builder.WriteString(", ")
	builder.WriteString("snap_lan_no=")
	builder.WriteString(fmt.Sprintf("%v", vd.SnapLanNo))
	builder.WriteString(", ")
	builder.WriteString("snap_open_strobe=")
	builder.WriteString(fmt.Sprintf("%v", vd.SnapOpenStrobe))
	builder.WriteString(", ")
	builder.WriteString("vehicle_bounding_box=")
	builder.WriteString(fmt.Sprintf("%v", vd.VehicleBoundingBox))
	builder.WriteString(", ")
	builder.WriteString("vehicle_series=")
	builder.WriteString(vd.VehicleSeries)
	builder.WriteByte(')')
	return builder.String()
}

// VehicleDataSlice is a parsable slice of VehicleData.
type VehicleDataSlice []*VehicleData
