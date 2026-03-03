package domain

import (
	"OperationAndMonitoring/controller/common"
	"OperationAndMonitoring/grafana"
	"OperationAndMonitoring/grafana/alterrules"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	convert2 "OperationAndMonitoring/utils/convert"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	copier "github.com/ybzhanghx/copier"
	"regexp"
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

	if !lo.IsEmpty(key) {
		v := "%" + key + "%"
		var arr []interface{}
		arr = append(arr, v)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "domain like ?", Value: arr})
	}

	if !lo.IsEmpty(c.Query("status")) {
		var arr []interface{}
		arr = append(arr, status)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = ?", Value: arr})
	}

	var total int64
	domains := []entity.Domain{}
	list := []vo.DomainVO{}
	err := utils.GetPage(&entity.Domain{}, &entity.Domain{}, &domains, page, limit, &total, whereOrder...)
	for _, domain := range domains {
		var domainVO vo.DomainVO
		copier.CopyByTag(&domainVO, &domain, "mson")
		fmt.Println(domainVO)
		// 解决时间没法复制
		domainVO.CreateTime = domain.CreateTime.Format("2006-01-02")

		list = append(list, domainVO)
	}
	if err != nil {
		common.ResErrSrv(c, err)
		return
	}
	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = total
	c.JSON(200, utils.SuccessRespon(pageR))
}

func Create(c *gin.Context) {
	// 获取传入的参数
	var domainVO vo.DomainVO
	var domain entity.Domain
	c.BindJSON(&domainVO)
	// 判断是否已存在相同域名
	has, _ := utils.First(&entity.Domain{Domain: domainVO.Domain}, &domain)
	if !has {
		c.JSON(200, utils.FailedRespon("该域名已经添加过，请勿重复添加！"))
		return
	}
	// 判断域名格式是否正确
	flag, _ := regexp.MatchString("^http(s)?://([a-zA-Z0-9][-a-zA-Z0-9]{0,62}\\.)+[a-zA-Z0-9]{2,6}$", domainVO.Domain)
	if !flag {
		c.JSON(200, utils.FailedRespon("输入的域名格式不对！"))
		return
	}
	// 写入数据库
	copier.Copy(&domain, &domainVO)
	domain.CreateTime = time.Now()
	domain.UpdateTime = time.Now()
	utils.Create(&domain)
	// 写入远程中心配置文件
	grafana.CmdWiterFile()
	// 添加grafana监控告警
	var nodes []entity.Node
	var status = 1
	utils.Find(&entity.Node{Status: &status}, &nodes)
	for _, node := range nodes {
		var nodeDomain entity.NodeDomain
		var domainF entity.Domain
		utils.First(&entity.Domain{Domain: domain.Domain}, &domainF)
		body, flag1 := alterrules.CreateDomainAlterRules(node.NodeIP, domainVO.Domain, convert2.ToString(domainVO.Code))
		if !flag1 {
			c.JSON(200, utils.FailedRespon("添加警报规则出错！"))
			return
		}
		nodeDomain.RuleID = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
		nodeDomain.DomainID = domainF.ID
		nodeDomain.NodeID = node.ID
		utils.Create(nodeDomain)
	}

	c.JSON(200, utils.SuccessRespon("添加域名成功！"))
}

func Delete(c *gin.Context) {
	var idUints []uint64
	param := c.Param("domainIds")
	ids := strings.Split(param, ",")
	for _, id := range ids {
		idUints = append(idUints, convert2.ToUint64(id))
		var nodes []entity.Node
		var status2 = 1
		utils.Find(&entity.Node{Status: &status2}, &nodes)
		for _, node := range nodes {
			var nodeDomain entity.NodeDomain
			utils.First(&entity.NodeDomain{NodeID: node.ID, DomainID: convert2.ToUint64(id)}, &nodeDomain)
			flag := alterrules.DeleteAlterRules(nodeDomain.RuleID)
			if !flag {
				c.JSON(200, utils.FailedRespon("监控删除失败！"))
				return
			}
		}
		utils.DeleteByWhere(&entity.Domain{}, &entity.Domain{ID: convert2.ToUint64(id)})
		utils.DeleteByWhere(&entity.NodeDomain{}, &entity.NodeDomain{DomainID: convert2.ToUint64(id)})
	}
	grafana.CmdWiterFile()
	c.JSON(200, utils.SuccessRespon("删除域名成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("domainId")
	domainId, _ := convert2.ToUint64E(param)
	var domain entity.Domain
	var domainVO vo.DomainVO
	utils.First(&entity.Domain{ID: domainId}, &domain)
	copier.CopyByTag(&domainVO, &domain, "mson")
	c.JSON(200, utils.SuccessRespon(domainVO))
}

func Update(c *gin.Context) {
	param := c.Param("domainId")
	domainId, _ := convert2.ToUint64E(param)
	var domainVO vo.DomainVO
	var domain entity.Domain
	c.BindJSON(&domainVO)
	flag, _ := regexp.MatchString("^http(s)?://([a-zA-Z0-9][-a-zA-Z0-9]{0,62}\\.)+[a-zA-Z0-9]{2,6}$", domainVO.Domain)
	if !flag {
		c.JSON(200, utils.FailedRespon("输入的域名格式不对！"))
		return
	}
	copier.Copy(&domain, &domainVO)
	domain.UpdateTime = time.Now()
	err := utils.Updates(&entity.Domain{ID: domainId}, &domain)
	if err != nil {
		c.JSON(200, utils.FailedRespon(err.Error()))
		return
	}
	if domainVO.Status == 1 {
		var nodes []entity.Node
		var status = 1
		utils.Find(&entity.Node{Status: &status}, &nodes)
		for _, node := range nodes {
			var nodeDomain entity.NodeDomain

			utils.First(&entity.NodeDomain{NodeID: node.ID, DomainID: domainId}, &nodeDomain)
			flag = alterrules.UpdateDomainAlterRules(node.NodeIP, domainVO.Domain, convert2.ToString(domainVO.Code), nodeDomain.RuleID)
			if !flag {
				c.JSON(200, utils.FailedRespon("域名更新失败！"))
				return
			}
		}
	}
	grafana.CmdWiterFile()
	c.JSON(200, utils.SuccessRespon("更新域名成功！"))
}

func Status(c *gin.Context) {
	param := c.Param("domainId")
	domainId, _ := convert2.ToUint64E(param)
	status := c.Query("status")
	intS := convert2.ToInt(status)
	var domain entity.Domain
	utils.First(&entity.Domain{ID: domainId}, &domain)
	err := utils.Updates(&entity.Domain{ID: domainId}, &entity.Domain{Status: &intS})
	if err != nil {
		c.JSON(200, utils.FailedRespon("更新域名监控状态失败！"))
		return
	}
	var nodes []entity.Node
	var status2 = 1
	if intS == 0 {
		utils.Find(&entity.Node{Status: &status2}, &nodes)
		for _, node := range nodes {
			var nodeDomain entity.NodeDomain

			utils.First(&entity.NodeDomain{NodeID: node.ID, DomainID: domainId}, &nodeDomain)
			flag := alterrules.DeleteAlterRules(nodeDomain.RuleID)
			if !flag {
				c.JSON(200, utils.FailedRespon("监控删除失败！"))
				return
			}
			utils.DeleteByWhere(&entity.NodeDomain{}, &entity.NodeDomain{DomainID: domainId, NodeID: node.ID})
		}
	} else {
		utils.Find(&entity.Node{Status: &status2}, &nodes)
		for _, node := range nodes {
			var nodeDomain entity.NodeDomain
			body, flag1 := alterrules.CreateDomainAlterRules(node.NodeIP, domain.Domain, convert2.ToString(domain.Code))
			if !flag1 {
				c.JSON(200, utils.FailedRespon("添加警报规则出错！"))
				return
			}
			nodeDomain.RuleID = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
			nodeDomain.DomainID = domainId
			nodeDomain.NodeID = node.ID
			utils.Create(nodeDomain)
		}
	}
	grafana.CmdWiterFile()
	c.JSON(200, utils.SuccessRespon("更新域名监控状态成功！"))
}
