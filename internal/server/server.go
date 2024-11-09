package server

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
    Echo   *echo.Echo
    Logger *slog.Logger
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

    // シンプルなエンドポイントの設定
    e.GET("/ping", func(c echo.Context) error {
        logger.Info("Ping endpoint called")
        return c.String(http.StatusOK, "pong")
    })

    e.GET("/hello", func(c echo.Context) error {
        logger.Info("Hello endpoint called")
        return c.JSON(http.StatusOK, map[string]string{"message": "Hello, world"})
    })

    return &Server{
        Echo:   e,
        Logger: logger,
    }
}

// Start runs the Echo server on the specified port
func (s *Server) Start(port string) error {
    s.Logger.Info("Starting server", "port", port)
    return s.Echo.Start(port)
}
