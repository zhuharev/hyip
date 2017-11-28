// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package advcash

import (
	"encoding/json"
	"time"

	"github.com/zhuharev/hyip/pkg/payment_system/store"
)

// Advcash ps manager
type Advcash struct {
	email       string
	apiPassword string
	apiName     string
	//usdWallet   string

	store.Store
	store.AmountConverter
}

// New returns new instance of advcash
func New(email, apiName, apiPassword string, stor store.Store) (a *Advcash) {
	a = new(Advcash)
	a.email = email
	a.apiPassword = apiPassword
	a.apiName = apiName
	a.Store = stor
	return
}

// PollEnabled represent what this PS needed update transactions via periodic
// http requests
func (a *Advcash) PollEnabled() bool {
	return true
}

// FetchNewTransactions returns last 10 transactions
// It need check is transactions new
func (a *Advcash) FetchNewTransactions() (res []store.Transaction, err error) {

	var lastFetchTime time.Time

	started := time.Now()

	state, err := a.LoadState()
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

	bts, err := json.Marshal(started)
	if err != nil {
		return
	}

	err = a.SaveState(bts)

	return
}
