package page

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/filetype"
	"history-engine/engine/service/host"
	"history-engine/engine/service/search"
	"history-engine/engine/setting"
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
			ignore := false

			if !filetype.Include(item.UserID, item.URL) && filetype.Exclude(item.UserID, item.URL) {
				logger.Zap().Info("ignore by suffix: " + item.URL)
				ignore = true
			}

			if !ignore && !host.Include(item.UserID, item.URL) && host.Exclude(item.UserID, item.URL) {
				logger.Zap().Info("ignore by rule: " + item.URL)
				ignore = true
			}

			if !ignore && !utils.FileExist(setting.SingleFile.HtmlPath+item.Path) {
				logger.Zap().Info("file not exist: " + item.Path)
				ignore = true
			}

			if ignore {
				x.DeleteOneID(item.ID).Exec(ctx.Context)
				docId := fmt.Sprintf("%s%d", item.UniqueID, item.Version)
				search.Engine().DelDocument(ctx.Context, item.UserID, docId)
			}
		}
	}

	return nil
}
