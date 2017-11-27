// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in wallet.go

package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/now"
)

// Wallet qiwi wallet credentials and setting
// gen:qs
type Wallet struct {
	gorm.Model

	Name string
	// Phone number
	WalletID    uint64
	Blocked     bool
	Token       string `gorm:"index"`
	TokenExpiry time.Time

	Balance float64
	Limit   uint

	TotalMonthIncoming float64
	TotalMonthOutgoing float64
	TotalSynced        time.Time

	// userID
	OwnerID uint

	GroupID uint

	WalletCounters WalletCounters `gorm:"-"`
}

func (w Wallet) String() string {
	if w.Name == "" {
		return fmt.Sprintf("+%d", w.WalletID)
	}
	return w.Name
}

// GroupWallets returns all group wallet
func GroupWallets(groupID uint) (res []Wallet, err error) {
	err = NewWalletQuerySet(db).GroupIDEq(groupID).All(&res)
	return
}

// GetAllWallets returns all wallets. Used for synchronizer
func GetAllWallets() (res []Wallet, err error) {
	err = NewWalletQuerySet(db).All(&res)
	return
}

// GetWallet returns wallet by their ID
func GetWallet(walletID uint, userIDs ...uint) (wallet *Wallet, err error) {
	wallet = new(Wallet)
	query := NewWalletQuerySet(db).IDEq(walletID)
	if len(userIDs) > 0 {
		query = query.OwnerIDEq(userIDs[0])
	}
	err = query.One(wallet)
	if err != nil {
		return
	}

	counters, err := GetWalletCounters(walletID)
	wallet.WalletCounters = counters

	return
}

// CreateWallet create an wallet
func CreateWallet(wallet *Wallet) (err error) {
	err = wallet.Create(db)
	if err != nil {
		return
	}

	return
}

// WalletCounters special counters (stat) for wallet
type WalletCounters struct {
	TodayTxnCount int     `gorm:"column:count"`
	TodayTxnSum   float64 `gorm:"column:sum"`
}

// GetWalletCounters returns wallet stat
func GetWalletCounters(walletID uint) (wc WalletCounters, err error) {
	sql := `select sum(txns.amount) as sum, count() as count from txns where wallet_id = ? and datetime(created_at) >= datetime(?) and txn_type = ?`
	err = db.Raw(sql, walletID, now.BeginningOfDay(), In).Scan(&wc).Error
	log.Println(wc)
	return
}
