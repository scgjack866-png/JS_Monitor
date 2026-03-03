package routers

import (
	host "OperationAndMonitoring/controller/host"
	"github.com/gin-gonic/gin"
)

func HostRouter(r *gin.RouterGroup) {
	hostRouter := r.Group("/hosts")
	{
		hostRouter.GET("/page", host.Page)
		hostRouter.POST("", host.Create)
		hostRouter.DELETE("/:ids", host.Delete)
		hostRouter.GET("/:hostId/form", host.Form)
		hostRouter.PUT("/:hostId", host.Update)
		hostRouter.PATCH("/:hostId/password", host.Password)
		hostRouter.PATCH("/:hostId/status", host.Status)
		hostRouter.GET("/:hostId/network", host.Network)
	}
}
