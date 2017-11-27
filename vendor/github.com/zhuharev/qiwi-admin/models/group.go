// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in group.go

package models

import "github.com/jinzhu/gorm"

// Group represent group of wallets
// gen:qs
type Group struct {
	gorm.Model

	Name    string
	OwnerID uint

	Counters GroupCounters `gorm:"-"`
}

// CreateGroup save new group in db
func CreateGroup(name string, ownerID uint) (group *Group, err error) {
	group = new(Group)
	group.Name = name
	group.OwnerID = ownerID

	err = group.Create(db)

	return
}

// GetUserGroups return groups where user has access
func GetUserGroups(userID uint) (res []Group, err error) {
	err = NewGroupQuerySet(db).OwnerIDEq(userID).All(&res)
	return
}

// GetGroup returns group by id
func GetGroup(id uint, userIDs ...uint) (group *Group, err error) {
	group = new(Group)
	query := NewGroupQuerySet(db).IDEq(id)
	if len(userIDs) > 0 {
		query = query.OwnerIDEq(userIDs[0])
	}
	err = query.One(group)
	if err != nil {
		return
	}

	counters, err := GetGroupCounters(id)
	group.Counters = counters

	return
}

// GroupCounters response of aggregate sql request
type GroupCounters struct {
	Balance float64 `gorm:"balance"`
	Count   int     `gorm:"count"`
}

// GetGroupCounters agregate stat
func GetGroupCounters(groupID uint) (res GroupCounters, err error) {
	sql := `select sum(wallets.balance) as balance,count() as count from wallets where "wallets"."deleted_at" IS NULL AND group_id = ?`
	err = db.Raw(sql, groupID).Scan(&res).Error
	return
}
