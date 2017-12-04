// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

type accountsModel int

// Accounts provide api acces for iterate with accounts
var Accounts accountsModel

// List returns list of accounts for specific user
func (accountsModel) List(userID uint) (res []Account, err error) {
	err = NewAccountQuerySet(db).UserIDEq(userID).All(&res)
	return
}

func (accountsModel) Increase(accountID uint, delta uint) error {
	sql := `UPDATE accounts
	SET amount = amount + ?
	WHERE id = ?`
	return db.Exec(sql, delta, accountID).Error
}
