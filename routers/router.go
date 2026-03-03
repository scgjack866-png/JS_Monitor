package routers

import (
	//"html/template"
	"net/http"

	"OperationAndMonitoring/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	//首页
	app.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	apiPrefix := "/api/v1"
	g := app.Group(apiPrefix)

	// 登录验证 jwt token 验证 及信息提取
	var notCheckLoginUrlArr []string
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/auth/login")
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/auth/logout")
	notCheckLoginUrlArr = append(notCheckLoginUrlArr, apiPrefix+"/auth/register")

	g.Use(middleware.UserAuthMiddleware(
		middleware.AllowPathPrefixSkipper(notCheckLoginUrlArr...),
	))

	AuthRouter(g)
	MenuRouter(g)
	DeptRouter(g)
	RoleRouter(g)
	GroupRouter(g)
	UserRouter(g)
	ServerRouter(g)
	DomainRouter(g)
	NodeRouter(g)
	HostRouter(g)
	FlowRouter(g)
	ProjectRouter(g)
	RoomRouter(g)
	ZoneRouter(g)
	OrderRouter(g)
	IpsecRouter(g)
	DomainStatusRouter(g)
}
