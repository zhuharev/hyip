// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in txn.go

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TxnType of txn
type TxnType uint

const (
	// In incominc txn
	In TxnType = iota + 1
	// Out outgoing txn
	Out
	// QiwiCard payment from Qiwi card
	QiwiCard
)

// Txn qiwi transaction
// gen:qs
type Txn struct {
	gorm.Model

	QiwiTxnID uint `gorm:"unique_index:idx_qiwi_txn"`

	TxnType       TxnType `gorm:"unique_index:idx_qiwi_txn"`
	ProviderID    uint    // ?
	Amount        float64
	QiwiCreatedAt time.Time `gorm:"index"`
	Fee           float64
	Status        Status

	Comment string

	WalletID uint `gorm:"unique_index:idx_qiwi_txn"`
}

// Status represent status of txn
type Status uint

const (
	// Waiting txn created but not processed
	Waiting Status = iota + 1
	// Success txn
	Success
	// Error represent txn with an error
	Error
)

// CreateMultipleTxns insert transactions in on txn
func CreateMultipleTxns(walletID uint, txns []Txn) (err error) {
	tx := db.Begin()
	for _, txn := range txns {
		txn.WalletID = walletID
		// ignore this errors
		_ = tx.Create(&txn).Error
	}
	err = tx.Commit().Error
	return
}

// GetWalletTxns get lasts wallet txns
func GetWalletTxns(walletID uint) (res []Txn, err error) {
	err = NewTxnQuerySet(db).WalletIDEq(walletID).OrderDescByQiwiTxnID().Limit(50).All(&res)
	return
}

// GetLastQiwiTxn return last txn on wallet with walletID
func GetLastQiwiTxn(walletID uint) (txnID uint, err error) {
	txn := new(Txn)
	err = NewTxnQuerySet(db).WalletIDEq(walletID).OrderDescByQiwiTxnID().One(txn)
	if err != nil {
		return
	}
	txnID = txn.QiwiTxnID
	return
}
