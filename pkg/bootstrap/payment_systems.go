// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bootstrap

import ps "github.com/zhuharev/hyip/pkg/payment_system"

func initPaymentSystems() (err error) {
	pss, err := ps.New()
	if err != nil {
		return
	}

	// stor, err := models.NewPaymentSystemStore(models.Advcash)
	// if err != nil {
	// 	return
	// }

	// for _, psSetting := range setting.App.PaymentSystems {
	// 	if psSetting.Enabled && psSetting.Name == "advcash" {
	// 		adv := advcash.New(psSetting.WalletID,
	// 			psSetting.APIName,
	// 			psSetting.APISecret, stor)
	//
	// 		adv.AmountConverter = exchange.DefaultTypeConverter
	//
	// 		color.Green("[payment systems] inited advcash ps %s %s %s", psSetting.WalletID,
	// 			psSetting.APIName,
	// 			psSetting.APISecret)
	//
	// 		if psSetting.WalletID != "" {
	// 			pss.Add(adv)
	// 		}
	// 	}
	// }

	go pss.Run()

	return
}
