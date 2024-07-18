package singlefile

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/net/webdav"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/service/readability"
	"history-engine/engine/service/singlefile"
	"history-engine/engine/utils"
	"io"
	"log"
	"net/http"
	"os"
)

type Endpoint struct {
	dir      string
	prefix   string
	internal webdav.Handler
}

func NewEndpoint(prefix, dir string) *Endpoint {
	return &Endpoint{
		dir:    dir,
		prefix: prefix,
		internal: webdav.Handler{
			Prefix:     prefix,
			LockSystem: webdav.NewMemLS(),
			FileSystem: Dir{Dir: webdav.Dir(dir)},
		},
	}
}

// Put 保存singlefile生成的html文件
func (e *Endpoint) Put(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("uid").(int64)

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if len(body) < 2048 { // todo 可配置
		return c.String(http.StatusOK, "")
	}

	url := readability.Parser().ExtractSingleFileUrl(body[:2048])
	if singlefile.CheckIgnore(url) {
		logger.Zap().Info("ignore: " + url)
		return c.String(http.StatusOK, "")
	}

	uniqueId := utils.Md5str(url) // todo 自定义
	version, created := page.NextVersion(ctx, uniqueId)
	if singlefile.CheckVersionInterval(created) {
		logger.Zap().Info("ignore: " + url)
		return c.JSON(http.StatusOK, nil)
	}

	logger.Zap().Debug("receive singleFile", zap.String("url", url), zap.String("uniqueId", uniqueId))

	// 检查并创建目录
	storagePath := fmt.Sprintf("/%s/%s", uniqueId[:2], uniqueId[2:4])
	_, err = e.internal.FileSystem.Stat(ctx, storagePath)
	if err != nil {
		if !os.IsNotExist(err) { // TODO 未知错误,记录日志
			return c.String(http.StatusMethodNotAllowed, err.Error())
		}
		if err := e.internal.FileSystem.(Dir).MkdirAll(ctx, storagePath, 0777); err != nil {
			// todo 这里可能有多种情况
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	// 文件写入
	file := fmt.Sprintf("%s/%s.%d.html", storagePath, uniqueId, version)
	f, err := e.internal.FileSystem.OpenFile(ctx, file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	_, err = io.Writer(f).Write(body)
	_ = f.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	// 入库
	_, err = page.SavePage(ctx, &model.Page{
		UserId:   userId,
		UniqueId: uniqueId,
		Version:  version,
		Url:      url,
		Size:     len(body),
		Path:     file,
	})
	if err != nil {
		logger.Zap().Fatal("save page error", zap.Error(err), zap.String("url", url))
		return c.String(http.StatusInternalServerError, "save page error")
	}

	// 后台分析HTML
	page.ParserPage(context.Background(), uniqueId)

	return c.String(http.StatusCreated, "ok")
}

func (e *Endpoint) Head(c echo.Context) error {
	http.NotFound(c.Response(), c.Request())
	return nil
}

// Cover PUT 以外的操作的兜底, 理论上不应该有逻辑走到这里
func (e *Endpoint) Cover(c echo.Context) error {
	log.Printf("webdav cover: %s", c.Request().URL.Path)
	e.internal.ServeHTTP(c.Response().Writer, c.Request())
	return nil
}
