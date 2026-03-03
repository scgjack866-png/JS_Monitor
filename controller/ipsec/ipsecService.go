package ipsec

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/req"
	"OperationAndMonitoring/mysql"
	"OperationAndMonitoring/utils"
	"github.com/samber/lo"
)

func GetIpsecListInfo(req req.IpsecReq) (error, []entity.Ipsec, int64) {
	var whereOrder []mysql.PageWhereOrder
	var ipsecs []entity.Ipsec
	var total int64

	whereOrder = append(whereOrder, mysql.PageWhereOrder{Order: "sort asc"})

	if lo.IsNotEmpty(req.Keywords) {
		v := "%" + req.Keywords + "%"
		var arr []interface{}
		arr = append(arr, v)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "domain like ? ", Value: arr})
	}
	if lo.IsNotEmpty(req.Status) {
		var arr []interface{}
		arr = append(arr, req.Status)
		whereOrder = append(whereOrder, mysql.PageWhereOrder{Where: "status = ?", Value: arr})
	}

	err := utils.GetPage(&entity.Ipsec{}, &entity.Ipsec{}, &ipsecs, req.PageNum, req.PageSize, &total, whereOrder...)
	return err, ipsecs, total
}
