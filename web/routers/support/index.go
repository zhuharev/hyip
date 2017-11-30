// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package support

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// Index is support.index controller
func Index(ctx *context.Context) {
	if ctx.Autorized() {
		tickets, err := models.Tickets.List(ctx.User.ID)
		if ctx.HasError(err) {
			return
		}
		ctx.Data["tickets"] = tickets
	}

	ctx.HTML(200, "support/index")
}
