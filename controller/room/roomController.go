package room

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
	var rooms []entity.Room
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = 1"})
	utils.Find(&entity.Room{}, &rooms, whereOrder...)

	list := recurDeptTreeOptions(0, rooms)
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
	var rooms []entity.Room
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

	utils.Find(&entity.Room{}, &rooms, whereOrder...)

	list := recurDept(0, rooms)
	c.JSON(200, utils.SuccessRespon(list))
}

func Add(c *gin.Context) {
	var roomVO vo.DeptVO
	var room entity.Room
	c.BindJSON(&roomVO)

	notfound, _ := utils.First(&entity.Room{Name: roomVO.Name}, &entity.Room{})
	if !notfound {
		c.JSON(200, utils.FailedRespon("分组名称相同，请重新输入！"))
		return
	}

	copier.CopyByTag(&room, &roomVO, "mson")

	room.TreePath = getTreePath(*room.ParentId)
	room.CreateTime = time.Now()
	room.UpdateTime = time.Now()

	err := utils.Create(&room)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("添加成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("roomId")
	deptId := convert.ToUint64(param)
	var room entity.Room
	var hostForm form.DeptForm
	utils.First(&entity.Room{ID: deptId}, &room)
	copier.Copy(&hostForm, &room)
	c.JSON(200, utils.SuccessRespon(hostForm))
}

func Update(c *gin.Context) {
	var roomVO vo.DeptVO
	var room entity.Room
	c.BindJSON(&roomVO)

	notfound, _ := utils.First(&entity.Room{Name: roomVO.Name}, &room)
	if !notfound {
		if room.ID != roomVO.ID {
			c.JSON(200, utils.FailedRespon("文件夹名称已存在，请重新输入！"))
			return
		}
	}

	if roomVO.ParentId == roomVO.ID {
		c.JSON(200, utils.FailedRespon("上级文件夹不能为自身，请重新编辑！"))
		return
	}
	utils.First(&entity.Room{ID: roomVO.ID}, &room)

	copier.CopyByTag(&room, &roomVO, "mson")
	room.TreePath = getTreePath(*room.ParentId)
	room.UpdateTime = time.Now()

	err := utils.Updates(&entity.Room{ID: roomVO.ID}, &room)
	if err != nil {
		c.JSON(200, utils.FailedRespon("服务端报错！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("更新成功！"))
}

func Delete(c *gin.Context) {
	deleteIds := c.Param("roomIds")
	roomIds := strings.Split(deleteIds, ",")

	for _, id := range roomIds {
		var host entity.Host
		var room entity.Room
		idUint := convert.ToUint64(id)
		utils.First(&entity.Room{ID: idUint}, &room)
		notfound, _ := utils.First(&entity.Host{GroupID: idUint}, &host)
		if !notfound {
			c.JSON(200, utils.FailedRespon(room.Name+"分组下存在服务器，无法删除！"))
			return
		}
		utils.DeleteByID(&entity.Room{}, idUint)
	}
	c.JSON(200, utils.SuccessRespon("删除成功！"))
}
