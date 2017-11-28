// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in store.go

package store

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Transaction unified interface of txn
// gen:qs
type Transaction struct {
	gorm.Model //sql

	ExternalID        string `gorm:"unique_index:idx_uniq_global"`
	StartedAt         time.Time
	Amount            uint
	CurrencyCode      string
	PaymentSystemName string `gorm:"unique_index:idx_uniq_global"`
	Incoming          bool   `gorm:"unique_index:idx_uniq_global"`
	SenderWalletID    string
	RecieverWalletID  string
	Status            string
}

// Store db interface
type Store interface {
	CreateTxn(*Transaction) error
	//HasTxn(string) (bool, error)
	SaveState([]byte) error
	LoadState() ([]byte, error)
	//UpdateTxnStatus(externalTxnID, newStatus string) error
}

// AmountConverter convert float to int
type AmountConverter interface {
	ConvertToUint(currencyCode string, amount float64) uint
	ConvertToFloat(currencyCode string, amount uint) float64
}
