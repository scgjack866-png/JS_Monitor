package entity

/**
 * 菜单实体对象
 *
 */
type UserRole struct {
	// 用户ID
	UserID uint64 `gorm:"Column:user_id;PRIMARY_KEY"`
	// 角色ID
	RoleID uint64 `gorm:"Column:role_id;PRIMARY_KEY"`
}

// 指定表名
func (UserRole) TableName() string {
	return "sys_user_role"
}
