package entity

/**
 * 菜单实体对象
 *
 */
type RoleMenu struct {
	// 角色ID
	RoleID uint64 `gorm:"Column:role_id;PRIMARY_KEY"`
	// 菜单ID
	MenuID uint64 `gorm:"Column:menu_id;PRIMARY_KEY"`
}

// 指定表名
func (RoleMenu) TableName() string {
	return "sys_role_menu"
}
