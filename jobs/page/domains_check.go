package page

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/page"
	"history-engine/engine/utils"
	"time"
)

var DomainsCheck = &cli.Command{
	Name:    "domain-check",
	Aliases: []string{"dc"},
	Usage:   "Check url domains",
	Action:  runDomainsCheck,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "process all page",
			Value:   false,
		},
	},
}

func runDomainsCheck(ctx *cli.Context) error {
	start := 0
	limit := 100
	x := db.GetEngine().Page
	for {
		list, err := page.Page(ctx.Context, start, limit)
		if err != nil {
			panic(err)
		}

		if len(list) == 0 {
			break
		}

		start += limit
		time.Sleep(time.Millisecond * 100)

		for _, item := range list {
			if len(item.Domains) > 0 && !ctx.Bool("all") {
				continue
			}

			domains := utils.ExtractDomains(item.URL)
			if len(domains) != 0 {
				_, err := x.UpdateOneID(item.ID).SetDomains(domains).Save(ctx.Context)
				if err != nil {
					logger.Zap().Warn("update page domains err", zap.Error(err), zap.String("unique_id", item.UniqueID), zap.Strings("domains", domains))
					continue
				}

				logger.Zap().Info("update page domains", zap.String("unique_id", item.UniqueID), zap.Strings("domains", domains))
			}
		}
	}

	return nil
}
