package server

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/routes"
)

type Server struct {
    Echo   *echo.Echo
}

func NewServer(logger *slog.Logger, enableAccessLog bool) *Server {
    e := echo.New()
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

    if enableAccessLog {
        e.Use(middleware.Logger())
    }

    // ロガーの設定
    if logger != nil {
        logger.Info("Logger initialized")
    }

    // ルーティングの設定
    routes.SetupRoutes(e)

    return &Server{
        Echo:   e,
    }
}

// Start runs the Echo server on the specified port
func (s *Server) Start(port string) error {
    logging.Logger().Info("Starting server", "port", port)
    return s.Echo.Start(port)
}
