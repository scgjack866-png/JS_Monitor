package prometheus

import (
	"OperationAndMonitoring/grafana/alterrules"
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/utils"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"strconv"
	"strings"
	"time"
)

var (
	request = *gorequest.New()
)

func GetActiveAgentHostname() (vo.PrometheusReq, error) {
	var req vo.PrometheusReq
	timeStr := time.Now().Add(-1 * time.Minute).In(time.UTC).Format(time.RFC3339)
	_, body, err := request.Get(initialize.Prometheus.ApiUrl + "/api/v1/label/agent_hostname/values?start=" + timeStr).End()
	if err != nil {
		return req, err[0]
	}

	errs := json.Unmarshal([]byte(body), &req)
	if errs != nil {
		return req, errs
	}

	var hosts []entity.Host

	utils.Find(&entity.Host{}, &hosts)

	for _, data := range req.Data {
		flag := true
		for _, s := range hosts {
			if s.IpAddr == data {
				flag = false
				break
			}
		}
		if flag {
			host := entity.Host{
				IpAddr:     data,
				DelayTime:  time.UnixMilli(0),
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			}
			errs = utils.Create(&host)
			if errs != nil {
				return req, errs
			}
		}
	}
	return req, nil
}

func GetNetworkName(ipAddr string) vo.PrometheusReq {
	timeStr := time.Now().Add(-1 * time.Minute).In(time.UTC).Format(time.RFC3339)
	_, body, err := request.Get(initialize.Prometheus.ApiUrl + `/api/v1/label/interface/values?match[]={agent_hostname="` + ipAddr + `"}&start=` + timeStr).End()

	if err != nil {
		fmt.Println(err)
	}
	var req vo.PrometheusReq
	errs := json.Unmarshal([]byte(body), &req)
	if errs != nil {
		fmt.Println(err)
	}
	return req
}

func GetActiveIpsecDomain() (vo.PrometheusQueryReq, error) {

	var req vo.PrometheusQueryReq
	_, body, err := request.Get(initialize.Prometheus.ApiUrl + "/api/v1/query?query=ipsec_online_num").End()
	if err != nil {
		return req, err[0]
	}

	errs := json.Unmarshal([]byte(body), &req)
	if errs != nil {
		return req, errs
	}

	var ipsecDomains []entity.Ipsec

	utils.Find(&entity.Ipsec{}, &ipsecDomains)

	for _, ipsecDomain := range req.Data.Result {
		flag := true
		fmt.Println(ipsecDomain)
		for _, s := range ipsecDomains {
			if s.Domain == ipsecDomain.Metric.Domain && s.AgentHostname == ipsecDomain.Metric.AgentHostname {
				flag = false
				break
			}
		}
		fmt.Println(flag)
		if flag {
			if err != nil {
				return req, errs
			}
			body, ok := alterrules.CreateIpsecDomainAlterRules(ipsecDomain.Metric.AgentHostname, ipsecDomain.Metric.Domain, 30)
			if !ok {
				return req, errs
			}
			fmt.Println(body)
			ruleID := strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
			statusSort := 1
			ipsec := entity.Ipsec{
				AgentHostname: ipsecDomain.Metric.AgentHostname,
				Domain:        ipsecDomain.Metric.Domain,
				RuleUID:       &ruleID,
				Status:        &statusSort,
				Sort:          &statusSort,
				CreateTime:    time.Now(),
				UpdateTime:    time.Now(),
			}
			errs = utils.Create(&ipsec)
			if errs != nil {
				return req, errs
			}
		}
	}

	return req, nil
}

func GetIpsecDomainOnlineNum() (vo.PrometheusQueryReq, error) {
	var req vo.PrometheusQueryReq
	_, body, err := request.Get(initialize.Prometheus.ApiUrl + "/api/v1/query?query=ipsec_online_num").End()
	if err != nil {
		fmt.Println(err)
		return req, err[0]
	}

	errs := json.Unmarshal([]byte(body), &req)
	if errs != nil {
		return req, errs
	}

	var ipsecDomains []entity.Ipsec

	utils.Find(&entity.Ipsec{}, &ipsecDomains)

	for _, ipsecDomain := range req.Data.Result {

		for _, s := range ipsecDomains {
			if s.Domain == ipsecDomain.Metric.Domain && s.AgentHostname == ipsecDomain.Metric.AgentHostname {
				for i, v := range ipsecDomain.Value {
					switch v.(type) {
					case int:
						fmt.Printf("slice[%d] is an int: %v\n", i, v.(int))
					case string:
						fmt.Printf("slice[%d] is a string: %v\n", i, v.(string))
						*s.OnlineNum, _ = strconv.Atoi(v.(string))
						errs = utils.Updates(&entity.Ipsec{ID: s.ID}, &s)
					case float64:
						fmt.Printf("slice[%d] is a float64: %v\n", i, v.(float64))
					}
				}
			}
		}
	}

	return req, nil
}

func GetQuery() ([]vo.DomainStatusVo, error) {
	var req vo.PrometheusQueryReq
	var domainStatuss []vo.DomainStatusVo

	seconds := time.Now().Unix()

	_, body, err := request.Post(initialize.Prometheus.ApiUrl + "/api/v1/query?query=http_response_response_code&time=" + strconv.FormatInt(seconds, 10)).Send(``).End()
	if err != nil {
		fmt.Println(err)
		return domainStatuss, err[0]
	}

	errs := json.Unmarshal([]byte(body), &req)

	if errs != nil {
		return domainStatuss, errs
	}

	var nodes []entity.Node

	utils.Find(&entity.Node{}, &nodes)

	var domains []entity.Domain

	status := 1
	utils.Find(&entity.Domain{Status: &status}, &domains)

	for _, domain := range domains {

		var domainStatus vo.DomainStatusVo
		domainStatus.Target = domain.Domain
		for _, node := range nodes {

			var values []vo.Value
			for _, domainR := range req.Data.Result {

				var value vo.Value
				value.Describe.AgentHostname = node.NodeIP

				if domain.Domain == domainR.Metric.Target {
					if node.NodeIP == domainR.Metric.AgentHostname {
						value.Values = domainR.Value
					}
				}
				values = append(values, value)
			}
			domainStatus.Value = values
		}

		domainStatuss = append(domainStatuss, domainStatus)
		fmt.Println(domainStatus)
	}

	return domainStatuss, nil
}
