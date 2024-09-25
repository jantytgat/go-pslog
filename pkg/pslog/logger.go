package pslog

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	return slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource:   false,
			Level:       LevelTrace,
			ReplaceAttr: ReplaceAttrs,
		}))
}
