package role

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/vo"
)

/**
 * 递归生成部门表格层级列表
 *
 * @param parentId
 * @param deptList
 * @return
 */
func recurDeptTreeOptions(roles []entity.Role) []vo.Option {
	var list []vo.Option

	for _, role := range roles {
		var option vo.Option
		option.Value = role.ID
		option.Label = role.Name
		list = append(list, option)
	}
	return list
}
