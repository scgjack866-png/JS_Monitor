package dept

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
func recurDeptTreeOptions(parentID uint64, depts []entity.Dept) []vo.Option {
	var list []vo.Option

	for _, dept := range depts {

		if *dept.ParentId == parentID {
			var option vo.Option
			option.Value = dept.ID
			option.Value = dept.ID
			option.Label = dept.Name
			option.Children = recurDeptTreeOptions(dept.ID, depts)
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
func recurDept(parentID uint64, depts []entity.Dept) []vo.DeptVO {
	var list []vo.DeptVO

	for _, dept := range depts {

		if *dept.ParentId == parentID {
			var deptVO vo.DeptVO
			copier.Copy(&deptVO, &dept)
			deptVO.CreateTime = dept.CreateTime.Format("2006-01-02")
			deptVO.UpdateTime = dept.UpdateTime.Format("2006-01-02")
			deptVO.Children = recurDept(dept.ID, depts)
			list = append(list, deptVO)
		}
	}
	return list
}

func getTreePath(parentId uint64) string {

	if parentId == 0 {
		return convert.ToString(parentId)
	} else {
		var dept entity.Dept
		utils.First(&entity.Dept{ID: parentId}, &dept)
		return dept.TreePath + "," + convert.ToString(parentId)
	}

}
