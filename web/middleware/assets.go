package middleware

import (
	"history-engine/engine/webui"
	"strings"

	"github.com/fengqi/lrace"
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

var spec = []string{
	"/",
	"/index.html",
	"/robots.txt",
	"/favicon.ico",
}

// Assets 静态文件，内嵌式
func Assets() echo.MiddlewareFunc {
	return em.StaticWithConfig(em.StaticConfig{
		Filesystem: webui.Dist("dist"),
		HTML5:      true,
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			if lrace.InArray(spec, path) || strings.HasPrefix(path, "/assets/") {
				c.Response().Header().Set(echo.HeaderCacheControl, "public, max-age=604800")
				return false
			}
			return true
		},
	})
}
