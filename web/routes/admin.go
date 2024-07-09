package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/web/handler/admin"
	"history-engine/engine/web/middleware"
)

func adminRouteRegister(r *echo.Group) {
	adminUserRouteRegister(r.Group("/user"))
}

func adminUserRouteRegister(r *echo.Group) {
	r.Use(middleware.JwtAuth)

	r.GET("/list", admin.UserList)
	r.POST("/create", admin.UserCreate)
	//r.PUT("/update", admin.UserUpdate)
	//r.DELETE("/delete", admin.UserDelete)
}
