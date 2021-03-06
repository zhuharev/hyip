package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set GroupQuerySet

// GroupQuerySet is an queryset type for Group
type GroupQuerySet struct {
	db *gorm.DB
}

// NewGroupQuerySet constructs new GroupQuerySet
func NewGroupQuerySet(db *gorm.DB) GroupQuerySet {
	return GroupQuerySet{
		db: db.Model(&Group{}),
	}
}

func (qs GroupQuerySet) w(db *gorm.DB) GroupQuerySet {
	return NewGroupQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) All(ret *[]Group) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Group) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) CreatedAtEq(createdAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) CreatedAtGt(createdAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) CreatedAtGte(createdAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) CreatedAtLt(createdAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) CreatedAtLte(createdAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) CreatedAtNe(createdAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) Delete() error {
	return qs.db.Delete(Group{}).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Group) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtEq(deletedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtGt(deletedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtGte(deletedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtIsNotNull() GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtIsNull() GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtLt(deletedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtLte(deletedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) DeletedAtNe(deletedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) GetUpdater() GroupUpdater {
	return NewGroupUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDEq(ID uint) GroupQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDGt(ID uint) GroupQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDGte(ID uint) GroupQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDIn(ID uint, IDRest ...uint) GroupQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDLt(ID uint) GroupQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDLte(ID uint) GroupQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDNe(ID uint) GroupQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) IDNotIn(ID uint, IDRest ...uint) GroupQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) Limit(limit int) GroupQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NameEq is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) NameEq(name string) GroupQuerySet {
	return qs.w(qs.db.Where("name = ?", name))
}

// NameIn is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) NameIn(name string, nameRest ...string) GroupQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name IN (?)", iArgs))
}

// NameNe is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) NameNe(name string) GroupQuerySet {
	return qs.w(qs.db.Where("name != ?", name))
}

// NameNotIn is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) NameNotIn(name string, nameRest ...string) GroupQuerySet {
	iArgs := []interface{}{name}
	for _, arg := range nameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("name NOT IN (?)", iArgs))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs GroupQuerySet) One(ret *Group) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderAscByCreatedAt() GroupQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderAscByDeletedAt() GroupQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderAscByID() GroupQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByOwnerID is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderAscByOwnerID() GroupQuerySet {
	return qs.w(qs.db.Order("owner_id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderAscByUpdatedAt() GroupQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderDescByCreatedAt() GroupQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderDescByDeletedAt() GroupQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderDescByID() GroupQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByOwnerID is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderDescByOwnerID() GroupQuerySet {
	return qs.w(qs.db.Order("owner_id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OrderDescByUpdatedAt() GroupQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OwnerIDEq is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDEq(ownerID uint) GroupQuerySet {
	return qs.w(qs.db.Where("owner_id = ?", ownerID))
}

// OwnerIDGt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDGt(ownerID uint) GroupQuerySet {
	return qs.w(qs.db.Where("owner_id > ?", ownerID))
}

// OwnerIDGte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDGte(ownerID uint) GroupQuerySet {
	return qs.w(qs.db.Where("owner_id >= ?", ownerID))
}

// OwnerIDIn is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDIn(ownerID uint, ownerIDRest ...uint) GroupQuerySet {
	iArgs := []interface{}{ownerID}
	for _, arg := range ownerIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("owner_id IN (?)", iArgs))
}

// OwnerIDLt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDLt(ownerID uint) GroupQuerySet {
	return qs.w(qs.db.Where("owner_id < ?", ownerID))
}

// OwnerIDLte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDLte(ownerID uint) GroupQuerySet {
	return qs.w(qs.db.Where("owner_id <= ?", ownerID))
}

// OwnerIDNe is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDNe(ownerID uint) GroupQuerySet {
	return qs.w(qs.db.Where("owner_id != ?", ownerID))
}

// OwnerIDNotIn is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) OwnerIDNotIn(ownerID uint, ownerIDRest ...uint) GroupQuerySet {
	iArgs := []interface{}{ownerID}
	for _, arg := range ownerIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("owner_id NOT IN (?)", iArgs))
}

// PreloadCounters is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) PreloadCounters() GroupQuerySet {
	return qs.w(qs.db.Preload("Counters"))
}

// SetCounters is an autogenerated method
// nolint: dupl
func (u GroupUpdater) SetCounters(counters GroupCounters) GroupUpdater {
	u.fields[string(GroupDBSchema.Counters)] = counters
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u GroupUpdater) SetCreatedAt(createdAt time.Time) GroupUpdater {
	u.fields[string(GroupDBSchema.CreatedAt)] = createdAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u GroupUpdater) SetID(ID uint) GroupUpdater {
	u.fields[string(GroupDBSchema.ID)] = ID
	return u
}

// SetName is an autogenerated method
// nolint: dupl
func (u GroupUpdater) SetName(name string) GroupUpdater {
	u.fields[string(GroupDBSchema.Name)] = name
	return u
}

// SetOwnerID is an autogenerated method
// nolint: dupl
func (u GroupUpdater) SetOwnerID(ownerID uint) GroupUpdater {
	u.fields[string(GroupDBSchema.OwnerID)] = ownerID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u GroupUpdater) SetUpdatedAt(updatedAt time.Time) GroupUpdater {
	u.fields[string(GroupDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u GroupUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) UpdatedAtEq(updatedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) UpdatedAtGt(updatedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) UpdatedAtGte(updatedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) UpdatedAtLt(updatedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) UpdatedAtLte(updatedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs GroupQuerySet) UpdatedAtNe(updatedAt time.Time) GroupQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ===== END of query set GroupQuerySet

// ===== BEGIN of Group modifiers

type groupDBSchemaField string

// GroupDBSchema stores db field names of Group
var GroupDBSchema = struct {
	ID        groupDBSchemaField
	CreatedAt groupDBSchemaField
	UpdatedAt groupDBSchemaField
	DeletedAt groupDBSchemaField
	Name      groupDBSchemaField
	OwnerID   groupDBSchemaField
	Counters  groupDBSchemaField
}{

	ID:        groupDBSchemaField("id"),
	CreatedAt: groupDBSchemaField("created_at"),
	UpdatedAt: groupDBSchemaField("updated_at"),
	DeletedAt: groupDBSchemaField("deleted_at"),
	Name:      groupDBSchemaField("name"),
	OwnerID:   groupDBSchemaField("owner_id"),
	Counters:  groupDBSchemaField("counters"),
}

// Update updates Group fields by primary key
func (o *Group) Update(db *gorm.DB, fields ...groupDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"name":       o.Name,
		"owner_id":   o.OwnerID,
		"counters":   o.Counters,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := string(f)
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Group %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// GroupUpdater is an Group updates manager
type GroupUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewGroupUpdater creates new Group updater
func NewGroupUpdater(db *gorm.DB) GroupUpdater {
	return GroupUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Group{}),
	}
}

// ===== END of Group modifiers

// ===== END of all query sets
