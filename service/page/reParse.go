package page

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/ent/page"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/icon"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

// ReParse 重复处理：分析HTML内容、下载ICON、提交索引
func ReParse(ctx context.Context, ident model.PageIdent) error {
	x := db.GetEngine()

	var err error
	var row *ent.Page
	if ident.Id > 0 {
		row, err = x.Page.Query().Where(page.ID(ident.Id)).First(ctx)
	} else {
		row, err = x.Page.Query().Where(page.UniqueID(ident.UniqueId), page.Version(ident.Version)).First(ctx)
	}

	if row == nil {
		return errors.New("page not found")
	}

	if err != nil {
		return err
	}

	if !utils.FileExist(setting.SingleFile.HtmlPath + row.Path) {
		go Delete(context.Background(), row)
		return errors.New("HTML file not exist")
	}

	go func() { //  TODO ICON强制重新下载，可能之前下载的有问题
		if err := ParserPageWithId(context.Background(), row.ID); err != nil {
			logger.Zap().Warn("parse page err", zap.Error(err), zap.Any("page", row))
			return
		}
		if err := PutIndexWithId(context.Background(), row.ID); err != nil {
			logger.Zap().Warn("put search index err", zap.Error(err), zap.Any("page", row))
		}
		if err := icon.DownloadIcon(context.Background(), row.URL, row.Path); err != nil {
			logger.Zap().Warn("download icon err", zap.Error(err), zap.Any("page", row))
		}
	}()

	return err
}
