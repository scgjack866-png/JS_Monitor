package domainStatus

import (
	"OperationAndMonitoring/prometheus"
	"OperationAndMonitoring/utils"
	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	domainStatuss, _ := prometheus.GetQuery()
	c.JSON(200, utils.SuccessRespon(domainStatuss))
}
