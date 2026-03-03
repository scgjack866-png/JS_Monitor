package auth

import (
	model "OperationAndMonitoring/controller/common"
	jwt "OperationAndMonitoring/middleware"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/cache"
	"OperationAndMonitoring/utils/convert"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"time"
)

// 设置jwt密钥secret

var (
	jwtSecret = []byte("123")
)

type Claims struct {
	UserID string `json:"userid"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user entity.User
	has, _ := utils.First(&entity.User{UserName: username}, &user)

	if has {
		c.JSON(200, utils.FailedRespon("该账号不存在，请先注册！"))
		return
	}

	uuid := utils.GetUUID()

	if lo.IsEmpty(*user.Status) {
		c.JSON(200, utils.FailedRespon("账号已经被禁用！"))
		return
	}

	if utils.EqualsPassword(password, user.Password) {
		userInfo := make(map[string]string)
		exp := convert.ToString(time.Now().Add(time.Hour * time.Duration(154)).Unix())
		userInfo["exp"] = exp // 1H
		userInfo["iat"] = convert.ToString(time.Now().Unix())
		userInfo["uuid"] = uuid
		aToken := jwt.CreateToken(userInfo)
		userInfo["tokenType"] = "Bearer"
		userInfo["refreshToken"] = ""
		userInfo["accessToken"] = aToken
		userInfo["expires"] = exp
		//initialize.CsbinAddRoleForUser(user.ID)

		cache.Set([]byte(uuid), []byte(convert.ToString(user.ID)), 60*60*168)

		c.JSON(200, utils.SuccessRespon(userInfo))
	} else {
		c.JSON(200, utils.FailedRespon("密码错误！请重新输入！"))
	}

}

func Logout(c *gin.Context) {
	tokenKey := c.GetHeader(model.TOKEN_KEY)
	if lo.IsEmpty(tokenKey) {
		utils.FailedRespon("token为空")
		return
	}
	u, ok := jwt.ParseToken(tokenKey)
	if !ok {
		utils.FailedRespon("解析失败")
		return
	}
	cid := u["uuid"]
	if lo.IsEmpty(cid) {
		utils.FailedRespon("uid为空")
		return
	}
	cache.Del([]byte(cid))
	c.JSON(200, utils.SuccessRespon("操作成功"))
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user entity.User
	var userRole entity.UserRole

	has, _ := utils.First(&entity.User{UserName: username}, &user)
	if !has {
		c.JSON(200, utils.FailedRespon("账号已存在，请前往登录！"))
		return
	}

	user.UserName = username
	user.NickName = username
	var err error
	user.Password, err = utils.EncryptPassword(password)
	if err != nil {
		c.JSON(200, utils.FailedRespon("系统发生错误！"))
	}

	user.Gender = 1
	user.DeptId = 3
	user.Avatar = "https://s2.loli.net/2022/04/07/gw1L2Z5sPtS8GIl.gif"

	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	utils.Create(&user)

	utils.First(&entity.User{UserName: username}, &user)
	userRole.UserID = user.ID
	userRole.RoleID = 3
	utils.Create(&userRole)
	c.JSON(200, utils.SuccessRespon(`{"username": `+username+`}`))

}
