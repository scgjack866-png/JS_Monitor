package zone

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/vo"
	"OperationAndMonitoring/utils"
	"OperationAndMonitoring/utils/convert"
	copier "github.com/ybzhanghx/copier"
)

/**
 * 递归生成部门表格层级列表
 *
 * @param parentId
 * @param deptList
 * @return
 */
func recurDeptTreeOptions(parentID uint64, zones []entity.Zone) []vo.Option {
	var list []vo.Option

	for _, zone := range zones {

		if *zone.ParentId == parentID {
			var option vo.Option
			option.Value = zone.ID
			option.Value = zone.ID
			option.Label = zone.Name
			option.Children = recurDeptTreeOptions(zone.ID, zones)
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
func recurDept(parentID uint64, zones []entity.Zone) []vo.ZoneVO {
	var list []vo.ZoneVO

	for _, zone := range zones {

		if *zone.ParentId == parentID {
			var zoneVO vo.ZoneVO
			copier.Copy(&zoneVO, &zone)
			zoneVO.CreateTime = zone.CreateTime.Format("2006-01-02")
			zoneVO.UpdateTime = zone.UpdateTime.Format("2006-01-02")
			zoneVO.Children = recurDept(zone.ID, zones)
			list = append(list, zoneVO)
		}
	}
	return list
}

func getTreePath(parentId uint64) string {

	if parentId == 0 {
		return convert.ToString(parentId)
	} else {
		var zone entity.Zone
		utils.First(&entity.Zone{ID: parentId}, &zone)
		return zone.TreePath + "," + convert.ToString(parentId)
	}

}
