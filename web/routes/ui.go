package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/setting"
)

// webui
func uiRouteRegister(e *echo.Group) {
	e.File("/", setting.Web.UiPath+"/index.html")
	e.File("/robots.txt", setting.Web.UiPath+"/robots.txt")
	e.File("/favicon.ico", setting.Web.UiPath+"/favicon.ico")
	e.Static("/assets", setting.Web.UiPath+"/assets")
}
