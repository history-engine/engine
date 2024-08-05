package routes

import (
	"history-engine/engine/web/handler/page"
	"history-engine/engine/web/middleware"

	"github.com/labstack/echo/v4"
)

func pageRouteRegister(r *echo.Group) {
	r.GET("/search", page.Search, middleware.JwtAuth)
	r.GET("/view/:path", page.View, middleware.JwtAuth)
	r.GET("/icon/:path", page.Icon, middleware.JwtAuth)
	r.POST("/save", page.RestSave, middleware.Token)
	r.HEAD("/save/:file", page.WebDavPreSave, middleware.BasicAuth())
	r.PUT("/save/:file", page.WebDavSave, middleware.BasicAuth())
}
