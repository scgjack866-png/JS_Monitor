package server

import (
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
	var serverReq req.ServerReq
	err := c.ShouldBindQuery(&serverReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("传入参数错误！"+err.Error()))
		return
	}

	err, servers, total := GetServerListInfo(serverReq)
	if err != nil {
		c.JSON(200, utils.FailedRespon("mysql查询错误！"+err.Error()))
		return
	}

	list := []vo.ServerVO{}
	for _, server := range servers {
		var serverVO vo.ServerVO
		var zone entity.Zone
		var project entity.Project
		var room entity.Room
		var order entity.Order
		err = copier.CopyByTag(&serverVO, &server, "mson")
		if err != nil {
			c.JSON(200, utils.FailedRespon("结构体复制错误！"))
			return
		}
		// 解决时间没法复制
		serverVO.CreateTime = server.CreateTime.Format("2006-01-02")
		var flag bool
		var ip_xiang []string
		serverVO.IpAddr, ip_xiang, flag = utils.String2Map(server.IpAddr)
		if !flag {
			c.JSON(200, utils.FailedRespon("string转map出错！"))
			return
		}
		serverVO.IpNum = len(ip_xiang)
		serverVO.MainIp = serverVO.IpAddr[0]
		fmt.Println(server.ZoneID, server.ProjectID, server.RoomID)
		if lo.IsNotEmpty(server.ZoneID) {
			utils.First(&entity.Zone{ID: server.ZoneID}, &zone)
			serverVO.ZoneName = zone.Name
		}
		if lo.IsNotEmpty(server.ProjectID) {
			utils.First(&entity.Project{ID: server.ProjectID}, &project)
			serverVO.ProjectName = project.Name
		}
		if lo.IsNotEmpty(server.RoomID) {
			utils.First(&entity.Room{ID: server.RoomID}, &room)
			serverVO.RoomName = room.Name
		}
		if lo.IsNotEmpty(server.OrderID) {
			utils.First(&entity.Order{ID: server.OrderID}, &order)
			serverVO.OrderName = order.Name
		}

		list = append(list, serverVO)
	}

	var pageR vo.PageResult
	pageR.Data = list
	pageR.Total = total
	c.JSON(200, utils.SuccessRespon(pageR))
}

func Create(c *gin.Context) {
	var serverForm form.ServerForm
	var server entity.Server

	c.BindJSON(&serverForm)
	notfound, _ := utils.First(&entity.Server{IpAddr: serverForm.IpAddr}, &server)

	if !notfound {
		c.JSON(200, utils.FailedRespon("已存在相同IP地址的服务器，请勿重复添加！"))
		return
	}

	err := copier.CopyByTag(&server, &serverForm, "mson")
	if err != nil {
		c.JSON(200, utils.FailedRespon("结构体复制失败！"))
		return
	}
	var status = 1
	server.Status = &status
	server.CreateTime = time.Now()
	server.UpdateTime = time.Now()

	_, ipMap, flag := utils.String2Map(server.IpAddr)

	if !flag {
		c.JSON(200, utils.FailedRespon("字符串转换数组失败！"))
		return
	}

	var allIP string
	for i, ip := range ipMap {
		allIP = allIP + ip
		if i != len(ipMap)-1 {
			allIP = allIP + ","
		}
	}
	server.AllIp = allIP

	err = utils.Create(&server)
	if err != nil {
		c.JSON(200, utils.FailedRespon("数据库添加失败！"))
		return
	}

	c.JSON(200, utils.SuccessRespon("添加服务器成功！"))
}

func Delete(c *gin.Context) {
	var idUints []uint64
	param := c.Param("ids")
	ids := strings.Split(param, ",")
	var server entity.Server
	for _, id := range ids {
		idUints = append(idUints, convert2.ToUint64(id))
		notFound, err := utils.First(&entity.Server{ID: convert2.ToUint64(id)}, &entity.Server{})
		if err != nil {
			c.JSON(200, utils.FailedRespon("查询mysql数据库报错！"))
			return
		}
		if notFound {
			c.JSON(200, utils.FailedRespon("数据库数据报错！"))
			return
		}
	}

	_, err := utils.DeleteByIDS(&server, idUints)
	if err != nil {
		c.JSON(200, utils.FailedRespon("删除服务器Mysqld发生错误！"))
		return
	}

	c.JSON(200, utils.SuccessRespon("删除服务器成功！"))
}

func Form(c *gin.Context) {
	param := c.Param("serverId")
	serverId, _ := convert2.ToUint64E(param)

	var server entity.Server
	var serverForm form.ServerForm

	utils.First(&entity.Server{ID: serverId}, &server)

	copier.CopyByTag(&serverForm, &server, "mson")

	c.JSON(200, utils.SuccessRespon(serverForm))
}

func Update(c *gin.Context) {
	param := c.Param("serverId")
	serverId, _ := convert2.ToUint64E(param)

	var serverForm form.ServerForm
	var server entity.Server

	c.BindJSON(&serverForm)
	utils.First(&entity.Server{ID: serverId}, &server)

	copier.CopyByTag(&server, &serverForm, "mson")

	_, ipMap, flag := utils.String2Map(server.IpAddr)

	if !flag {
		c.JSON(200, utils.FailedRespon("字符串转换数组失败！"))
		return
	}
	var allIP string
	for i, ip := range ipMap {
		allIP = allIP + ip
		if i != len(ipMap)-1 {
			allIP = allIP + ","
		}
	}
	server.AllIp = allIP
	server.UpdateTime = time.Now()

	utils.Updates(&entity.Server{ID: serverId}, &server)
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
	param := c.Param("serverId")
	serverId, _ := convert2.ToUint64E(param)

	var server entity.Server
	utils.First(&entity.Server{ID: serverId}, &server)

	c.JSON(200, utils.SuccessRespon("开启服务器告警成功！"))
}

func Network(c *gin.Context) {
	param := c.Param("serverId")

	serverId, _ := convert2.ToUint64E(param)
	var server entity.Server
	utils.First(&entity.Server{ID: serverId}, &server)
	req := prometheus.GetNetworkName(server.IpAddr)

	var list []vo.Option

	for i, s := range req.Data {
		var option vo.Option
		option.Value = i
		option.Label = s
		list = append(list, option)
	}
	c.JSON(200, utils.SuccessRespon(list))
}
