package server

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/req"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"strings"
)

func GetServerListInfo(req req.ServerReq) (error, []entity.Server, int64) {
	var whereOrder []mysql.PageWhereOrder
	var servers []entity.Server
	var total int64

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})

	if !utils.IsNull(req.Keywords) {
		keywords := strings.Replace(req.Keywords, " ", "", -1)
		v := "%" + keywords + "%"
		var arr []interface{}
		arr = append(arr, v)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "all_ip like ?", Value: arr})
	}
	if !utils.IsNull(req.Status) {
		var arr []interface{}
		arr = append(arr, req.Status)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = ?", Value: arr})
	}

	//有个全部组的根目录
	if req.ZoneID != 0 && req.ZoneID != 1 {
		var arr []interface{}
		arr = append(arr, req.ZoneID)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "zone_id = ?", Value: arr})
	}

	if req.ProjectID != 0 && req.ProjectID != 1 {
		var arr []interface{}
		arr = append(arr, req.ProjectID)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "project_id = ?", Value: arr})
	}

	if req.RoomID != 0 && req.RoomID != 1 {
		var arr []interface{}
		arr = append(arr, req.RoomID)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "room_id = ?", Value: arr})
	}

	err := utils.GetPage(&entity.Server{}, &entity.Server{}, &servers, req.PageNum, req.PageSize, &total, whereOrder...)
	return err, servers, total
}
