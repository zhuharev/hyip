// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dash

import (
	"fmt"
	"path/filepath"

	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/base"
	"github.com/zhuharev/hyip/pkg/setting"
	"github.com/zhuharev/hyip/web/context"
)

// Partners is dash.partners controller
func Partners(ctx *context.Context) {
	var (
		inviteHash = base.HashNumber(int(ctx.User.ID))
		inviteURL  = fmt.Sprintf(filepath.Join(setting.App.Web.Base, "r", inviteHash))

		paidRefs int
	)

	refs, err := models.Users.AllReferrals(ctx.User.ID)
	if ctx.HasError(err) {
		return
	}

	for _, ref := range refs {
		if ref.Paid {
			paidRefs++
		}
	}

	ctx.Data["paidRefs"] = paidRefs
	ctx.Data["refsCount"] = len(refs)
	ctx.Data["refs"] = refs
	ctx.Data["inviteURL"] = inviteURL
	ctx.HTML(200, "dash/partners")
}
