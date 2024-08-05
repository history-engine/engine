package user

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

func Avatar(c echo.Context) error {
	path := c.Param("path")

	defer func() {
		if err := recover(); err != nil {
			logger.Zap().Info("view avatar err", zap.Any("recover", err), zap.String("path", path))
		}
	}()

	file := setting.Common.DataPath + "/avatar/" + path
	if !utils.FileExist(file) {
		return c.String(404, "file not found")
	}

	return c.File(file)
}
