// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bootstrap

import (
	"github.com/fatih/color"
	"github.com/zhuharev/hyip/models"

	"github.com/zhuharev/hyip/pkg/setting"

	"github.com/zhuharev/hyip/pkg/exchange"
	ps "github.com/zhuharev/hyip/pkg/payment_system"
	"github.com/zhuharev/hyip/pkg/payment_system/adapters/advcash"
)

func initPaymentSystems() (err error) {
	pss, err := ps.New()
	if err != nil {
		return
	}

	stor, err := models.NewPaymentSystemStore(models.Advcash)
	if err != nil {
		return
	}

	adv := advcash.New(setting.App.PaymentSystems.AdvcashEmail,
		setting.App.PaymentSystems.AdvcashAPIName,
		setting.App.PaymentSystems.AdvcashAPIPassword, stor)

	adv.AmountConverter = exchange.DefaultTypeConverter

	color.Green("[payment systems] inited advcash ps %s %s %s", setting.App.PaymentSystems.AdvcashEmail,
		setting.App.PaymentSystems.AdvcashAPIName,
		setting.App.PaymentSystems.AdvcashAPIPassword)

	pss.Add(adv)

	go pss.Run()

	return
}
