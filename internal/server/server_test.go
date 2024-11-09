package server

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/testutils"
)

var (
	s *Server
)

func TestMain(m *testing.M) {
	// テスト用のロガーを初期化
	logging.InitLogger(slog.New(&testutils.DiscardHandler{}))

	// サーバーを初期化し、ロガーを設定
	s = NewServer(logging.Logger(), false)

	// テストの実行
	code := m.Run()

	// 終了
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

func TestTodosEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	rec := httptest.NewRecorder()
	c := s.Echo.NewContext(req, rec)

	s.Echo.ServeHTTP(c.Response(), c.Request())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `["todo1", "todo2", "todo3"]`, rec.Body.String())
}
