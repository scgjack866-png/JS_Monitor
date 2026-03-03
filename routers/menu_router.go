package routers

import (
	"OperationAndMonitoring/controller/menu"
	"github.com/gin-gonic/gin"
)

func MenuRouter(r *gin.RouterGroup) {
	menuRouter := r.Group("/menus")
	{
		menuRouter.GET("/routes", menu.Routes)
	}
}
