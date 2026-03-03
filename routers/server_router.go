package routers

import (
	"OperationAndMonitoring/controller/server"
	"github.com/gin-gonic/gin"
)

func ServerRouter(r *gin.RouterGroup) {
	serverRouter := r.Group("/servers")
	{
		serverRouter.GET("/page", server.Page)
		serverRouter.POST("", server.Create)
		serverRouter.DELETE("/:ids", server.Delete)
		serverRouter.GET("/:serverId/form", server.Form)
		serverRouter.PUT("/:serverId", server.Update)
		serverRouter.PATCH("/:serverId/password", server.Password)
		serverRouter.PATCH("/:serverId/status", server.Status)
		serverRouter.GET("/:serverId/network", server.Network)
	}
}
