// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "fmt"

// Role role of user
type Role int

const (
	// RoleGuest unconfirmed user
	RoleGuest Role = 1 << iota
	// RoleUser confirmed user
	RoleUser
	// RoleManager manager user
	RoleManager
	// RoleAdmin admin user
	RoleAdmin
)

func (r Role) String() string {
	switch r {
	case RoleUser:
		return "role.user"
	case RoleAdmin:
		return "role.admin"
	case RoleGuest:
		return "role.guest"
	case RoleManager:
		return "role.manager"
	}
	return fmt.Sprintf("unknown role %d", r)
}

// IsAdmin check is user role == Admin
func (u User) IsAdmin() bool {
	return u.Role&RoleAdmin != 0
}
