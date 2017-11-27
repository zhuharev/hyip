// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package webhooks

import (
	"log"
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/exchange"
	"github.com/zhuharev/hyip/web/context"

	"github.com/Unknwon/com"
	"github.com/fatih/color"
)

// Qiwi handle webhook from qiwi-admin
func Qiwi(ctx *context.Context, payload models.QiwiTxn) {
	color.Green("Incoming webhook %v", payload)

	amountF, err := exchange.ConvertAmount(models.RUB.Code, models.USD.Code, payload.Amount)
	if err != nil {
		log.Printf("[webhook] Error converting amount from RUB to USD: %s",
			err)
		return
	}

	var (
		userID = uint(com.StrTo(payload.Comment).MustInt())
		amount = uint(amountF * 100.0)
	)

	user, err := models.Users.Get(userID)
	if err != nil {
		log.Printf("[webhook] Error getting user by ID from payment comment (%d): %s",
			userID,
			err)
		return
	}

	plan, err := models.Plans.GetByAmount(amount, models.USD.ID)
	if err != nil {
		log.Printf("[webhook] error getting plan for payment currency: %s", err)
		return
	}

	_, err = models.NewInvestment(plan.ID, user.ID, amount)
	if err != nil {
		log.Printf("[webhook] error creating investment: %s", err)
		return
	}

	ctx.WriteHeader(200)
}
