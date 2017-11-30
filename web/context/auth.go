// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import macaron "gopkg.in/macaron.v1"

// ToggleOptions options for Toggle
type ToggleOptions struct {
	SignInRequired  bool
	SignOutRequired bool
	AdminRequired   bool
	DisableCSRF     bool
}

// Toggle make helper func
func Toggle(options *ToggleOptions) macaron.Handler {
	return func(c *Context) {
		// Check non-logged users landing page.
		if c.Autorized() && c.Req.RequestURI == "/" {
			c.Redirect("/dash")
			return
		}

		// Redirect to dashboard if user tries to visit any non-login page.
		if options.SignOutRequired && c.Autorized() && c.Req.RequestURI != "/" {
			c.Redirect("/")
			return
		}

		if options.SignInRequired {
			if !c.Autorized() {

				//c.SetCookie("redirect_to", url.QueryEscape(setting.AppSubURL+c.Req.RequestURI), 0, setting.AppSubURL)
				c.Redirect("/login")
				return
			}
		}

		// if options.AdminRequired {
		// 	if !c.User.IsAdmin {
		// 		c.Error(403)
		// 		return
		// 	}
		// 	c.Data["PageIsAdmin"] = true
		// }
	}
}
