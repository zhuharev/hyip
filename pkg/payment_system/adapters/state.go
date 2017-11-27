// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package advcash

import "time"

type State struct {
	LastFetched time.Time
	LastTxnID   string
}
