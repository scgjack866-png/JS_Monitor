package order

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/form"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/convert"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/ybzhanghx/copier"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/**
 * 部门下拉选项
 *
 * @return
 */
func Option(c *gin.Context) {
	var orders []entity.Order
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = 1"})
	utils.Find(&entity.Order{}, &orders, whereOrder...)

	list := recurDeptTreeOptions(0, orders)
	c.JSON(200, utils.SuccessRespon(list))
}

func Page(c *gin.Context) {
	key := c.Query("keywords")
	status, _ := strconv.Atoi(c.Query("status"))

	keywords, err := url.QueryUnescape(key)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	var orders []entity.Order
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})

	if lo.IsNotEmpty(keywords) {
		v := "%" + keywords + "%"
		var arr []interface{}
		arr = append(arr, v)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "name like ?", Value: arr})
	}
	if lo.IsNotEmpty(c.Query("status")) {
		var arr []interface{}
		arr = append(arr, status)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = ?", Value: arr})
	}

	utils.Find(&entity.Order{}, &orders, whereOrder...)

	list := recurDept(0, orders)
	c.JSON(200, utils.SuccessRespon(list))
}

func Add(c *gin.Context) {
	var orderVO vo.DeptVO
	var order entity.Order
	c.BindJSON(&orderVO)

	notfound, _ := utils.First(&entity.Order{Name: orderVO.Name}, &entity.Order{})
	if !notfound {
		c.JSON(200, utils.FailedRespon("分组名称相同，请重新输入！"))
		return
	}

	copier.CopyByTag(&order, &orderVO, "mson")

	order.TreePath = getTreePath(*order.ParentId)
	order.CreateTime = time.Now()
	order.UpdateTime = time.Now()

	err := utils.Create(&order)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("添加成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("orderId")
	deptId := convert.ToUint64(param)
	var order entity.Order
	var hostForm form.DeptForm
	utils.First(&entity.Order{ID: deptId}, &order)
	copier.Copy(&hostForm, &order)
	c.JSON(200, utils.SuccessRespon(hostForm))
}

func Update(c *gin.Context) {
	var orderVO vo.DeptVO
	var order entity.Order
	c.BindJSON(&orderVO)

	notfound, _ := utils.First(&entity.Order{Name: orderVO.Name}, &order)
	if !notfound {
		if order.ID != orderVO.ID {
			c.JSON(200, utils.FailedRespon("文件夹名称已存在，请重新输入！"))
			return
		}
	}

	if orderVO.ParentId == orderVO.ID {
		c.JSON(200, utils.FailedRespon("上级文件夹不能为自身，请重新编辑！"))
		return
	}
	utils.First(&entity.Order{ID: orderVO.ID}, &order)

	copier.CopyByTag(&order, &orderVO, "mson")
	order.TreePath = getTreePath(*order.ParentId)
	order.UpdateTime = time.Now()

	err := utils.Updates(&entity.Order{ID: orderVO.ID}, &order)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("更新成功！"))
}

func Delete(c *gin.Context) {
	deleteIds := c.Param("orderIds")
	orderIds := strings.Split(deleteIds, ",")

	for _, id := range orderIds {
		var host entity.Host
		var order entity.Order
		idUint := convert.ToUint64(id)
		utils.First(&entity.Order{ID: idUint}, &order)
		notfound, _ := utils.First(&entity.Host{GroupID: idUint}, &host)
		if !notfound {
			c.JSON(200, utils.FailedRespon(order.Name+"分组下存在服务器，无法删除！"))
			return
		}
		utils.DeleteByID(&entity.Order{}, idUint)
	}
	c.JSON(200, utils.SuccessRespon("删除成功！"))
}
