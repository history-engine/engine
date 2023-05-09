package jobs

import "github.com/urfave/cli"

var Jobs = cli.Command{
	Name:        "jobs",
	Usage:       "start engine jobs server",
	Description: "run engine jobs server and listen for requests",
	Before:      before,
	After:       after,
	Action:      runJobs,
	Flags:       []cli.Flag{},
}

func before(ctx *cli.Context) error {
	return nil
}

func after(ctx *cli.Context) error {
	return nil
}

func runJobs(c *cli.Context) error {
	return nil
}
