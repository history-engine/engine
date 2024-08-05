package page

import (
	"history-engine/engine/library/logger"
	"history-engine/engine/setting"
	"history-engine/engine/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Icon(c echo.Context) error {
	defer func() {
		if err := recover(); err != nil {
			logger.Zap().Info("view html err", zap.Any("recover", err), zap.String("path", c.Param("path")))
		}
	}()

	path := c.Param("path")
	file := setting.Common.IconPath + "/" + path
	if !utils.FileExist(file) {
		return c.String(404, "file not found")
	}

	return c.File(file)
}
