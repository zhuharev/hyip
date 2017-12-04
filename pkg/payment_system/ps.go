// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ps

import (
	"log"
	"time"

	"github.com/asdine/storm"
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/payment_system/store"

	"github.com/fatih/color"
)

var (
	interval = 10 * time.Minute
)

// PaymentSystem iface of payment system
type PaymentSystem interface {
	// PollEnabled if poll disabled, new transactions not be fetched .
	// It used if payment sustem use webhooks form getting new transactions
	PollEnabled() bool
	FetchNewTransactions() ([]store.Transaction, error)
	// used for webhook
	//ConvertToTxn(interface{}) (store.Transactioner, error)
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

// Add ps to internal slice
func (pss *PaymentSystems) Add(ps PaymentSystem) {
	pss.systems = append(pss.systems, ps)
}

// Run payment systems poller. It fetch new txn from PS api by given interval
func (pss *PaymentSystems) Run() {
	for {
		for _, ps := range pss.systems {
			if ps.PollEnabled() {
				color.Cyan("[payment systems] fetch updates %v", ps)
				transactins, err := ps.FetchNewTransactions()
				if err != nil {
					color.Red("Error fetching transactions: %s", err)
					continue
				}
				for _, txn := range transactins {
					//if has, err := ps.HasTxn(txn.ExternalID); !has && err == nil {
					color.Green("Create new txn, %s", txn.ExternalID)
					err = ps.CreateTxn(&txn)
					if err != nil {
						color.Red("Error creating txn: %s", err)
					}
					// TODO: create investment here
					// } else {
					// 	if err != nil {
					// 		color.Red("Error checking txn is exists: %s", err)
					// 	}
					// }

					us, err := models.UserSettings.GetByField("Advcash", txn.SenderWalletID)
					if err != nil {
						if err == storm.ErrNotFound {
							color.Green("New transaction without sender! %d", txn.ID)
							continue
						}
						log.Printf("[payment systems] error detecting sender: %s", err)
						continue
					}

					user, err := models.Users.Get(us.UserID)
					if err != nil {
						log.Printf("[payment systems] error getting user by id: %s", err)
						continue
					}

					plan, err := models.Plans.GetByAmount(txn.Amount, models.USD.ID)
					if err != nil {
						log.Printf("[payment systems] error getting plan for payment currency: %s", err)
						continue
					}

					_, err = models.NewInvestment(plan.ID, user.ID, txn.Amount, txn.ID)
					if err != nil {
						log.Printf("[payment systems] error creating investment: %s", err)
						continue
					}
				}

			}
		}
		time.Sleep(interval)
	}
}
