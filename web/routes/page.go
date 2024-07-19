package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/web/handler/page"
	"history-engine/engine/web/middleware"
)

func pageRouteRegister(r *echo.Group) {
	r.GET("/search", page.Search, middleware.JwtAuth)
	r.GET("/view/:path", page.View, middleware.JwtAuth)
	r.POST("/save", page.Save, middleware.Token)
}
