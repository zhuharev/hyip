// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate goqueryset -in user.go

package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User represent usually user object
// gen:qs
type User struct {
	gorm.Model

	Username       string `gorm:"unique_index"`
	HashedPassword []byte
}

// AuthForm used in binding
type AuthForm struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

// CreateUser save user in db
func CreateUser(af AuthForm) (user *User, err error) {
	user = new(User)
	user.HashedPassword, err = bcrypt.GenerateFromPassword([]byte(af.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Username = af.Login

	err = SaveUser(user)
	if err != nil {
		if IsErrUnniqueConstraintFailed(err) {
			err = fmt.Errorf("Пользователь с таким именем уже существует")
		}
		return
	}

	_, err = CreateGroup("Главная", user.ID)

	return user, err
}

// GetUser returns user by id
func GetUser(id uint) (user *User, err error) {
	user = new(User)
	err = NewUserQuerySet(db).IDEq(id).One(user)
	return
}

// GetUserByAuthForm check password and returns user
func GetUserByAuthForm(af AuthForm) (user *User, err error) {
	user = new(User)
	err = NewUserQuerySet(db).UsernameEq(af.Login).One(user)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(af.Password))
	return
}

// SaveUser update or create user
func SaveUser(user *User) (err error) {
	if user.ID == 0 {
		err = user.Create(db)
		return
	}
	err = user.Update(db)
	return
}
