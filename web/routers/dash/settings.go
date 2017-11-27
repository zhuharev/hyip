// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"pure/api/socs/telegram/trinity/models"
	"pure/api/socs/telegram/trinity/web/context"
)

// Settings is dash.settings controller
func Settings(ctx *context.Context) {
	if ctx.User.IsAdmin() {
		plans, err := models.Plans.All()
		if err != nil {
			return
		}
		ctx.Data["plans"] = plans
	}

	ctx.Data["currencies"] = models.Currencies
	ctx.HTML(200, "dash/settings")
}
