package entity

import (
	"time"
)

/**
 * 菜单实体对象
 *
 */
type Role struct {
	// 主键
	ID uint64 `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT"`
	// 角色名称
	Name string `gorm:"Column:name"`
	// 角色编码
	Code string `gorm:"Column:code"`
	// 显示顺序
	Sort string `gorm:"Column:sort"`
	// 角色状态(1-正常；0-停用)
	Status string `gorm:"Column:status"`
	// 逻辑删除标识(0-未删除；1-已删除)
	Deleted string `gorm:"Column:deleted"`
	// 数据权限
	DataScope string `gorm:"Column:data_scope"`

	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`

	Users []User `gorm:"many2many:sys_user_role"`
	Menus []Menu `gorm:"many2many:sys_role_menu"`
}

// 指定表名
func (Role) TableName() string {
	return "sys_role"
}
