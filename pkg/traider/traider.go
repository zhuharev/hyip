// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package traider

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/zhuharev/hyip/models"
)

var (
	ProfitPeriod = time.Minute
)

func NewContext() error {
	go func() {
		for {
			start := time.Now()
			err := trade()
			if err != nil {
				color.Red("[traider] trade failed: %s", err)
			}
			color.Green("[traider] Done %s", time.Since(start))
			time.Sleep(ProfitPeriod)
		}
	}()
	return nil
}

// TODO: optimize it
func trade() (err error) {
	investments, err := models.Investments.NeedsProfit()
	if err != nil {
		return
	}

	for _, invest := range investments {
		user, err := models.Users.Get(invest.UserID)
		if err != nil {
			return err
		}
		plan, err := models.Plans.Get(invest.PlanID)
		if err != nil {
			return err
		}

		accounts, err := models.Accounts.List(user.ID)
		if err != nil {
			return err
		}

		color.Cyan("Not a curr %+v", accounts)

		for _, acc := range accounts {
			if acc.CurrencyID == plan.Currency {
				log.Println("profit")
				days := uint(time.Since(invest.LastPayment).Hours() / 24.0)
				if days == 0 || invest.LastPayment.IsZero() {
					days = 1
				}
				profit := uint(float64(invest.Amount*days) * plan.Profit)
				color.Green("[traider] days: %d, invested:%d,plan_profit:%f profit: %d",
					days, invest.Amount, plan.Profit, profit)

				err = models.MakeProfit(acc.ID, profit, acc.CurrencyID, invest.ID, models.ProfitInvestment)
				if err != nil {
					return err
				}
				break
			} else {
				color.Cyan("Not a curr %+v %+v", acc, plan)
			}
		}

	}

	return
}
