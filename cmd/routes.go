// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"pure/api/socs/telegram/trinity/models"
	"pure/api/socs/telegram/trinity/web/routers/auth"
	"pure/api/socs/telegram/trinity/web/routers/dash"
	"pure/api/socs/telegram/trinity/web/routers/pages"
	"pure/api/socs/telegram/trinity/web/routers/plans"
	"pure/api/socs/telegram/trinity/web/routers/support"
	"pure/api/socs/telegram/trinity/web/routers/tools"
	"pure/api/socs/telegram/trinity/web/routers/webhooks"

	"github.com/go-macaron/binding"

	"gopkg.in/macaron.v1"
)

func registreRoutes(m *macaron.Macaron) {
	m.Get("/about/:slug", pages.Show)
	m.Post("/about/:slug", pages.Update)
	m.Any("/about/:slug/edit", pages.Edit)
	m.Get("/", pages.Index)
	m.Post("/tools/calc", tools.Calc)
	m.Get("/dash", dash.Index)
	m.Get("/dash/partners", dash.Partners)
	m.Get("/dash/contracts", dash.Contracts)
	m.Get("/dash/transactions", dash.Transactions)
	m.Get("/dash/settings", dash.Settings)
	m.Post("/dash/settings", dash.UpdateSettings)
	m.Post("/dash/settings/plans/create", plans.Create)
	m.Get("/support", support.Index)
	m.Post("/support", support.CreateTicket)
	m.Any("/login", auth.Login)
	m.Any("/reg", auth.Reg)
	m.Get("/logout", auth.Logout)

	m.Post("/external/webhooks/qiwi", binding.Bind(models.QiwiTxn{}), webhooks.Qiwi)
}
