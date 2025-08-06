package web

import (
	"github.com/urfave/cli/v2"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
	"history-engine/engine/library/logger"
	"history-engine/engine/library/server"
	"history-engine/engine/web/routes"
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
	logger.EnableLogger()
	db.EnableDb()
	localcache.EnableLocalCache()
	return nil
}

func runWeb(c *cli.Context) error {
	server.New(routes.RegisterRoute).Run()
	return nil
}
