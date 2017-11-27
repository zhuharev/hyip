// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "github.com/jinzhu/gorm"

type pageModel int

// Pages for shortcat models calling
var Pages pageModel

func (p pageModel) Get(slug string) (page *Page, err error) {
	page = new(Page)
	err = NewPageQuerySet(db).SlugEq(slug).One(page)
	if err != nil {
		return nil, err
	}
	return
}

func (p pageModel) Save(slug string, title string, body string) (err error) {
	page, err := p.Get(slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			page = new(Page)
		} else {
			return err
		}
	}
	page.Title = title
	page.Slug = slug
	page.Body = body

	if page.ID == 0 {
		return page.Create(db)
	}
	return page.Update(db, PageDBSchema.Title, PageDBSchema.Body)
}
