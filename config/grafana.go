package config

// grafana配置
type Grafana struct {
	ApiUrl                   string
	Authorization            string
	PrometheusUid            string
	HostRuleFolderUid        string
	DomainRuleFolderUid      string
	IpsecDomainRuleFolderUid string
	ConfigPath               string
	SnapshotDomain           string
}
