// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in app.go

package models

import "github.com/jinzhu/gorm"

// AppLevel type of app
type AppLevel int

const (
	// AppWallet app setting affect only on wallet
	AppWallet AppLevel = iota + 1
	// AppGroup setting affect all wallets in group
	AppGroup
	// AppAccount setting affect all wallets in account
	AppAccount
)

// App used for webhooks and developers
// gen:qs
type App struct {
	gorm.Model

	Name string

	Token      string
	WebHookURL string
	Level      AppLevel

	// not used now
	WalletID uint
	// not used now
	GroupID   uint
	AccountID uint
}
