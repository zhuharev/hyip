// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

func ChoosePaymentSystem(ctx *context.Context) {
	pss, err := models.GetPaymentSystemsByCurrency(uint(ctx.QueryInt("currency_id")))
	if ctx.HasError(err) {
		return
	}

	accounts, err := models.Accounts.List(ctx.User.ID)
	if ctx.HasError(err) {
		return
	}

	ctx.Data["accounts"] = accounts
	ctx.Data["payment_systems"] = pss
	ctx.HTML(200, "dash/choose_payment_system")
}
