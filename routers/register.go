package routers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register api path handler
func Register(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body User

		// bind request body
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				map[string]any{"message": "failed to bind request body"},
			)
			ctx.Abort()
			totalFailedRegisterCounter.Inc()
			return
		}

		// guard when email and password empty
		// send response error
		if body.Email == "" || body.Password == "" {
			ctx.JSON(
				http.StatusBadRequest,
				map[string]any{"message": "email and password required"},
			)
			ctx.Abort()
			totalFailedRegisterCounter.Inc()
			return
		}

		var findUser = User{}

		for _, u := range registeredUser {
			if u.Email == body.Email {
				findUser = u
				break
			}
		}

		// guard when user email found
		if findUser.Email != "" {
			ctx.JSON(
				http.StatusBadRequest,
				map[string]any{"message": "email already registered"},
			)
			ctx.Abort()
			totalFailedRegisterCounter.Inc()
			return
		}
		traceId := ctx.GetString("traceId")
		logger.Info("someone try to register", slog.Any("traceId", traceId))
		ctx.JSON(http.StatusOK, map[string]any{"message": "OK"})
		totalSuccessRegisterCounter.Inc()
	}
}
