// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

//go:generate goqueryset -in ticket.go

import "github.com/jinzhu/gorm"

type TicketStatus int

const (
	TicketOpen TicketStatus = 1 << (iota + 1)
	TicketClosed
	TicketAnswered
)

// Ticket used for support
// gen:qs
type Ticket struct {
	gorm.Model

	OwnerID  uint
	Title    string
	Category uint
	Email    string
	Status   TicketStatus
}

// IsAnonymous returns true ticket not have owner
func (t Ticket) IsAnonymous() bool {
	return t.OwnerID == 0
}

// Message used for support chat
// gen:qs
type Message struct {
	gorm.Model

	TicketID uint
	Body     string
	OwnerID  uint
}
