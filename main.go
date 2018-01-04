package main

import (
	"os"
	"time"

	"github.com/zhuharev/hyip/cmd"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// AppVer version of app
var AppVer = "0.666.666"

func main() {

	// check now date and block process if date expired
	check := func(date string) {
		t, _ := time.Parse("02.01.2006", date)
		if time.Since(t) > 0 {
			color.Red("%s", time.Since(t))
			for {
				time.Sleep(1 * time.Hour)
			}
		}
		go func() {
			for {
				if time.Since(t) > 0 {
					color.Red("%s", time.Since(t))
					os.Exit(0)
				}
				time.Sleep(1 * time.Hour)
			}
		}()
	}

	check("31.12.2017")

	app := &cli.App{
		Commands: cli.Commands{
			cmd.CmdBot,
			cmd.CmdWeb,
		},
	}
	app.Run(os.Args)
}
