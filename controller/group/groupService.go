package group

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
func recurDeptTreeOptions(parentID uint64, groups []entity.Group) []vo.Option {
	var list []vo.Option

	for _, group := range groups {

		if *group.ParentId == parentID {
			var option vo.Option
			option.Value = group.ID
			option.Label = group.Name
			option.Children = recurDeptTreeOptions(group.ID, groups)
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
func recurDept(parentID uint64, groups []entity.Group) []vo.GroupVO {
	var list []vo.GroupVO

	for _, group := range groups {

		if *group.ParentId == parentID {
			var groupVO vo.GroupVO
			copier.Copy(&groupVO, &group)
			groupVO.CreateTime = group.CreateTime.Format("2006-01-02")
			groupVO.UpdateTime = group.UpdateTime.Format("2006-01-02")
			groupVO.Children = recurDept(group.ID, groups)
			list = append(list, groupVO)
		}
	}
	return list
}

func getTreePath(parentId uint64) string {

	if parentId == 0 {
		return convert.ToString(parentId)
	} else {
		var group entity.Group
		utils.First(&entity.Group{ID: parentId}, &group)
		return group.TreePath + "," + convert.ToString(parentId)
	}

}
