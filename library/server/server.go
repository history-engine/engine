package server

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/library/shutdown"
	"history-engine/engine/library/wait"
	"history-engine/engine/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	http       *echo.Echo
	runError   chan error
	isShutdown bool
}

func New(routeRegister func(r *echo.Echo)) *Server {
	wait.Wait()

	s := &Server{
		http:       echo.New(),
		runError:   make(chan error, 5),
		isShutdown: false,
	}

	s.http.HideBanner = true
	s.http.HidePort = true
	s.http.Debug = setting.Common.Env == "dev"
	s.http.Validator = NewCustomValidator()
	routeRegister(s.http)

	return s
}

func (s *Server) Run() {
	go s.error()
	go s.boot()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-quit
	log.Printf("get shutdown: %v\n", sig)

	s.shutdown()
}

func (s *Server) boot() {
	listen := fmt.Sprintf(":%d", setting.Web.Port)
	log.Printf("server start at %s\n", listen)
	if err := s.http.Start(listen); err != nil {
		s.runError <- err
		s.shutdown()
	}
}

func (s *Server) shutdown() {
	if s.isShutdown {
		return
	}

	s.isShutdown = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // todo 可配置
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		s.runError <- err
	}

	shutdown.ShutdownComponent(s.runError)

	log.Println("server shutdown")
	os.Exit(0)
}

func (s *Server) error() {
	for err := range s.runError {
		if errors.Is(err, http.ErrServerClosed) {
			continue
		}
		log.Printf("run err: %v\n", err)
	}
}
