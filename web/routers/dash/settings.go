// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
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

	setting, err := models.UserSettings.Get(ctx.User.ID)
	if ctx.HasError(err) {
		return
	}

	ctx.Data["userSetting"] = setting
	ctx.Data["currencies"] = models.Currencies
	ctx.HTML(200, "dash/settings")
}

// MakeSaveField returns handler func
func MakeSaveField(fieldName string) func(ctx *context.Context) {
	return func(ctx *context.Context) {
		var (
			value = ctx.Query(fieldName)
		)
		if ctx.HasError(models.UserSettings.SetField(ctx.User.ID, fieldName, value)) {
			return
		}
		ctx.Flash.Success("Настройки сохранены")
		ctx.Redirect("/dash/settings")
	}
}
