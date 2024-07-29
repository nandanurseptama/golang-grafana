package routers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	registerEndpointCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_register_request",
		Help: "Total requested to register endpoint",
	})
)

func Register(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info("someone try to register")
		registerEndpointCounter.Inc()
		ctx.JSON(http.StatusOK, map[string]any{"message": "OK"})
	}
}
