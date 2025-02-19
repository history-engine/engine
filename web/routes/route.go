package routes

import (
	"log"
	"strings"

	em "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	e.Use(em.Recover())
	e.OnAddRouteHandler = onAddRouteHandler

	uiRouteRegister(e.Group(""))
	pageViewRegister(e.Group("/page"))
	
	userRouteRegister(e.Group("/api/user"))
	pageRouteRegister(e.Group("/api/page"))
	adminRouteRegister(e.Group("/api/admin"))
	settingRouteRegister(e.Group("/api/setting"))
	miscRegister(e.Group("/api/misc"))
}

func onAddRouteHandler(host string, route echo.Route, handler echo.HandlerFunc, middleware []echo.MiddlewareFunc) {
	if strings.Contains(route.Name, "history-engine") {
		log.Printf("add route: %s %s\n", route.Method, route.Path)
	}
}
