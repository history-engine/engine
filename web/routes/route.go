package routes

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.OnAddRouteHandler = onAddRouteHandler

	userRouteRegister(e.Group("/user"))
	singleFileRouteRegister(e.Group("/singlefile"))
	pageRouteRegister(e.Group("/page"))
	adminRouteRegister(e.Group("/admin"))
}

func onAddRouteHandler(host string, route echo.Route, handler echo.HandlerFunc, middleware []echo.MiddlewareFunc) {
	if strings.Contains(route.Name, "history-engine") {
		log.Printf("add route: %s %s\n", route.Method, route.Path)
	}
}
