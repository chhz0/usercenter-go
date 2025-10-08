package httpmw

import (
	"time"

	"github.com/chhz0/usercenter-go/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(baseLogger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requsetID := ctx.GetHeader("X-Request-Id")
		if requsetID == "" {
			requsetID = genRequestID()
		}

		requestLogger := baseLogger.With(
			zap.String("request_id", requsetID),
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("remote_addr", ctx.Request.RemoteAddr),
			zap.String("user_agent", ctx.Request.UserAgent()),
		)

		_ = logger.WithLogger(ctx, requestLogger)

		ctx.Next()
	}
}

func genRequestID() string {
	return time.Now().Format("20060102150405")
}
