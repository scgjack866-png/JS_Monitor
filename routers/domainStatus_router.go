package routers

import (
	"OperationAndMonitoring/controller/domainStatus"
	"github.com/gin-gonic/gin"
)

func DomainStatusRouter(r *gin.RouterGroup) {
	domainRouter := r.Group("/domains/status")
	{
		domainRouter.GET("/query", domainStatus.Query)
	}
}
