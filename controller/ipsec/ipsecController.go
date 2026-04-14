package ipsec

import (
	"OperationAndMonitoring/grafana/alterrules"
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/form"
	"OperationAndMonitoring/model/req"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/prometheus"
	"OperationAndMonitoring/utils"
	convert2 "OperationAndMonitoring/utils/convert"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/ybzhanghx/copier"
	"strings"
	"time"
)

func Page(c *gin.Context) {
	var ipsecReq req.IpsecReq
	err := c.ShouldBindQuery(&ipsecReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("传入参数错误！"+err.Error()))
		return
	}

	err, ipsecs, total := GetIpsecListInfo(ipsecReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("mysql查询错误！"+err.Error()))
		return
	}

	list := []vo.IpsecVO{}
	for _, ipsec := range ipsecs {
		var ipsecVO vo.IpsecVO

		err = copier.CopyByTag(&ipsecVO, &ipsec, "mson")
		if err != nil {
			c.JSON(200, utils.FailedRespon("结构体复制错误！"))
			return
		}
		// 解决时间没法复制
		ipsecVO.CreateTime = ipsec.CreateTime.Format("2006-01-02")

		// 用户组为空则不返回分组名
		//if lo.IsEmpty(ipsec.GroupID) {
		//	utils.First(&entity.Group{ID: ipsec.GroupID}, &group)
		//	ipsecVO.GroupName = group.Name
		//}

		list = append(list, ipsecVO)
	}

	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = total
	c.JSON(200, utils.SuccessRespon(pageR))

}

func Refresh(c *gin.Context) {
	fmt.Println("刷新！")
	_, err := prometheus.GetActiveIpsecDomain()
	if err != nil {
		c.JSON(200, utils.FailedRespon(err.Error()))
	}
	c.JSON(200, utils.SuccessRespon("刷新服务器成功"))
}

func RefreshOnlineNum(c *gin.Context) {
	fmt.Println("刷新在线数！")
	prometheus.GetIpsecDomainOnlineNum()
	c.JSON(200, utils.SuccessRespon("刷新服务器成功"))
}

func Delete(c *gin.Context) {
	var idUints []uint64
	param := c.Param("ids")
	ids := strings.Split(param, ",")

	for _, id := range ids {
		var ipsec entity.Ipsec
		idUints = append(idUints, convert2.ToUint64(id))
		notFound, err := utils.First(&entity.Ipsec{ID: convert2.ToUint64(id)}, &ipsec)

		if err != nil {
			c.JSON(200, utils.FailedRespon("查询mysql数据库报错！"))
			return
		}
		if notFound {
			c.JSON(200, utils.FailedRespon("数据库数据报错！"))
			return
		}

		if *ipsec.Status != 0 {
			ok, errMsg := alterrules.DeleteAlterRules(*ipsec.RuleUID)
			if !ok {
				c.JSON(200, utils.FailedRespon("删除服务器告警策略失败！"+errMsg))
				return
			}
		}

	}

	_, err := utils.DeleteByIDS(&entity.Ipsec{}, idUints)
	if err != nil {
		c.JSON(200, utils.FailedRespon("删除服务器Mysqld发生错误！"))
		return
	}

	c.JSON(200, utils.SuccessRespon("删除服务器成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("ipsecId")
	ipsecId, _ := convert2.ToUint64E(param)

	var ipsec entity.Ipsec
	var ipsecForm form.IpsecForm
	utils.First(&entity.Ipsec{ID: ipsecId}, &ipsec)
	if lo.IsEmpty(*ipsec.Status) {

	} else {
		copier.CopyByTag(&ipsecForm, &ipsec, "mson")
		c.JSON(200, utils.SuccessRespon(ipsecForm))
	}
}

func Update(c *gin.Context) {
	param := c.Param("ipsecId")
	ipsecId, _ := convert2.ToUint64E(param)

	var ipsecForm form.IpsecForm
	var ipsec entity.Ipsec

	c.BindJSON(&ipsecForm)
	utils.First(&entity.Ipsec{ID: ipsecId}, &ipsec)

	copier.CopyByTag(&ipsec, &ipsecForm, "mson")

	ipsec.UpdateTime = time.Now()

	if lo.IsNotEmpty(*ipsec.Status) {
		ok := alterrules.UpdateIpsecAlterRules(ipsec.AgentHostname, ipsec.Domain, *ipsec.RuleUID, *ipsec.AlterNum)
		if !ok {
			c.JSON(200, utils.FailedRespon("更新服务器失败！"))
			return
		}
	}

	utils.Updates(&entity.Ipsec{ID: ipsecId}, &ipsec)
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
	param := c.Param("ipsecId")
	ipsecId, _ := convert2.ToUint64E(param)
	status := convert2.ToInt(c.Query("status"))
	var ipsec entity.Ipsec
	utils.First(&entity.Ipsec{ID: ipsecId}, &ipsec)
	var ruleID string
	if lo.IsEmpty(status) {
		ok, errMsg := alterrules.DeleteAlterRules(*ipsec.RuleUID)
		if !ok {
			c.JSON(200, utils.FailedRespon("删除ipsec域名告警策略失败！"+errMsg))
			return
		}
		ruleID = ""
	} else {
		body, ok := alterrules.CreateIpsecDomainAlterRules(ipsec.AgentHostname, ipsec.Domain, *ipsec.AlterNum)
		if !ok {
			c.JSON(200, utils.FailedRespon("创建ipsec域名告警策略失败！"))
			return
		}
		ruleID = strings.Split(strings.Split(body, "\"uid\"")[1], "\"")[1]
	}
	err := utils.Updates(&entity.Ipsec{ID: ipsecId}, &entity.Ipsec{Status: &status, RuleUID: &ruleID})
	if err != nil {
		c.JSON(200, utils.FailedRespon("开启ipsec域名告警失败！"))
		return
	}
	c.JSON(200, utils.SuccessRespon("开启ipsec域名告警成功！"))
}

//func Network(c *gin.Context) {
//	param := c.Param("ipsecId")
//
//	ipsecId, _ := convert2.ToUint64E(param)
//	var ipsec entity.Ipsec
//	utils.First(&entity.Ipsec{ID: ipsecId}, &ipsec)
//	req := prometheus.GetNetworkName(ipsec.IpAddr)
//
//	var list []vo.Option
//
//	for i, s := range req.Data {
//		var option vo.Option
//		option.Value = i
//		option.Label = s
//		list = append(list, option)
//	}
//	c.JSON(200, utils.SuccessRespon(list))
//}
