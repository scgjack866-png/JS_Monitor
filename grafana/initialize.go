package grafana

import (
	"OperationAndMonitoring/grafana/dashboards"
	"OperationAndMonitoring/grafana/datasources"
	"OperationAndMonitoring/grafana/folders"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 引入MySQL驱动
	"strings"
	"time"
)

// 添加 Prometheus 数据源
func CreatePrometheus(PrometheusDatasourceUid string) string {

	// 服务器告警策略
	body, ok := datasources.AddDatasource()
	if !ok {

		return ""
	}
	var setting entity.Setting
	setting.Name = PrometheusDatasourceUid
	setting.Value = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
	setting.CreateTime = time.Now()
	setting.UpdateTime = time.Now()
	utils.Create(&setting)
	return setting.Value
}

// 创建初始化时所需要的文件夹
func CreateGrafanaFolder(ServerFolderUid, title string) string {
	// 服务器告警策略
	body, ok := folders.CreateFolder(title)
	if !ok {
		return ""
	}
	fmt.Println(body)

	var setting entity.Setting
	setting.Name = ServerFolderUid
	setting.Value = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
	setting.CreateTime = time.Now()
	setting.UpdateTime = time.Now()
	utils.Create(&setting)
	return setting.Value
}

// 创建初始化时所需要的文件夹
func CreateIpsecDomainDashboards() {
	// 服务器告警策略
	body, ok := dashboards.CreateIpsecDomainDashboards("null")
	if !ok {
		return
	}
	fmt.Println(body)
}
