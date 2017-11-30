// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package context

import "github.com/mailru/easyjson"

// EasyJSON fast json serializer
func (ctx *Context) EasyJSON(v easyjson.Marshaler) {
	easyjson.MarshalToHTTPResponseWriter(v, ctx.Context)
}
