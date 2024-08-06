package page

import (
	"github.com/urfave/cli/v2"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/utils"
	"time"
)

var ExcludeCheck = &cli.Command{
	Name:    "exclude-check",
	Aliases: []string{"ec"},
	Usage:   "Check Exclude Document",
	Action:  runExcludeCheck,
}

func runExcludeCheck(ctx *cli.Context) error {
	offset := 0
	limit := 100
	x := db.GetEngine().Page
	for {
		list, err := x.Query().Offset(offset).Limit(limit).All(ctx.Context)
		if err != nil || len(list) == 0 {
			break
		}

		offset += limit
		time.Sleep(time.Millisecond * 100)

		for _, item := range list {
			hi := &model.HtmlInfo{
				Url:    item.URL,
				Suffix: utils.FileSuffix(item.URL),
				Size:   item.Size,
				UserId: item.UserID,
				Path:   item.Path,
			}
			if ok, err := page.Filter(hi); !ok {
				logger.Zap().Info(err.Error())
				page.Delete(ctx.Context, item)
				continue
			}
		}
	}

	return nil
}
