// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package advcash

import (
	"encoding/json"
	"github.com/zhuharev/hyip/pkg/payment_system/store"
	"time"
)

// Advcash ps manager
type Advcash struct {
	email       string
	apiPassword string
	apiName     string
	//usdWallet   string

	store store.Store
	store.AmountConverter
}

// New returns new instance of advcash
func New(email, apiName, apiPassword string) (a *Advcash) {
	a = new(Advcash)
	a.email = email
	a.apiPassword = apiPassword
	a.apiName = apiName
	return
}

// FetchNewTransactions returns last 10 transactions
// It need check is transactions new
func (a *Advcash) FetchNewTransactions() (res []store.Transaction, err error) {

	var lastFetchTime time.Time

	state, err := a.store.LoadState()
	if err != nil {
		return
	}

	// skip error handling
	// if lastFetchtime zero value
	// api just recieve last 10 txns
	json.Unmarshal(state, &lastFetchTime)

	atxns, err := a.hist(lastFetchTime)
	if err != nil {
		return
	}
	for _, atxn := range atxns {
		txn := a.ConvertToTxn(atxn)
		if lastFetchTime.IsZero() || lastFetchTime.Before(txn.CreatedAt) {
			res = append(res, txn)
		}
	}
	return
}
