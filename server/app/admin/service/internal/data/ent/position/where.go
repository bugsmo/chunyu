// Code generated by ent, DO NOT EDIT.

package position

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bugsmo/chunyu/app/admin/service/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUpdateTime, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDeleteTime, v))
}

// CreateBy applies equality check predicate on the "create_by" field. It's identical to CreateByEQ.
func CreateBy(v uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreateBy, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldRemark, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCode, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldParentID, v))
}

// OrderNo applies equality check predicate on the "order_no" field. It's identical to OrderNoEQ.
func OrderNo(v int32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldOrderNo, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldUpdateTime))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldDeleteTime))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldStatus))
}

// CreateByEQ applies the EQ predicate on the "create_by" field.
func CreateByEQ(v uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCreateBy, v))
}

// CreateByNEQ applies the NEQ predicate on the "create_by" field.
func CreateByNEQ(v uint32) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldCreateBy, v))
}

// CreateByIn applies the In predicate on the "create_by" field.
func CreateByIn(vs ...uint32) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldCreateBy, vs...))
}

// CreateByNotIn applies the NotIn predicate on the "create_by" field.
func CreateByNotIn(vs ...uint32) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldCreateBy, vs...))
}

// CreateByGT applies the GT predicate on the "create_by" field.
func CreateByGT(v uint32) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldCreateBy, v))
}

// CreateByGTE applies the GTE predicate on the "create_by" field.
func CreateByGTE(v uint32) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldCreateBy, v))
}

// CreateByLT applies the LT predicate on the "create_by" field.
func CreateByLT(v uint32) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldCreateBy, v))
}

// CreateByLTE applies the LTE predicate on the "create_by" field.
func CreateByLTE(v uint32) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldCreateBy, v))
}

// CreateByIsNil applies the IsNil predicate on the "create_by" field.
func CreateByIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldCreateBy))
}

// CreateByNotNil applies the NotNil predicate on the "create_by" field.
func CreateByNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldCreateBy))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldRemark))
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldRemark))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldRemark, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldName, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Position {
	return predicate.Position(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Position {
	return predicate.Position(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Position {
	return predicate.Position(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Position {
	return predicate.Position(sql.FieldContainsFold(FieldCode, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v uint32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v uint32) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...uint32) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...uint32) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.Position {
	return predicate.Position(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.Position {
	return predicate.Position(sql.FieldNotNull(FieldParentID))
}

// OrderNoEQ applies the EQ predicate on the "order_no" field.
func OrderNoEQ(v int32) predicate.Position {
	return predicate.Position(sql.FieldEQ(FieldOrderNo, v))
}

// OrderNoNEQ applies the NEQ predicate on the "order_no" field.
func OrderNoNEQ(v int32) predicate.Position {
	return predicate.Position(sql.FieldNEQ(FieldOrderNo, v))
}

// OrderNoIn applies the In predicate on the "order_no" field.
func OrderNoIn(vs ...int32) predicate.Position {
	return predicate.Position(sql.FieldIn(FieldOrderNo, vs...))
}

// OrderNoNotIn applies the NotIn predicate on the "order_no" field.
func OrderNoNotIn(vs ...int32) predicate.Position {
	return predicate.Position(sql.FieldNotIn(FieldOrderNo, vs...))
}

// OrderNoGT applies the GT predicate on the "order_no" field.
func OrderNoGT(v int32) predicate.Position {
	return predicate.Position(sql.FieldGT(FieldOrderNo, v))
}

// OrderNoGTE applies the GTE predicate on the "order_no" field.
func OrderNoGTE(v int32) predicate.Position {
	return predicate.Position(sql.FieldGTE(FieldOrderNo, v))
}

// OrderNoLT applies the LT predicate on the "order_no" field.
func OrderNoLT(v int32) predicate.Position {
	return predicate.Position(sql.FieldLT(FieldOrderNo, v))
}

// OrderNoLTE applies the LTE predicate on the "order_no" field.
func OrderNoLTE(v int32) predicate.Position {
	return predicate.Position(sql.FieldLTE(FieldOrderNo, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Position) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Position) predicate.Position {
	return predicate.Position(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Position) predicate.Position {
	return predicate.Position(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Position) predicate.Position {
	return predicate.Position(sql.NotPredicates(p))
}
