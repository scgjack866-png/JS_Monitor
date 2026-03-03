package routers

import (
	"OperationAndMonitoring/controller/ipsec"
	"github.com/gin-gonic/gin"
)

func IpsecRouter(r *gin.RouterGroup) {
	ipsecRouter := r.Group("/ipsecs")
	{
		ipsecRouter.GET("/page", ipsec.Page)
		ipsecRouter.POST("", ipsec.Refresh)
		ipsecRouter.POST("/online", ipsec.RefreshOnlineNum)
		ipsecRouter.DELETE("/:ids", ipsec.Delete)
		ipsecRouter.GET("/:ipsecId/form", ipsec.Form)
		ipsecRouter.PUT("/:ipsecId", ipsec.Update)
		ipsecRouter.PATCH("/:ipsecId/password", ipsec.Password)
		ipsecRouter.PATCH("/:ipsecId/status", ipsec.Status)

	}
}
