// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"pure/api/socs/telegram/trinity/models"
	"pure/api/socs/telegram/trinity/pkg/qiwi"
	"pure/api/socs/telegram/trinity/web/context"
)

// Contracts is dash.partners controller
func Contracts(ctx *context.Context) {
	_ = models.User{}
	paymentURL := qiwi.MakePaymentURL("+79997651151", ctx.User.ID)
	ctx.Data["paymentURL"] = paymentURL

	investments, err := models.GetInvestments(ctx.User.ID)
	if ctx.HasError(err) {
		return
	}
	ctx.Data["investments"] = investments

	ctx.HTML(200, "dash/contracts")
}
