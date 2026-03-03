package menu

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"github.com/gin-gonic/gin"
)

func Routes(c *gin.Context) {
	var menus []entity.Menu
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})

	utils.Find(&entity.Menu{}, &menus, whereOrder...)
	list := RecurRoutes(0, menus)
	c.JSON(200, utils.SuccessRespon(list))

}
