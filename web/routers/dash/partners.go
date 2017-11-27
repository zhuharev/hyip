// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// Partners is dash.partners controller
func Partners(ctx *context.Context) {
	_ = models.User{}
	ctx.HTML(200, "dash/partners")
}
