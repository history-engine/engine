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

	UserRouteRegister(e.Group("/user"))
	SingleFileRouteRegister(e.Group("/singlefile"))
}

func onAddRouteHandler(host string, route echo.Route, handler echo.HandlerFunc, middleware []echo.MiddlewareFunc) {
	if strings.Contains(route.Name, "history-engine") {
		log.Printf("add route: %s %s\n", route.Method, route.Path)
	}
}
