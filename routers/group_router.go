package routers

import (
	"OperationAndMonitoring/controller/group"
	"github.com/gin-gonic/gin"
)

func GroupRouter(r *gin.RouterGroup) {
	groupRouter := r.Group("/groups")
	{
		groupRouter.GET("/options", group.Option)
		groupRouter.GET("", group.Page)
		groupRouter.GET("/:groupId/form", group.Form)
		groupRouter.POST("", group.Add)
		groupRouter.PUT("/:groupId", group.Update)
		groupRouter.DELETE("/:groupIds", group.Delete)
	}
}
