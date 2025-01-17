// Code generated by ent, DO NOT EDIT.

package car

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the car type in the database.
	Label = "car"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMake holds the string denoting the make field in the database.
	FieldMake = "make"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldRegistration holds the string denoting the registration field in the database.
	FieldRegistration = "registration"
	// FieldColor holds the string denoting the color field in the database.
	FieldColor = "color"
	// Table holds the table name of the car in the database.
	Table = "cars"
)

// Columns holds all SQL columns for car fields.
var Columns = []string{
	FieldID,
	FieldMake,
	FieldModel,
	FieldYear,
	FieldRegistration,
	FieldColor,
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
	// MakeValidator is a validator for the "make" field. It is called by the builders before save.
	MakeValidator func(string) error
	// ModelValidator is a validator for the "model" field. It is called by the builders before save.
	ModelValidator func(string) error
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int) error
	// RegistrationValidator is a validator for the "registration" field. It is called by the builders before save.
	RegistrationValidator func(string) error
	// ColorValidator is a validator for the "color" field. It is called by the builders before save.
	ColorValidator func(string) error
)

// OrderOption defines the ordering options for the Car queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMake orders the results by the make field.
func ByMake(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMake, opts...).ToFunc()
}

// ByModel orders the results by the model field.
func ByModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModel, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByRegistration orders the results by the registration field.
func ByRegistration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegistration, opts...).ToFunc()
}

// ByColor orders the results by the color field.
func ByColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldColor, opts...).ToFunc()
}
