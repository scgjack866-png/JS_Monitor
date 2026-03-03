package routers

import (
	"OperationAndMonitoring/controller/order"
	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup) {
	roomRouter := r.Group("/orders")
	{
		roomRouter.GET("/options", order.Option)
		roomRouter.GET("", order.Page)
		roomRouter.GET("/:orderId/form", order.Form)
		roomRouter.POST("", order.Add)
		roomRouter.PUT("/:orderId", order.Update)
		roomRouter.DELETE("/:orderIds", order.Delete)
	}
}
