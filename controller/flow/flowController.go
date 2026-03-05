package host

import (
	"OperationAndMonitoring/grafana/ds"
	"OperationAndMonitoring/initialize"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/req"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/convert"
	"sync"

	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"strconv"
	"strings"
	"time"
)

func All(c *gin.Context) {
	var flowReq req.FlowReq
	err := c.BindJSON(&flowReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("传入参数错误！"+err.Error()))
		return
	}

	var list []vo.FlowVO
	now := time.Now().UnixNano() / 1e6

	nowStr := strconv.FormatInt(now, 10)

	var safeMap sync.Map

	for _, ip := range flowReq.Ips {
		var flowVO vo.FlowVO
		var host entity.Host
		//notFound, _ := utils.First(&entity.Host{AllIp: "%" + ip + ",%"}, &host)
		v := "%" + ip + ",%"
		var arr []interface{}
		arr = append(arr, v)

		err = utils.Find(&entity.Host{}, &host, mysql.PageWhereOrder{Where: "all_ip like ?", Value: arr})

		if err != nil {
			return
		}
		flowVO.Ip = ip

		if lo.IsNil(host.MachineCode) || lo.IsEmpty(host.MachineCode) {

			flowVO.Out = -1
			flowVO.In = -1
		} else {

			if value, ok := safeMap.Load(host.MachineCode); ok {
				if valueF, ok := value.(vo.FlowVO); ok {
					flowVO = valueF
				} else {
					c.JSON(200, utils.FailedRespon("断言错误！"))
					return
				}
			} else {
				body, ok := ds.QueryDsFlow(host, nowStr, nowStr)

				if !ok {
					c.JSON(200, utils.FailedRespon("请求Grafana错误！"))
					return
				}

				var dsQueryFlowBody vo.DsQueryFlowBody
				err = json.Unmarshal([]byte(body), &dsQueryFlowBody)
				if err != nil {
					c.JSON(200, utils.FailedRespon("解析字符串错误！"))
					return
				}
				if lo.IsEmpty(len(dsQueryFlowBody.Results.Out.Frames[0].Data.Values)) {

					flowVO.Out = -2
					flowVO.In = -2
				} else {
					flowVO.Out = dsQueryFlowBody.Results.Out.Frames[0].Data.Values[1][0]
					flowVO.In = dsQueryFlowBody.Results.In.Frames[0].Data.Values[1][0]
					flowVO.OutUnit = humanize.IBytes(convert.ToUint64(flowVO.Out))
					flowVO.InUnit = humanize.IBytes(convert.ToUint64(flowVO.In))
				}

				safeMap.Store(host.MachineCode, flowVO)

			}
		}
		list = append(list, flowVO)

	}

	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = int64(len(flowReq.Ips))
	c.JSON(200, utils.SuccessRespon(pageR))
}

func GetSnapshotUrl(c *gin.Context) {

	var flowSnapshotReq req.FlowSnapshotReq
	err := c.BindJSON(&flowSnapshotReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("传入参数错误！"+err.Error()))
		return
	}

	t := time.Now()

	now := t.UnixNano() / 1e6
	beforeNow := t.Add(-3*time.Hour).UnixNano() / 1e6

	nowStr := strconv.FormatInt(now, 10)
	beforeNowStr := strconv.FormatInt(beforeNow, 10)

	nowFormatStr := t.UTC().Format("2006-01-02T15:04:05Z")
	beforeNowFormatStr := t.Add(-3 * time.Hour).UTC().Format("2006-01-02T15:04:05Z")
	// "2025-05-14T04:36:24.691Z"
	var host entity.Host

	v := "%" + flowSnapshotReq.Ip + ",%"
	var arr []interface{}
	arr = append(arr, v)

	err = utils.Find(&entity.Host{}, &host, mysql.PageWhereOrder{Where: "all_ip like ?", Value: arr})

	if err != nil {
		c.JSON(200, utils.FailedRespon("查询数据库失败！"))
		return
	}

	if lo.IsNil(host.MachineCode) || lo.IsEmpty(host.MachineCode) {
		c.JSON(200, utils.FailedRespon("不存在该IP的主机！"))
		return
	}

	flowBody, ok := ds.QueryDsFlow(host, nowStr, beforeNowStr)

	cpuBody, ok := ds.QueryDsCpu(host, nowStr, beforeNowStr)

	freeBody, ok := ds.QueryDsFree(host, nowStr, beforeNowStr)

	loadBody, ok := ds.QueryDsLoad(host, nowStr, beforeNowStr)

	if !ok {
		c.JSON(200, utils.FailedRespon("请求Grafana错误！"+flowBody+"请求Grafana错误！"+cpuBody+"请求Grafana错误！"+freeBody+"请求Grafana错误！"+loadBody))
		return
	}

	var dsQueryFlowBody vo.DsQueryFlowBody
	var dsQueryCpuBody vo.DsQueryCpuBody
	var dsQueryFreeBody vo.DsQueryFreeBody
	var dsQueryLoadBody vo.DsQueryLoadBody

	err = json.Unmarshal([]byte(flowBody), &dsQueryFlowBody)

	err = json.Unmarshal([]byte(cpuBody), &dsQueryCpuBody)

	err = json.Unmarshal([]byte(freeBody), &dsQueryFreeBody)

	err = json.Unmarshal([]byte(loadBody), &dsQueryLoadBody)

	if err != nil {
		c.JSON(200, utils.FailedRespon("解析字符串错误！"))
		return
	}

	OutValue := dsQueryFlowBody.Results.Out.Frames[0].Data.Values

	InValue := dsQueryFlowBody.Results.In.Frames[0].Data.Values

	CpuValue := dsQueryCpuBody.Results.CPU.Frames[0].Data.Values

	FreeValue := dsQueryFreeBody.Results.Free.Frames[0].Data.Values

	LoadOneValue := dsQueryLoadBody.Results.LoadOne.Frames[0].Data.Values

	LoadFiveValue := dsQueryLoadBody.Results.LoadFive.Frames[0].Data.Values

	LoadFifteenValue := dsQueryLoadBody.Results.LoadFifteen.Frames[0].Data.Values

	snapshotUrlBody, ok := ds.GetSnapshotUrl(host, nowFormatStr, beforeNowFormatStr, OutValue, InValue, CpuValue, FreeValue, LoadOneValue, LoadFiveValue, LoadFifteenValue)

	if !ok {
		c.JSON(200, utils.FailedRespon("请求Grafana错误！"))
		return
	}

	var snapshotVO vo.SnapshotVO

	snapshotVO.Url = strings.ReplaceAll(strings.Split(strings.Split(snapshotUrlBody, "url")[1], "\"")[2], "localhost:3000", initialize.Grafana.SnapshotDomain)

	c.JSON(200, utils.SuccessRespon(snapshotVO))

}

func FilterUnderOneM(c *gin.Context) {
	threshold := 102400.0
	if thresholdStr := c.Query("threshold"); thresholdStr != "" {
		parsedThreshold, err := strconv.ParseFloat(thresholdStr, 64)
		if err != nil || parsedThreshold < 0 {
			c.JSON(200, utils.FailedRespon("threshold 参数错误！"))
			return
		}
		threshold = parsedThreshold
	}

	var list []vo.FlowVO
	now := time.Now().UnixNano() / 1e6

	nowStr := strconv.FormatInt(now, 10)

	var safeMap sync.Map
	var hosts []entity.Host

	err := utils.Find(&entity.Host{}, &hosts)
	if err != nil {
		c.JSON(200, utils.FailedRespon("数据库报错！"))
		return
	}

	for _, host := range hosts {
		var flowVO vo.FlowVO
		flowVO.Ip = host.IpAddr
		if lo.IsNil(host.MachineCode) || lo.IsEmpty(host.MachineCode) {

			flowVO.Out = -1
			flowVO.In = -1
		} else {

			if value, ok := safeMap.Load(host.MachineCode); ok {
				if valueF, ok := value.(vo.FlowVO); ok {
					flowVO = valueF
				} else {
					c.JSON(200, utils.FailedRespon("断言错误！"))
					return
				}
			} else {
				body, ok := ds.QueryDsFlow(host, nowStr, nowStr)

				if !ok {
					c.JSON(200, utils.FailedRespon("请求Grafana错误！"))
					return
				}

				var dsQueryFlowBody vo.DsQueryFlowBody
				err = json.Unmarshal([]byte(body), &dsQueryFlowBody)
				if err != nil {
					c.JSON(200, utils.FailedRespon("解析字符串错误！"))
					return
				}
				if lo.IsEmpty(len(dsQueryFlowBody.Results.Out.Frames[0].Data.Values)) {

					flowVO.Out = -2
					flowVO.In = -2
				} else {
					flowVO.Out = dsQueryFlowBody.Results.Out.Frames[0].Data.Values[1][0]
					flowVO.In = dsQueryFlowBody.Results.In.Frames[0].Data.Values[1][0]
					if flowVO.Out >= threshold || flowVO.In >= threshold {
						continue
					}
					flowVO.OutUnit = humanize.IBytes(convert.ToUint64(flowVO.Out))
					flowVO.InUnit = humanize.IBytes(convert.ToUint64(flowVO.In))
				}
				safeMap.Store(host.MachineCode, flowVO)
			}
		}
		list = append(list, flowVO)

	}

	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = int64(len(list))
	c.JSON(200, utils.SuccessRespon(pageR))
}
