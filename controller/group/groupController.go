package group

import (
	"OperationAndMonitoring/grafana/folders"
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
	var groups []entity.Group
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = 1"})
	utils.Find(&entity.Group{}, &groups, whereOrder...)

	list := recurDeptTreeOptions(0, groups)
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
	var groups []entity.Group
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

	utils.Find(&entity.Group{}, &groups, whereOrder...)

	list := recurDept(0, groups)
	c.JSON(200, utils.SuccessRespon(list))
}

func Add(c *gin.Context) {
	var groupVO vo.DeptVO
	var group entity.Group
	c.BindJSON(&groupVO)

	notfound, _ := utils.First(&entity.Group{Name: groupVO.Name}, &entity.Group{})
	if !notfound {
		c.JSON(200, utils.FailedRespon("分组名称相同，请重新输入！"))
		return
	}

	copier.CopyByTag(&group, &groupVO, "mson")

	body, ok := folders.CreateFolder(group.Name)
	if !ok {
		c.JSON(200, utils.FailedRespon("监控服务端报错！"))
		return
	}
	group.FolderID = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
	group.TreePath = getTreePath(*group.ParentId)
	group.CreateTime = time.Now()
	group.UpdateTime = time.Now()

	err := utils.Create(&group)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("添加成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("groupId")
	deptId := convert.ToUint64(param)
	var group entity.Group
	var hostForm form.DeptForm
	utils.First(&entity.Group{ID: deptId}, &group)
	copier.Copy(&hostForm, &group)
	c.JSON(200, utils.SuccessRespon(hostForm))
}

func Update(c *gin.Context) {
	var groupVO vo.DeptVO
	var group entity.Group
	c.BindJSON(&groupVO)

	notfound, _ := utils.First(&entity.Group{Name: groupVO.Name}, &group)
	if !notfound {
		if group.ID != groupVO.ID {
			c.JSON(200, utils.FailedRespon("文件夹名称已存在，请重新输入！"))
			return
		}
	}

	if groupVO.ParentId == groupVO.ID {
		c.JSON(200, utils.FailedRespon("上级文件夹不能为自身，请重新编辑！"))
		return
	}
	utils.First(&entity.Group{ID: groupVO.ID}, &group)
	ok := folders.UpdateFolder(group.FolderID, groupVO.Name)
	if !ok {
		c.JSON(200, utils.FailedRespon("监控服务端报错！"))
		return
	}

	copier.CopyByTag(&group, &groupVO, "mson")
	group.TreePath = getTreePath(*group.ParentId)
	group.UpdateTime = time.Now()

	err := utils.Updates(&entity.Group{ID: groupVO.ID}, &group)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("更新成功！"))
}

func Delete(c *gin.Context) {
	deleteIds := c.Param("groupIds")
	groupIds := strings.Split(deleteIds, ",")

	for _, id := range groupIds {
		var host entity.Host
		var group entity.Group
		idUint := convert.ToUint64(id)
		utils.First(&entity.Group{ID: idUint}, &group)
		notfound, _ := utils.First(&entity.Host{GroupID: idUint}, &host)
		if !notfound {
			c.JSON(200, utils.FailedRespon(group.Name+"分组下存在服务器，无法删除！"))
			return
		}
		ok := folders.DeleteFolder(group.FolderID)
		if !ok {
			c.JSON(200, utils.FailedRespon("删除报错！"))
			return
		}
		utils.DeleteByID(&entity.Group{}, idUint)
	}
	c.JSON(200, utils.SuccessRespon("删除成功！"))
}
