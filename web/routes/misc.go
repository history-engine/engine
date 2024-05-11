package routes

import (
	"github.com/labstack/echo/v4"
	"history-engine/engine/web/handler/misc"
)

func miscRegister(r *echo.Group) {
	//r.GET("/ping", misc.Ping)
	//r.GET("/version", misc.Version)
	//r.GET("/health", misc.Health)
	r.GET("/config", misc.Config)
}
