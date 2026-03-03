package role

import (
	"github.com/gin-gonic/gin"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/utils"
)

func Options(c *gin.Context) {
	var roles []entity.Role
	utils.Find(&entity.Role{}, &roles)

	list := recurDeptTreeOptions(roles)
	c.JSON(200, utils.SuccessRespon(list))
}
