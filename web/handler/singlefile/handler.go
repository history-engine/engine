package singlefile

import (
	"fmt"
	"history-engine/engine/model"
	"history-engine/engine/service/page"
	"history-engine/engine/service/readability"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/webdav"
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

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if len(body) < 2048 { // todo 可配置
		return c.String(http.StatusOK, "")
	}

	url := readability.Parser().ExtractSingleFileUrl(body[:2048])
	uniqueId := utils.Md5str(url) // todo 自定义

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
	version := page.NextVersion(ctx, uniqueId)
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

	// 内容分析
	article := readability.Parser().Parse(setting.SingleFile.Path + file)

	// 入库
	_, err = page.SavePage(ctx, &model.Page{
		UserId:   c.Get("uid").(int64),
		UniqueId: uniqueId,
		Version:  version,
		Url:      url,
		FullSize: len(body),
		FullPath: file,
	})
	if err != nil {
		panic(err)
	}

	err = page.AddIndex(uniqueId, &model.ZincDocument{
		FilePath: file,
		Url:      url,
		Title:    article.Title,
		Content:  article.TextContent,
		Size:     len(body),
	})
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusCreated, "")
}

// Cover PUT 以外的操作的兜底, 理论上不应该有逻辑走到这里
func (e *Endpoint) Cover(c echo.Context) error {
	log.Printf("webdav cover: %s", c.Request().URL.Path)
	e.internal.ServeHTTP(c.Response().Writer, c.Request())
	return nil
}
