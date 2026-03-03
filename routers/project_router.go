package routers

import (
	"OperationAndMonitoring/controller/project"
	"github.com/gin-gonic/gin"
)

func ProjectRouter(r *gin.RouterGroup) {
	projectRouter := r.Group("/projects")
	{
		projectRouter.GET("/options", project.Option)
		projectRouter.GET("", project.Page)
		projectRouter.GET("/:projectId/form", project.Form)
		projectRouter.POST("", project.Add)
		projectRouter.PUT("/:projectId", project.Update)
		projectRouter.DELETE("/:projectIds", project.Delete)
	}
}
