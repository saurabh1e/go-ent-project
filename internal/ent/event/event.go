// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPlateBoundingBox holds the string denoting the plate_bounding_box field in the database.
	FieldPlateBoundingBox = "plate_bounding_box"
	// FieldPlateChannel holds the string denoting the plate_channel field in the database.
	FieldPlateChannel = "plate_channel"
	// FieldPlateIsExist holds the string denoting the plate_is_exist field in the database.
	FieldPlateIsExist = "plate_is_exist"
	// FieldPlateColor holds the string denoting the plate_color field in the database.
	FieldPlateColor = "plate_color"
	// FieldPlateNumber holds the string denoting the plate_number field in the database.
	FieldPlateNumber = "plate_number"
	// FieldPlateType holds the string denoting the plate_type field in the database.
	FieldPlateType = "plate_type"
	// FieldPlateRegion holds the string denoting the plate_region field in the database.
	FieldPlateRegion = "plate_region"
	// FieldPlateUploadNum holds the string denoting the plate_upload_num field in the database.
	FieldPlateUploadNum = "plate_upload_num"
	// FieldSnapAllowUser holds the string denoting the snap_allow_user field in the database.
	FieldSnapAllowUser = "snap_allow_user"
	// FieldSnapAllowUserEndTime holds the string denoting the snap_allow_user_end_time field in the database.
	FieldSnapAllowUserEndTime = "snap_allow_user_end_time"
	// FieldSnapDefenceCode holds the string denoting the snap_defence_code field in the database.
	FieldSnapDefenceCode = "snap_defence_code"
	// FieldSnapDeviceID holds the string denoting the snap_device_id field in the database.
	FieldSnapDeviceID = "snap_device_id"
	// FieldSnapInCarPeopleNum holds the string denoting the snap_in_car_people_num field in the database.
	FieldSnapInCarPeopleNum = "snap_in_car_people_num"
	// FieldSnapLanNo holds the string denoting the snap_lan_no field in the database.
	FieldSnapLanNo = "snap_lan_no"
	// FieldSnapOpenStrobe holds the string denoting the snap_open_strobe field in the database.
	FieldSnapOpenStrobe = "snap_open_strobe"
	// FieldVehicleBoundingBox holds the string denoting the vehicle_bounding_box field in the database.
	FieldVehicleBoundingBox = "vehicle_bounding_box"
	// FieldVehicleSeries holds the string denoting the vehicle_series field in the database.
	FieldVehicleSeries = "vehicle_series"
	// Table holds the table name of the event in the database.
	Table = "events"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPlateBoundingBox,
	FieldPlateChannel,
	FieldPlateIsExist,
	FieldPlateColor,
	FieldPlateNumber,
	FieldPlateType,
	FieldPlateRegion,
	FieldPlateUploadNum,
	FieldSnapAllowUser,
	FieldSnapAllowUserEndTime,
	FieldSnapDefenceCode,
	FieldSnapDeviceID,
	FieldSnapInCarPeopleNum,
	FieldSnapLanNo,
	FieldSnapOpenStrobe,
	FieldVehicleBoundingBox,
	FieldVehicleSeries,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Event queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByPlateChannel orders the results by the plate_channel field.
func ByPlateChannel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlateChannel, opts...).ToFunc()
}

// ByPlateIsExist orders the results by the plate_is_exist field.
func ByPlateIsExist(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlateIsExist, opts...).ToFunc()
}

// ByPlateColor orders the results by the plate_color field.
func ByPlateColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlateColor, opts...).ToFunc()
}

// ByPlateNumber orders the results by the plate_number field.
func ByPlateNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlateNumber, opts...).ToFunc()
}

// ByPlateType orders the results by the plate_type field.
func ByPlateType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlateType, opts...).ToFunc()
}

// ByPlateRegion orders the results by the plate_region field.
func ByPlateRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlateRegion, opts...).ToFunc()
}

// ByPlateUploadNum orders the results by the plate_upload_num field.
func ByPlateUploadNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlateUploadNum, opts...).ToFunc()
}

// BySnapAllowUser orders the results by the snap_allow_user field.
func BySnapAllowUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapAllowUser, opts...).ToFunc()
}

// BySnapAllowUserEndTime orders the results by the snap_allow_user_end_time field.
func BySnapAllowUserEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapAllowUserEndTime, opts...).ToFunc()
}

// BySnapDefenceCode orders the results by the snap_defence_code field.
func BySnapDefenceCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapDefenceCode, opts...).ToFunc()
}

// BySnapDeviceID orders the results by the snap_device_id field.
func BySnapDeviceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapDeviceID, opts...).ToFunc()
}

// BySnapInCarPeopleNum orders the results by the snap_in_car_people_num field.
func BySnapInCarPeopleNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapInCarPeopleNum, opts...).ToFunc()
}

// BySnapLanNo orders the results by the snap_lan_no field.
func BySnapLanNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapLanNo, opts...).ToFunc()
}

// BySnapOpenStrobe orders the results by the snap_open_strobe field.
func BySnapOpenStrobe(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSnapOpenStrobe, opts...).ToFunc()
}

// ByVehicleSeries orders the results by the vehicle_series field.
func ByVehicleSeries(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVehicleSeries, opts...).ToFunc()
}
