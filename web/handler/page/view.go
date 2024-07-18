package page

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

func View(c echo.Context) error {
	file := setting.SingleFile.HtmlPath + "/" + c.Param("path")
	if !utils.FileExist(file) {
		return c.String(404, "file not found")
	}

	return c.File(file)
}
