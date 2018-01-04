package base

import (
	"strconv"

	"github.com/zhuharev/hyip/pkg/setting"

	base62 "github.com/pilu/go-base62"
)

func FmtAmount(amount float64) string {
	return strconv.FormatFloat(amount, 'f', -1, 32)
}

// func ProfitUSD(amount float64) float64 {
//   if amount>= setting.
// }
func Percents(amount float64) float64 {
	return amount / 100.0
}

// HashNumber base62 hash multiple id with app secret number (from config)
func HashNumber(id int) string {
	return base62.Encode(id * int(setting.App.SecretNumber))
}

// DecodeHash returns int by hash
func DecodeHash(hash string) int {
	return base62.Decode(hash) / setting.App.SecretNumber
}
