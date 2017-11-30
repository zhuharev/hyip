// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package support

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// SendMessage sends message to ticket
func SendMessage(ctx *context.Context) {
	var (
		ticketID = uint(ctx.QueryInt("ticket_id"))
	)

	_, err := models.Messages.Create(ticketID, ctx.Query("body"), ctx.User.ID)
	if ctx.HasError(err) {
		return
	}

	ctx.Flash.Success("Сообщение отправлено")
	ctx.Redirect("/support/tickets/" + ctx.Query("ticket_id"))
}
