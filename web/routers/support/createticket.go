// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package support

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// CreateTicket is support.createticket controller
func CreateTicket(ctx *context.Context) {
	var (
		body     = ctx.Query("body")
		title    = ctx.Query("title")
		email    = ctx.Query("email")
		ownerIDs []uint
	)

	if ctx.Autorized() {
		ownerIDs = append(ownerIDs, ctx.User.ID)
	}

	_, err := models.Tickets.Create(title, body, email, ownerIDs...)
	if ctx.HasError(err) {
		return
	}

	ctx.Flash.Success("Тикет успешно создан!")
	ctx.Redirect("/support")
}
