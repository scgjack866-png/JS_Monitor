package initialize

import (
	"OperationAndMonitoring/config"
)

var Grafana *config.Grafana

func InitGrafana(config *config.Config) {
	Grafana = &config.Grafana

}
