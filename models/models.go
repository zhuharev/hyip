package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	db, err = gorm.Open("sqlite3", "/storage/db.sqlite")
	if err != nil {
		return err
	}

	db.LogMode(true)
	db.Debug()

	db.AutoMigrate(&User{}, &TelegramBind{}, &Page{},
		&Plan{},
		&Investment{},
		&Account{})

	// key value
	if err = initBolt(); err != nil {
		return err
	}

	return
}
