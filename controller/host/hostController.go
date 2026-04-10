package host

import (
	"OperationAndMonitoring/grafana/alterrules"
	"OperationAndMonitoring/grafana/dashboards"
	"OperationAndMonitoring/grafana/ds"
	"OperationAndMonitoring/grafana/silence"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/form"
	"OperationAndMonitoring/model/req"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/prometheus"
	"OperationAndMonitoring/utils"
	convert2 "OperationAndMonitoring/utils/convert"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/ybzhanghx/copier"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Page(c *gin.Context) {
	var hostReq req.HostReq
	err := c.ShouldBindQuery(&hostReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("传入参数错误！"+err.Error()))
		return
	}

	err, hosts, total := GetHostListInfo(hostReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("mysql查询错误！"+err.Error()))
		return
	}

	list := []vo.HostVO{}
	for _, host := range hosts {
		var hostVO vo.HostVO
		var group entity.Group
		err = copier.CopyByTag(&hostVO, &host, "mson")
		if err != nil {
			c.JSON(200, utils.FailedRespon("结构体复制错误！"))
			return
		}
		// 解决时间没法复制
		hostVO.CreateTime = host.CreateTime.Format("2006-01-02")
		allIps := strings.Split(host.AllIp, ",")
		validIps := make([]string, 0, len(allIps))
		seen := make(map[string]struct{}, len(allIps)+1)
		for _, ip := range allIps {
			if ip != "" {
				if _, ok := seen[ip]; ok {
					continue
				}
				seen[ip] = struct{}{}
				validIps = append(validIps, ip)
			}
		}
		if len(validIps) == 0 {
			validIps = append(validIps, host.IpAddr)
			seen[host.IpAddr] = struct{}{}
		}
		if host.IpAddr != "" {
			if _, ok := seen[host.IpAddr]; !ok {
				validIps = append([]string{host.IpAddr}, validIps...)
			}
		}
		hostVO.OtherIp = validIps
		hostVO.IpNum = len(validIps)

		// 用户组为空则不返回分组名
		if lo.IsNotEmpty(host.GroupID) {
			utils.First(&entity.Group{ID: host.GroupID}, &group)
			hostVO.GroupName = group.Name
		}

		list = append(list, hostVO)
	}

	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = total
	c.JSON(200, utils.SuccessRespon(pageR))

}

func Create(c *gin.Context) {
	//prometheus.GetActiveAgentHostname()

	var hostForm form.HostForm
	var host entity.Host
	var group entity.Group
	if err := c.BindJSON(&hostForm); err != nil {
		fmt.Println(err)
		c.JSON(200, utils.FailedRespon("获取json数据失败！"))
		return
	}

	notFound, _ := utils.First(&entity.Host{MachineCode: hostForm.MachineCode}, &host)
	hostID := host.ID
	if err := copier.CopyByTag(&host, &hostForm, "mson"); err != nil {
		c.JSON(200, utils.FailedRespon("复制结构体失败！"))
		return
	}

	delayTime, _ := time.ParseInLocation("2006-01-02 15:04:05", hostForm.DelayTime, time.Local)

	host.UpdateTime = time.Now()
	if _, err := utils.First(&entity.Group{ID: hostForm.GroupID}, &group); err != nil {
		return
	}
	var body string
	var ok bool
	var errs []error
	dashboardUID := "null"

	if lo.IsNil(host.Status) || lo.IsEmpty(*host.Status) {
		body, ok, errs = dashboards.CreateDashboards(dashboardUID, hostForm.IpAddr, hostForm.Name, group.FolderID, hostForm.NetworkName)
		var status = 1
		host.Status = &status
	} else {
		dashboardUID = "\"" + host.UID + "\""
		body, ok, errs = dashboards.CreateDashboards(dashboardUID, hostForm.IpAddr, hostForm.Name, group.FolderID, hostForm.NetworkName)
	}

	if !ok || errs != nil {
		c.JSON(200, utils.FailedRespon("请求Grafana失败！"))
		return
	}
	if delayTime.After(time.Now()) && (host.DelayTime.After(delayTime) || host.DelayTime.Before(delayTime)) {
		host.DelayTime = delayTime
		body1, ok1 := silence.CreateSilences(host.IpAddr, host.DelayTime)
		if !reflect.ValueOf(host.SilenceUID).IsNil() {
			silence.DeleteSilences(*host.SilenceUID)
		}
		if !ok1 {
			c.JSON(200, utils.FailedRespon("创建服务器失败！"))
			return
		}
		host.SilenceUID = &strings.Split(strings.Split(body1, "\"silenceID\"")[1], "\"")[1]
	}

	if !(lo.IsNil(host.IsAlter) || lo.IsEmpty(*host.IsAlter)) {

		ok = alterrules.UpdateAlterRules(host)
	}
	if !ok {
		c.JSON(200, utils.FailedRespon("创建服务器失败！"))
		return
	}

	host.UID = strings.Split(strings.Split(body, "uid")[1], "\"")[2]
	var err error
	if notFound {
		err = utils.Create(&host)
	} else {
		err = utils.Updates(&entity.Host{ID: hostID}, &host)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(200, utils.FailedRespon("创建数据库失败！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("创建服务器成功！"))
}

func Delete(c *gin.Context) {
	var idUints []uint64
	param := c.Param("ids")
	ids := strings.Split(param, ",")

	for _, id := range ids {
		var host entity.Host
		idUints = append(idUints, convert2.ToUint64(id))
		notFound, err := utils.First(&entity.Host{ID: convert2.ToUint64(id)}, &host)

		if err != nil {
			c.JSON(200, utils.FailedRespon("查询mysql数据库报错！"))
			return
		}
		if notFound {
			c.JSON(200, utils.FailedRespon("数据库数据报错！"))
			return
		}

		if *host.Status != 0 {
			ok := dashboards.DeleteDashboards(host.UID)
			if !ok {
				c.JSON(200, utils.FailedRespon("删除服务器Grafana发生错误！"))
				return
			}
		}
		if *host.IsAlter != 0 {
			ok := alterrules.DeleteAlterRules(*host.RuleUID)
			if !ok {
				c.JSON(200, utils.FailedRespon("删除服务器告警策略失败！"))
				return
			}
		}

	}

	_, err := utils.DeleteByIDS(&entity.Host{}, idUints)
	if err != nil {
		c.JSON(200, utils.FailedRespon("删除服务器Mysqld发生错误！"))
		return
	}

	c.JSON(200, utils.SuccessRespon("删除服务器成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("hostId")
	hostId, _ := convert2.ToUint64E(param)

	var host entity.Host
	var hostForm form.HostForm
	var hostFormNoGroup form.HostFormNoGroup
	utils.First(&entity.Host{ID: hostId}, &host)
	if lo.IsEmpty(*host.Status) {
		copier.CopyByTag(&hostFormNoGroup, &host, "mson")
		c.JSON(200, utils.SuccessRespon(hostFormNoGroup))
	} else {
		copier.CopyByTag(&hostForm, &host, "mson")
		hostForm.DelayTime = host.DelayTime.Format("2006-01-02 15:04:05")
		c.JSON(200, utils.SuccessRespon(hostForm))
	}
}

func Update(c *gin.Context) {
	param := c.Param("hostId")
	hostId, _ := convert2.ToUint64E(param)

	var hostForm form.HostForm
	var host entity.Host
	var group entity.Group

	c.BindJSON(&hostForm)
	utils.First(&entity.Host{ID: hostId}, &host)

	copier.CopyByTag(&host, &hostForm, "mson")

	delayTime, _ := time.ParseInLocation("2006-01-02 15:04:05", hostForm.DelayTime, time.Local)

	host.UpdateTime = time.Now()

	utils.First(&entity.Group{ID: hostForm.GroupID}, &group)
	var body string
	var ok bool
	var errs []error
	if lo.IsNil(host.Status) || lo.IsEmpty(*host.Status) {
		body, ok, errs = dashboards.CreateDashboards("null", hostForm.IpAddr, hostForm.Name, group.FolderID, hostForm.NetworkName)
		var status = 1
		host.Status = &status
	} else {
		body, ok, errs = dashboards.CreateDashboards("\""+host.UID+"\"", hostForm.IpAddr, hostForm.Name, group.FolderID, hostForm.NetworkName)
	}
	if !ok || errs != nil {
		c.JSON(200, utils.FailedRespon("更新mysql失败！"))
		return
	}

	if delayTime.After(time.Now()) && (host.DelayTime.After(delayTime) || host.DelayTime.Before(delayTime)) {
		host.DelayTime = delayTime
		body1, ok1 := silence.CreateSilences(host.IpAddr, host.DelayTime)
		if !reflect.ValueOf(host.SilenceUID).IsNil() {
			silence.DeleteSilences(*host.SilenceUID)
		}
		if !ok1 {
			c.JSON(200, utils.FailedRespon("更新服务器失败！"))
			return
		}
		host.SilenceUID = &strings.Split(strings.Split(body1, "\"silenceID\"")[1], "\"")[1]
	}
	if !ok {
		c.JSON(200, utils.FailedRespon("更新mysql失败！"))
		return
	}
	if lo.IsNil(host.IsAlter) || lo.IsEmpty(*host.IsAlter) {
		ok = alterrules.UpdateAlterRules(host)
	}
	if !ok {
		c.JSON(200, utils.FailedRespon("更新服务器失败！"))
		return
	}
	host.UID = strings.Split(strings.Split(body, "uid")[1], "\"")[2]
	utils.Updates(&entity.Host{ID: hostId}, &host)
	c.JSON(200, utils.SuccessRespon("更新服务器成功！"))
}

func Password(c *gin.Context) {
	param := c.Param("userId")
	userId, _ := convert2.ToUint64E(param)
	password, err := utils.EncryptPassword(c.Query("password"))
	if err != nil {
		c.JSON(200, utils.FailedRespon("重置密码失败！"))
	}
	err = utils.Updates(&entity.User{ID: userId}, &entity.User{Password: password})
	if err != nil {
		c.JSON(200, utils.FailedRespon("重置密码失败！"))
	}
	c.JSON(200, utils.SuccessRespon("重置密码成功！"))
}

func Status(c *gin.Context) {
	param := c.Param("hostId")
	hostId, _ := convert2.ToUint64E(param)
	status := convert2.ToInt(c.Query("status"))
	var host entity.Host
	utils.First(&entity.Host{ID: hostId}, &host)
	var ruleID string
	if lo.IsEmpty(status) {
		ok := alterrules.DeleteAlterRules(*host.RuleUID)
		if !ok {
			c.JSON(200, utils.FailedRespon("删除服务器告警策略失败！"))
			return
		}
		ruleID = ""
	} else {
		body, ok := alterrules.CreateAlterRules(host)
		if !ok {
			c.JSON(200, utils.FailedRespon("创建服务器告警策略失败！"))
			return
		}
		ruleID = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
	}
	err := utils.Updates(&entity.Host{ID: hostId}, &entity.Host{IsAlter: &status, RuleUID: &ruleID})
	if err != nil {
		c.JSON(200, utils.FailedRespon("开启服务器告警失败！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("开启服务器告警成功！"))
}

func Network(c *gin.Context) {
	param := c.Param("hostId")

	hostId, _ := convert2.ToUint64E(param)
	var host entity.Host
	utils.First(&entity.Host{ID: hostId}, &host)
	req := prometheus.GetNetworkName(host.IpAddr)

	var list []vo.Option

	for i, s := range req.Data {
		var option vo.Option
		option.Value = i
		option.Label = s
		list = append(list, option)
	}
	c.JSON(200, utils.SuccessRespon(list))
}

func UpdateFlow(c *gin.Context) {
	now := time.Now().UnixNano() / 1e6
	nowStr := strconv.FormatInt(now, 10)

	var hosts []entity.Host
	if err := utils.Find(&entity.Host{}, &hosts); err != nil {
		c.JSON(200, utils.FailedRespon("数据库报错！"))
		return
	}

	for _, host := range hosts {
		if lo.IsEmpty(host.MachineCode) {
			continue
		}
		var flowIn float64
		var flowOut float64

		body, ok := ds.QueryDsFlow(host, nowStr, nowStr)

		if !ok {
			c.JSON(200, utils.FailedRespon("请求Grafana错误！"))
			return
		}

		var dsQueryFlowBody vo.DsQueryFlowBody
		if err := json.Unmarshal([]byte(body), &dsQueryFlowBody); err != nil {
			c.JSON(200, utils.FailedRespon("解析字符串错误！"))
			return
		}

		if len(dsQueryFlowBody.Results.Out.Frames) == 0 ||
			len(dsQueryFlowBody.Results.In.Frames) == 0 ||
			len(dsQueryFlowBody.Results.Out.Frames[0].Data.Values) < 2 ||
			len(dsQueryFlowBody.Results.In.Frames[0].Data.Values) < 2 ||
			len(dsQueryFlowBody.Results.Out.Frames[0].Data.Values[1]) == 0 ||
			len(dsQueryFlowBody.Results.In.Frames[0].Data.Values[1]) == 0 {
			continue
		}

		flowOut = dsQueryFlowBody.Results.Out.Frames[0].Data.Values[1][0]
		flowIn = dsQueryFlowBody.Results.In.Frames[0].Data.Values[1][0]
		
		if err := utils.Updates(&entity.Host{ID: host.ID}, &entity.Host{
			FlowIn:  flowIn,
			FlowOut: flowOut,
		}); err != nil {
			c.JSON(200, utils.FailedRespon("更新流量失败！"))
			return
		}

	}

	c.JSON(200, utils.SuccessRespon("更新全部主机流量成功！"))
}
