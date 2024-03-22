package web

import (
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/library/server"
	"history-engine/engine/web/routes"

	"github.com/urfave/cli/v2"
)

var Web = &cli.Command{
	Name:        "web",
	Usage:       "start engine web server",
	Description: "run engine web server and listen for requests",
	Before:      before,
	Action:      runWeb,
	Flags:       []cli.Flag{},
}

func before(c *cli.Context) error {
	db.EnableDb()
	logger.EnableLogger()
	return nil
}

func runWeb(c *cli.Context) error {
	server.New(routes.RegisterRoute).Run()
	return nil
}
