package base

import "strconv"

func FmtAmount(amount float64) string {
	return strconv.FormatFloat(amount, 'f', -1, 32)
}

// func ProfitUSD(amount float64) float64 {
//   if amount>= setting.
// }
func Percents(amount float64) float64 {
	return amount / 100.0
}
