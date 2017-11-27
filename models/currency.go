// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in currency.go

// Currency fiat or cryptocurrency representation
// easyjson:json
// gen:qs
type Currency struct {
	gorm.Model

	Code   string
	Number uint

	// digits after the decimal separator
	Digits        uint
	MinimalAmount uint

	// used for i18n
	Message string

	Symbol      string
	SymbolRight bool
}

// FormatAmount returns floating string
func (c Currency) FormatAmount(amount uint) string {
	amoStr := strconv.FormatFloat(float64(amount)/math.Pow10(int(c.Digits)), 'f', -1, 32)
	if c.SymbolRight {
		return fmt.Sprintf("%s%s", amoStr, c.Symbol)
	}
	return fmt.Sprintf("%s%s", c.Symbol, amoStr)
}

var (
	// BTC is bitcoin currency
	BTC = Currency{
		Model:         gorm.Model{ID: 1},
		Code:          "XBT",
		Digits:        8,
		MinimalAmount: 13000,
		Message:       "bitcoin",

		Symbol: "BTC",
	}

	// RUB is russian ruble currency
	RUB = Currency{
		Model:         gorm.Model{ID: 2},
		Code:          "RUB",
		Digits:        2,
		MinimalAmount: 10000,
		Message:       "rub",
		Symbol:        "руб.",
		SymbolRight:   true,
	}

	// USD is russian ruble currency
	USD = Currency{
		Model:         gorm.Model{ID: 3},
		Code:          "USD",
		Digits:        2,
		MinimalAmount: 1000,
		Message:       "usd",
		Symbol:        "$",
	}
)

type currencies []Currency

func (c currencies) Message(currID uint) string {
	for _, v := range c {
		if v.ID == currID {
			return v.Message
		}
	}
	return ""
}

func (c currencies) FormatAmount(currID, amount uint) string {
	for _, v := range c {
		if v.ID == currID {
			return v.FormatAmount(amount)
		}
	}
	return ""
}

var (
	// Currencies list of currencies
	Currencies = currencies{
		BTC, USD,
	}
)

// GetCurrency returns currency
func GetCurrency(id uint) Currency {
	return Currencies[int(id)-1]
}
