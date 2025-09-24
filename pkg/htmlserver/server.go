// Package httpserver implements HTTP server.
package htmlserver

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"time"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/pkg"
	"github.com/mechiko/telebot_v4/pkg/zaplog"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(app entity.Application, e *echo.Echo, opts ...Option) *Server {
	// HTTP Server
	mime.AddExtensionType(".js", "application/javascript")
	e.Use(middleware.Recover())
	e.Use(echozap.ZapLogger(zaplog.EchoSugar))

	app.GetLogger().Infof("htmlserver entity.Mode=%s", entity.Mode)
	if entity.Mode == "development" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"https://tgx.egais-help.ru", "http://localhost:3788", "http://localhost:3650", "http://localhost:3550"},
			AllowHeaders:     []string{"authorization", "Content-Type"},
			AllowCredentials: true,
			AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"https://tgx.egais-help.ru"},
			AllowHeaders:     []string{"authorization", "Content-Type"},
			AllowCredentials: true,
			AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))
	}

	e.Pre(middleware.Rewrite(map[string]string{
		"/login*":        "/",
		"/logout*":       "/",
		"/examens*":      "/",
		"/missions*":     "/",
		"/users*":        "/",
		"/telebotusers*": "/",
		"/masters*":      "/",
	}))

	httpServer := &http.Server{
		Handler:      e,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	e.GET("/shutdown", func(c echo.Context) error {
		// e.Shutdown(context.Background())
		c.String(http.StatusOK, "Shutdown in process")
		return s.Shutdown()
	})

	assetHandler := http.FileServer(pkg.GetFileSystem())

	e.GET("/*", echo.WrapHandler(assetHandler))

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		fmt.Printf("Получен сигнал s.notify <- s.server.ListenAndServe()\n")
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	s.server.Shutdown(ctx)
	return nil
}

// Shutdown with context
func (s *Server) ShutdownCtx(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
