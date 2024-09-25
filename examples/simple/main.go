package main

import (
	"context"
	"log/slog"

	"github.com/jantytgat/go-pslog/pkg/pslog"
)

func main() {
	slog.SetDefault(pslog.New())
	slog.Log(context.Background(), pslog.LevelTrace, "Hello from PSLOG")
}
