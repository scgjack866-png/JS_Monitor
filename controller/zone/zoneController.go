package zone

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/form"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/convert"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	copier "github.com/ybzhanghx/copier"
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
	var zones []entity.Zone
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = 1"})
	utils.Find(&entity.Zone{}, &zones, whereOrder...)

	list := recurDeptTreeOptions(0, zones)
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
	var zones []entity.Zone
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

	utils.Find(&entity.Zone{}, &zones, whereOrder...)

	list := recurDept(0, zones)
	c.JSON(200, utils.SuccessRespon(list))
}

func Add(c *gin.Context) {
	var zoneVO vo.DeptVO
	var zone entity.Zone
	c.BindJSON(&zoneVO)

	notfound, _ := utils.First(&entity.Zone{Name: zoneVO.Name}, &entity.Zone{})
	if !notfound {
		c.JSON(200, utils.FailedRespon("分组名称相同，请重新输入！"))
		return
	}

	copier.CopyByTag(&zone, &zoneVO, "mson")

	zone.TreePath = getTreePath(*zone.ParentId)
	zone.CreateTime = time.Now()
	zone.UpdateTime = time.Now()

	err := utils.Create(&zone)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("添加成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("zoneId")
	deptId := convert.ToUint64(param)
	var zone entity.Zone
	var hostForm form.DeptForm
	utils.First(&entity.Zone{ID: deptId}, &zone)
	copier.Copy(&hostForm, &zone)
	c.JSON(200, utils.SuccessRespon(hostForm))
}

func Update(c *gin.Context) {
	var zoneVO vo.DeptVO
	var zone entity.Zone
	c.BindJSON(&zoneVO)

	notfound, _ := utils.First(&entity.Zone{Name: zoneVO.Name}, &zone)
	if !notfound {
		if zone.ID != zoneVO.ID {
			c.JSON(200, utils.FailedRespon("文件夹名称已存在，请重新输入！"))
			return
		}
	}

	if zoneVO.ParentId == zoneVO.ID {
		c.JSON(200, utils.FailedRespon("上级文件夹不能为自身，请重新编辑！"))
		return
	}
	utils.First(&entity.Zone{ID: zoneVO.ID}, &zone)

	copier.CopyByTag(&zone, &zoneVO, "mson")
	zone.TreePath = getTreePath(*zone.ParentId)
	zone.UpdateTime = time.Now()

	err := utils.Updates(&entity.Zone{ID: zoneVO.ID}, &zone)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("更新成功！"))
}

func Delete(c *gin.Context) {
	deleteIds := c.Param("zoneIds")
	zoneIds := strings.Split(deleteIds, ",")

	for _, id := range zoneIds {
		var host entity.Host
		var zone entity.Zone
		idUint := convert.ToUint64(id)
		utils.First(&entity.Zone{ID: idUint}, &zone)
		notfound, _ := utils.First(&entity.Host{GroupID: idUint}, &host)
		if !notfound {
			c.JSON(200, utils.FailedRespon(zone.Name+"分组下存在服务器，无法删除！"))
			return
		}
		utils.DeleteByID(&entity.Zone{}, idUint)
	}
	c.JSON(200, utils.SuccessRespon("删除成功！"))
}
