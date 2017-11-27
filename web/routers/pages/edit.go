// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pages

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
)

// Edit is pages.index controller
func Edit(ctx *context.Context) {
	var (
		title = ctx.Query("title")
		body  = ctx.Query("body")
		slug  = ctx.Params(":slug")
	)

	if ctx.Req.Method == "POST" {
		if ctx.HasError(models.Pages.Save(slug, title, body), "/about/"+slug+"/edit") {
			return
		}
	}

	page, _ := models.Pages.Get(slug)
	ctx.Data["page"] = page
	ctx.HTML(200, "pages/edit")
}
