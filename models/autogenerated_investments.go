package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set InvestmentQuerySet

// InvestmentQuerySet is an queryset type for Investment
type InvestmentQuerySet struct {
	db *gorm.DB
}

// NewInvestmentQuerySet constructs new InvestmentQuerySet
func NewInvestmentQuerySet(db *gorm.DB) InvestmentQuerySet {
	return InvestmentQuerySet{
		db: db.Model(&Investment{}),
	}
}

func (qs InvestmentQuerySet) w(db *gorm.DB) InvestmentQuerySet {
	return NewInvestmentQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) All(ret *[]Investment) error {
	return qs.db.Find(ret).Error
}

// AmountEq is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountEq(amount uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("amount = ?", amount))
}

// AmountGt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountGt(amount uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("amount > ?", amount))
}

// AmountGte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountGte(amount uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("amount >= ?", amount))
}

// AmountIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountIn(amount uint, amountRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{amount}
	for _, arg := range amountRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("amount IN (?)", iArgs))
}

// AmountLt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountLt(amount uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("amount < ?", amount))
}

// AmountLte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountLte(amount uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("amount <= ?", amount))
}

// AmountNe is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountNe(amount uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("amount != ?", amount))
}

// AmountNotIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) AmountNotIn(amount uint, amountRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{amount}
	for _, arg := range amountRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("amount NOT IN (?)", iArgs))
}

// Count is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Investment) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) CreatedAtEq(createdAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) CreatedAtGt(createdAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) CreatedAtGte(createdAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) CreatedAtLt(createdAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) CreatedAtLte(createdAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) CreatedAtNe(createdAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Investment) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) Delete() error {
	return qs.db.Delete(Investment{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtEq(deletedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtGt(deletedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtGte(deletedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtIsNotNull() InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtIsNull() InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtLt(deletedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtLte(deletedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) DeletedAtNe(deletedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) GetUpdater() InvestmentUpdater {
	return NewInvestmentUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDEq(ID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDGt(ID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDGte(ID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDIn(ID uint, IDRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDLt(ID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDLte(ID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDNe(ID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) IDNotIn(ID uint, IDRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) Limit(limit int) InvestmentQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs InvestmentQuerySet) One(ret *Investment) error {
	return qs.db.First(ret).Error
}

// OrderAscByAmount is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderAscByAmount() InvestmentQuerySet {
	return qs.w(qs.db.Order("amount ASC"))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderAscByCreatedAt() InvestmentQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderAscByDeletedAt() InvestmentQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderAscByID() InvestmentQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByPlanID is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderAscByPlanID() InvestmentQuerySet {
	return qs.w(qs.db.Order("plan_id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderAscByUpdatedAt() InvestmentQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderAscByUserID is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderAscByUserID() InvestmentQuerySet {
	return qs.w(qs.db.Order("user_id ASC"))
}

// OrderDescByAmount is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderDescByAmount() InvestmentQuerySet {
	return qs.w(qs.db.Order("amount DESC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderDescByCreatedAt() InvestmentQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderDescByDeletedAt() InvestmentQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderDescByID() InvestmentQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByPlanID is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderDescByPlanID() InvestmentQuerySet {
	return qs.w(qs.db.Order("plan_id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderDescByUpdatedAt() InvestmentQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OrderDescByUserID is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) OrderDescByUserID() InvestmentQuerySet {
	return qs.w(qs.db.Order("user_id DESC"))
}

// PlanIDEq is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDEq(planID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("plan_id = ?", planID))
}

// PlanIDGt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDGt(planID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("plan_id > ?", planID))
}

// PlanIDGte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDGte(planID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("plan_id >= ?", planID))
}

// PlanIDIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDIn(planID uint, planIDRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{planID}
	for _, arg := range planIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("plan_id IN (?)", iArgs))
}

// PlanIDLt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDLt(planID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("plan_id < ?", planID))
}

// PlanIDLte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDLte(planID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("plan_id <= ?", planID))
}

// PlanIDNe is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDNe(planID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("plan_id != ?", planID))
}

// PlanIDNotIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) PlanIDNotIn(planID uint, planIDRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{planID}
	for _, arg := range planIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("plan_id NOT IN (?)", iArgs))
}

// SetAmount is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) SetAmount(amount uint) InvestmentUpdater {
	u.fields[string(InvestmentDBSchema.Amount)] = amount
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) SetCreatedAt(createdAt time.Time) InvestmentUpdater {
	u.fields[string(InvestmentDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) SetDeletedAt(deletedAt *time.Time) InvestmentUpdater {
	u.fields[string(InvestmentDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) SetID(ID uint) InvestmentUpdater {
	u.fields[string(InvestmentDBSchema.ID)] = ID
	return u
}

// SetPlanID is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) SetPlanID(planID uint) InvestmentUpdater {
	u.fields[string(InvestmentDBSchema.PlanID)] = planID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) SetUpdatedAt(updatedAt time.Time) InvestmentUpdater {
	u.fields[string(InvestmentDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUserID is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) SetUserID(userID uint) InvestmentUpdater {
	u.fields[string(InvestmentDBSchema.UserID)] = userID
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u InvestmentUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UpdatedAtEq(updatedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UpdatedAtGt(updatedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UpdatedAtGte(updatedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UpdatedAtLt(updatedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UpdatedAtLte(updatedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UpdatedAtNe(updatedAt time.Time) InvestmentQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UserIDEq is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDEq(userID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("user_id = ?", userID))
}

// UserIDGt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDGt(userID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("user_id > ?", userID))
}

// UserIDGte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDGte(userID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("user_id >= ?", userID))
}

// UserIDIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDIn(userID uint, userIDRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{userID}
	for _, arg := range userIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("user_id IN (?)", iArgs))
}

// UserIDLt is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDLt(userID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("user_id < ?", userID))
}

// UserIDLte is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDLte(userID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("user_id <= ?", userID))
}

// UserIDNe is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDNe(userID uint) InvestmentQuerySet {
	return qs.w(qs.db.Where("user_id != ?", userID))
}

// UserIDNotIn is an autogenerated method
// nolint: dupl
func (qs InvestmentQuerySet) UserIDNotIn(userID uint, userIDRest ...uint) InvestmentQuerySet {
	iArgs := []interface{}{userID}
	for _, arg := range userIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("user_id NOT IN (?)", iArgs))
}

// ===== END of query set InvestmentQuerySet

// ===== BEGIN of Investment modifiers

type investmentDBSchemaField string

func (f investmentDBSchemaField) String() string {
	return string(f)
}

// InvestmentDBSchema stores db field names of Investment
var InvestmentDBSchema = struct {
	ID        investmentDBSchemaField
	CreatedAt investmentDBSchemaField
	UpdatedAt investmentDBSchemaField
	DeletedAt investmentDBSchemaField
	PlanID    investmentDBSchemaField
	UserID    investmentDBSchemaField
	Amount    investmentDBSchemaField
}{

	ID:        investmentDBSchemaField("id"),
	CreatedAt: investmentDBSchemaField("created_at"),
	UpdatedAt: investmentDBSchemaField("updated_at"),
	DeletedAt: investmentDBSchemaField("deleted_at"),
	PlanID:    investmentDBSchemaField("plan_id"),
	UserID:    investmentDBSchemaField("user_id"),
	Amount:    investmentDBSchemaField("amount"),
}

// Update updates Investment fields by primary key
func (o *Investment) Update(db *gorm.DB, fields ...investmentDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"plan_id":    o.PlanID,
		"user_id":    o.UserID,
		"amount":     o.Amount,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Investment %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// InvestmentUpdater is an Investment updates manager
type InvestmentUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewInvestmentUpdater creates new Investment updater
func NewInvestmentUpdater(db *gorm.DB) InvestmentUpdater {
	return InvestmentUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Investment{}),
	}
}

// ===== END of Investment modifiers

// ===== END of all query sets