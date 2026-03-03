package alterrules

import (
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/utils/convert"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

var (
	request = *gorequest.New()
)

func CreateAlterRules(host entity.Host) (string, bool) {
	res, body, _ := request.Post(initialize.Grafana.ApiUrl+"/api/v1/provisioning/alert-rules").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getJson(host.IpAddr, host.Name, host.NetworkName, host.FlowLimit, host.LoadLimit, host.DiskLimit, host.CpuLimit, host.MemLimit)).
		End()
	if res.StatusCode == 201 {
		return body, true
	}
	return body, false
}

func CreateDomainAlterRules(nodeIP, domain, code string) (string, bool) {
	res, body, _ := request.Post(initialize.Grafana.ApiUrl+"/api/v1/provisioning/alert-rules").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getDomainJson(nodeIP, domain, code)).
		End()
	if res.StatusCode == 201 {
		return body, true
	}
	return body, false
}

func CreateIpsecDomainAlterRules(agentHostname, domain string, alterNum int) (string, bool) {
	res, body, _ := request.Post(initialize.Grafana.ApiUrl+"/api/v1/provisioning/alert-rules").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getIpsecDomainJson(agentHostname, domain, alterNum)).
		End()
	if res.StatusCode == 201 {
		return body, true
	}
	return body, false
}

func UpdateAlterRules(host entity.Host) bool {
	res, _, _ := request.Put(initialize.Grafana.ApiUrl+"/api/v1/provisioning/alert-rules/"+*host.RuleUID).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getJson(host.IpAddr, host.Name, host.NetworkName, host.FlowLimit, host.LoadLimit, host.DiskLimit, host.CpuLimit, host.MemLimit)).
		End()
	return res.StatusCode == 200
}

func UpdateDomainAlterRules(nodeIP, domain, code, ruleUID string) bool {
	res, _, _ := request.Put(initialize.Grafana.ApiUrl+"/api/v1/provisioning/alert-rules/"+ruleUID).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getDomainJson(nodeIP, domain, code)).
		End()
	return res.StatusCode == 200
}

func UpdateIpsecAlterRules(agentHostname, domain, ruleUID string, alterNum int) bool {
	res, _, _ := request.Put(initialize.Grafana.ApiUrl+"/api/v1/provisioning/alert-rules/"+ruleUID).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getIpsecDomainJson(agentHostname, domain, alterNum)).
		End()
	return res.StatusCode == 200
}

func DeleteAlterRules(uid string) bool {
	res, _, _ := request.Delete(initialize.Grafana.ApiUrl+"/api/v1/provisioning/alert-rules/"+uid).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		End()
	return res.StatusCode == 204
}

func getJson(IP, name, networkName string, flowLimit, loadLimit, diskLimit, cpuLimit, memLimit *int) string {
	return `{
	"id": null,
	"uid": null,
	"orgID": 1,
	"folderUID": "` + initialize.Grafana.HostRuleFolderUid + `",
	"title": "` + name + `-` + IP + `",
	"condition": "G",
	"data": [
		{
			"refId": "A",
			"queryType": "",
			"relativeTimeRange": {
				"from": 60,
				"to": 0
			},
			"datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
			"model": {
				"editorMode": "builder",
				"expr": "idelta(net_bits_recv{agent_hostname=\"` + IP + `\", interface=\"` + networkName + `\"}[1m]) / 15",
				"hide": false,
				"intervalMs": 1000,
				"legendFormat": "__auto",
				"maxDataPoints": 43200,
				"range": true,
				"refId": "A"
			}
		},
		{
			"refId": "B",
			"queryType": "",
			"relativeTimeRange": {
				"from": 60,
				"to": 0
			},
			"datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
			"model": {
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"editorMode": "builder",
				"expr": "idelta(net_bits_sent{agent_hostname=\"` + IP + `\", interface=\"` + networkName + `\"}[1m]) / 15",
				"hide": false,
				"intervalMs": 1000,
				"legendFormat": "__auto",
				"maxDataPoints": 43200,
				"range": true,
				"refId": "B"
			}
		},
		{
			"refId": "C",
			"queryType": "",
			"relativeTimeRange": {
				"from": 60,
				"to": 0
			},
			"datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
			"model": {
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"editorMode": "builder",
				"expr": "cpu_usage_user{agent_hostname=\"` + IP + `\"}",
				"hide": false,
				"intervalMs": 1000,
				"legendFormat": "__auto",
				"maxDataPoints": 43200,
				"range": true,
				"refId": "C"
			}
		},
		{
			"refId": "D",
			"queryType": "",
			"relativeTimeRange": {
				"from": 60,
				"to": 0
			},
			"datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
			"model": {
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"editorMode": "builder",
				"expr": "mem_used_percent{agent_hostname=\"` + IP + `\"}",
				"hide": false,
				"intervalMs": 1000,
				"legendFormat": "__auto",
				"maxDataPoints": 43200,
				"range": true,
				"refId": "D"
			}
		},
		{
			"refId": "E",
			"queryType": "",
			"relativeTimeRange": {
				"from": 60,
				"to": 0
			},
			"datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
			"model": {
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"editorMode": "builder",
				"expr": "disk_used_percent{agent_hostname=\"` + IP + `\"}",
				"hide": false,
				"intervalMs": 1000,
				"legendFormat": "__auto",
				"maxDataPoints": 43200,
				"range": true,
				"refId": "E"
			}
		},
		{
			"refId": "F",
			"queryType": "",
			"relativeTimeRange": {
				"from": 60,
				"to": 0
			},
			"datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
			"model": {
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"editorMode": "builder",
				"expr": "system_load1{agent_hostname=\"` + IP + `\"}",
				"hide": false,
				"intervalMs": 1000,
				"legendFormat": "__auto",
				"maxDataPoints": 43200,
				"range": true,
				"refId": "F"
			}
		},
		{
			"refId": "G",
			"queryType": "",
			"relativeTimeRange": {
				"from": 60,
				"to": 0
			},
			"datasourceUid": "-100",
			"model": {
				"conditions": [
					{
						"evaluator": {
							"params": [
								` + convert.ToString(*flowLimit*1024*1024) + `,
								0
							],
							"type": "gt"
						},
						"operator": {
							"type": "and"
						},
						"query": {
							"params": [
								"A"
							]
						},
						"reducer": {
							"params": [],
							"type": "last"
						},
						"type": "query"
					},
					{
						"evaluator": {
							"params": [
								` + convert.ToString(*flowLimit*1024*1024) + `,
								0
							],
							"type": "gt"
						},
						"operator": {
							"type": "or"
						},
						"query": {
							"params": [
								"B"
							]
						},
						"reducer": {
							"params": [],
							"type": "last"
						},
						"type": "query"
					},
					{
						"evaluator": {
							"params": [
								` + convert.ToString(cpuLimit) + `,
								0
							],
							"type": "gt"
						},
						"operator": {
							"type": "or"
						},
						"query": {
							"params": [
								"C"
							]
						},
						"reducer": {
							"params": [],
							"type": "last"
						},
						"type": "query"
					},
					{
						"evaluator": {
							"params": [
								` + convert.ToString(memLimit) + `,
								0
							],
							"type": "gt"
						},
						"operator": {
							"type": "or"
						},
						"query": {
							"params": [
								"D"
							]
						},
						"reducer": {
							"params": [],
							"type": "last"
						},
						"type": "query"
					},
					{
						"evaluator": {
							"params": [
								` + convert.ToString(diskLimit) + `,
								0
							],
							"type": "gt"
						},
						"operator": {
							"type": "or"
						},
						"query": {
							"params": [
								"E"
							]
						},
						"reducer": {
							"params": [],
							"type": "last"
						},
						"type": "query"
					},
					{
						"evaluator": {
							"params": [
								` + convert.ToString(loadLimit) + `,
								0
							],
							"type": "gt"
						},
						"operator": {
							"type": "or"
						},
						"query": {
							"params": [
								"F"
							]
						},
						"reducer": {
							"params": [],
							"type": "last"
						},
						"type": "query"
					}
				],
				"datasource": {
					"name": "Expression",
					"type": "__expr__",
					"uid": "__expr__"
				},
				"expression": "H",
				"hide": false,
				"intervalMs": 1000,
				"maxDataPoints": 43200,
				"refId": "F",
				"type": "classic_conditions"
			}
		}
	],
	"noDataState": "Alerting",
	"execErrState": "Alerting",
	"for": "1m",
	"labels": {
		"IP": "` + IP + `"
	}
}`
}

func getDomainJson(nodeIP, domain, code string) string {
	return `{
    "id": null,
    "uid": null,
    "orgID": 1,
    "folderUID": "` + initialize.Grafana.DomainRuleFolderUid + `",
    "ruleGroup": "domain",
    "title": "` + nodeIP + `-` + domain + `-域名监控",
    "condition": "B",
    "data": [
        {
            "refId": "A",
            "queryType": "",
            "relativeTimeRange": {
                "from": 60,
                "to": 0
            },
            "datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
            "model": {
                "datasource": {
                    "type": "prometheus",
                    "uid": "` + initialize.Grafana.PrometheusUid + `"
                },
                "editorMode": "builder",
                "expr": "http_response_response_code{agent_hostname=\"` + nodeIP + `\", target=\"` + domain + `\"}",
                "instant": false,
                "interval": "",
                "intervalMs": 15000,
                "maxDataPoints": 43200,
                "range": true,
                "refId": "A"
            }
        },
        {
            "refId": "B",
            "queryType": "",
            "relativeTimeRange": {
                "from": 10800,
                "to": 0
            },
            "datasourceUid": "__expr__",
            "model": {
                "conditions": [
                    {
                        "evaluator": {
                            "params": [
                                ` + code + `,
                                ` + code + `
                            ],
                            "type": "outside_range"
                        },
                        "operator": {
                            "type": "and"
                        },
                        "query": {
                            "params": [
                                "A"
                            ]
                        },
                        "reducer": {
                            "params": [],
                            "type": "last"
                        },
                        "type": "query"
                    }
                ],
                "datasource": {
                    "name": "Expression",
                    "type": "__expr__",
                    "uid": "__expr__"
                },
                "expression": "",
                "intervalMs": 1000,
                "maxDataPoints": 43200,
                "refId": "B",
                "type": "classic_conditions"
            }
        }
    ],
    "noDataState": "Alerting",
    "execErrState": "Alerting",
    "for": "30s",
    "labels": {
        "IP": "` + nodeIP + `",
        "domain": "` + domain + `"
    },
    "isPaused": false
}`
}

func getIpsecDomainJson(agentHostname, domain string, alterNum int) string {
	return `{
    "id": null,
    "uid": null,
    "orgID": 1,
    "folderUID": "` + initialize.Grafana.IpsecDomainRuleFolderUid + `",
    "ruleGroup": "IPsec域名监控",
    "title": "` + agentHostname + `-` + domain + `-IPsec域名监控",
    "condition": "B",
    "data": [
        {
            "refId": "A",
            "queryType": "",
            "relativeTimeRange": {
                "from": 60,
                "to": 0
            },
            "datasourceUid": "` + initialize.Grafana.PrometheusUid + `",
            "model": {
                "editorMode": "builder",
                "expr": "ipsec_online_num{agent_hostname=\"` + agentHostname + `\", domain=\"` + domain + `\"}",
                "instant": false,
                "interval": "",
                "intervalMs": 15000,
                "maxDataPoints": 43200,
                "range": true,
                "refId": "A"
            }
        },
        {
            "refId": "B",
            "queryType": "",
            "relativeTimeRange": {
                "from": 10800,
                "to": 0
            },
            "datasourceUid": "__expr__",
            "model": {
                "conditions": [
                    {
                        "evaluator": {
                            "params": [
                                ` + strconv.Itoa(alterNum) + `,
                                0
                            ],
                            "type": "lt"
                        },
                        "operator": {
                            "type": "and"
                        },
                        "query": {
                            "params": [
                                "A"
                            ]
                        },
                        "reducer": {
                            "params": [],
                            "type": "last"
                        },
                        "type": "query"
                    }
                ],
                "datasource": {
                    "name": "Expression",
                    "type": "__expr__",
                    "uid": "__expr__"
                },
                "expression": "",
                "intervalMs": 1000,
                "maxDataPoints": 43200,
                "refId": "B",
                "type": "classic_conditions"
            }
        }
    ],
    "noDataState": "Alerting",
    "execErrState": "Alerting",
    "for": "30s",
    "labels": {
        "IP": "` + agentHostname + `",
        "domain": "` + domain + `"
    },
    "isPaused": false
}`
}
