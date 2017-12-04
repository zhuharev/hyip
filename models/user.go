// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in user.go

// User is main struct
// easyjson:json
// gen:qs
type User struct {
	gorm.Model

	BalanceUSD float64
	BalanceBTC float64

	DepositUSD float64
	DepositBTC float64

	ReinvestUSD float64
	ReinvestBTC float64

	EarnUSD             float64
	EarnBTC             float64
	EarnWithPartnersUSD float64
	EarnWithPartnersBTC float64

	// 0 - russian
	// 1 = en
	Lang int

	Ref1 uint
	Ref2 uint
	Ref3 uint
	Ref4 uint
	Ref5 uint

	DepositWallet  string
	WithdrawWallet string

	Paid bool

	LastReceivedProfitAt time.Time

	Role Role

	Name           string `gorm:"unique_index"`
	HashedPassword []byte
}

// LangString returns user lang
func (u User) LangString() string {
	switch u.Lang {
	case 1:
		return "en-US"
	default:
		return "ru-RU"
	}
}
