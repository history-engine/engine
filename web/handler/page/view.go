package page

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

func View(c echo.Context) error {
	defer func() {
		if err := recover(); err != nil {
			logger.Zap().Info("view html err", zap.Any("recover", err), zap.String("path", c.Param("path")))
		}
	}()

	path := c.Param("path")
	uid := c.Get("uid")
	file := fmt.Sprintf("/%d/%s/%s/%s", uid, path[:2], path[2:4], path)

	file = setting.Common.HtmlPath + file
	if !utils.FileExist(file) {
		return c.String(404, "file not found")
	}

	return c.File(file)
}
