package page

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/library/db"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/icon"
	"history-engine/engine/service/readability"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"io"
	"os"
)

// SavePage 保存页面
func SavePage(ctx context.Context, hi *model.HtmlInfo) error {
	version, _ := NextVersion(ctx, hi.Sha1)

	// 检查并创建目录
	storagePath := fmt.Sprintf("/%d/%s/%s", hi.UserId, hi.Sha1[:2], hi.Sha1[2:4])
	if _, err := os.Stat(setting.Common.HtmlPath + storagePath); err != nil {
		if !os.IsNotExist(err) { // TODO 未知错误,记录日志
			return err
		}
		if err := os.MkdirAll(setting.Common.HtmlPath+storagePath, 0775); err != nil {
			// todo 这里可能有多种情况
			return err
		}
	}

	// 文件写入
	storageFile := fmt.Sprintf("%s/%s.%d.html", storagePath, hi.Sha1, version)
	f, err := os.OpenFile(setting.Common.HtmlPath+storageFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, hi.IoReader)
	if err != nil {
		return err
	}

	// 补全Url，WebDav下可能为空
	if hi.Url == "" {
		_, err = f.Seek(0, io.SeekStart)
		if err == nil {
			comment := make([]byte, 2048)
			_, err = f.Read(comment)
			if err == nil {
				hi.Url = readability.Parser().ExtractSingleFileUrl(comment)
			}
		}
	}

	logger.Zap().Info("rest receive singleFile",
		zap.String("url", utils.Ternary[string](hi.Url == "", hi.Host, hi.Url)),
		zap.String("path", setting.Common.HtmlPath+storageFile),
		zap.String("uniqueId", hi.Sha1),
		zap.Int("version", version))

	_ = f.Close()
	_ = hi.IoReader.Close()

	row, err := db.GetEngine().Page.Create().
		SetUserID(hi.UserId).
		SetUniqueID(hi.Sha1).
		SetVersion(version).
		SetURL(hi.Url).
		SetPath(storageFile).
		SetSize(hi.Size).
		SetDomains(utils.ExtractDomains(hi.Url)).
		SetStatus(0).
		Save(ctx)
	if err != nil {
		return err
	}

	// 后台分析HTML、更新索引、下载icon
	go func() {
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
	// 清理历史版本
	go func() {
		if err := CleanHistory(context.Background(), hi.UserId, hi.Sha1, version); err != nil {
			logger.Zap().Warn("clean history err", zap.Error(err), zap.Any("page", row))
		}
	}()

	return nil
}
