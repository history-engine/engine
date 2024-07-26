package page

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	entPage "history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/filetype"
	"history-engine/engine/service/host"
	"history-engine/engine/service/page"
	"history-engine/engine/service/zincsearch"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"time"
)

var IntegrityCheck = &cli.Command{
	Name:    "integrity-check",
	Aliases: []string{"ic"},
	Usage:   "Check database, HTML files, ZincSearch Document consistency",
	Action:  runIntegrityCheck,
}

func runIntegrityCheck(ctx *cli.Context) error {
	start := 0
	limit := 100
	x := db.GetEngine().Page
	for {
		list, err := x.Query().Order(ent.Desc(entPage.FieldID)).Offset(start).Limit(limit).All(ctx.Context)
		if err != nil {
			panic(err)
		}

		if len(list) == 0 {
			break
		}

		start += limit
		time.Sleep(time.Millisecond * 100)

		for _, item := range list {
			if !filetype.Include(item.UserID, item.URL) && filetype.Exclude(item.UserID, item.URL) {
				logger.Zap().Info("ignore by suffix: " + item.URL)
				continue
			}

			if !host.Include(item.UserID, item.URL) {
				if host.Exclude(item.UserID, item.URL) {
					logger.Zap().Info("ignore by rule: " + item.URL)
					x.DeleteOneID(item.ID).Exec(ctx.Context)
					_ = zincsearch.DelDocument(item.UserID, item.UniqueID, item.Version)
					continue
				}
			}

			if !utils.FileExist(setting.SingleFile.HtmlPath + item.Path) {
				logger.Zap().Info("HTML file not exist", zap.String("html", setting.SingleFile.HtmlPath+item.Path))
				x.DeleteOneID(item.ID).Exec(ctx.Context)
				_ = zincsearch.DelDocument(item.UserID, item.UniqueID, item.Version)
				continue
			}

			if item.Title == "" || item.IndexedAt.IsZero() {
				err := page.ParserPageWithId(item.ID)
				if err != nil {
					panic(err)
				}
				logger.Zap().Info("HTML file not parse", zap.Int64("id", item.ID), zap.String("url", item.URL))
			}
		}
	}
	return nil
}
