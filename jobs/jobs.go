package jobs

import (
	"github.com/urfave/cli/v2"
	"history-engine/engine/jobs/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/localcache"
	"history-engine/engine/library/logger"
)

var Jobs = &cli.Command{
	Name:        "jobs",
	Usage:       "start engine jobs server",
	Description: "run engine jobs server and listen for requests",
	Before:      before,
	Subcommands: cli.Commands{
		page.Analyse,
		page.ParseCheck,
		page.LostCheck,
		page.IndexCheck,
		page.ExcludeCheck,
	},
}

func before(ctx *cli.Context) error {
	logger.EnableLogger()
	db.EnableDb()
	localcache.EnableLocalCache()
	return nil
}
