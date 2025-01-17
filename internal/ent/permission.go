// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-ent-project/internal/ent/permission"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CanRead holds the value of the "can_read" field.
	CanRead bool `json:"can_read,omitempty"`
	// CanCreate holds the value of the "can_create" field.
	CanCreate bool `json:"can_create,omitempty"`
	// CanUpdate holds the value of the "can_update" field.
	CanUpdate bool `json:"can_update,omitempty"`
	// CanDelete holds the value of the "can_delete" field.
	CanDelete        bool `json:"can_delete,omitempty"`
	role_permissions *int
	selectValues     sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldCanRead, permission.FieldCanCreate, permission.FieldCanUpdate, permission.FieldCanDelete:
			values[i] = new(sql.NullBool)
		case permission.FieldID:
			values[i] = new(sql.NullInt64)
		case permission.FieldName:
			values[i] = new(sql.NullString)
		case permission.ForeignKeys[0]: // role_permissions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case permission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case permission.FieldCanRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_read", values[i])
			} else if value.Valid {
				pe.CanRead = value.Bool
			}
		case permission.FieldCanCreate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_create", values[i])
			} else if value.Valid {
				pe.CanCreate = value.Bool
			}
		case permission.FieldCanUpdate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_update", values[i])
			} else if value.Valid {
				pe.CanUpdate = value.Bool
			}
		case permission.FieldCanDelete:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field can_delete", values[i])
			} else if value.Valid {
				pe.CanDelete = value.Bool
			}
		case permission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field role_permissions", value)
			} else if value.Valid {
				pe.role_permissions = new(int)
				*pe.role_permissions = int(value.Int64)
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Permission.
// This includes values selected through modifiers, order, etc.
func (pe *Permission) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("can_read=")
	builder.WriteString(fmt.Sprintf("%v", pe.CanRead))
	builder.WriteString(", ")
	builder.WriteString("can_create=")
	builder.WriteString(fmt.Sprintf("%v", pe.CanCreate))
	builder.WriteString(", ")
	builder.WriteString("can_update=")
	builder.WriteString(fmt.Sprintf("%v", pe.CanUpdate))
	builder.WriteString(", ")
	builder.WriteString("can_delete=")
	builder.WriteString(fmt.Sprintf("%v", pe.CanDelete))
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
