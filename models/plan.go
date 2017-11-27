// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in plan.go

// Plan investments
// easyjson:json
// gen:qs
type Plan struct {
	gorm.Model

	Currency            uint
	MinInvestmentAmount uint
	Duration            time.Duration
	Profit              float64
}

// FormatMinInvestmentAmount formats min investment amount
func (p Plan) FormatMinInvestmentAmount() string {
	return GetCurrency(p.Currency).FormatAmount(p.MinInvestmentAmount)
}
