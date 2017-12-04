// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var (
	confFile = "conf/app.yaml"

	// Dev represent is dev mode running
	Dev bool

	// App app.ini will be mapped to this var
	App struct {
		SecretNumber int `yaml:"secret_number"`

		Telegram struct {
			BotToken    string `yaml:"bot_token"`
			BotUsername string `yaml:"bot_username"`
		}

		Web struct {
			Port int
		}

		Languages Languages

		PaymentSystems PaymentSystems `yaml:"payment_systems"`

		Socials []struct {
			Link  string
			Image string
			Alt   string
		} `yaml:"social_networks"`
	}

	//iniFile *ini.File
)

type Languages []struct {
	Name string
	Code string
}

type PaymentSystems []struct {
	Enabled   bool
	Name      string
	APISecret string `yaml:"api_secret"`
	APIName   string `yaml:"api_name"`
	WalletID  string `yaml:"wallet_id"`
}

func (l Languages) Names() (res []string) {
	for _, v := range l {
		res = append(res, v.Name)
	}
	return
}

func (l Languages) Codes() (res []string) {
	for _, v := range l {
		res = append(res, v.Code)
	}
	return
}

func NewContext(ops ...func()) (err error) {

	for _, v := range ops {
		v()
	}

	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &App)
	if err != nil {
		return
	}

	// iniFile, err = ini.Load(confFile)
	// if err != nil {
	// 	return
	// }
	// iniFile.NameMapper = mapper
	// err = iniFile.MapTo(&App)

	return
}

func CustomLocation(path string) func() {
	return func() {
		confFile = path
	}
}
