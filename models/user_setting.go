// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"log"

	"github.com/asdine/storm"
	"github.com/zhuharev/hyip/pkg/setting"
)

var (
	stormDB *storm.DB
)

// UserSetting storm model for saving user setting
type UserSetting struct {
	UserID uint `storm:"id" json:"user_id"`

	// wallets
	Advcash string `storm:"unique" json:"advcash"`
	Qiwi    string `json:"qiwi"`
}

func newStormContext() (err error) {
	if setting.Dev {
		stormDB, err = storm.Open("data/storm.db")
	} else {
		stormDB, err = storm.Open("/storage/storm.db")
	}

	//defer db.Close()

	return
}

type userSetting int

var (
	// UserSettings like API endpoint
	UserSettings userSetting

	userSettingsBucketName = "user_setting"
)

func (us userSetting) Get(userID uint) (res UserSetting, err error) {
	err = stormDB.One("UserID", userID, &res)
	if err == storm.ErrNotFound {
		err = nil
		return
	}
	return
}

func (us userSetting) SetField(userID uint, fieldName string, value string) (err error) {
	log.Println("Update field", fieldName, value)
	err = stormDB.UpdateField(&UserSetting{UserID: userID}, fieldName, value)
	return
}

func (us userSetting) GetByField(fieldName, value string) (res UserSetting, err error) {
	err = stormDB.One(fieldName, value, &res)
	return
}
