// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in profit.go

package models

import "github.com/jinzhu/gorm"

type ProfitType int

const (
	ProfitInvestment ProfitType = iota + 1
	ProfitBonus
	ProfitReferralBonus
)

// Profit represent model of users profit for logging all trader payment
// gen:qs
type Profit struct {
	gorm.Model

	ProfitType ProfitType

	ObjectID uint
	Amount   uint
	Currency uint
}

// MakeProfit create Profit with ProfitType objectType, ObjectID objectID,
// Amount amount, Currency currency.
// Also increase user account with accountID by amount
// All changes in single transaction. If an error will happen, transaction will
// rollback
func MakeProfit(accountID uint, amount uint, currency uint, objectID uint, objectType ProfitType) (err error) {
	txn := db.Begin()
	defer txn.Rollback()

	profit := &Profit{
		ProfitType: objectType,
		ObjectID:   objectID,
		Amount:     amount,
		Currency:   currency,
	}

	err = profit.Create(txn)
	if err != nil {
		return
	}

	sql := `UPDATE accounts
	SET amount = amount + ?
	WHERE id = ?`
	err = txn.Exec(sql, amount, accountID).Error
	if err != nil {
		return
	}

	return txn.Commit().Error
}

type profitModel int

var Profits profitModel

func (profitModel) Lasts(userID uint) (profits []Profit, err error) {
	sql := `SELECT *
  FROM profits
  WHERE id
  IN (
    SELECT id
    FROM investments
    WHERE user_id = ?
  )
  AND profit_type = ?
  `
	err = db.Raw(sql, userID, ProfitInvestment).Scan(&profits).Error
	return
}
