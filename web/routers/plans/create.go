// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plans

import (
	"fmt"
	"log"
	"math"

	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// Create plan in database
func Create(ctx *context.Context) {
	var (
		currencyID    = ctx.QueryInt("currency")
		minimalAmount = ctx.QueryFloat64("amount")
		profit        = ctx.QueryFloat64("profit")
	)

	curr := models.GetCurrency(uint(currencyID))
	log.Println(curr)
	amount := int(math.Pow10(int(curr.Digits)) * minimalAmount)

	plan := &models.Plan{
		MinInvestmentAmount: uint(amount),
		Currency:            curr.ID,
		Profit:              profit,
	}

	err := models.Plans.CreatePlan(plan)
	if err != nil {
		ctx.Flash.Error(err.Error())
		ctx.Redirect("/dash/settings")
		return
	}

	ctx.Flash.Success(fmt.Sprintf("cur: %d, amount: %d", currencyID, amount))
	ctx.Redirect("/dash/settings")
}
