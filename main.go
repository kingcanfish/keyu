package main

import (
	"github.com/kingcanfish/keyu/time"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{{
			Name:        "time",
			Aliases:     nil,
			Usage:       "",
			UsageText:   "",
			Description: "",
			ArgsUsage:   "",
			Category:    "",
			Action:      time.StampAction,
		},
		}}
	err := app.Run(os.Args)
	if err != nil {
		panic("start fail")
	}
}
