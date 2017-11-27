// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import (
	"github.com/zhuharev/hyip/models"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

// Context will be used in routers
type Context struct {
	*macaron.Context

	Flash   *session.Flash
	Session session.Store
	User    *models.User
}

// Contexter wrap macaron.Context
func Contexter() macaron.Handler {
	return func(c *macaron.Context, sess session.Store, f *session.Flash) {
		ctx := &Context{
			Context: c,
			Flash:   f,
			Session: sess,
		}

		// set current request uri
		// used for menu active links
		ctx.Data["requestURI"] = ctx.Req.URL.RequestURI()

		if userIface := sess.Get("user_id"); userIface != nil {
			if userID, ok := userIface.(uint); ok {
				user, err := models.Users.Get(userID)
				if err != nil {
					sess.Delete("user_id")
					c.Redirect("/login")
					return
				}
				ctx.User = user
				c.Data["User"] = user
			}
		}

		c.Map(ctx)
	}
}

// Autorized just hellper
func (ctx *Context) Autorized() bool {
	return ctx.User != nil
}

// HTML overwrite macaron.HTML method
func (ctx *Context) HTML(code int, tmplName string, other ...interface{}) {
	layoutName := "layout"
	if !ctx.Autorized() {
		layoutName = "unauth-layout"
	}
	ctx.Context.HTML(code, tmplName, ctx.Data, macaron.HTMLOptions{Layout: layoutName})
}

// HasError check passed err and write resposne if err!=nil
func (ctx *Context) HasError(err error, redirects ...string) bool {
	location := "/"
	if len(redirects) > 0 {
		location = redirects[0]
	}
	if err != nil {
		ctx.Flash.Error(err.Error())
		if ctx.User != nil {
			ctx.Redirect(location)
		} else {
			ctx.Redirect(location)
		}

		//ctx.Error(200, err.Error())
		return true
	}
	return false
}
