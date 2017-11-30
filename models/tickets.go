// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

type ticketModel int

// Tickets helper for tickets api
var Tickets ticketModel

func (t ticketModel) Create(title string, body string, email string, ownerIDs ...uint) (tic *Ticket, err error) {
	tic = new(Ticket)
	tic.Title = title
	tic.Email = email
	tic.Status = TicketOpened
	if len(ownerIDs) > 0 {
		tic.OwnerID = ownerIDs[0]
	}
	err = tic.Create(db)
	if err != nil {
		return
	}

	_, err = Messages.Create(tic.ID, body, ownerIDs...)

	return
}

func (ticketModel) Get(ticketID uint) (*Ticket, error) {
	var tic Ticket
	err := NewTicketQuerySet(db).IDEq(ticketID).One(&tic)
	return &tic, err
}

func (ticketModel) List(userID uint) (res []Ticket, err error) {
	err = NewTicketQuerySet(db).OwnerIDEq(userID).OrderDescByID().All(&res)
	return
}

func (ticketModel) AdminList(startIDs ...uint) (res []Ticket, err error) {
	qs := NewTicketQuerySet(db).OrderDescByID()
	qs = qs.StatusEq(TicketOpened) // TODO: change this to bit operation
	if len(startIDs) > 0 {
		qs = qs.IDLt(startIDs[0])
	}
	err = qs.All(&res)
	return
}

type messagesModel int

// Messages helper for tickets messages api
var Messages messagesModel

func (messagesModel) Create(ticketID uint, body string, ownerIDs ...uint) (msg *Message, err error) {
	msg = new(Message)
	msg.Body = body
	msg.TicketID = ticketID
	if len(ownerIDs) > 0 {
		msg.OwnerID = ownerIDs[0]
	}
	err = msg.Create(db)
	return
}

func (messagesModel) List(ticketID uint) (msgs []Message, err error) {
	err = NewMessageQuerySet(db).TicketIDEq(ticketID).OrderAscByID().All(&msgs)
	return
}
