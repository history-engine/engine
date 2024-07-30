package page

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	entPage "history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/service/filetype"
	"history-engine/engine/service/host"
	"history-engine/engine/service/page"
	"history-engine/engine/service/search"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"time"
)

var ParseCheck = &cli.Command{
	Name:    "parse-check",
	Aliases: []string{"pc"},
	Usage:   "Check database, HTML files consistency",
	Action:  runParseCheck,
}

func runParseCheck(ctx *cli.Context) error {
	var maxId int64 = 0
	limit := 100
	x := db.GetEngine().Page
	for {
		list, err := x.Query().
			Where(entPage.And(entPage.IDGT(maxId), entPage.ParsedAtEQ(timeZero))).
			Order(ent.Asc(entPage.FieldID)).
			Limit(limit).
			All(ctx.Context)
		if err != nil {
			panic(err)
		}

		if len(list) == 0 {
			break
		}

		maxId = list[len(list)-1].ID
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
				logger.Zap().Info("HTML file not exist", zap.String("html", setting.SingleFile.HtmlPath+item.Path))
				ignore = true
			}

			if ignore {
				x.DeleteOneID(item.ID).Exec(ctx.Context)
				docId := fmt.Sprintf("%s%d", item.UniqueID, item.Version)
				search.Engine().DelDocument(ctx.Context, item.UserID, docId)
				continue
			}

			logger.Zap().Info("pend page", zap.Int64("id", item.ID), zap.String("path", item.Path), zap.String("url", item.URL))
			if err := page.ParserPageWithId(item.ID); err != nil {
				logger.Zap().Warn("parse HTML err", zap.Error(err), zap.Int64("id", item.ID), zap.String("url", item.URL))
				continue
			}
		}
	}
	return nil
}
