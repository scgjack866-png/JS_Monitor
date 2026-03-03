package routers

import (
	"OperationAndMonitoring/controller/zone"
	"github.com/gin-gonic/gin"
)

func ZoneRouter(r *gin.RouterGroup) {
	zoneRouter := r.Group("/zones")
	{
		zoneRouter.GET("/options", zone.Option)
		zoneRouter.GET("", zone.Page)
		zoneRouter.GET("/:zoneId/form", zone.Form)
		zoneRouter.POST("", zone.Add)
		zoneRouter.PUT("/:zoneId", zone.Update)
		zoneRouter.DELETE("/:zoneIds", zone.Delete)
	}
}
