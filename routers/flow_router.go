package routers

import (
	flow "OperationAndMonitoring/controller/flow"
	"github.com/gin-gonic/gin"
)

func FlowRouter(r *gin.RouterGroup) {
	flowRouter := r.Group("/flows")
	{
		flowRouter.POST("/all", flow.All)
		flowRouter.POST("/snapshot", flow.GetSnapshotUrl)
		flowRouter.POST("/filter", flow.FilterUnderOneM)
		//hostRouter.DELETE("/:ids", host.Delete)
		//hostRouter.GET("/:hostId/form", host.Form)
		//hostRouter.PUT("/:hostId", host.Update)
		//hostRouter.PATCH("/:hostId/password", host.Password)
		//hostRouter.PATCH("/:hostId/status", host.Status)
		//hostRouter.GET("/:hostId/network", host.Network)
	}
}
