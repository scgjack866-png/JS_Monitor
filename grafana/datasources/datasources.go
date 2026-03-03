package datasources

import (
	"OperationAndMonitoring/initialize"
	"github.com/parnurzeal/gorequest"
)

var (
	request = *gorequest.New()
)

func AddDatasource() (string, bool) {
	res, body, _ := request.
		Post(initialize.Grafana.ApiUrl+"/api/datasources").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(`{
 "name": "my_prometheus",
 "type": "prometheus",
 "url": "` + initialize.Prometheus.ApiUrl + `",
 "access": "proxy",
 "basicAuth": false
}`).
		End()

	if res.StatusCode == 200 {
		return body, true
	}

	if res.StatusCode == 409 {
		res1, body1, _ := request.
			Get(initialize.Grafana.ApiUrl+"/api/datasources/name/my_prometheus").
			Set("Authorization", initialize.Grafana.Authorization).
			Set("Content-Type", "application/json").
			Set("Accept", "application/json").
			End()
		if res1.StatusCode == 200 {
			return body1, true
		}
		return body1, false
	}
	return body, false
}
