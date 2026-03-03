package user

import (
	"OperationAndMonitoring/controller/common"
	"OperationAndMonitoring/controller/menu"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/form"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	convert2 "OperationAndMonitoring/utils/convert"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	copier "github.com/ybzhanghx/copier"
	"strconv"
	"strings"
	"time"
)

func Me(c *gin.Context) {
	var user entity.User
	var userInfo vo.UserInfo
	userID, _ := c.Get(common.USER_ID_Key)
	userIDU, _ := convert2.ToUint64E(userID)

	_, err := utils.First(&entity.User{ID: userIDU}, &user)
	if err != nil {
		fmt.Println("qq")
	}
	copier.CopyByTag(&userInfo, &user, "mson")

	var roles []entity.Role

	d := entity.User{
		ID: userIDU,
	}

	utils.Test(&d, &roles, "Users", "Roles")

	for _, role := range roles {
		var menus []entity.Menu
		m := entity.Role{
			ID: role.ID,
		}
		utils.Test(&m, &menus, "Roles", "Menus")

		for _, menu := range menus {
			if menu.Perm != "" {
				userInfo.Perms = append(userInfo.Perms, menu.Perm)
			}
		}

		userInfo.Roles = append(userInfo.Roles, role.Code)
	}

	c.JSON(200, utils.SuccessRespon(userInfo))
}

func Page(c *gin.Context) {
	page, _ := convert2.ToIntE(c.Query("pageNum"))
	limit, _ := convert2.ToIntE(c.Query("pageSize"))
	deptId, _ := convert2.ToUint64E(c.Query("deptId"))
	key := c.Query("keywords")
	status, _ := strconv.Atoi(c.Query("status"))
	var whereOrder []mysql.PageWhereOrder

	if lo.IsNotEmpty(key) {
		v := "%" + key + "%"
		var arr []interface{}
		arr = append(arr, v)
		arr = append(arr, v)
		arr = append(arr, v)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "username like ? or nickname like ? or mobile like ?", Value: arr})
	}
	if lo.IsNotEmpty(c.Query("status")) {
		var arr []interface{}
		arr = append(arr, status)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = ?", Value: arr})
	}

	if deptId > 1 {
		var arr []interface{}
		arr = append(arr, deptId)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "dept_id = ?", Value: arr})
	}

	var total int64
	users := []entity.User{}
	list := []vo.Admin{}
	err := utils.GetPage(&entity.User{}, &entity.User{}, &users, page, limit, &total, whereOrder...)
	for _, user := range users {

		var admin vo.Admin
		var dept entity.Dept
		copier.CopyByTag(&admin, &user, "mson")
		// 解决时间没法复制
		admin.CreateTime = user.CreateTime.Format("2006-01-02")

		if user.Gender == 1 {
			admin.Gender = "男"
		} else if user.Gender == 2 {
			admin.Gender = "女"
		} else {
			admin.Gender = "未知"
		}
		utils.First(&entity.Dept{ID: user.DeptId}, &dept)
		admin.DeptName = dept.Name
		admin.RoleNames = menu.GetRoles(user.ID)

		list = append(list, admin)
	}
	if err != nil {
		c.JSON(200, utils.SuccessRespon("获取数据库数据失败"))
		return
	}
	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = total
	c.JSON(200, utils.SuccessRespon(pageR))

}

func Create(c *gin.Context) {
	var userForm form.UserForm
	var user entity.User
	c.BindJSON(&userForm)
	has, _ := utils.First(&entity.User{UserName: userForm.UserName}, &entity.User{})
	if !has {
		c.JSON(200, utils.FailedRespon("用户名重复，请重新输入！"))
		return
	}
	copier.CopyByTag(&user, &userForm, "mson")
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	utils.Create(&user)
	if len(userForm.RoleIds) != 0 {
		utils.First(&entity.User{UserName: userForm.UserName}, &user)
		for _, id := range userForm.RoleIds {
			var userRole entity.UserRole
			userRole.UserID = user.ID
			userRole.RoleID = id
			utils.Create(&userRole)
		}
	}
	c.JSON(200, utils.SuccessRespon("添加新用户成功"))
}

func Delete(c *gin.Context) {
	var idUints []uint64
	param := c.Param("ids")
	ids := strings.Split(param, ",")
	for _, id := range ids {
		idUints = append(idUints, convert2.ToUint64(id))
	}
	var user entity.User
	_, err := utils.DeleteByIDS(&user, idUints)
	if err != nil {
		c.JSON(200, utils.FailedRespon("删除用户发生错误！"))
	}
	c.JSON(200, utils.SuccessRespon("删除用户成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("userId")
	userId, _ := convert2.ToUint64E(param)

	var user entity.User
	var admin vo.Admin
	utils.First(&entity.User{ID: userId}, &user)

	var dept entity.Dept
	copier.CopyByTag(&admin, &user, "mson")
	// 解决时间没法复制
	admin.CreateTime = convert2.TimeToString(user.CreateTime)
	if user.Gender == 1 {
		admin.Gender = "男"
	} else if user.Gender == 2 {
		admin.Gender = "女"
	} else {
		admin.Gender = "未知"
	}
	utils.First(&entity.Dept{ID: user.DeptId}, &dept)
	admin.DeptName = dept.Name
	admin.RoleNames = menu.GetRoles(user.ID)

	c.JSON(200, utils.SuccessRespon(admin))
}

func Update(c *gin.Context) {
	param := c.Param("userId")
	userId, _ := convert2.ToUint64E(param)

	var userForm form.UserForm
	var user entity.User

	c.BindJSON(&userForm)
	copier.CopyByTag(&user, &userForm, "mson")
	user.UpdateTime = time.Now()

	utils.Updates(&entity.User{ID: userId}, &user)

	utils.DeleteByWhere(&entity.UserRole{}, &entity.UserRole{UserID: userId})

	if len(userForm.RoleIds) != 0 {
		utils.First(&entity.User{UserName: userForm.UserName}, &user)
		for _, id := range userForm.RoleIds {
			var userRole entity.UserRole
			userRole.UserID = user.ID
			userRole.RoleID = id
			utils.Create(&userRole)
		}
	}

	c.JSON(200, utils.SuccessRespon("更新用户成功！"))
}

func Password(c *gin.Context) {
	param := c.Param("userId")
	userId, _ := convert2.ToUint64E(param)
	password, err := utils.EncryptPassword(c.Query("password"))
	if err != nil {
		c.JSON(200, utils.FailedRespon("重置密码失败！"))
	}
	err = utils.Updates(&entity.User{ID: userId}, &entity.User{Password: password})
	if err != nil {
		c.JSON(200, utils.FailedRespon("重置密码失败！"))
	}
	c.JSON(200, utils.SuccessRespon("重置密码成功！"))
}

func Status(c *gin.Context) {
	param := c.Param("userId")
	userId, _ := convert2.ToUint64E(param)
	status := c.Query("status")
	intS := convert2.ToInt(status)
	err := utils.Updates(&entity.User{ID: userId}, &entity.User{Status: &intS})
	if err != nil {
		c.JSON(200, utils.FailedRespon("更新用户状态失败！"))
	}
	c.JSON(200, utils.SuccessRespon("更新用户状态成功！"))
}
