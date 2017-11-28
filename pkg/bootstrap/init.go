// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bootstrap

import (
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/qiwi"
	"github.com/zhuharev/hyip/pkg/setting"
	"github.com/zhuharev/hyip/pkg/traider"
)

// GlobalInit inits all packages, who have global depends funcs and vars (db
// conections, setting reads and etc).
// It's must be called before all actions.
func GlobalInit(dev bool) (err error) {

	if dev {
		err = setting.NewContext()
		if err != nil {
			return err
		}
	} else {
		err = setting.NewContext(setting.CustomLocation("/storage/app.ini"))
		if err != nil {
			return err
		}
	}

	setting.Dev = dev

	err = models.NewContext()
	if err != nil {
		return err
	}

	err = qiwi.NewContext()
	if err != nil {
		return err
	}

	err = traider.NewContext()
	if err != nil {
		return err
	}

	err = initPaymentSystems()
	if err != nil {
		return err
	}

	return
}
