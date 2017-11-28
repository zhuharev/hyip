// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"

	"github.com/zhuharev/boltutils"
	"github.com/zhuharev/hyip/pkg/payment_system/store"
)

type PaymentSystemType string

const (
	Coinbase PaymentSystemType = "coinbase"
	Qiwi                       = "qiwi"
	Advcash                    = "advcash"
)

func paymentSystemMakeKey(userID uint, typ PaymentSystemType) string {
	return fmt.Sprintf("%s_%d", typ, userID)
}

func PaymentSystemSave(userID uint, typ PaymentSystemType, walletID string) error {
	key := paymentSystemMakeKey(userID, typ)
	return SaveValue(key, []byte(walletID))
}

func PaymentSystemGet(userID uint, typ PaymentSystemType) (string, error) {
	key := paymentSystemMakeKey(userID, typ)
	bts, err := GetValue(key)
	if err != nil {
		return "", err
	}
	if len(bts) == 0 {
		return "", boltutils.ErrNotFound
	}
	return string(bts), err
}

type simpleStore string

func (s simpleStore) CreateTxn(txn *store.Transaction) error {
	return txn.Create(db)
}

var (
	psStorePrefix = "ps_store_"
)

func makePsStoreKey(psName string) string {
	return psStorePrefix + psName
}

func (s simpleStore) SaveState(data []byte) error {
	return SaveValue(makePsStoreKey(string(s)), data)
}

func (s simpleStore) LoadState() ([]byte, error) {
	data, err := GetValue(makePsStoreKey(string(s)))
	if err != nil {
		if err == boltutils.ErrNotFound {
			return []byte("null"), nil
		}
		return nil, err
	}
	return data, nil
}

// NewPaymentSystemStore returns store.Store
func NewPaymentSystemStore(psName PaymentSystemType) (store.Store, error) {
	return simpleStore(psName), nil
}
