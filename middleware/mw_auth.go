package middleware

import (
	"OperationAndMonitoring/controller/common"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/cache"
	"OperationAndMonitoring/utils/convert"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		var uuid string
		if t := c.GetHeader(common.TOKEN_KEY); t != "" {

			// 用于服务器直接通过api添加服务器，更改服务器
			// 不用时可以注释
			if t == "BccNQJzLJIaboc2tvkGlop1z66H3ROnf" {
				c.Next()
				return
			}

			userInfo, ok := ParseToken(t)
			if !ok {
				c.JSON(200, utils.FailedTokenRespon("token 无效,返回登录界面！"))
				c.Abort()
				return
			}
			exptimestamp, _ := strconv.ParseInt(userInfo["exp"], 10, 64)
			exp := time.Unix(exptimestamp, 0)
			ok = exp.After(time.Now())
			if !ok {

				c.JSON(200, utils.FailedTokenRespon("token 过期,返回登录界面！"))
				c.Abort()
				return
			}
			uuid = userInfo["uuid"]

		}
		if uuid != "" {
			//查询用户ID
			val, err := cache.Get([]byte(uuid))
			if err != nil {

				c.JSON(200, utils.FailedTokenRespon("token 无效,返回登录界面！"))
				c.Abort()
				return
			}
			userID := convert.ToUint(string(val))
			c.Set(common.USER_UUID_Key, uuid)
			c.Set(common.USER_ID_Key, userID)
		}
		if uuid == "" {

			c.JSON(200, utils.FailedTokenRespon("用户未登录,返回登录界面！"))
			c.Abort()
			return
		}
		c.Next()
	}
}
