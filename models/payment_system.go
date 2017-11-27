// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"

	"github.com/zhuharev/boltutils"
)

type PaymentSystemType string

const (
	Bitcoin PaymentSystemType = "btc"
	Qiwi                      = "qiwi"
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
