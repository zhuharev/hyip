// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in investments.go

// Investment table when stored all users investments
// gen:qs
type Investment struct {
	gorm.Model

	PlanID uint
	UserID uint

	Amount uint

	TxnID uint `gorm:"unique_index"`
}

// NewInvestment create investment in db
func NewInvestment(planID, userID, amount, txnID uint) (*Investment, error) {
	inv := &Investment{
		PlanID: planID,
		UserID: userID,
		Amount: amount,
		TxnID:  txnID,
	}
	err := inv.Create(db)
	return inv, err
}

type InvestmentView struct {
	PlanID    uint
	Currency  uint
	Amount    uint
	CreatedAt time.Time
	UserID    uint
	Profit    float64

	Curr Currency `gorm:"-"`
}

// GetInvestments retuns user investments
func GetInvestments(userID uint) (res []InvestmentView, err error) {
	sql := `select plans.id as plan_id, plans.currency, investments.amount, investments.created_at,
	 users.id as user_id, plans.profit from plans, investments, users where plans.id = investments.plan_id and users.id = investments.user_id and users.id = ?`
	err = db.Raw(sql, userID).Scan(&res).Error
	if err != nil {
		return
	}

	for i, plan := range res {
		res[i].Curr = GetCurrency(plan.Currency)
	}

	return
}
