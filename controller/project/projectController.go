package project

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
	var projects []entity.Project
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = 1"})
	utils.Find(&entity.Project{}, &projects, whereOrder...)

	list := recurDeptTreeOptions(0, projects)
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
	var projects []entity.Project
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

	utils.Find(&entity.Project{}, &projects, whereOrder...)

	list := recurDept(0, projects)
	c.JSON(200, utils.SuccessRespon(list))
}

func Add(c *gin.Context) {
	var projectVO vo.DeptVO
	var project entity.Project
	c.BindJSON(&projectVO)

	notfound, _ := utils.First(&entity.Project{Name: projectVO.Name}, &entity.Project{})
	if !notfound {
		c.JSON(200, utils.FailedRespon("分组名称相同，请重新输入！"))
		return
	}

	copier.CopyByTag(&project, &projectVO, "mson")

	project.TreePath = getTreePath(*project.ParentId)
	project.CreateTime = time.Now()
	project.UpdateTime = time.Now()

	err := utils.Create(&project)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("添加成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("projectId")
	deptId := convert.ToUint64(param)
	var project entity.Project
	var hostForm form.DeptForm
	utils.First(&entity.Project{ID: deptId}, &project)
	copier.Copy(&hostForm, &project)
	c.JSON(200, utils.SuccessRespon(hostForm))
}

func Update(c *gin.Context) {
	var projectVO vo.DeptVO
	var project entity.Project
	c.BindJSON(&projectVO)

	notfound, _ := utils.First(&entity.Project{Name: projectVO.Name}, &project)
	if !notfound {
		if project.ID != projectVO.ID {
			c.JSON(200, utils.FailedRespon("文件夹名称已存在，请重新输入！"))
			return
		}
	}

	if projectVO.ParentId == projectVO.ID {
		c.JSON(200, utils.FailedRespon("上级文件夹不能为自身，请重新编辑！"))
		return
	}
	utils.First(&entity.Project{ID: projectVO.ID}, &project)

	copier.CopyByTag(&project, &projectVO, "mson")
	project.TreePath = getTreePath(*project.ParentId)
	project.UpdateTime = time.Now()

	err := utils.Updates(&entity.Project{ID: projectVO.ID}, &project)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("更新成功！"))
}

func Delete(c *gin.Context) {
	deleteIds := c.Param("projectIds")
	projectIds := strings.Split(deleteIds, ",")

	for _, id := range projectIds {
		var host entity.Host
		var project entity.Project
		idUint := convert.ToUint64(id)
		utils.First(&entity.Project{ID: idUint}, &project)
		notfound, _ := utils.First(&entity.Host{GroupID: idUint}, &host)
		if !notfound {
			c.JSON(200, utils.FailedRespon(project.Name+"分组下存在服务器，无法删除！"))
			return
		}
		utils.DeleteByID(&entity.Project{}, idUint)
	}
	c.JSON(200, utils.SuccessRespon("删除成功！"))
}
