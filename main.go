package main

import (
	"OperationAndMonitoring/config"
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/log/zap"
	"OperationAndMonitoring/middleware"
	"OperationAndMonitoring/routers"
	"OperationAndMonitoring/setting"
	"OperationAndMonitoring/utils/convert"
	"fmt"
	"github.com/gin-gonic/gin"
	log "go.uber.org/zap"
	"net/http"
	"time"
)

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

var (
	HostFolderUid           = "HostFolderUid"
	DomainFolderUid         = "DomainFolderUid"
	IpsecDomainFolderUid    = "IpsecDomainFolderUid"
	PrometheusDatasourceUid = "PrometheusDatasourceUid"
)

func main() {

	configPath := "./config/config.yaml"

	// 加载配置
	config, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	// 加载日志
	if err = zap.InitLogger(config); err != nil {
		panic(err)
	}

	initialize.InitDB(config)
	//initialize.InitCsbinEnforcer()

	initialize.InitGrafana(config)
	initialize.InitPrometheus(config)

	setting.InitSetting(HostFolderUid, DomainFolderUid, IpsecDomainFolderUid, PrometheusDatasourceUid)

	initWeb(config)
	log.L().Debug(config.Web.Domain + "站点已启动...")

}

func initWeb(config *config.Config) {
	if config.Zap.Debug || config.Gorm.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	// 注册zap相关中间件
	app.Use(zap.GinLogger(), zap.GinRecovery(true))
	//app.LoadHTMLGlob(config.Web.StaticPath + "dist/*.html")
	app.Static("/static", config.Web.StaticPath+"dist/static")
	app.Static("/resource", config.Web.StaticPath+"resource")
	app.StaticFile("/favicon.ico", config.Web.StaticPath+"dist/favicon.ico")
	// 注册路由
	app.Use(middleware.Cors())
	routers.RegisterRouter(app)
	initHTTPServer(config, app)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config, handler http.Handler) {
	srv := &http.Server{
		Addr:         ":" + convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
