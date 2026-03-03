package dept

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
	var depts []entity.Dept
	utils.Find(&entity.Dept{}, &depts)

	list := recurDeptTreeOptions(0, depts)
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

	var depts []entity.Dept
	var whereOrder []mysql.PageWhereOrder

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

	utils.Find(&entity.Dept{}, &depts, whereOrder...)

	list := recurDept(0, depts)
	c.JSON(200, utils.SuccessRespon(list))
}

func Add(c *gin.Context) {
	var deptVO vo.DeptVO
	var dept entity.Dept
	c.BindJSON(&deptVO)

	copier.CopyByTag(&dept, &deptVO, "mson")
	dept.TreePath = getTreePath(*dept.ParentId)
	dept.CreateTime = time.Now()
	dept.UpdateTime = time.Now()

	err := utils.Create(&dept)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("添加成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("deptId")
	deptId := convert.ToUint64(param)
	var dept entity.Dept
	var deptForm form.DeptForm
	utils.First(&entity.Dept{ID: deptId}, &dept)
	copier.Copy(&deptForm, &dept)
	c.JSON(200, utils.SuccessRespon(deptForm))
}

func Update(c *gin.Context) {
	var deptVO vo.DeptVO
	var dept entity.Dept
	c.BindJSON(&deptVO)

	if deptVO.ParentId == deptVO.ID {
		c.JSON(200, utils.FailedRespon("上级文件夹不能为自身，请重新编辑！"))
		return
	}
	copier.CopyByTag(&dept, &deptVO, "mson")
	dept.TreePath = getTreePath(*dept.ParentId)
	dept.UpdateTime = time.Now()
	err := utils.Updates(&entity.Dept{ID: deptVO.ID}, &dept)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("更新成功！"))
}

func Delete(c *gin.Context) {
	deleteIds := c.Param("deptIds")

	deptIds := strings.Split(deleteIds, ",")
	for _, id := range deptIds {
		idUint := convert.ToUint64(id)
		utils.DeleteByID(&entity.Dept{}, idUint)
	}
	c.JSON(200, utils.SuccessRespon("删除成功！"))
}
