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
	userRouteRegister(e.Group("/user"))
	pageRouteRegister(e.Group("/page"))
	adminRouteRegister(e.Group("/admin"))
	miscRegister(e.Group("/misc"))
}

func onAddRouteHandler(host string, route echo.Route, handler echo.HandlerFunc, middleware []echo.MiddlewareFunc) {
	if strings.Contains(route.Name, "history-engine") {
		log.Printf("add route: %s %s\n", route.Method, route.Path)
	}
}
