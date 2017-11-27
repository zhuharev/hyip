// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"pure/api/socs/telegram/trinity/models"
	"pure/api/socs/telegram/trinity/web/context"
)

// Transactions is dash.transactions controller
func Transactions(ctx *context.Context) {
	_ = models.User{}
	ctx.HTML(200, "dash/transactions")
}
