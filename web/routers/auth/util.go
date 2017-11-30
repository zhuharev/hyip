// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/zhuharev/hyip/web/context"
)

// Toggle returns handler, which redirect to redirectTo when user authorized or
// not.
func Toggle(redirectTo string, needAuth bool) func(*context.Context) {
	return func(ctx *context.Context) {
		if !ctx.Autorized() && needAuth {
			ctx.Redirect(redirectTo)
			ctx.Next()
			return
		}

		if ctx.Autorized() && !needAuth {
			ctx.Redirect(redirectTo)
			ctx.Next()
			return
		}
	}
}
