package routers

import (
	"OperationAndMonitoring/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/users")
	{
		userRouter.GET("/me", user.Me)
		userRouter.GET("/page", user.Page)
		userRouter.POST("", user.Create)
		userRouter.DELETE("/:ids", user.Delete)
		userRouter.GET("/:userId/form", user.Form)
		userRouter.PUT("/:userId", user.Update)
		userRouter.PATCH("/:userId/password", user.Password)
		userRouter.PATCH("/:userId/status", user.Status)
	}
}
