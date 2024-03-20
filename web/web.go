package web

import (
	"context"
	"fmt"
	"history-engine/engine/library/db"
	"history-engine/engine/service"
	"history-engine/engine/setting"
	"history-engine/engine/web/handlers"
	"history-engine/engine/web/task"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

var Web = &cli.Command{
	Name:        "web",
	Usage:       "start engine web server",
	Description: "run engine web server and listen for requests",
	Before:      before,
	After:       after,
	Action:      runWeb,
	Flags:       []cli.Flag{},
}

func before(c *cli.Context) error {
	log.Printf("web server start at %s:%d\n", setting.Web.Addr, setting.Web.Port)
	_ = db.InitEngine(context.TODO())
	return nil
}

func after(c *cli.Context) error {
	go task.RunPageVersionCheck()

	return nil
}

// runWeb starts engine web server by echo framework
func runWeb(c *cli.Context) error {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.OnAddRouteHandler = onAddRouteHandler

	// 初始化依赖
	svr := service.NewServiceImpl(context.Background(), db.GetEngine())
	handler := handlers.NewHandlerImpl(svr)
	// 初始化路由
	InitRoutes(e.Group("/api"), handler)

	listen := fmt.Sprintf("%s:%d", setting.Web.Addr, setting.Web.Port)
	return e.Start(listen)
}

func onAddRouteHandler(host string, route echo.Route, handler echo.HandlerFunc, middleware []echo.MiddlewareFunc) {
	if strings.Contains(route.Name, "history-engine") {
		log.Printf("add route: %s %s\n", route.Method, route.Path)
	}
}
