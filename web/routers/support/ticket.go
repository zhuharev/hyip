// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package support

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// Ticket shows messages and other info about ticket with ticketID
func Ticket(ctx *context.Context) {
	var (
		ticketID = uint(ctx.ParamsInt64(":ticketID"))
	)
	ticket, err := models.Tickets.Get(ticketID)
	if ctx.HasError(err) {
		return
	}

	messages, err := models.Messages.List(ticketID)
	if ctx.HasError(err) {
		return
	}

	ctx.Data["ticket"] = ticket
	ctx.Data["messages"] = messages
	ctx.HTML(200, "support/tickets/index")
}
