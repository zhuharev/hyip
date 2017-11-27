// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

var (
	// AppVer current app version
	AppVer string
	// App main config
	App struct {
		DataDir string
		// Db struct {
		// 	Driver string
		// 	Config string
		// }
	}
)
