// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/qiwi"
	"github.com/zhuharev/hyip/web/context"
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

	profits, err := models.Profits.Lasts(ctx.User.ID)
	if ctx.HasError(err) {
		return
	}
	ctx.Data["profits"] = profits

	ctx.HTML(200, "dash/contracts")
}
