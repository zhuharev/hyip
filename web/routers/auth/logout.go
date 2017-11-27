// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import "pure/api/socs/telegram/trinity/web/context"

// Logout delete req cookies and session user_id
func Logout(ctx *context.Context) {
	ctx.Session.Delete("user_id")
	ctx.SetCookie("s", "")
	ctx.Redirect("/login")
}
