// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/web/context"
	paymentSystem "github.com/zhuharev/hyip/web/routers/admin/payment_system"
	"github.com/zhuharev/hyip/web/routers/auth"
	"github.com/zhuharev/hyip/web/routers/dash"
	"github.com/zhuharev/hyip/web/routers/pages"
	"github.com/zhuharev/hyip/web/routers/plans"
	"github.com/zhuharev/hyip/web/routers/support"
	"github.com/zhuharev/hyip/web/routers/tools"
	"github.com/zhuharev/hyip/web/routers/webhooks"

	"github.com/go-macaron/binding"

	"gopkg.in/macaron.v1"
)

func registreRoutes(m *macaron.Macaron) {
	m.Get("/about/:slug", pages.Show)
	m.Post("/about/:slug", pages.Update)
	m.Any("/about/:slug/edit", pages.Edit)

	m.Get("/", context.Toggle(&context.ToggleOptions{SignOutRequired: true}), pages.Index)
	m.Post("/tools/calc", tools.Calc)

	m.Group("/dash", func() {
		m.Get("/", dash.Index)
		m.Get("/partners", dash.Partners)
		m.Get("/contracts", dash.Contracts)
		m.Get("/transactions", dash.Transactions)
		m.Get("/settings", dash.Settings)
		m.Post("/settings", dash.UpdateSettings)
		m.Post("/settings/plans/create", plans.Create)
		m.Post("/settings/advcash", dash.MakeSaveField("Advcash"))

		m.Get("/invest", dash.ChoosePaymentSystem)
	}, context.Toggle(&context.ToggleOptions{SignInRequired: true}))

	m.Get("/support", support.Index)
	m.Post("/support/tickets/create", support.CreateTicket)
	m.Get("/support/tickets/:ticketID", support.Ticket)
	m.Post("/support/messages/send", support.SendMessage)
	m.Get("/support/admin", support.Admin)
	m.Any("/login", context.Toggle(&context.ToggleOptions{SignOutRequired: true}), auth.Login)
	m.Any("/reg", context.Toggle(&context.ToggleOptions{SignOutRequired: true}), auth.Reg)
	m.Get("/logout", context.Toggle(&context.ToggleOptions{SignInRequired: true}), auth.Logout)

	m.Get("/r/:hash", context.Toggle(&context.ToggleOptions{SignOutRequired: true}), auth.Invite)

	m.Group("/admin", func() {
		m.Group("/ps", func() {
			m.Get("/", paymentSystem.List)
			m.Post("/create", binding.Bind(models.PaymentSystem{}), paymentSystem.Create)
		})
	}, context.Toggle(&context.ToggleOptions{SignInRequired: true, AdminRequired: true}))

	m.Post("/external/webhooks/qiwi", binding.Bind(models.QiwiTxn{}), webhooks.Qiwi)
}
