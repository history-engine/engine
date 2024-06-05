package jobs

import (
	"github.com/urfave/cli/v2"
	"history-engine/engine/jobs/handler"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
)

var Jobs = &cli.Command{
	Name:        "jobs",
	Usage:       "start engine jobs server",
	Description: "run engine jobs server and listen for requests",
	Before:      before,
	Subcommands: cli.Commands{handler.Analyse},
}

func before(ctx *cli.Context) error {
	logger.EnableLogger()
	db.EnableDb()
	return nil
}
