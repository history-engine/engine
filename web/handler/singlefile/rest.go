package singlefile

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/service/singlefile"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"io"
	"net/http"
	"os"
)

func RestSave(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("uid").(int64)

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

	if html.Size < 2048 { // todo 可配置
		return c.JSON(http.StatusBadRequest, nil)
	}

	if singlefile.CheckIgnore(url) {
		logger.Zap().Info("ignore by rule: " + url)
		return c.JSON(http.StatusOK, nil)
	}

	uniqueId := utils.Md5str(url) // todo 自定义
	version, created := page.NextVersion(ctx, uniqueId)
	if singlefile.CheckVersionInterval(created) {
		logger.Zap().Info("ignore by interval: " + url)
		return c.JSON(http.StatusOK, nil)
	}

	logger.Zap().Debug("rest receive singleFile",
		zap.String("url", url),
		zap.String("uniqueId", uniqueId),
		zap.Int("version", version))

	// 检查并创建目录
	storagePath := fmt.Sprintf("/%s/%s", uniqueId[:2], uniqueId[2:4])
	if _, err = os.Stat(setting.SingleFile.Path + storagePath); err != nil {
		if !os.IsNotExist(err) { // TODO 未知错误,记录日志
			return c.JSON(http.StatusInternalServerError, nil)
		}
		if err := os.MkdirAll(setting.SingleFile.Path+storagePath, 0775); err != nil {
			// todo 这里可能有多种情况
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}

	// 文件写入
	storageFile := fmt.Sprintf("%s/%s.%d.html", storagePath, uniqueId, version)
	f, err := os.OpenFile(setting.SingleFile.Path+storageFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	_, err = io.Copy(f, src)
	_ = f.Close()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// 入库
	_, err = page.SavePage(ctx, &model.Page{
		UserId:   userId,
		UniqueId: uniqueId,
		Version:  version,
		Url:      url,
		Size:     int(html.Size),
		Path:     storageFile,
	})
	if err != nil {
		logger.Zap().Fatal("save page error", zap.Error(err), zap.String("url", url))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	// 后台分析HTML
	go page.ParserPage(context.Background(), uniqueId)

	return c.JSON(http.StatusCreated, nil)
}
