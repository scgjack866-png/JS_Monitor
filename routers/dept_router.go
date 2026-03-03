package routers

import (
	"OperationAndMonitoring/controller/dept"
	"github.com/gin-gonic/gin"
)

func DeptRouter(r *gin.RouterGroup) {
	deptRouter := r.Group("/dept")
	{
		deptRouter.GET("/options", dept.Option)
		deptRouter.GET("", dept.Page)
		deptRouter.GET("/:deptId/form", dept.Form)
		deptRouter.POST("", dept.Add)
		deptRouter.PUT("/:deptId", dept.Update)
		deptRouter.DELETE("/:deptIds", dept.Delete)
	}
}
