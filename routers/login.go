package routers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	loginEndpointCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_login_request",
		Help: "Total requested to login endpoint",
	})
)

func Login(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := ctx.GetString("traceId")
		logger.Info("someone try to login", slog.Any("traceId", traceId))
		loginEndpointCounter.Inc()
		ctx.JSON(http.StatusOK, map[string]any{"message": "OK"})
	}
}
