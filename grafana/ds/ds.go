package ds

import (
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/model/entity"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"sync"
)

var (
	request = *gorequest.New()
	mu      sync.Mutex
)

func QueryDsFlow(host entity.Host, now string, beforeNow string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	res, body, err := request.
		Post(initialize.Grafana.ApiUrl+"/api/ds/query").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		//Send(`{"queries":[{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_recv{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"instant":false,"legendFormat":"流入-{{interface}}","range":true,"refId":"In","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流入流量","utcOffsetSec":28800,"interval":"","datasourceId":1,"intervalMs":20000,"maxDataPoints":518},{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_sent{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"interval":"","legendFormat":"流出-{{interface}}","range":true,"refId":"Out","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流出流量","utcOffsetSec":28800,"datasourceId":1,"intervalMs":20000,"maxDataPoints":518}],"from":"` + now + `","to":"` + now + `"}`).
		Send(`{
  "queries": [
    {
      "editorMode": "builder",
      "exemplar": false,
      "expr": "idelta(net_bits_recv{interface=\"` + host.NetworkName + `\", agent_hostname=\"` + host.IpAddr + `\"}[1m]) / 15",
      "format": "heatmap",
      "hide": false,
      "instant": false,
      "legendFormat": "流入-{{interface}}",
      "range": true,
      "refId": "In",
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "requestId": "2流入流量",
      "utcOffsetSec": 28800,
      "interval": "",
      "datasourceId": 1,
      "intervalMs": 20000,
      "maxDataPoints": 518
    },
    {
      "editorMode": "builder",
      "exemplar": false,
      "expr": "idelta(net_bits_sent{interface=\"` + host.NetworkName + `\", agent_hostname=\"` + host.IpAddr + `\"}[1m]) / 15",
      "format": "heatmap",
      "hide": false,
      "interval": "",
      "legendFormat": "流出-{{interface}}",
      "range": true,
      "refId": "Out",
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "requestId": "2流出流量",
      "utcOffsetSec": 28800,
      "datasourceId": 1,
      "intervalMs": 20000,
      "maxDataPoints": 518
    }
  ],
  "from": "` + beforeNow + `",
  "to": "` + now + `"
}`).
		End()
	if err != nil {
		fmt.Println("flow")
		fmt.Println(err)
		return err[0].Error(), false
	}
	if res.StatusCode == 200 {
		return body, true
	}

	return body, false
}

func QueryDsCpu(host entity.Host, now string, beforeNow string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	res, body, err := request.
		Post(initialize.Grafana.ApiUrl+"/api/ds/query").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		//Send(`{"queries":[{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_recv{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"instant":false,"legendFormat":"流入-{{interface}}","range":true,"refId":"In","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流入流量","utcOffsetSec":28800,"interval":"","datasourceId":1,"intervalMs":20000,"maxDataPoints":518},{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_sent{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"interval":"","legendFormat":"流出-{{interface}}","range":true,"refId":"Out","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流出流量","utcOffsetSec":28800,"datasourceId":1,"intervalMs":20000,"maxDataPoints":518}],"from":"` + now + `","to":"` + now + `"}`).
		Send(`{
  "queries": [
    {
      "editorMode": "builder",
      "expr": "cpu_usage_user{agent_hostname=\"` + host.IpAddr + `\"}",
      "legendFormat": "__auto",
      "range": true,
      "refId": "Cpu",
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "exemplar": false,
      "requestId": "5cpu空闲",
      "utcOffsetSec": 0,
      "interval": "",
      "datasourceId": 1,
      "intervalMs": 20000,
      "maxDataPoints": 466
    }
  ],
  "from": "` + beforeNow + `",
  "to": "` + now + `"
}`).
		End()
	if err != nil {
		fmt.Println("cpu")
		fmt.Println(err)
		return err[0].Error(), false
	}
	if res.StatusCode == 200 {
		return body, true
	}

	return body, false
}

func QueryDsFree(host entity.Host, now string, beforeNow string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	res, body, err := request.
		Post(initialize.Grafana.ApiUrl+"/api/ds/query").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		//Send(`{"queries":[{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_recv{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"instant":false,"legendFormat":"流入-{{interface}}","range":true,"refId":"In","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流入流量","utcOffsetSec":28800,"interval":"","datasourceId":1,"intervalMs":20000,"maxDataPoints":518},{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_sent{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"interval":"","legendFormat":"流出-{{interface}}","range":true,"refId":"Out","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流出流量","utcOffsetSec":28800,"datasourceId":1,"intervalMs":20000,"maxDataPoints":518}],"from":"` + now + `","to":"` + now + `"}`).
		Send(`{
  "queries": [
    {
      "editorMode": "builder",
      "expr": "mem_used_percent{agent_hostname=\"` + host.IpAddr + `\"}",
      "legendFormat": "__auto",
      "range": true,
      "refId": "Free",
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "exemplar": false,
      "requestId": "7内存使用百分比",
      "utcOffsetSec": 0,
      "interval": "",
      "datasourceId": 1,
      "intervalMs": 20000,
      "maxDataPoints": 466
    }
  ],
  "from": "` + beforeNow + `",
  "to": "` + now + `"
}`).
		End()

	if err != nil {
		fmt.Println("free")
		fmt.Println(err)
		return err[0].Error(), false
	}
	if res.StatusCode == 200 {
		return body, true
	}

	return body, false
}

func QueryDsLoad(host entity.Host, now string, beforeNow string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	res, body, err := request.
		Post(initialize.Grafana.ApiUrl+"/api/ds/query").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		//Send(`{"queries":[{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_recv{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"instant":false,"legendFormat":"流入-{{interface}}","range":true,"refId":"In","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流入流量","utcOffsetSec":28800,"interval":"","datasourceId":1,"intervalMs":20000,"maxDataPoints":518},{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_sent{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"interval":"","legendFormat":"流出-{{interface}}","range":true,"refId":"Out","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流出流量","utcOffsetSec":28800,"datasourceId":1,"intervalMs":20000,"maxDataPoints":518}],"from":"` + now + `","to":"` + now + `"}`).
		Send(`{
  "queries": [
    {
      "editorMode": "builder",
      "expr": "system_load1{agent_hostname=\"` + host.IpAddr + `\"}",
      "legendFormat": "__auto",
      "range": true,
      "refId": "LoadOne",
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "exemplar": false,
      "requestId": "9A",
      "utcOffsetSec": 0,
      "interval": "",
      "datasourceId": 1,
      "intervalMs": 20000,
      "maxDataPoints": 466
    },
    {
      "editorMode": "builder",
      "expr": "system_load5{agent_hostname=\"` + host.IpAddr + `\"}",
      "hide": false,
      "legendFormat": "__auto",
      "range": true,
      "refId": "LoadFive",
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "exemplar": false,
      "requestId": "9B",
      "utcOffsetSec": 0,
      "interval": "",
      "datasourceId": 1,
      "intervalMs": 20000,
      "maxDataPoints": 466
    },
    {
      "editorMode": "builder",
      "expr": "system_load15{agent_hostname=\"` + host.IpAddr + `\"}",
      "hide": false,
      "legendFormat": "__auto",
      "range": true,
      "refId": "LoadFifteen",
      "datasource": {
        "type": "prometheus",
        "uid": "` + initialize.Grafana.PrometheusUid + `"
      },
      "exemplar": false,
      "requestId": "9C",
      "utcOffsetSec": 0,
      "interval": "",
      "datasourceId": 1,
      "intervalMs": 20000,
      "maxDataPoints": 466
    }
  ],

  "from": "` + beforeNow + `",
  "to": "` + now + `"
}`).
		End()
	if err != nil {
		fmt.Println("load")
		fmt.Println(err)
		return err[0].Error(), false
	}
	if res.StatusCode == 200 {
		return body, true
	}

	return body, false
}

func GetSnapshotUrl(host entity.Host, nowFormatStr string, beforeNowFormatStr string, OutValue, InValue, CpuValue, FreeValue, Load1Value, Load5Value, Load15Value [][]float64) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	inTime, err := formatMapToString(InValue[0])
	inValue, err := formatMapToString(InValue[1])
	outTime, err := formatMapToString(OutValue[0])
	outValue, err := formatMapToString(OutValue[1])
	cpuTime, err := formatMapToString(CpuValue[0])
	cpuValue, err := formatMapToString(CpuValue[1])
	freeTime, err := formatMapToString(FreeValue[0])
	freeValue, err := formatMapToString(FreeValue[1])
	loadOneTime, err := formatMapToString(Load1Value[0])
	loadOneValue, err := formatMapToString(Load1Value[1])
	loadFiveTime, err := formatMapToString(Load5Value[0])
	loadFiveValue, err := formatMapToString(Load5Value[1])
	loadFifteenTime, err := formatMapToString(Load15Value[0])
	loadFifteenValue, err := formatMapToString(Load15Value[1])

	if err != nil {
		return "", false
	}

	res, body, _ := request.
		Post(initialize.Grafana.ApiUrl+"/api/snapshots").
		Set("Authorization", initialize.Grafana.Authorization).
		Set("Content-Type", "application/json").
		Set("Accept", "application/json").
		//Send(`{"queries":[{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_recv{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"instant":false,"legendFormat":"流入-{{interface}}","range":true,"refId":"In","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流入流量","utcOffsetSec":28800,"interval":"","datasourceId":1,"intervalMs":20000,"maxDataPoints":518},{"editorMode":"builder","exemplar":false,"expr":"idelta(net_bits_sent{interface="` + host.NetworkName + `", agent_hostname="` + host.IpAddr + `"}[1m]) / 15","format":"heatmap","hide":false,"interval":"","legendFormat":"流出-{{interface}}","range":true,"refId":"Out","datasource":{"type":"prometheus","uid":"` + initialize.Grafana.PrometheusUid + `"},"requestId":"2流出流量","utcOffsetSec":28800,"datasourceId":1,"intervalMs":20000,"maxDataPoints":518}],"from":"` + now + `","to":"` + now + `"}`).
		Send(`{
  "dashboard": {
    "annotations": {
      "list": [
        {
          "name": "Annotations & Alerts",
          "enable": true,
          "iconColor": "rgba(0, 211, 255, 1)",
          "snapshotData": [],
          "type": "dashboard",
          "builtIn": 1,
          "hide": true
        }
      ]
    },
    "editable": true,
    "fiscalYearStartMonth": 0,
    "graphTooltip": 0,
    "id": 20,
    "links": [],
    "liveNow": false,
    "panels": [
      {
        "datasource": null,
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
            },
            "unit": "binbps"
          },
          "overrides": []
        },
        "gridPos": {
          "h": 9,
          "w": 24,
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
        "snapshotData": [
          {
            "fields": [
              {
                "config": {
                  "color": {
                    "mode": "palette-classic"
                  },
                  "custom": {
                    "axisCenteredZero": false,
                    "axisColorMode": "text",
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
                    "lineWidth": 1,
                    "pointSize": 5,
                    "showPoints": "auto",
                    "thresholdsStyle": {
                      "mode": "off"
                    }
                  },
                  "interval": 20000,
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
                  },
                  "unit": "binbps"
                },
                "name": "Time",
                "type": "time",
                "values": ` + inTime + `
              },
              {
                "config": {
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
                  "displayNameFromDS": "流入-` + host.NetworkName + `",
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
                  },
                  "unit": "binbps"
                },
                "labels": {
                  "agent_hostname": "` + host.IpAddr + `",
                  "interface": "+` + host.NetworkName + `+"
                },
                "name": "流入-+` + host.NetworkName + `+",
                "type": "number",
                "values": ` + inValue + `
              }
            ],
            "meta": {
              "custom": {
                "resultType": "matrix"
              },
              "executedQueryString": "Expr: idelta(net_bits_recv{interface=\"` + host.NetworkName + `\", agent_hostname=\"+` + host.NetworkName + `+\"}[1m]) / 15\nStep: 20s",
              "type": "heatmap-rows",
              "typeVersion": [
                0,
                1
              ]
            },
            "refId": "流入流量"
          },
          {
            "fields": [
              {
                "config": {
                  "color": {
                    "mode": "palette-classic"
                  },
                  "custom": {
                    "axisCenteredZero": false,
                    "axisColorMode": "text",
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
                    "lineWidth": 1,
                    "pointSize": 5,
                    "showPoints": "auto",
                    "thresholdsStyle": {
                      "mode": "off"
                    }
                  },
                  "interval": 20000,
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
                  },
                  "unit": "binbps"
                },
                "name": "Time",
                "type": "time",
                "values": ` + outTime + `
              },
              {
                "config": {
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
                  "displayNameFromDS": "流出-` + host.NetworkName + `",
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
                  },
                  "unit": "binbps"
                },
                "labels": {
                  "agent_hostname": "38.148.214.2",
                  "interface": "ens33"
                },
                "name": "流出-` + host.NetworkName + `",
                "type": "number",
                "values": ` + outValue + `
              }
            ],
            "meta": {
              "custom": {
                "resultType": "matrix"
              },
              "executedQueryString": "Expr: idelta(net_bits_sent{interface=\"` + host.NetworkName + `\", agent_hostname=\"` + host.IpAddr + `\"}[1m]) / 15\nStep: 20s",
              "type": "heatmap-rows",
              "typeVersion": [
                0,
                1
              ]
            },
            "refId": "流出流量"
          }
        ],
        "targets": [],
        "title": "` + host.IpAddr + ` 进出流量",
        "type": "timeseries",
        "links": []
      },
      {
        "datasource": null,
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
          "w": 24,
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
        "snapshotData": [
          {
            "fields": [
              {
                "config": {
                  "color": {
                    "mode": "palette-classic"
                  },
                  "custom": {
                    "axisCenteredZero": false,
                    "axisColorMode": "text",
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
                    "showPoints": "auto",
                    "thresholdsStyle": {
                      "mode": "off"
                    }
                  },
                  "interval": 20000,
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
                "name": "Time",
                "type": "time",
                "values": ` + loadOneTime + `
              },
              {
                "config": {
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
                  "displayNameFromDS": "{__name__=\"system_load1\", agent_hostname=\"` + host.IpAddr + `\"}",
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
                "labels": {
                  "__name__": "system_load1",
                  "agent_hostname": "` + host.IpAddr + `"
                },
                "name": "system_load1",
                "type": "number",
                "values": ` + loadOneValue + `
              }
            ],
            "meta": {
              "custom": {
                "resultType": "matrix"
              },
              "executedQueryString": "Expr: system_load1{agent_hostname=\"` + host.IpAddr + `\"}\nStep: 20s",
              "preferredVisualisationType": "graph",
              "type": "timeseries-multi",
              "typeVersion": [
                0,
                1
              ]
            },
            "refId": "A"
          },
          {
            "fields": [
              {
                "config": {
                  "color": {
                    "mode": "palette-classic"
                  },
                  "custom": {
                    "axisCenteredZero": false,
                    "axisColorMode": "text",
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
                    "showPoints": "auto",
                    "thresholdsStyle": {
                      "mode": "off"
                    }
                  },
                  "interval": 20000,
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
                "name": "Time",
                "type": "time",
                "values": ` + loadFiveTime + `
              },
              {
                "config": {
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
                  "displayNameFromDS": "{__name__=\"system_load5\", agent_hostname=\"` + host.IpAddr + `\"}",
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
                "labels": {
                  "__name__": "system_load5",
                  "agent_hostname": "` + host.IpAddr + `"
                },
                "name": "system_load5",
                "type": "number",
                "values": ` + loadFiveValue + `
              }
            ],
            "meta": {
              "custom": {
                "resultType": "matrix"
              },
              "executedQueryString": "Expr: system_load5{agent_hostname=\"` + host.IpAddr + `\"}\nStep: 20s",
              "preferredVisualisationType": "graph",
              "type": "timeseries-multi",
              "typeVersion": [
                0,
                1
              ]
            },
            "refId": "B"
          },
          {
            "fields": [
              {
                "config": {
                  "color": {
                    "mode": "palette-classic"
                  },
                  "custom": {
                    "axisCenteredZero": false,
                    "axisColorMode": "text",
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
                    "showPoints": "auto",
                    "thresholdsStyle": {
                      "mode": "off"
                    }
                  },
                  "interval": 20000,
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
                "name": "Time",
                "type": "time",
                "values": ` + loadFifteenTime + `
              },
              {
                "config": {
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
                  "displayNameFromDS": "{__name__=\"system_load15\", agent_hostname=\"` + host.IpAddr + `\"}",
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
                "labels": {
                  "__name__": "system_load15",
                  "agent_hostname": "` + host.IpAddr + `"
                },
                "name": "system_load15",
                "type": "number",
                "values": ` + loadFifteenValue + `
              }
            ],
            "meta": {
              "custom": {
                "resultType": "matrix"
              },
              "executedQueryString": "Expr: system_load15{agent_hostname=\"` + host.IpAddr + `\"}\nStep: 20s",
              "preferredVisualisationType": "graph",
              "type": "timeseries-multi",
              "typeVersion": [
                0,
                1
              ]
            },
            "refId": "C"
          }
        ],
        "targets": [],
        "title": "` + host.IpAddr + ` 系统负载",
        "type": "timeseries",
        "links": []
      },
      {
        "datasource": null,
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
            },
            "unit": "percent"
          },
          "overrides": []
        },
        "gridPos": {
          "h": 8,
          "w": 24,
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
        "snapshotData": [
          {
            "fields": [
              {
                "config": {
                  "color": {
                    "mode": "palette-classic"
                  },
                  "custom": {
                    "axisCenteredZero": false,
                    "axisColorMode": "text",
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
                    "showPoints": "auto",
                    "thresholdsStyle": {
                      "mode": "off"
                    }
                  },
                  "interval": 20000,
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
                  },
                  "unit": "percent"
                },
                "name": "Time",
                "type": "time",
                "values": ` + cpuTime + `
              },
              {
                "config": {
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
                  "displayNameFromDS": "{__name__=\"cpu_usage_user\", agent_hostname=\"` + host.IpAddr + `\", cpu=\"cpu-total\"}",
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
                  },
                  "unit": "percent"
                },
                "labels": {
                  "__name__": "cpu_usage_user",
                  "agent_hostname": "` + host.IpAddr + `",
                  "cpu": "cpu-total"
                },
                "name": "cpu_usage_user",
                "type": "number",
                "values": ` + cpuValue + `
              }
            ],
            "meta": {
              "custom": {
                "resultType": "matrix"
              },
              "executedQueryString": "Expr: cpu_usage_user{agent_hostname=\"` + host.IpAddr + `\"}\nStep: 20s",
              "preferredVisualisationType": "graph",
              "type": "timeseries-multi",
              "typeVersion": [
                0,
                1
              ]
            },
            "refId": "cpu空闲"
          }
        ],
        "targets": [],
        "title": "` + host.IpAddr + ` CPU使用百分比",
        "type": "timeseries",
        "links": []
      },
      {
        "datasource": null,
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
            },
            "unit": "percent"
          },
          "overrides": []
        },
        "gridPos": {
          "h": 8,
          "w": 24,
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
        "snapshotData": [
          {
            "fields": [
              {
                "config": {
                  "color": {
                    "mode": "palette-classic"
                  },
                  "custom": {
                    "axisCenteredZero": false,
                    "axisColorMode": "text",
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
                    "showPoints": "auto",
                    "thresholdsStyle": {
                      "mode": "off"
                    }
                  },
                  "interval": 20000,
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
                  },
                  "unit": "percent"
                },
                "name": "Time",
                "type": "time",
                "values": ` + freeTime + `
              },
              {
                "config": {
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
                  "displayNameFromDS": "{__name__=\"mem_used_percent\", agent_hostname=\"` + host.IpAddr + `\"}",
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
                  },
                  "unit": "percent"
                },
                "labels": {
                  "__name__": "mem_used_percent",
                  "agent_hostname": "` + host.IpAddr + `"
                },
                "name": "mem_used_percent",
                "type": "number",
                "values": ` + freeValue + `
              }
            ],
            "meta": {
              "custom": {
                "resultType": "matrix"
              },
              "executedQueryString": "Expr: mem_used_percent{agent_hostname=\"` + host.IpAddr + `\"}\nStep: 20s",
              "preferredVisualisationType": "graph",
              "type": "timeseries-multi",
              "typeVersion": [
                0,
                1
              ]
            },
            "refId": "内存使用百分比"
          }
        ],
        "targets": [],
        "title": "` + host.IpAddr + ` 内存使用百分比",
        "type": "timeseries",
        "links": []
      }
    ],
    "refresh": "5s",
    "schemaVersion": 38,
    "snapshot": {
      "timestamp": "` + nowFormatStr + `"
    },
    "style": "dark",
    "tags": [
      "24"
    ],
    "templating": {
      "list": []
    },
    "time": {
      "from": "` + beforeNowFormatStr + `",
      "to": "` + nowFormatStr + `",
      "raw": {
        "from": "now-3h",
        "to": "now"
      }
    },
    "timepicker": {},
    "timezone": "Asia/Shanghai",
    "title": "` + host.Name + `-` + host.IpAddr + `",
    "uid": "` + initialize.Grafana.PrometheusUid + `",
    "version": 1,
    "weekStart": ""
  },
  "name": "` + host.Name + `-` + host.IpAddr + `",
  "expires": 600
}`).
		End()
	if res.StatusCode == 200 {
		return body, true
	}

	return body, false
}

func formatMapToString(v any) (string, error) {
	mV, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(mV), nil
}
