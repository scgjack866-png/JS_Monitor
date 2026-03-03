package project

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
func recurDeptTreeOptions(parentID uint64, projects []entity.Project) []vo.Option {
	var list []vo.Option

	for _, project := range projects {

		if *project.ParentId == parentID {
			var option vo.Option
			option.Value = project.ID
			option.Value = project.ID
			option.Label = project.Name
			option.Children = recurDeptTreeOptions(project.ID, projects)
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
func recurDept(parentID uint64, projects []entity.Project) []vo.ProjectVO {
	var list []vo.ProjectVO

	for _, project := range projects {

		if *project.ParentId == parentID {
			var projectVO vo.ProjectVO
			copier.Copy(&projectVO, &project)
			projectVO.CreateTime = project.CreateTime.Format("2006-01-02")
			projectVO.UpdateTime = project.UpdateTime.Format("2006-01-02")
			projectVO.Children = recurDept(project.ID, projects)
			list = append(list, projectVO)
		}
	}
	return list
}

func getTreePath(parentId uint64) string {

	if parentId == 0 {
		return convert.ToString(parentId)
	} else {
		var project entity.Project
		utils.First(&entity.Project{ID: parentId}, &project)
		return project.TreePath + "," + convert.ToString(parentId)
	}

}
