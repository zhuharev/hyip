// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package advcash

import "github.com/zhuharev/hyip/pkg/payment_system/store"
import "github.com/zhuharev/hyip/models"

// ConvertToTxn convert adv cash responsed txn to our txn
func (a *Advcash) ConvertToTxn(at AdvTransaction) (txn store.Transaction) {
	txn.CreatedAt = at.StartTime
	txn.CurrencyCode = models.USD.Code // at.Currency
	txn.ExternalID = at.Id
	txn.Incoming = at.Direction == "INCOMING"
	txn.PaymentSystemName = "advcash"
	txn.RecieverWalletID = at.ReceiverEmail
	txn.SenderWalletID = at.SenderEmail

	txn.Amount = a.ConvertToUint(models.USD.Code, at.AmountInUSD)
	return
}
