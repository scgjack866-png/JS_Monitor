package initialize

import "OperationAndMonitoring/config"

var Prometheus *config.Prometheus

func InitPrometheus(config *config.Config) {
	Prometheus = &config.Prometheus
}
