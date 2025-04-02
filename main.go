package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"history-engine/engine/jobs"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"history-engine/engine/web"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var buildVersion = "dev-master"
var buildTime = "unknown"

func main() {
	app := cli.NewApp()
	app.Name = "history engine"
	app.Usage = "history engine"
	app.Description = "history engine"
	app.Before = loadSetting
	app.Commands = []*cli.Command{web.Web, jobs.Jobs}
	app.Version = fmt.Sprintf("%s, build with: %s, time: %s", buildVersion, runtime.Version(), buildTime)
	app.DefaultCommand = web.Web.Name

	app.Flags = append(app.Flags, []cli.Flag{
		&cli.StringFlag{
			Name:    "config, c",
			Aliases: []string{"c"},
			Usage:   "Custom configuration file path",
		},
	}...)

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to run app with %s: %v", os.Args, err)
	}
}

func loadSetting(c *cli.Context) error {
	name := "setting.toml"
	file := c.String("config")

	if file == "" {
		pwd, _ := os.Executable()
		file = filepath.Dir(pwd) + "/" + name
		if !utils.FileExist(file) {
			pwd, _ := os.Getwd()
			file = pwd + "/" + name
		}
	}

	if !utils.FileExist(file) {
		log.Fatalln("setting file not exist")
	}

	return setting.Load(file)
}
