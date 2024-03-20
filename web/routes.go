package web

import (
	"history-engine/engine/setting"
	"history-engine/engine/web/handlers"
	"history-engine/engine/web/handlers/singlefile"
	"history-engine/engine/web/middleware"

	"github.com/labstack/echo/v4"
)

// InitRoutes 初始化路由
func InitRoutes(r *echo.Group, handler handlers.Handler) {
	// 注册页面相关路由
	e := singlefile.NewEndpoint("/", setting.SingleFile.Path)

	//r.Use(middleware.BasicAuth(basicAuth))
	r.Use(middleware.BasicAuth())
	r.Add("PUT", "/html/:file", e.Put)
	r.Add("OPTIONS", "/html/:file", e.Cover)
	r.Add("MKCOL", "/html/:file", e.Cover)
	r.Add("PROPFIND", "/html/:file", e.Cover)

	// 用户路由
	r.POST("/login", handler.Login)
	r.POST("/register", handler.RegisterUser)
}
