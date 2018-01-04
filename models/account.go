// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in account.go

package models

import "github.com/jinzhu/gorm"

// Account user balance for specific currency
// gen:qs
type Account struct {
	gorm.Model

	CurrencyID uint `gorm:"unique_index:curr_user_idx"`
	UserID     uint `gorm:"unique_index:curr_user_idx"`

	Amount uint

	WithdrawWalletID string
	RefillWalletID   string
}
