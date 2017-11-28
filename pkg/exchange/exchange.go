// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exchange

import (
	"fmt"
	"math"

	"github.com/zhuharev/hyip/models"
)

// ConvertAmount from currency code to currency code
func ConvertAmount(from, to string, amount float64) (result float64, err error) {
	// TODO:
	return 0, fmt.Errorf("not implemented")
}

type TypeConverter struct {
}

var DefaultTypeConverter = &TypeConverter{}

func NewTypeConverter() *TypeConverter {
	return &TypeConverter{}
}

// ConvertToUint magick happen
func (tc *TypeConverter) ConvertToUint(currencyCode string, amount float64) uint {
	curr, found := models.GetCurrencyByCode(currencyCode)
	if !found {
		return 0
	}
	return uint(math.Pow10(int(curr.Digits)) * amount)
}

// ConvertToFloat antogonist ConvertToUint
func (tc *TypeConverter) ConvertToFloat(currencyCode string, amount uint) float64 {
	curr, found := models.GetCurrencyByCode(currencyCode)
	if !found {
		return 0.0
	}
	return float64(amount) / math.Pow10(int(curr.Digits))
}
