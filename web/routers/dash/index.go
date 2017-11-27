// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// Index is dash.index controller
func Index(ctx *context.Context) {
	_ = models.User{}

	accounts, err := models.Accounts.List(ctx.User.ID)
	if err != nil {
		return
	}

	ctx.Data["currencies"] = models.Currencies
	ctx.Data["accounts"] = accounts
	ctx.HTML(200, "dash/index")
}
