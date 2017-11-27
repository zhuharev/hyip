// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import "strings"

// IsErrUnniqueConstraintFailed check is sql error unique constrint failed
func IsErrUnniqueConstraintFailed(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), "UNIQUE constraint failed:")
}
