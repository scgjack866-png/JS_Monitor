package setting

import (
	"OperationAndMonitoring/grafana"
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/utils"
)

func InitSetting(HostFolderUid, DomainFolderUid, IpsecDomainFolderUid, PrometheusDatasourceUid string) {

	var hostFolderSetting entity.Setting
	var domainFolderSetting entity.Setting
	var ipsecDomainFolderSetting entity.Setting
	var prometheusDatasourceSetting entity.Setting

	notFound, _ := utils.First(&entity.Setting{Name: HostFolderUid}, &hostFolderSetting)
	if notFound {
		initialize.Grafana.HostRuleFolderUid = grafana.CreateGrafanaFolder(HostFolderUid, "服务器告警策略")
	} else {
		initialize.Grafana.HostRuleFolderUid = hostFolderSetting.Value
	}

	notFound, _ = utils.First(&entity.Setting{Name: DomainFolderUid}, &domainFolderSetting)
	if notFound {
		initialize.Grafana.DomainRuleFolderUid = grafana.CreateGrafanaFolder(DomainFolderUid, "域名告警策略")
	} else {
		initialize.Grafana.DomainRuleFolderUid = domainFolderSetting.Value
	}

	notFound, _ = utils.First(&entity.Setting{Name: IpsecDomainFolderUid}, &ipsecDomainFolderSetting)
	if notFound {
		initialize.Grafana.IpsecDomainRuleFolderUid = grafana.CreateGrafanaFolder(IpsecDomainFolderUid, "IPsec域名告警策略")
		grafana.CreateIpsecDomainDashboards()
	} else {
		initialize.Grafana.IpsecDomainRuleFolderUid = ipsecDomainFolderSetting.Value
	}

	notFound, _ = utils.First(&entity.Setting{Name: PrometheusDatasourceUid}, &prometheusDatasourceSetting)
	if notFound {
		initialize.Grafana.PrometheusUid = grafana.CreatePrometheus(PrometheusDatasourceUid)
	} else {
		initialize.Grafana.PrometheusUid = prometheusDatasourceSetting.Value
	}

}

func InitAlterSetting() {
	// 创建模板 -> 创建联络点 -> 创建策略

}
