package router

import (
	"net/http"

	usersrouter "github.com/chhz0/usercenter-go/internal/usercenter/router/users"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":   200,
			"status": "ok",
		})
	})

	// register
	usersrouter.Register(r.Group("/users"))
}
