// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in plan.go

// Plan investments
// easyjson:json
// gen:qs
type Plan struct {
	gorm.Model

	Name                string
	Currency            uint
	MinInvestmentAmount uint
	Duration            time.Duration
	Profit              float64
}

func (p Plan) String() string {
	if p.Name != "" {
		return p.Name
	}
	return fmt.Sprintf("#%d от %s (%f)", p.ID, p.FormatMinInvestmentAmount(), p.Profit)
}

// FormatMinInvestmentAmount formats min investment amount
func (p Plan) FormatMinInvestmentAmount() string {
	return GetCurrency(p.Currency).FormatAmount(p.MinInvestmentAmount)
}
