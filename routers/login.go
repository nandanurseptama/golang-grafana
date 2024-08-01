package routers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login api path handler
func Login(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body User

		// bind request body
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				map[string]any{"message": "failed to bind request body"},
			)
			ctx.Abort()
			totalFailedLoginCounter.Inc()
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
			totalFailedLoginCounter.Inc()
			return
		}

		var findUser = User{}

		for _, u := range registeredUser {
			if u.Email == body.Email {
				findUser = u
				break
			}
		}

		// guard when user not registered
		if findUser.Email == "" {
			ctx.JSON(
				http.StatusBadRequest,
				map[string]any{"message": "email or password not match"},
			)
			ctx.Abort()
			totalFailedLoginCounter.Inc()
			return
		}

		// guard when password not match
		if findUser.Password != body.Password {
			ctx.JSON(
				http.StatusBadRequest,
				map[string]any{"message": "email or password not match"},
			)
			ctx.Abort()
			totalFailedLoginCounter.Inc()
			return
		}
		traceId := ctx.GetString("traceId")
		logger.Info("someone try to login", slog.Any("traceId", traceId))
		ctx.JSON(http.StatusOK, map[string]any{"message": "OK"})
		ctx.Abort()
		totalSuccessLoginCounter.Inc()
	}
}
