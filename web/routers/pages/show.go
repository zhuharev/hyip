// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pages

import (
	"pure/api/socs/telegram/trinity/models"
	"pure/api/socs/telegram/trinity/web/context"
)

// Show is pages.show controller
func Show(ctx *context.Context) {
	slug := ctx.Params(":slug")
	page, _ := models.Pages.Get(slug)

	ctx.Data["slug"] = slug
	ctx.Data["page"] = page
	ctx.HTML(200, "pages/show")
}
