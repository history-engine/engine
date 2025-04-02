package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/setting"
	"strings"
)

// webui
func uiRouteRegister(e *echo.Group) {
	e.File("/robots.txt", setting.Web.UiPath+"/robots.txt")
	e.File("/favicon.ico", setting.Web.UiPath+"/favicon.ico")
	e.Static("/avatar", setting.Common.DataPath+"/avatar")
	e.Static("/assets", setting.Web.UiPath+"/assets")

	uiRoutePrefix := []string{"/", "/admin", "/setting"}
	echo.NotFoundHandler = func(ctx echo.Context) error {
		for _, v := range uiRoutePrefix {
			if strings.HasPrefix(ctx.Request().URL.Path, v) {
				return ctx.File(setting.Web.UiPath + "/index.html")
			}
		}
		return echo.ErrNotFound
	}
}
