// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package exchange

import (
	"math"

	"github.com/zhuharev/hyip/models"
)

// ConvertAmount from currency code to currency code
func ConvertAmount(from, to string, amount float64) (result float64, err error) {
	if from == to {
		return amount, nil
	}

	var (
		curr map[string]float64
		ok   bool
	)

	cacheMu.RLock()
	curr, ok = cache[from]
	cacheMu.RUnlock()
	if !ok {
		curr, err = getchCurr(from)
		if err != nil {
			return 0, err
		}
	}

	//color.Cyan("Convert from %s to %s %f * %f = %f ", from, to, amount, cache[from][to], cache[from][to]*amount)

	return curr[to] * amount, nil
}

// TypeConverter converts uint amount to float and back
type TypeConverter struct {
}

// DefaultTypeConverter initialized type TypeConverter
var DefaultTypeConverter = &TypeConverter{}

// NewTypeConverter returns emty type converter
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
