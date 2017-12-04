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

type InvestmentWithLastPayment struct {
	Investment
	LastPayment time.Time
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

type investmentsModel int

// Investments tool models calling Investments api
// TODO: check investment duration
var Investments investmentsModel

func (investmentsModel) NeedsProfit() (res []InvestmentWithLastPayment, err error) {
	var (
		// todo move this
		profitPeriod  = 1 * time.Minute
		needUpdatedAt = time.Now().Add(-profitPeriod)
	)

	sql := `SELECT i.id, i.created_at, i.updated_at, i.deleted_at, i.plan_id, i.user_id, i.amount, i.txn_id, p.created_at as last_payment, max(p.id) as noop
          FROM investments i
					LEFT JOIN profits p
					ON i.id = p.object_id
					WHERE (p.profit_type = ?
					  OR p.profit_type IS NULL)
					  AND i.deleted_at IS NULL
					  AND (p.created_at <= ?
							OR p.created_at IS NULL
						)
					GROUP BY i.id
						`

	err = db.Raw(sql, ProfitInvestment, needUpdatedAt).Scan(&res).Error
	if err != nil {
		return
	}

	return
}
