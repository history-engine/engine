package main

import (
	"github.com/urfave/cli"
	"history-engine/engine/jobs"
	"history-engine/engine/setting"
	"history-engine/engine/web"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "history engine"
	app.Usage = "history engine"
	app.Description = "history engine"
	app.Before = loadSetting
	app.Commands = []cli.Command{web.Web, jobs.Jobs}
	app.EnableBashCompletion = true

	app.Flags = append(app.Flags, []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: setting.CustomFile,
			Usage: "Custom configuration file path",
		},
	}...)

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to run app with %s: %v", os.Args, err)
	}
}

func loadSetting(c *cli.Context) error {
	file := c.String("config")
	if file == "" {
		pwd, _ := os.Getwd()
		file = pwd + "/" + setting.CustomFile
	}

	return setting.Load(file)
}
