// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"path/filepath"

	"github.com/zhuharev/qiwi-admin/pkg/setting"

	"github.com/jinzhu/gorm"

	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB

	dbName = "db.sqlite"
)

// DB returns db for global usage
func DB() *gorm.DB {
	return db
}

// NewContext init db instance
func NewContext() (err error) {
	db, err = gorm.Open("sqlite3", filepath.Join(setting.App.DataDir, dbName))
	if err != nil {
		return
	}
	db.LogMode(true)

	err = db.AutoMigrate(&User{},
		&Txn{},
		&Wallet{},
		&Group{},
		&App{}).Error
	return
}
