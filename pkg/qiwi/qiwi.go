// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package qiwi

import (
	"fmt"
	"net/url"
)

var (
//	client *qiwi.Client

//lastTxnKVKey = "last_txn"
)

// func NewContext() (err error) {
// 	if len(setting.App.PaymentSystems.QiwiTokens) == 0 {
// 		return nil
// 	}
// 	color.Cyan("Init qiwi wallet with token %s", setting.App.PaymentSystems.QiwiTokens[0])
// 	client = qiwi.New(setting.App.PaymentSystems.QiwiTokens[0])
// 	profile, err := client.Profile.Current()
// 	if err != nil {
// 		return err
// 	}
// 	color.Cyan("Qiwi authorized %d", profile.AuthInfo.PersonID)
// 	client.SetWallet(fmt.Sprint(profile.AuthInfo.PersonID))
//
// 	payments, err := getLastPayments()
// 	if err != nil {
// 		return err
// 	}
// 	color.Cyan("payments: %+v", payments)
//
// 	return
// }
//
// func getLastTxnID() (txnID int64) {
// 	models.UnmarshalValue(lastTxnKVKey, &txnID)
// 	return
// }
//
// func saveLastTxnID(txnID int64) error {
// 	return models.SaveValue(lastTxnKVKey, txnID)
// }
//
// func getLastPayments() (resp []Payment, err error) {
// 	ph, err := client.Payments.History(50, url.Values{"operation": {"IN"}})
// 	if err != nil {
// 		return
// 	}
// 	lastTxn := getLastTxnID()
// 	for _, payment := range ph.Data {
// 		userID := uint(com.StrTo(strings.Split(payment.Comment, " ")[0]).MustInt64())
// 		if userID == 0 || payment.TxnID <= lastTxn {
// 			continue
// 		}
// 		resp = append(resp, Payment{
// 			CreatedAt:     payment.Date,
// 			Amount:        payment.Sum.Amount,
// 			Currency:      payment.Sum.Currency,
// 			UserID:        userID,
// 			TransactionID: payment.TxnID,
// 		})
// 	}
// 	return
// }
//
// type Payment struct {
// 	CreatedAt     time.Time
// 	Amount        float64
// 	Currency      int
// 	UserID        uint
// 	TransactionID int64
// }

func MakePaymentURL(walletID string, userID uint) string {
	var params = url.Values{
		"extra['comment']": {fmt.Sprint(userID)},
		"extra['account']": {walletID},
	}
	return fmt.Sprintf("%s?%s", "https://qiwi.com/payment/form/99", params.Encode())
}
