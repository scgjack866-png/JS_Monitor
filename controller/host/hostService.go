package host

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/req"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"github.com/samber/lo"
)

func GetHostListInfo(req req.HostReq) (error, []entity.Host, int64) {
	var whereOrder []mysql.PageWhereOrder
	var hosts []entity.Host
	var total int64

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})

	if lo.IsNotEmpty(req.Keywords) {
		v := "%" + req.Keywords + "%"
		var arr []interface{}
		arr = append(arr, v)
		arr = append(arr, v)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "ip_addr like ? or name like ?", Value: arr})
	}
	if lo.IsNotEmpty(req.Status) {
		var arr []interface{}
		arr = append(arr, req.Status)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = ?", Value: arr})
	}

	//有个全部组的根目录
	if req.GroupID > 1 {
		var arr []interface{}
		arr = append(arr, req.GroupID)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "group_id = ?", Value: arr})
	}

	err := utils.GetPage(&entity.Host{}, &entity.Host{}, &hosts, req.PageNum, req.PageSize, &total, whereOrder...)
	return err, hosts, total
}
