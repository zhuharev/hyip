// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in project_payment_system.go
package models

import "github.com/jinzhu/gorm"

type PaymentDirection string

const (
	PaymentDirectionIn   = "in"
	PaymentDirectionOut  = "out"
	PaymentDirectionBoth = "bi"
)

// PaymentSystem model for db
// easyjson:json
// gen:qs
type PaymentSystem struct {
	gorm.Model
	Key              string
	Secret           string
	AccountID        string            `form:"account_id"`
	PSType           PaymentSystemType `form:"ps_type" binding:"ps_type"`
	PaymentDirection PaymentDirection
	Currencies       string
	Direction        PaymentDirection
	Name             string
}

func (ps PaymentSystem) RefillAddress(accs []Account) string {
	switch ps.PSType {
	case Qiwi:
		return ps.AccountID
	case Coinbase:
		for _, acc := range accs {
			if acc.CurrencyID == BTC.ID {
				return acc.RefillWalletID
			}
		}
	case Advcash:
		return ps.AccountID
	}
	return ""
}

func (ps PaymentSystem) PSTypeName() string {
	switch ps.PSType {
	case Qiwi:
		return "Qiwi"
	case Coinbase:
		return "Биткоин"
	case Advcash:
		return "Advcash"
	}
	return ""
}

func CreatePaymentSystem(ps *PaymentSystem) (err error) {
	err = db.Create(ps).Error
	return
}

// GetPaymentSystems returns all payment system
func GetPaymentSystems() (pss []PaymentSystem, err error) {
	err = NewPaymentSystemQuerySet(db).All(&pss)
	return
}

// TODO:
// type PaymentSystemCurrencies struct {
// 	PaymentSystemID uint
// 	CurrencyID      uint
// }

// GetPaymentSystemsByCurrency returns payment systems for spec curr
func GetPaymentSystemsByCurrency(currID uint) (pss []PaymentSystem, err error) {
	qs := NewPaymentSystemQuerySet(db)
	switch currID {
	case RUB.ID:
		qs = qs.PSTypeEq(Qiwi)
	case USD.ID:
		qs = qs.PSTypeIn(Qiwi, Advcash)
	case BTC.ID:
		qs = qs.PSTypeEq(Coinbase)
	}
	err = qs.All(&pss)
	return
}
