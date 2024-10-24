package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/web/handler/setting"
	"history-engine/engine/web/middleware"
)

func settingRouteRegister(r *echo.Group) {
	r.GET("/profile", setting.GetProfile, middleware.JwtAuth)
	r.POST("/profile", setting.SaveProfile, middleware.JwtAuth)

	r.GET("/host", setting.GetHost, middleware.JwtAuth)
	r.PUT("/host", setting.AddHost, middleware.JwtAuth)
	r.POST("/host", setting.SaveHost, middleware.JwtAuth)
	r.DELETE("/host", setting.DeleteHost, middleware.JwtAuth)

	r.GET("/alias", setting.GetAlias, middleware.JwtAuth)
	r.POST("/alias", setting.SaveAlias, middleware.JwtAuth)

	r.GET("/storage", setting.GetStorage, middleware.JwtAuth)
	r.POST("/storage", setting.SaveStorage, middleware.JwtAuth)

	r.GET("/filetype", setting.GetFiletype, middleware.JwtAuth)
	r.POST("/filetype", setting.SaveFiletype, middleware.JwtAuth)
}
