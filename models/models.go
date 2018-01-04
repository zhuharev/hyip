// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/zhuharev/hyip/pkg/payment_system/store"
	"github.com/zhuharev/hyip/pkg/setting"
)

var (
	db *gorm.DB

	// ErrNotFound returns in all funcs if obj not found in db
	ErrNotFound = fmt.Errorf("not found")
)

// DB helper
func DB() *gorm.DB {
	return db
}

// NewContext open db
func NewContext() (err error) {
	if setting.Dev {
		db, err = gorm.Open("sqlite3", "data/db.sqlite")
		if err != nil {
			return err
		}
	} else {
		db, err = gorm.Open("sqlite3", "/storage/db.sqlite")
		if err != nil {
			return err
		}
	}

	db.LogMode(true)
	db.Debug()

	db.AutoMigrate(&User{}, &TelegramBind{}, &Page{},
		&Plan{},
		&Investment{},
		&Account{},
		&store.Transaction{},
		&Ticket{},
		&Message{},
		&Profit{},
		&PaymentSystem{},
	)

	// key value
	if err = initBolt(); err != nil {
		return err
	}

	if err = newStormContext(); err != nil {
		return err
	}

	return
}
