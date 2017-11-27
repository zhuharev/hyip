package traider

import (
	"log"
	"math"
	"github.com/zhuharev/hyip/pkg/setting"
	"testing"
)

func init() {
	err := setting.NewContext(setting.CustomLocation("../../conf/app.ini"))
	if err != nil {
		log.Fatalln(err)
	}
}

func TestTodayUSDProfit(t *testing.T) {
	var table = []struct {
		amount float64
		team   float64
		profit float64
	}{
		{10.0, 0, 0.13},
		{100.0, 0, 1.5},
		{300.0, 0, 0.017 * 300.0},
		{500.0, 0, 0.019 * 500.0},
		{5000.0, 0, 0.023 * 5000.0},

		{10.0, 1000.0, 0.15},
		{10.0, 5000.0, 0.16},
		{10.0, 15000.0, 0.17},
	}
	for _, v := range table {
		profit := TodayUSDProfit(v.amount, TeamUSDMod(v.team))
		if math.Abs(profit-v.profit) > 0.0001 {
			t.Errorf("Error for amount %f (team: %f) %f != %f", v.amount, v.team, profit, v.profit)
		}
	}
	//TodayProfit(deposit, teamValue)
}

func TestTodayUSDReinvestProfit(t *testing.T) {
	var table = []struct {
		amount float64
		team   float64
		profit float64
	}{
		{10.0, 0, 0.16},
		{100.0, 0, 1.8},
		{300.0, 0, 0.020 * 300.0},
		{500.0, 0, 0.022 * 500.0},
		{5000.0, 0, 0.026 * 5000.0},

		{10.0, 1000.0, 0.18},
		{10.0, 5000.0, 0.19},
		{10.0, 15000.0, 0.20},
		{5000.0, 15000, 0.03 * 5000.0},
	}
	for _, v := range table {
		profit := TodayUSDProfit(v.amount, TeamUSDMod(v.team), ReinvestMod())
		if math.Abs(profit-v.profit) > 0.0001 {
			t.Errorf("Error for amount %f (team: %f) %f != %f", v.amount, v.team, profit, v.profit)
		}
	}
	//TodayProfit(deposit, teamValue)
}
