package dashboards

import (
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

var (
	request = *gorequest.New()
)

func CreateDashboards(uid string, IP string, tag string, folderUid string, networkName string) (string, bool, []error) {
	res, body, errs := request.
		Post(initialize.Grafana.ApiUrl+"/api/dashboards/db").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getSendJson(uid, IP, tag, folderUid, networkName)).
		End()
	if errs != nil || res.StatusCode == 200 {
		return body, true, errs
	}

	return body, false, errs
}

func CreateDomainDashboards(uid string) (string, bool) {
	res, body, _ := request.
		Post(initialize.Grafana.ApiUrl+"/api/dashboards/db").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getDomainSendJson(uid)).
		End()

	if res.StatusCode == 200 {
		return body, true
	}
	return body, false
}

func CreateIpsecDomainDashboards(uid string) (string, bool) {
	res, body, _ := request.
		Post(initialize.Grafana.ApiUrl+"/api/dashboards/db").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		Send(getIpsecDomainSendJson(uid)).
		End()
	fmt.Println(body, res)
	if res.StatusCode == 200 {
		return body, true
	}
	return body, false
}

func DeleteDashboards(uid string) bool {

	res, _, _ := request.
		Delete(initialize.Grafana.ApiUrl+"/api/dashboards/uid/"+uid).
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		End()
	return res.StatusCode == 200
}

func getSendJson(uid string, IP string, tag string, folderUid string, networkName string) interface{} {
	return `{
	"dashboard": {
		"id": null,
		"uid": ` + uid + `,
		"title": "` + tag + `-` + IP + `",
		"tags": ["` + tag + `"],

		"timezone": "browser",
		"version": 0,
		"refresh": "15s",

		"editable": true,
		"fiscalYearStartMonth": 0,
		"graphTooltip": 0,

		"links": [],
		"liveNow": false,

		"panels": [
			{
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"fieldConfig": {
					"defaults": {
						"color": {
							"mode": "palette-classic"
						},

						"custom": {
							"axisCenteredZero": false,
							"axisColorMode": "text",
							"axisLabel": "",
							"axisPlacement": "auto",
							"barAlignment": 0,
							"drawStyle": "line",
							"fillOpacity": 0,
							"gradientMode": "none",
							"hideFrom": {
								"legend": false,
								"tooltip": false,
								"viz": false
							},
							"lineInterpolation": "smooth",
							"lineStyle": {
								"fill": "solid"
							},
							"lineWidth": 1,
							"pointSize": 5,
							"scaleDistribution": {
								"type": "linear"
							},
							"showPoints": "auto",
							"spanNulls": true,
							"stacking": {
								"group": "A",
								"mode": "none"
							},
							"thresholdsStyle": {
								"mode": "off"
							}
						},

						"mappings": [],
						"thresholds": {
							"mode": "absolute",
							"steps": [{
									"color": "green",
									"value": null
								},
								{
									"color": "red",
									"value": 80
								}
							]
						},
						"unit": "binbps"
					},
					"overrides": []
				},
				"gridPos": {
					"h": 9,
					"w": 12,
					"x": 0,
					"y": 0
				},
				"id": 2,
				"options": {
					"legend": {
						"calcs": [
							"last",
							"max",
							"min"
						],
						"displayMode": "table",
						"placement": "bottom",
						"showLegend": true
					},
					"tooltip": {
						"mode": "multi",
						"sort": "none"
					}
				},
				"pluginVersion": "9.2.4",
				"targets": [{
						"editorMode": "builder",
						"exemplar": false,
						"expr": "idelta(net_bits_recv{interface=\"` + networkName + `\", agent_hostname=\"` + IP + `\"}[1m]) / 15",
						"format": "heatmap",
						"hide": false,
						"instant": false,
						"legendFormat": "流入-{{interface}}",
						"range": true,
						"refId": "流入流量"
					},
					{
						"editorMode": "builder",
						"exemplar": false,
						"expr": "idelta(net_bits_sent{interface=\"` + networkName + `\", agent_hostname=\"` + IP + `\"}[1m]) / 15",
						"format": "heatmap",
						"hide": false,
						"interval": "",
						"legendFormat": "流出-{{interface}}",
						"range": true,
						"refId": "流出流量"
					}
				],
				"title": "` + IP + ` 进出流量",
				"type": "timeseries"
			},

			{
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"fieldConfig": {
					"defaults": {
						"color": {
							"mode": "palette-classic"
						},
						"custom": {
							"axisCenteredZero": false,
							"axisColorMode": "text",
							"axisLabel": "",
							"axisPlacement": "auto",
							"barAlignment": 0,
							"drawStyle": "line",
							"fillOpacity": 0,
							"gradientMode": "none",
							"hideFrom": {
								"legend": false,
								"tooltip": false,
								"viz": false
							},
							"lineInterpolation": "linear",
							"lineWidth": 1,
							"pointSize": 5,
							"scaleDistribution": {
								"type": "linear"
							},
							"showPoints": "auto",
							"spanNulls": false,
							"stacking": {
								"group": "A",
								"mode": "none"
							},
							"thresholdsStyle": {
								"mode": "off"
							}
						},
						"mappings": [],
						"thresholds": {
							"mode": "absolute",
							"steps": [{
									"color": "green",
									"value": null
								},
								{
									"color": "red",
									"value": 80
								}
							]
						}
					},
					"overrides": []
				},
				"gridPos": {
					"h": 8,
					"w": 12,
					"x": 0,
					"y": 9
				},
				"id": 9,
				"options": {
					"legend": {
						"calcs": [],
						"displayMode": "list",
						"placement": "bottom",
						"showLegend": true
					},
					"tooltip": {
						"mode": "single",
						"sort": "none"
					}
				},
				"targets": [{
						"editorMode": "builder",
						"expr": "system_load1{agent_hostname=\"` + IP + `\"}",
						"legendFormat": "__auto",
						"range": true,
						"refId": "A"
					},
					{
						"editorMode": "builder",
						"expr": "system_load5{agent_hostname=\"` + IP + `\"}",
						"hide": false,
						"legendFormat": "__auto",
						"range": true,
						"refId": "B"
					},
					{
						"editorMode": "builder",
						"expr": "system_load15{agent_hostname=\"` + IP + `\"}",
						"hide": false,
						"legendFormat": "__auto",
						"range": true,
						"refId": "C"
					}
				],
				"title": "` + IP + ` 系统负载",
				"type": "timeseries"
			},

			{
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"fieldConfig": {
					"defaults": {
						"color": {
							"mode": "palette-classic"
						},
						"custom": {
							"axisCenteredZero": false,
							"axisColorMode": "text",
							"axisLabel": "",
							"axisPlacement": "auto",
							"barAlignment": 0,
							"drawStyle": "line",
							"fillOpacity": 0,
							"gradientMode": "none",
							"hideFrom": {
								"legend": false,
								"tooltip": false,
								"viz": false
							},
							"lineInterpolation": "linear",
							"lineWidth": 1,
							"pointSize": 5,
							"scaleDistribution": {
								"type": "linear"
							},
							"showPoints": "auto",
							"spanNulls": false,
							"stacking": {
								"group": "A",
								"mode": "none"
							},
							"thresholdsStyle": {
								"mode": "off"
							}
						},
						"mappings": [],
						"thresholds": {
							"mode": "absolute",
							"steps": [{
									"color": "green",
									"value": null
								},
								{
									"color": "red",
									"value": 80
								}
							]
						},
						"unit": "percent"
					},
					"overrides": []
				},
				"gridPos": {
					"h": 8,
					"w": 12,
					"x": 0,
					"y": 17
				},
				"id": 5,
				"options": {
					"legend": {
						"calcs": [],
						"displayMode": "list",
						"placement": "bottom",
						"showLegend": true
					},
					"tooltip": {
						"mode": "single",
						"sort": "none"
					}
				},
				"targets": [{
					"editorMode": "builder",
					"expr": "cpu_usage_user{agent_hostname=\"` + IP + `\"}",
					"legendFormat": "__auto",
					"range": true,
					"refId": "cpu空闲"
				}],
				"title": "` + IP + ` CPU使用百分比",
				"type": "timeseries"
			},

			{
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"fieldConfig": {
					"defaults": {
						"color": {
							"mode": "palette-classic"
						},
						"custom": {
							"axisCenteredZero": false,
							"axisColorMode": "text",
							"axisLabel": "",
							"axisPlacement": "auto",
							"barAlignment": 0,
							"drawStyle": "line",
							"fillOpacity": 0,
							"gradientMode": "none",
							"hideFrom": {
								"legend": false,
								"tooltip": false,
								"viz": false
							},
							"lineInterpolation": "linear",
							"lineWidth": 1,
							"pointSize": 5,
							"scaleDistribution": {
								"type": "linear"
							},
							"showPoints": "auto",
							"spanNulls": false,
							"stacking": {
								"group": "A",
								"mode": "none"
							},
							"thresholdsStyle": {
								"mode": "off"
							}
						},
						"mappings": [],
						"thresholds": {
							"mode": "absolute",
							"steps": [{
									"color": "green",
									"value": null
								},
								{
									"color": "red",
									"value": 80
								}
							]
						},
						"unit": "percent"
					},
					"overrides": []
				},
				"gridPos": {
					"h": 8,
					"w": 12,
					"x": 0,
					"y": 25
				},
				"id": 7,
				"options": {
					"legend": {
						"calcs": [],
						"displayMode": "list",
						"placement": "bottom",
						"showLegend": true
					},
					"tooltip": {
						"mode": "single",
						"sort": "none"
					}
				},
				"targets": [{
					"editorMode": "builder",
					"expr": "mem_used_percent{agent_hostname=\"` + IP + `\"}",
					"legendFormat": "__auto",
					"range": true,
					"refId": "内存使用百分比"
				}],
				"title": "` + IP + ` 内存使用百分比",
				"type": "timeseries"
			},

			{
				"datasource": {
					"type": "prometheus",
					"uid": "` + initialize.Grafana.PrometheusUid + `"
				},
				"fieldConfig": {
					"defaults": {
						"color": {
							"mode": "palette-classic"
						},
						"custom": {
							"axisCenteredZero": false,
							"axisColorMode": "text",
							"axisLabel": "",
							"axisPlacement": "auto",
							"barAlignment": 0,
							"drawStyle": "line",
							"fillOpacity": 0,
							"gradientMode": "none",
							"hideFrom": {
								"legend": false,
								"tooltip": false,
								"viz": false
							},
							"lineInterpolation": "linear",
							"lineWidth": 1,
							"pointSize": 5,
							"scaleDistribution": {
								"type": "linear"
							},
							"showPoints": "auto",
							"spanNulls": false,
							"stacking": {
								"group": "A",
								"mode": "none"
							},
							"thresholdsStyle": {
								"mode": "off"
							}
						},
						"mappings": [],
						"thresholds": {
							"mode": "absolute",
							"steps": [{
									"color": "green",
									"value": null
								},
								{
									"color": "red",
									"value": 80
								}
							]
						},
						"unit": "percent"
					},
					"overrides": []
				},
				"gridPos": {
					"h": 8,
					"w": 12,
					"x": 0,
					"y": 33
				},
				"id": 4,
				"options": {
					"legend": {
						"calcs": [],
						"displayMode": "list",
						"placement": "bottom",
						"showLegend": true
					},
					"tooltip": {
						"mode": "single",
						"sort": "none"
					}
				},
				"targets": [{
					"editorMode": "builder",
					"expr": "disk_used_percent{agent_hostname=\"` + IP + `\"}",
					"legendFormat": "__auto",
					"range": true,
					"refId": "A"
				}],
				"title": "` + IP + ` 硬盘使用百分比",
				"type": "timeseries"
			}
		],

		"style": "dark",
		"templating": {
			"list": []
		},
		"time": {
			"from": "now-3h",
			"to": "now"
		},
		"timepicker": {},
		"weekStart": ""

	},
	"folderUid": "` + folderUid + `",
	"message": "Made changes to xyz",
	"overwrite": true
}`
}

func getDomainSendJson(uid string) interface{} {
	return `{
	"dashboard": {
		"id": null,
		"uid": ` + uid + `,
		"title": "域名监控",
		"tags": ["域名监控"],

		"timezone": "browser",
		"version": 0,
		"refresh": "15s",

		"editable": true,
		"fiscalYearStartMonth": 0,
		"graphTooltip": 0,

		"links": [],
		"liveNow": false,

		"panels": [
			` + getDomainDashboards() + `
		],

		"style": "dark",
		"templating": {
			"list": []
		},
		"time": {
			"from": "now-3h",
			"to": "now"
		},
		"timepicker": {},
		"weekStart": ""

	},
	"folderUid": "` + initialize.Grafana.DomainRuleFolderUid + `",
	"message": "Made changes to xyz",
	"overwrite": true
}`
}

func getIpsecDomainSendJson(uid string) interface{} {
	return `{
	"dashboard": {
		"id": null,
		"uid": ` + uid + `,
		"title": "IPsec域名监控",
		"tags": ["IPsec域名监控"],

		"timezone": "browser",
		"version": 0,
		"refresh": "15s",

		"editable": true,
		"fiscalYearStartMonth": 0,
		"graphTooltip": 0,

		"links": [],
		"liveNow": false,

		"panels": [
			{
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 1,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "` + initialize.Grafana.PrometheusUid + `"
          },
          "editorMode": "builder",
          "expr": "ipsec_online_num",
          "instant": false,
          "range": true,
          "refId": "A"
        }
      ],
      "title": "Ipsec域名监控",
      "type": "timeseries"
    }
		],

		"style": "dark",
		"templating": {
			"list": []
		},
		"time": {
			"from": "now-3h",
			"to": "now"
		},
		"timepicker": {},
		"weekStart": ""

	},
	"folderUid": "` + initialize.Grafana.IpsecDomainRuleFolderUid + `",
	"message": "Made changes to xyz",
	"overwrite": true
}`
}

func getDomainDashboards() string {

	var nodes []entity.Node
	var whereOrder []mysql.PageWhereOrder
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})
	whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = 1"})
	err := utils.Find(&entity.Node{}, &nodes, whereOrder...)
	if err != nil {
		return ``
	}

	var dashboards []byte
	for i, node := range nodes {
		dashboards = append(dashboards, []byte(`{
      "datasource": {
        "type": "prometheus",
        "uid": "`+initialize.Grafana.PrometheusUid+`"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 2,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "`+initialize.Grafana.PrometheusUid+`"
          },
          "editorMode": "builder",
          "expr": "http_response_response_code{agent_hostname=\"`+node.NodeIP+`\"}",
          "instant": false,
          "range": true,
          "refId": "A"
        }
      ],
      "title": "`+node.NodeIP+"-"+node.NodeName+`-域名监控",
      "type": "timeseries"
    }`)...)
		if i != len(nodes)-1 {
			dashboards = append(dashboards, []byte(",")...)
		}
	}
	return string(dashboards)
}
