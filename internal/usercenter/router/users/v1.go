package usersrouter

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	_ = r.Group("/v1", v1RootMw()...)
	{
		// user v1 router
	}
}
