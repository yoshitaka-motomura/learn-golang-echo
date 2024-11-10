package server

import (
	"log/slog"
	"net/http"

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
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

    if enableAccessLog {
        e.Use(middleware.Logger())
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
