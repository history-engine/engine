package icon

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/icon"
	"history-engine/engine/service/page"
	"time"
)

var DownloadIcon = &cli.Command{
	Name:    "download-icon",
	Aliases: []string{"di"},
	Usage:   "download host icon",
	Action:  runDownloadIcon,
}

func runDownloadIcon(ctx *cli.Context) error {
	start := 0
	limit := 100
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
			err := icon.DownloadIcon(ctx.Context, item.URL, item.Path)
			if err != nil {
				logger.Zap().Warn("download icon err", zap.String("url", item.URL))
				continue
			}
			logger.Zap().Info("download icon", zap.String("url", item.URL))
		}
	}

	return nil
}
