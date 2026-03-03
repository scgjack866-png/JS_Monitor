package routers

import (
	"OperationAndMonitoring/controller/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	authRouter := r.Group("/auth")
	{
		authRouter.POST("/login", auth.Login)
		authRouter.DELETE("/logout", auth.Logout)
		authRouter.POST("/register", auth.Register)
	}
}
