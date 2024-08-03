package page

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"history-engine/engine/ent"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/service/readability"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"io"
	"net/http"
	"os"
)

func RestSave(c echo.Context) error {
	err := c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	url := c.FormValue("url")
	html, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	src, err := html.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	hi := &model.HtmlInfo{
		Url:      url,
		Suffix:   utils.FileSuffix(url),
		Size:     int(html.Size),
		Sha1:     utils.Sha1Str(url),
		UserId:   c.Get("uid").(int64),
		IoReader: src,
	}

	if ok, err := page.Filter(hi); !ok {
		logger.Zap().Info(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	return save(hi, c)
}

func WebDavPreSave(c echo.Context) error {
	hi := page.ParseHtmlInfo(c.Param("file"))
	hi.UserId = c.Get("uid").(int64)

	if ok, err := page.Filter(hi); !ok {
		logger.Zap().Info(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNotFound, nil)
}

func WebDavSave(c echo.Context) error {
	hi := page.ParseHtmlInfo(c.Param("file"))
	if hi == nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	hi.UserId = c.Get("uid").(int64)
	hi.Size = int(c.Request().ContentLength)
	hi.IoReader = c.Request().Body

	return save(hi, c)
}

func save(hi *model.HtmlInfo, c echo.Context) error {
	ctx := c.Request().Context()
	version, _ := page.NextVersion(ctx, hi.Sha1)

	logger.Zap().Debug("rest receive singleFile",
		zap.String("url", hi.Url),
		zap.String("uniqueId", hi.Sha1),
		zap.Int("version", version))

	// 检查并创建目录
	storagePath := fmt.Sprintf("/%d/%s/%s", hi.UserId, hi.Sha1[:2], hi.Sha1[2:4])
	if _, err := os.Stat(setting.SingleFile.HtmlPath + storagePath); err != nil {
		if !os.IsNotExist(err) { // TODO 未知错误,记录日志
			return c.JSON(http.StatusInternalServerError, nil)
		}
		if err := os.MkdirAll(setting.SingleFile.HtmlPath+storagePath, 0775); err != nil {
			// todo 这里可能有多种情况
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}

	// 文件写入
	storageFile := fmt.Sprintf("%s/%s.%d.html", storagePath, hi.Sha1, version)
	f, err := os.OpenFile(setting.SingleFile.HtmlPath+storageFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	_, err = io.Copy(f, hi.IoReader)
	if err != nil {
		logger.Zap().Error("write html err", zap.Error(err), zap.Any("hi", hi))
		return c.JSON(http.StatusInternalServerError, err.Error())
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

	_ = f.Close()
	_ = hi.IoReader.Close()

	// 入库
	row, err := page.SavePage(ctx, &ent.Page{
		UserID:   hi.UserId,
		UniqueID: hi.Sha1,
		Version:  version,
		URL:      hi.Url,
		Size:     hi.Size,
		Path:     storageFile,
	})
	if err != nil {
		logger.Zap().Fatal("save page error", zap.Error(err), zap.String("url", hi.Url))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// 后台分析HTML、清理历史版本
	go func() {
		if err := page.ParserPageWithId(row.ID); err != nil {
			logger.Zap().Warn("parse page err", zap.Error(err), zap.Any("page", row))
			return
		}
		if err := page.PutIndexWithId(row.ID); err != nil {
			logger.Zap().Warn("put search index err", zap.Error(err), zap.Any("page", row))
		}
	}()
	go func() {
		if err := page.CleanHistory(context.Background(), hi.UserId, hi.Sha1, version); err != nil {
			logger.Zap().Warn("clean history err", zap.Error(err), zap.Any("page", row))
		}
	}()

	return c.JSON(http.StatusCreated, nil)
}
