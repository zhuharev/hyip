// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package store

import "time"

// Transaction unified interface of txn
type Transaction struct {
	ID                uint //sql
	ExternalID        string
	CreatedAt         time.Time
	Amount            uint
	CurrencyCode      string
	PaymentSystemName string
	Incoming          bool
	SenderWalletID    string
	RecieverWalletID  string
	Status            string
}

// Store db interface
type Store interface {
	CreateTxn(Transaction) error
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
