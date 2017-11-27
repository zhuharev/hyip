// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "github.com/jinzhu/gorm"

//go:generate goqueryset -in page.go

// Page text or article on website
// easyjson:json
// gen:qs
type Page struct {
	gorm.Model

	Slug  string `gorm:"unique_index"`
	Title string
	Body  string

	Published bool
}
