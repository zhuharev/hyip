// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"golang.org/x/crypto/bcrypt"

	"pure/api/socs/telegram/trinity/models"
	"pure/api/socs/telegram/trinity/web/context"
)

var (
	errStr = "Пользователя с таким именем или паролем не существует"
)

// Login is dash.index controller
func Login(ctx *context.Context) {
	if ctx.Req.Method == "POST" {
		login := ctx.Query("login")
		pass := ctx.Query("password")

		user := new(models.User)
		err := models.NewUserQuerySet(models.DB()).NameEq(login).One(user)
		if err != nil {
			ctx.Flash.Error(errStr)
			ctx.Redirect("/login")
			return
		}

		err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(pass))
		if err != nil {
			ctx.Flash.Error(errStr)
			ctx.Redirect("/login")
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
	ctx.HTML(200, "auth/login")
}
