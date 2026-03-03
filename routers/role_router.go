package routers

import (
	"OperationAndMonitoring/controller/role"
	"github.com/gin-gonic/gin"
)

func RoleRouter(r *gin.RouterGroup) {
	roleRouter := r.Group("/roles")
	{
		roleRouter.GET("/options", role.Options)

	}
}
