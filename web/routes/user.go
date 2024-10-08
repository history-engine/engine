package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/web/handler/user"
	"history-engine/engine/web/middleware"
)

// 注册用户相关路由
func userRouteRegister(r *echo.Group) {
	r.GET("/info", user.Info, middleware.JwtAuth)
	r.GET("/logout", user.Logout, middleware.JwtAuth)
	r.GET("/avatar/:path", user.Avatar, middleware.JwtAuth)
	r.POST("/login", user.PasswordLogin)
	r.POST("/register", user.Register)
}
