package testutils

import (
	"context"

	"log/slog"
)

// DiscardHandler is a handler that discards log output during testing
// DiscardHandlerはテスト時のログ出力を無視するためのハンドラー
type DiscardHandler struct{}

func (d *DiscardHandler) Handle(ctx context.Context, r slog.Record) error {
	return nil // ログ出力を無視
}

func (d *DiscardHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return false // ログを無効化
}

func (d *DiscardHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return d // 属性を無視して返す
}

func (d *DiscardHandler) WithGroup(name string) slog.Handler {
	return d // グループを無視して返す
}
