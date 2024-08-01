package logger

import (
	"io"
	"log/slog"
	"os"
)

func New(logFile string) *slog.Logger {
	f, err := os.OpenFile(logFile, os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(os.Stdout, f)

	return slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}
