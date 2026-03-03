package order

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/convert"
	"github.com/ybzhanghx/copier"
)

/**
 * 递归生成部门表格层级列表
 *
 * @param parentId
 * @param deptList
 * @return
 */
func recurDeptTreeOptions(parentID uint64, orders []entity.Order) []vo.Option {
	var list []vo.Option

	for _, order := range orders {

		if *order.ParentId == parentID {
			var option vo.Option
			option.Value = order.ID
			option.Value = order.ID
			option.Label = order.Name
			option.Children = recurDeptTreeOptions(order.ID, orders)
			list = append(list, option)
		}

	}
	return list
}

/**
 * 递归生成部门表格层级列表
 *
 * @param parentId
 * @param deptList
 * @return
 */
func recurDept(parentID uint64, orders []entity.Order) []vo.OrderVO {
	var list []vo.OrderVO

	for _, order := range orders {

		if *order.ParentId == parentID {
			var orderVO vo.OrderVO
			copier.Copy(&orderVO, &order)
			orderVO.CreateTime = order.CreateTime.Format("2006-01-02")
			orderVO.UpdateTime = order.UpdateTime.Format("2006-01-02")
			orderVO.Children = recurDept(order.ID, orders)
			list = append(list, orderVO)
		}
	}
	return list
}

func getTreePath(parentId uint64) string {

	if parentId == 0 {
		return convert.ToString(parentId)
	} else {
		var order entity.Order
		utils.First(&entity.Order{ID: parentId}, &order)
		return order.TreePath + "," + convert.ToString(parentId)
	}

}
