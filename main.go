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
	"github.com/nandanurseptama/golang_grafana/logger"
	"github.com/nandanurseptama/golang_grafana/routers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func safeEnv(name string, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}
func main() {
	logFilePath := safeEnv("LOG_FILE_PATH", "volumes/var/app.log")
	logClient := logger.New(logFilePath)
	port := safeEnv("PORT", "8080")
	promServerPort := safeEnv("PROMETHEUS_SERVER_PORT", "2221")
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		logClient.Info("new request come", slog.Any("path", ctx.FullPath()))
		ctx.Next()
		logClient.Info(
			"request end",
			slog.Any("path", ctx.FullPath()),
			slog.Any("status", ctx.Writer.Status()),
		)
	})

	r.POST("/api/login", routers.Login(logClient))
	r.POST("/api/register", routers.Register(logClient))

	promServer := gin.Default()
	promServer.GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	go func() {
		<-sig
		slog.Info("shutting down application")
		cancel()
	}()

	go func() {
		if err := promServer.Run(fmt.Sprintf(":%s", promServerPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed running prom server", slog.Any("err", err))
			panic(err)
		}
	}()

	go func() {
		if err := r.Run(fmt.Sprintf(":%s", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("failed demo app server", slog.Any("err", err))
			panic(err)
		}
	}()

	<-ctx.Done()
}
