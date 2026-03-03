package routers

import (
	"OperationAndMonitoring/controller/node"
	"github.com/gin-gonic/gin"
)

func NodeRouter(r *gin.RouterGroup) {
	domainRouter := r.Group("/nodes")
	{
		domainRouter.GET("/page", node.Page)
		domainRouter.POST("", node.Create)
		domainRouter.DELETE("/:ids", node.Delete)
		domainRouter.GET("/:nodeId/form", node.Form)
		domainRouter.PUT("/:nodeId", node.Update)
		domainRouter.PATCH("/:nodeId/status", node.Status)
	}
}
