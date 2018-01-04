// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import "github.com/zhuharev/hyip/web/context"

const (
	refCookieName = "ref"
)

// Invite set ref cookie
func Invite(ctx *context.Context) {
	ctx.SetCookie(refCookieName, ctx.Params(":hash"), 60*60*24*365)
	ctx.Redirect("/")
}
