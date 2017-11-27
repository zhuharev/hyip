// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

//go:generate goqueryset -in telegram.go

// TelegramBind is bind telegram IDs
// easyjson:json
// gen:qs
type TelegramBind struct {
	UserID     uint  `gorm:"unique_index"`
	TelegramID int64 `gorm:"unique_index"`
}

// BindTelegramID binds userID and telegram ID
func BindTelegramID(userID uint, telegramID int64) (err error) {
	tb := &TelegramBind{
		UserID:     userID,
		TelegramID: telegramID,
	}

	err = tb.Create(db)
	return
}
