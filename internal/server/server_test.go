package server

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/testutils"
)

var (
	logger *slog.Logger
	s      *Server
)

func TestMain(m *testing.M) {
	logger = slog.New(&testutils.DiscardHandler{})
	s = NewServer(logger, false)

	code := m.Run()

	os.Exit(code)
}

func TestPingEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	c := s.Echo.NewContext(req, rec)

	s.Echo.ServeHTTP(c.Response(), c.Request())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "pong", rec.Body.String())
}

func TestHelloEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := s.Echo.NewContext(req, rec)

	s.Echo.ServeHTTP(c.Response(), c.Request())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "Hello, world"}`, rec.Body.String())
}
