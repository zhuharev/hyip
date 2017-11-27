package setting

import ini "gopkg.in/ini.v1"

var (
	confFile = "/storage/app.ini"

	// App app.ini will be mapped to this var
	App struct {
		SecretNumber int
		BotName      string
		Token        string

		Web struct {
			Port int
		}

		PaymentSystems struct {
			QiwiTokens     []string
			QiwiAdminToken string
		}

		Account struct {
			MinimalUSDDeposit float64
			MinimalBTCDeposit float64

			MinimalUSDReinvestAmount float64
			MinimalBTCReinvestAmount float64
		}

		Deposit struct {
			ReinvestAdditional float64

			Lvl1Percent float64
			Lvl2Percent float64
			Lvl3Percent float64
			Lvl4Percent float64
			Lvl5Percent float64
			Lvl6Percent float64

			Lvl1USDMinSum float64
			Lvl2USDMinSum float64
			Lvl3USDMinSum float64
			Lvl4USDMinSum float64
			Lvl5USDMinSum float64
			Lvl6USDMinSum float64

			Lvl1BTCMinSum float64
			Lvl2BTCMinSum float64
			Lvl3BTCMinSum float64
			Lvl4BTCMinSum float64
			Lvl5BTCMinSum float64
			Lvl6BTCMinSum float64
		}

		Team struct {
			Lvl1Bonus float64
			Lvl2Bonus float64
			Lvl3Bonus float64

			Lvl1USDMinValue float64
			Lvl2USDMinValue float64
			Lvl3USDMinValue float64

			Lvl1BTCMinValue float64
			Lvl2BTCMinValue float64
			Lvl3BTCMinValue float64
		}

		Net struct {
			Ref1 float64
			Ref2 float64
			Ref3 float64
			Ref4 float64
			Ref5 float64
		}
	}

	iniFile *ini.File
)

func NewContext(ops ...func()) (err error) {

	for _, v := range ops {
		v()
	}

	iniFile, err = ini.Load(confFile)
	if err != nil {
		return
	}
	err = iniFile.MapTo(&App)

	return
}

func CustomLocation(path string) func() {
	return func() {
		confFile = path
	}
}
