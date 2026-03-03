package entity

import "time"

/**
 * 菜单实体对象
 *
 */
type Menu struct {
	// 主键
	ID uint64 `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT"`
	// 父菜单ID
	ParentId uint64 `gorm:"Column:parent_id"`
	// 菜单名称
	Name string `gorm:"Column:name"`
	// 菜单类型(1-菜单；2-目录；3-外链；4-按钮权限)
	MenuType int `gorm:"Column:type"`
	// 路由路径(浏览器地址栏路径)
	Path string `gorm:"Column:path"`
	// 组件路径(vue页面完整路径，省略.vue后缀)
	Component string `gorm:"Column:component"`
	// 权限标识
	Perm string `gorm:"Column:perm"`
	// 显示状态(1:显示;0:隐藏)
	Visible bool `gorm:"Column:visible"`
	// 排序
	Sort string `gorm:"Column:sort"`
	// 菜单图标
	Icon string `gorm:"Column:icon"`
	// 跳转路径
	Redirect string `gorm:"Column:redirect"`

	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`

	Roles []Role `gorm:"many2many:sys_role_menu"`
}

// 指定表名
func (Menu) TableName() string {
	return "sys_menu"
}
