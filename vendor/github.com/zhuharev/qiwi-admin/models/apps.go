// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "github.com/zhuharev/guse/random"

type appsModel int

// Apps used for models helper
var Apps appsModel

func (a appsModel) List(accountID uint) (apps []App, err error) {
	err = NewAppQuerySet(db).AccountIDEq(accountID).All(&apps)
	return
}

func (a appsModel) Create(accountID uint, name string) (app *App, err error) {
	app = new(App)
	app.Token = random.String(10)
	app.AccountID = accountID
	app.Level = AppAccount
	app.Name = name
	err = app.Create(db)
	return
}

func (a appsModel) SetWebhook(ID uint, uri string) (err error) {
	err = NewAppUpdater(db).SetID(ID).SetWebHookURL(uri).Update()
	return
}
