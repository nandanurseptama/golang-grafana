package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nandanurseptama/golang_grafana/routers"
)

// get env variable, if not found use `defaultValue`
func safeEnv(name string, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}
func main() {
	logClient := slog.Default()
	port := safeEnv("PORT", "8080")
	r := gin.New()

	r.Use(func(ctx *gin.Context) {
		traceId := uuid.NewString()
		ctx.Set("traceId", traceId)
		logClient.Info("new request come", slog.Any("traceId", traceId), slog.Any("path", ctx.FullPath()))
		ctx.Next()
		logClient.Info(
			"request end",
			slog.Any("traceId", traceId),
			slog.Any("path", ctx.FullPath()),
			slog.Any("status", ctx.Writer.Status()),
		)
	})

	r.POST("/api/login", routers.Login(logClient))
	r.POST("/api/register", routers.Register(logClient))

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sig
		slog.Info("receive signal ", slog.Any("signal", sig))
		slog.Info("shutting down application")
		cancel()
	}()

	go func() {
		if err := r.Run(fmt.Sprintf(":%s", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed demo app server", slog.Any("err", err))
			panic(err)
		}
	}()

	<-ctx.Done()
}
