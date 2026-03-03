package routers

import (
	"OperationAndMonitoring/controller/domain"
	"github.com/gin-gonic/gin"
)

func DomainRouter(r *gin.RouterGroup) {
	domainRouter := r.Group("/domains")
	{
		domainRouter.GET("/page", domain.Page)
		domainRouter.POST("", domain.Create)
		domainRouter.DELETE("/:domainIds", domain.Delete)
		domainRouter.GET("/:domainId/form", domain.Form)
		domainRouter.PUT("/:domainId", domain.Update)
		domainRouter.PATCH("/:domainId/status", domain.Status)
	}
}
