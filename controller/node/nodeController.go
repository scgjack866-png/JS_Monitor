package node

import (
	"OperationAndMonitoring/grafana/alterrules"
	"OperationAndMonitoring/grafana/dashboards"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	convert2 "OperationAndMonitoring/utils/convert"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	copier "github.com/ybzhanghx/copier"
	"strconv"
	"strings"
	"time"
)

func Page(c *gin.Context) {
	page, _ := convert2.ToIntE(c.Query("pageNum"))
	limit, _ := convert2.ToIntE(c.Query("pageSize"))
	key := c.Query("keywords")
	status, _ := strconv.Atoi(c.Query("status"))
	var whereOrder []mysql.PageWhereOrder

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})

	if lo.IsNotEmpty(key) {
		v := "%" + key + "%"
		var arr []interface{}
		arr = append(arr, v)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "node_ip like ?", Value: arr})
	}

	if lo.IsNotEmpty(c.Query("status")) {
		var arr []interface{}
		arr = append(arr, status)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = ?", Value: arr})
	}

	var total int64
	nodes := []entity.Node{}
	list := []vo.NodeVO{}
	err := utils.GetPage(&entity.Node{}, &entity.Node{}, &nodes, page, limit, &total, whereOrder...)
	for _, node := range nodes {

		var nodeVO vo.NodeVO
		copier.CopyByTag(&nodeVO, &node, "mson")
		// 解决时间没法复制
		nodeVO.CreateTime = node.CreateTime.Format("2006-01-02")

		list = append(list, nodeVO)
	}
	if err != nil {
		c.JSON(200, utils.SuccessRespon(err))
		return
	}
	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = total
	c.JSON(200, utils.SuccessRespon(pageR))
}

func Create(c *gin.Context) {
	// 校验先
	var nodeVO vo.NodeVO
	var node entity.Node
	c.BindJSON(&nodeVO)
	has, _ := utils.First(&entity.Node{NodeIP: nodeVO.NodeIP}, &node)
	if !has {
		c.JSON(200, utils.FailedRespon("该监控点IP已经添加过，请勿重复添加！"))
		return
	}

	copier.Copy(&node, &nodeVO)
	node.CreateTime = time.Now()
	node.UpdateTime = time.Now()
	utils.Create(&node)

	_, flag := dashboards.CreateDomainDashboards("null")
	if !flag {
		c.JSON(200, utils.FailedRespon("添加域名监控面板失败！"))
		return
	}
	var domains []entity.Domain
	var status = 1
	utils.Find(&entity.Domain{Status: &status}, &domains)
	for _, domain := range domains {
		var nodeDomain entity.NodeDomain
		var nodeF entity.Node
		utils.First(&entity.Node{NodeIP: node.NodeIP}, &nodeF)
		body, flag1 := alterrules.CreateDomainAlterRules(node.NodeIP, domain.Domain, convert2.ToString(domain.Code))
		if !flag1 {
			c.JSON(200, utils.FailedRespon("添加警报规则出错！"))
			return
		}
		nodeDomain.RuleID = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
		nodeDomain.DomainID = domain.ID
		nodeDomain.NodeID = nodeF.ID
		utils.Create(nodeDomain)
	}

	c.JSON(200, utils.SuccessRespon("添加域名成功！"))
}

func Delete(c *gin.Context) {
	var idUints []uint64
	param := c.Param("ids")
	ids := strings.Split(param, ",")
	for _, id := range ids {
		idUints = append(idUints, convert2.ToUint64(id))

		var domains []entity.Domain
		var status2 = 1
		utils.Find(&entity.Domain{Status: &status2}, &domains)
		for _, domain := range domains {
			var nodeDomain entity.NodeDomain
			utils.First(&entity.NodeDomain{NodeID: convert2.ToUint64(id), DomainID: domain.ID}, &nodeDomain)
			flag, errMsg := alterrules.DeleteAlterRules(nodeDomain.RuleID)
			if !flag {
				c.JSON(200, utils.FailedRespon("监控删除失败！"+errMsg))
				return
			}
		}

		utils.DeleteByWhere(&entity.Node{}, &entity.Node{ID: convert2.ToUint64(id)})
		utils.DeleteByWhere(&entity.NodeDomain{}, &entity.NodeDomain{NodeID: convert2.ToUint64(id)})
	}
	_, flag := dashboards.CreateDomainDashboards("null")
	if !flag {
		c.JSON(200, utils.FailedRespon("删除域名监控面板失败！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("删除监控点成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("nodeId")
	nodeId, _ := convert2.ToUint64E(param)

	var node entity.Node
	var nodeVO vo.NodeVO
	utils.First(&entity.Node{ID: nodeId}, &node)

	copier.CopyByTag(&nodeVO, &node, "mson")

	c.JSON(200, utils.SuccessRespon(nodeVO))
}

func Update(c *gin.Context) {
	param := c.Param("nodeId")
	nodeId, _ := convert2.ToUint64E(param)

	var nodeVO vo.NodeVO
	var node entity.Node
	c.BindJSON(&nodeVO)

	copier.Copy(&node, &nodeVO)
	node.UpdateTime = time.Now()

	utils.Updates(&entity.Node{ID: nodeId}, &node)
	_, flag := dashboards.CreateDomainDashboards("null")
	if !flag {
		c.JSON(200, utils.FailedRespon("创建监控图失败！"))
		return
	}
	if *node.Status == 1 {
		var domains []entity.Domain
		var status = 1
		utils.Find(&entity.Domain{Status: &status}, &domains)
		for _, domain := range domains {
			var nodeDomain entity.NodeDomain

			utils.First(&entity.NodeDomain{NodeID: nodeId, DomainID: domain.ID}, &nodeDomain)
			flag = alterrules.UpdateDomainAlterRules(node.NodeIP, domain.Domain, convert2.ToString(domain.Code), nodeDomain.RuleID)
			if !flag {
				c.JSON(200, utils.FailedRespon("域名更新失败！"))
				return
			}
		}
	}

	c.JSON(200, utils.SuccessRespon("更新域名成功！"))
}

func Status(c *gin.Context) {
	param := c.Param("nodeId")
	nodeId, _ := convert2.ToUint64E(param)
	status := convert2.ToInt(c.Query("status"))
	var node entity.Node
	utils.First(&entity.Node{ID: nodeId}, &node)

	err := utils.Updates(&entity.Node{ID: nodeId}, &entity.Node{Status: &status})
	if err != nil {
		c.JSON(200, utils.FailedRespon("开启监控点监控失败！"))
		return
	}
	_, flag := dashboards.CreateDomainDashboards("null")
	if !flag {
		c.JSON(200, utils.FailedRespon("域名更新失败！"))
		return
	}

	var domains []entity.Domain
	var domainStatus = 1
	if lo.IsEmpty(status) {
		utils.Find(&entity.Domain{Status: &domainStatus}, &domains)
		for _, domain := range domains {
			var nodeDomain entity.NodeDomain

			utils.First(&entity.NodeDomain{NodeID: nodeId, DomainID: domain.ID}, &nodeDomain)
			flag, errMsg := alterrules.DeleteAlterRules(nodeDomain.RuleID)
			if !flag {
				c.JSON(200, utils.FailedRespon("监控删除失败！"+errMsg))
				return
			}
			utils.DeleteByWhere(&entity.NodeDomain{}, &entity.NodeDomain{DomainID: domain.ID, NodeID: nodeId})
		}
	} else {
		utils.Find(&entity.Domain{Status: &domainStatus}, &domains)
		for _, domain := range domains {
			var nodeDomain entity.NodeDomain
			body, flag1 := alterrules.CreateDomainAlterRules(node.NodeIP, domain.Domain, convert2.ToString(domain.Code))
			if !flag1 {
				c.JSON(200, utils.FailedRespon("添加警报规则出错！"))
				return
			}
			nodeDomain.RuleID = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
			nodeDomain.DomainID = domain.ID
			nodeDomain.NodeID = nodeId
			utils.Create(nodeDomain)
		}
	}

	c.JSON(200, utils.SuccessRespon("开启监控点监控成功！"))
}
