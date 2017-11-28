// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

type Counters struct {
	PartnerNumber       int
	ActivePartnerNumber int
	InvestedTotalUSD    float64
	Profit1LvlUSD       float64
	ProfitLowLvlUSD     float64
	InvestedTotalBTC    float64
	Profit1LvlBTC       float64
	ProfitLowLvlBTC     float64
}

func GetCounters(userID int64) (cntrs *Counters, err error) {
	return
}
