package traider

import (
	"log"
	"github.com/zhuharev/hyip/models"
	"github.com/zhuharev/hyip/pkg/base"
	"github.com/zhuharev/hyip/pkg/setting"
	"time"

	"github.com/fatih/color"
)

var (
	ProfitPeriod = time.Hour * 24
)

func NewContext() error {
	go func() {
		for {
			start := time.Now()
			users, err := models.Users.NeedsProfit(ProfitPeriod)
			if err != nil {
				log.Fatalln(err)
			}
			tx := models.Tx()
			for _, user := range users {
				profitUSD := TodayUSDProfit(user.DepositUSD)
				profitUSD += TodayUSDProfit(user.ReinvestUSD, ReinvestMod())
				if profitUSD != 0 {
					err = models.Users.Inc(models.BalanceUSD, profitUSD, user.ID, tx)
					if err != nil {
						log.Println(err)
					}
				}

				profitBTC := TodayBTCProfit(user.DepositBTC)
				profitBTC += TodayBTCProfit(user.ReinvestBTC)
				if profitBTC != 0 {
					err = models.Users.Inc(models.BalanceBTC, profitBTC, user.ID, tx)
					if err != nil {
						log.Println(err)
					}
				}
			}
			tx.Commit()
			color.Green("Done %s", time.Since(start))
			time.Sleep(ProfitPeriod)
		}
	}()
	return nil
}

func TodayUSDProfit(deposit float64, mods ...func(float64) float64) (res float64) {
	var levels = []float64{
		setting.App.Deposit.Lvl1USDMinSum,
		setting.App.Deposit.Lvl2USDMinSum,
		setting.App.Deposit.Lvl3USDMinSum,
		setting.App.Deposit.Lvl4USDMinSum,
		setting.App.Deposit.Lvl5USDMinSum,
		setting.App.Deposit.Lvl6USDMinSum,
	}
	var modificators = []float64{
		setting.App.Deposit.Lvl1Percent,
		setting.App.Deposit.Lvl2Percent,
		setting.App.Deposit.Lvl3Percent,
		setting.App.Deposit.Lvl4Percent,
		setting.App.Deposit.Lvl5Percent,
		setting.App.Deposit.Lvl6Percent,
	}

	percent := 0.0

	for i, level := range levels {
		if deposit >= level {
			percent = base.Percents(modificators[i])
		} else {
			break
		}
	}

	for _, fn := range mods {
		percent += fn(deposit)
	}
	return deposit * percent
}

func TeamUSDMod(teamValue float64) func(float64) float64 {
	return func(depo float64) (team float64) {
		if teamValue >= setting.App.Team.Lvl3USDMinValue {
			team = setting.App.Team.Lvl3Bonus
		} else if teamValue >= setting.App.Team.Lvl2USDMinValue {
			team = setting.App.Team.Lvl2Bonus
		} else if teamValue >= setting.App.Team.Lvl1USDMinValue {
			team = setting.App.Team.Lvl1Bonus
		}
		return base.Percents(team)
	}
}

func ReinvestMod() func(float64) float64 {
	return func(depo float64) (percent float64) {
		return base.Percents(setting.App.Deposit.ReinvestAdditional)
	}
}

// btc

func TodayBTCProfit(deposit float64, mods ...func(float64) float64) (res float64) {
	var levels = []float64{
		setting.App.Deposit.Lvl1BTCMinSum,
		setting.App.Deposit.Lvl2BTCMinSum,
		setting.App.Deposit.Lvl3BTCMinSum,
		setting.App.Deposit.Lvl4BTCMinSum,
		setting.App.Deposit.Lvl5BTCMinSum,
		setting.App.Deposit.Lvl6BTCMinSum,
	}
	var modificators = []float64{
		setting.App.Deposit.Lvl1Percent,
		setting.App.Deposit.Lvl2Percent,
		setting.App.Deposit.Lvl3Percent,
		setting.App.Deposit.Lvl4Percent,
		setting.App.Deposit.Lvl5Percent,
		setting.App.Deposit.Lvl6Percent,
	}

	percent := 0.0

	for i, level := range levels {
		if deposit >= level {
			percent = base.Percents(modificators[i])
		} else {
			break
		}
	}

	for _, fn := range mods {
		percent += fn(deposit)
	}
	return deposit * percent
}

func TeamBTCMod(teamValue float64) func(float64) float64 {
	return func(depo float64) (team float64) {
		if teamValue >= setting.App.Team.Lvl3BTCMinValue {
			team = setting.App.Team.Lvl3Bonus
		} else if teamValue >= setting.App.Team.Lvl2BTCMinValue {
			team = setting.App.Team.Lvl2Bonus
		} else if teamValue >= setting.App.Team.Lvl1BTCMinValue {
			team = setting.App.Team.Lvl1Bonus
		}
		return base.Percents(team)
	}
}
