// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exchange

import (
	"sync"
	"time"

	"github.com/fadion/gofixerio"
)

var (
	cache   = map[string]map[string]float64{}
	cacheMu sync.RWMutex

	exchange = fixerio.New()
)

func init() {
	go func() {
		tick := time.NewTicker(15 * time.Minute)
		for range tick.C {
			cacheMu.Lock()
			for key := range cache {
				delete(cache, key)
			}
			cacheMu.Unlock()
		}
	}()
}

func getchCurr(curr string) (map[string]float64, error) {
	exchange.Base(curr)
	rates, err := exchange.GetRates()
	if err != nil {
		return nil, err
	}

	res := make(map[string]float64)

	for k, v := range rates {
		res[k] = float64(v)
	}

	cacheMu.Lock()
	cache[curr] = res
	cacheMu.Unlock()

	return res, nil
}
