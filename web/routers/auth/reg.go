// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/base"
	"github.com/zhuharev/hyip/web/context"
)

// Reg is dash.index controller
func Reg(ctx *context.Context) {
	if ctx.Req.Method == "POST" {
		login := ctx.Query("login")
		pass := ctx.Query("password")
		passRepeat := ctx.Query("password_repeat")
		if pass != passRepeat {
			ctx.Flash.Error("Пароли не совпадают")
			ctx.Redirect("/reg")
			return
		}
		user := new(models.User)
		user.Name = login
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			ctx.Flash.Error("Ошибка пароля")
			ctx.Redirect("/reg")
			return
		}
		user.HashedPassword = hashedPassword

		refHash := ctx.GetCookie(refCookieName)
		if refHash != "" {
			refID := base.DecodeHash(refHash)
			user.Ref1 = uint(refID)
		}

		err = models.Users.Create(user)
		if err != nil {
			ctx.Flash.Error("Ошибка создания пользователя")
			ctx.Redirect("/reg")
			return
		}
		err = ctx.Session.Set("user_id", user.ID)
		if err != nil {
			ctx.Flash.Error("Ошибка сессии")
			ctx.Redirect("/reg")
			return
		}

		ctx.Redirect("/dash")
		return
	}

	ctx.HTML(200, "auth/reg")
}
