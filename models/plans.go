// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

type plansModel int

// Plans helper for plan methods
var Plans plansModel

func (p plansModel) All() (res []Plan, err error) {
	err = NewPlanQuerySet(db).OrderAscByCurrency().OrderAscByProfit().All(&res)
	return
}

// CreatePlan inserts new plan into database
func (p plansModel) CreatePlan(plan *Plan) (err error) {
	err = plan.Create(db)
	return
}

// GetPlanByAmount returns plan by provided amount
func (p plansModel) GetByAmount(amount uint, currency uint) (plan *Plan, err error) {
	plan = new(Plan)
	err = NewPlanQuerySet(db).
		MinInvestmentAmountLte(amount).
		OrderDescByMinInvestmentAmount().
		CurrencyEq(currency).
		One(plan)
	return
}
