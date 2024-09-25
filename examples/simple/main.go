package main

import (
	"log/slog"

	"github.com/jantytgat/go-pslog/pkg/pslog"
)

func main() {
	slog.SetDefault(pslog.New())
	slog.Info("Hello from PSLOG")
}
