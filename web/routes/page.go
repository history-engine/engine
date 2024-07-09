package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/web/handler/page"
	"history-engine/engine/web/middleware"
)

func pageRouteRegister(r *echo.Group) {
	r.Use(middleware.JwtAuth)

	r.GET("/search", page.Search)
	r.GET("/preview/:path", page.View)
}
