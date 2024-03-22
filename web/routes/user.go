package routes

import (
	"history-engine/engine/web/handler/auth"
	"history-engine/engine/web/handler/user"

	"github.com/labstack/echo/v4"
)

// 注册用户相关路由
func userRouteRegister(r *echo.Group) {
	r.POST("/login", auth.Password)
	r.POST("/register", user.Register)
}
