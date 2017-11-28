// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type usersModel int

// Users helper for iterating with db
var Users usersModel

func (u usersModel) Get(id uint) (*User, error) {
	var user User
	err := NewUserQuerySet(db).IDEq(id).One(&user)
	return &user, err
}

func (u usersModel) GetByTelegramID(telegramID int64) (user *User, err error) {
	var bind = new(TelegramBind)
	err = NewTelegramBindQuerySet(db).TelegramIDEq(telegramID).One(bind)
	if err != nil {
		return
	}
	if bind.UserID == 0 {
		err = gorm.ErrRecordNotFound
		return
	}
	user = new(User)
	err = NewUserQuerySet(db).IDEq(bind.UserID).One(user)
	return user, err
}

func (u usersModel) Save(us *User, fields ...userDBSchemaField) error {
	return us.Update(db, fields...)
}

func (u usersModel) Create(user *User) (err error) {

	if user.Ref1 != 0 {
		refUser, err := Users.Get(user.Ref1)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				user.Ref1 = 0
			} else {
				return err
			}
		} else {
			user.Ref2 = refUser.Ref1
			user.Ref3 = refUser.Ref2
			user.Ref4 = refUser.Ref3
			user.Ref5 = refUser.Ref4
		}
	}

	err = user.Create(db)
	if err != nil {
		return err
	}

	if user.ID == 1 {
		user.Role = RoleAdmin
		err = user.Update(db, UserDBSchema.Role)
		if err != nil {
			return err
		}
	}

	var us = UserSetting{UserID: user.ID}
	err = stormDB.Save(&us)

	return
}

func (u usersModel) NeedsProfit(period time.Duration) (users []User, err error) {
	start := time.Now().Add(-period)
	err = NewUserQuerySet(db).w(db.Where("deposit_usd > 0 or deposit_btc > 0")).LastReceivedProfitAtLte(start).All(&users)
	return
}

const (
	BalanceUSD  = "balance_usd"
	BalanceBTC  = "balance_btc"
	DepositUSD  = "deposit_usd"
	ReinvestUSD = "reinvest_usd"
	DepositBTC  = "deposit_btc"
	ReinvestBTC = "reinvest_btc"
)

func Tx() *gorm.DB {
	return db.Begin()
}

func (u usersModel) Inc(fieldName string, amount float64, userID uint, txs ...*gorm.DB) error {
	tx := db
	if len(txs) > 0 {
		tx = txs[0]
	}
	var autoCommit = len(txs) == 0
	sqlfmt := `update users set %[1]s = %[1]s + ?, last_received_profit_at = ? where id = ?`
	sql := fmt.Sprintf(sqlfmt, fieldName)
	err := tx.Exec(sql, amount, time.Now(), userID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if autoCommit {
		return tx.Commit().Error
	}

	return err
}

func (u usersModel) Dec(fieldName string, amount float64, userID uint) error {
	sqlfmt := `update users set %[1]s = %[1]s - ?, last_received_profit_at = ? where id = ? and %[1]s - ? >= 0 `
	sql := fmt.Sprintf(sqlfmt, fieldName)
	err := db.Exec(sql, amount, time.Now(), userID, amount).Error
	return err
}

func (u usersModel) Move(fieldNameFrom, fieldNameTo string, userID uint) error {
	sqlfmt := `update users set %[2]s = %[2]s + %[1]s, %[1]s = 0 where id = ? and %[1]s <> 0 `
	sql := fmt.Sprintf(sqlfmt, fieldNameFrom, fieldNameTo)
	err := db.Exec(sql, userID).Error
	return err
}

func (u usersModel) AllReferrals(userID uint) (users []User, err error) {
	err = db.Where("ref1 = ?", userID).
		Or("ref2 = ?", userID).
		Or("ref3 = ?", userID).
		Or("ref4 = ?", userID).
		Or("ref5 = ?", userID).
		Find(&users).Error
	return
}
