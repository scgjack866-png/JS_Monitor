package alterrules

import (
	"OperationAndMonitoring/initialize"
)

// 创建通知模板
func CreateAlterTemplates(name, jsonBody string) (string, bool) {
	res, body, _ := request.Put(initialize.Grafana.ApiUrl+"/api/v1/provisioning/templates/"+name).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(jsonBody).
		End()
	if res.StatusCode == 201 {
		return body, true
	}
	return body, false
}

// 创建TG联络点
func CreateAlterContactPoints(jsonBody string) (string, bool) {
	res, body, _ := request.Post(initialize.Grafana.ApiUrl+"/api/v1/provisioning/contact-points").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(jsonBody).
		End()
	if res.StatusCode == 201 {
		return body, true
	}
	return body, false
}

// 创建通知策略
func CreateAlterPolicy(jsonBody string) (string, bool) {
	res, body, _ := request.Post(initialize.Grafana.ApiUrl+"/api/v1/provisioning/contact-points").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(jsonBody).
		End()
	if res.StatusCode == 201 {
		return body, true
	}
	return body, false
}
