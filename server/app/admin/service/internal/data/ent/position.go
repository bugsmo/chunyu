// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent/position"
)

// Position is the model entity for the Position schema.
type Position struct {
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
	// 状态
	Status *position.Status `json:"status,omitempty"`
	// 创建者ID
	CreateBy *uint32 `json:"create_by,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty"`
	// 职位名称
	Name string `json:"name,omitempty"`
	// 职位标识
	Code string `json:"code,omitempty"`
	// 上一层职位ID
	ParentID uint32 `json:"parent_id,omitempty"`
	// 排序ID
	OrderNo int32 `json:"order_no,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PositionQuery when eager-loading is set.
	Edges        PositionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PositionEdges holds the relations/edges for other nodes in the graph.
type PositionEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Position `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Position `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PositionEdges) ParentOrErr() (*Position, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: position.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e PositionEdges) ChildrenOrErr() ([]*Position, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Position) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case position.FieldID, position.FieldCreateBy, position.FieldParentID, position.FieldOrderNo:
			values[i] = new(sql.NullInt64)
		case position.FieldStatus, position.FieldRemark, position.FieldName, position.FieldCode:
			values[i] = new(sql.NullString)
		case position.FieldCreateTime, position.FieldUpdateTime, position.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Position fields.
func (po *Position) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case position.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = uint32(value.Int64)
		case position.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				po.CreateTime = new(time.Time)
				*po.CreateTime = value.Time
			}
		case position.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				po.UpdateTime = new(time.Time)
				*po.UpdateTime = value.Time
			}
		case position.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				po.DeleteTime = new(time.Time)
				*po.DeleteTime = value.Time
			}
		case position.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				po.Status = new(position.Status)
				*po.Status = position.Status(value.String)
			}
		case position.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				po.CreateBy = new(uint32)
				*po.CreateBy = uint32(value.Int64)
			}
		case position.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				po.Remark = new(string)
				*po.Remark = value.String
			}
		case position.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				po.Name = value.String
			}
		case position.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				po.Code = value.String
			}
		case position.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				po.ParentID = uint32(value.Int64)
			}
		case position.FieldOrderNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				po.OrderNo = int32(value.Int64)
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Position.
// This includes values selected through modifiers, order, etc.
func (po *Position) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Position entity.
func (po *Position) QueryParent() *PositionQuery {
	return NewPositionClient(po.config).QueryParent(po)
}

// QueryChildren queries the "children" edge of the Position entity.
func (po *Position) QueryChildren() *PositionQuery {
	return NewPositionClient(po.config).QueryChildren(po)
}

// Update returns a builder for updating this Position.
// Note that you need to call Position.Unwrap() before calling this method if this Position
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Position) Update() *PositionUpdateOne {
	return NewPositionClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Position entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Position) Unwrap() *Position {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Position is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Position) String() string {
	var builder strings.Builder
	builder.WriteString("Position(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	if v := po.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := po.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := po.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := po.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.CreateBy; v != nil {
		builder.WriteString("create_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.Remark; v != nil {
		builder.WriteString("remark=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(po.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(po.Code)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", po.ParentID))
	builder.WriteString(", ")
	builder.WriteString("order_no=")
	builder.WriteString(fmt.Sprintf("%v", po.OrderNo))
	builder.WriteByte(')')
	return builder.String()
}

// Positions is a parsable slice of Position.
type Positions []*Position
