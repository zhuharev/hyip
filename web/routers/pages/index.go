// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pages

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
	macaron "gopkg.in/macaron.v1"
)

// Index is pages.index controller
func Index(ctx *context.Context) {
	_ = models.User{}
	ctx.HTML(200, "pages/index", nil, macaron.HTMLOptions{Layout: ""})
}
