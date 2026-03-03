package room

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
func recurDeptTreeOptions(parentID uint64, rooms []entity.Room) []vo.Option {
	var list []vo.Option

	for _, room := range rooms {

		if *room.ParentId == parentID {
			var option vo.Option
			option.Value = room.ID
			option.Value = room.ID
			option.Label = room.Name
			option.Children = recurDeptTreeOptions(room.ID, rooms)
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
func recurDept(parentID uint64, rooms []entity.Room) []vo.RoomVO {
	var list []vo.RoomVO

	for _, room := range rooms {

		if *room.ParentId == parentID {
			var roomVO vo.RoomVO
			copier.Copy(&roomVO, &room)
			roomVO.CreateTime = room.CreateTime.Format("2006-01-02")
			roomVO.UpdateTime = room.UpdateTime.Format("2006-01-02")
			roomVO.Children = recurDept(room.ID, rooms)
			list = append(list, roomVO)
		}
	}
	return list
}

func getTreePath(parentId uint64) string {

	if parentId == 0 {
		return convert.ToString(parentId)
	} else {
		var room entity.Room
		utils.First(&entity.Room{ID: parentId}, &room)
		return room.TreePath + "," + convert.ToString(parentId)
	}

}
