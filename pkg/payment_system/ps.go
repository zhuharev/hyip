// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ps

import (
	"pure/api/socs/telegram/trinity/pkg/payment_system/store"
	"time"

	"github.com/fatih/color"
)

// PaymentSystem iface of payment system
type PaymentSystem interface {
	// PollEnabled if poll disabled, new transactions not be fetched .
	// It used if payment sustem use webhooks form getting new transactions
	PollEnabled() bool
	FetchNewTransactions() ([]store.Transaction, error)
	// used for webhook
	//ConvertToTxn(interface{}) (store.Transactioner, error)
	SetStore(store.Store)
	store.Store
	//SendMoney(walletID string, currencyCode string, amount uint) (store.Transactioner, error)
}

// PaymentSystems main struct
type PaymentSystems struct {
	systems []PaymentSystem
}

// New returns new instance of payment system
func New() (ps *PaymentSystems, err error) {
	ps = new(PaymentSystems)
	return
}

// Run payment systems poller. It fetch new txn from PS api by given interval
func (pss *PaymentSystems) Run() {
	for {
		for _, ps := range pss.systems {
			if ps.PollEnabled() {
				transactins, err := ps.FetchNewTransactions()
				if err != nil {
					color.Red("Error fetching transactions: %s", err)
					continue
				}
				for _, txn := range transactins {
					//if has, err := ps.HasTxn(txn.ExternalID); !has && err == nil {
					color.Green("Create new txn, %s", txn.ExternalID)
					err = ps.CreateTxn(txn)
					if err != nil {
						color.Red("Error creating txn: %s", err)
					}
					// TODO: create investment here
					// } else {
					// 	if err != nil {
					// 		color.Red("Error checking txn is exists: %s", err)
					// 	}
					// }
				}

			}
		}
		time.Sleep(1 * time.Minute)
	}
}
