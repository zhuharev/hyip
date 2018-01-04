// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paymentSystem

import (
	"fmt"
	"log"

	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// Create creates ps
func Create(ctx *context.Context, ps models.PaymentSystem) {
	log.Printf("%+v", ps)
	err := models.CreatePaymentSystem(&ps)
	if ctx.HasError(err) {
		return
	}
	ctx.Flash.Success("Платёжная система добавлена")
	ctx.Redirect(fmt.Sprintf("/admin/ps/%d", ps.ID))
}

// List show all payment systems
func List(ctx *context.Context) {
	list, err := models.GetPaymentSystems()
	if ctx.HasError(err) {
		return
	}
	ctx.Data["list"] = list
	ctx.HTML(200, "admin/payment_system/list")
}
