package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/web/handler/page"
)

func pageRouteRegister(r *echo.Group) {
	r.GET("/search", page.Search)
	r.GET("/preview/:path", page.View)
}
