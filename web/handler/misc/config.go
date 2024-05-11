package misc

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/setting"
	"history-engine/engine/utils"
)

func Config(c echo.Context) error {
	data := map[string]interface{}{
		"enable_register": setting.Common.EnableRegister,
		"lang":            setting.Common.Lang,
		//"role":            "admin",
		//"login":           false,
	}
	return utils.ApiSuccess(c, data)
}
