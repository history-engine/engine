package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/setting"
)

// webui
func uiRouteRegister(e *echo.Group) {
	echo.NotFoundHandler = func(c echo.Context) error { return c.File(setting.Web.UiPath + "/index.html") }
	echo.MethodNotAllowedHandler = func(c echo.Context) error { return c.File(setting.Web.UiPath + "/index.html") }
	e.File("/", setting.Web.UiPath+"/index.html")
	e.File("/robots.txt", setting.Web.UiPath+"/robots.txt")
	e.File("/favicon.ico", setting.Web.UiPath+"/favicon.ico")
	e.Static("/assets", setting.Web.UiPath+"/assets")
}
