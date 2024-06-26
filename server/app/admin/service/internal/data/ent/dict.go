// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent/dict"
)

// Dict is the model entity for the Dict schema.
type Dict struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建时间
	CreateTime *time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 创建者ID
	CreateBy *uint32 `json:"create_by,omitempty"`
	// 字典名称
	Name *string `json:"name,omitempty"`
	// 描述
	Description  *string `json:"description,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dict) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dict.FieldID, dict.FieldCreateBy:
			values[i] = new(sql.NullInt64)
		case dict.FieldName, dict.FieldDescription:
			values[i] = new(sql.NullString)
		case dict.FieldCreateTime, dict.FieldUpdateTime, dict.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dict fields.
func (d *Dict) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dict.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = uint32(value.Int64)
		case dict.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = new(time.Time)
				*d.CreateTime = value.Time
			}
		case dict.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				d.UpdateTime = new(time.Time)
				*d.UpdateTime = value.Time
			}
		case dict.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				d.DeleteTime = new(time.Time)
				*d.DeleteTime = value.Time
			}
		case dict.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				d.CreateBy = new(uint32)
				*d.CreateBy = uint32(value.Int64)
			}
		case dict.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = new(string)
				*d.Name = value.String
			}
		case dict.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = new(string)
				*d.Description = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Dict.
// This includes values selected through modifiers, order, etc.
func (d *Dict) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// Update returns a builder for updating this Dict.
// Note that you need to call Dict.Unwrap() before calling this method if this Dict
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dict) Update() *DictUpdateOne {
	return NewDictClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dict entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dict) Unwrap() *Dict {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dict is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dict) String() string {
	var builder strings.Builder
	builder.WriteString("Dict(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	if v := d.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := d.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := d.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := d.CreateBy; v != nil {
		builder.WriteString("create_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := d.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := d.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Dicts is a parsable slice of Dict.
type Dicts []*Dict
