package middleware

import (
	"OperationAndMonitoring/controller/common"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/cache"
	"OperationAndMonitoring/utils/convert"
	"crypto/subtle"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
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
			if allowByAPIKey(c, t) {
				c.Set(common.USER_UUID_Key, "internal-api")
				c.Set(common.USER_ID_Key, uint64(0))
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

// allowByAPIKey 使用环境变量 INTERNAL_API_KEY 控制无登录调用。
// 请求头支持：
// 1) Authorization: ApiKey <key>
// 2) X-API-Key: <key>
func allowByAPIKey(c *gin.Context, authorizationHeader string) bool {
	expected := strings.TrimSpace(os.Getenv("INTERNAL_API_KEY"))
	if expected == "" {
		return false
	}

	provided := strings.TrimSpace(c.GetHeader("X-API-Key"))
	if provided == "" {
		fields := strings.Fields(strings.TrimSpace(authorizationHeader))
		if len(fields) == 2 && strings.EqualFold(fields[0], "ApiKey") {
			provided = fields[1]
		}
	}

	if provided == "" {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(provided), []byte(expected)) == 1
}
