package entity

import (
	"time"
)

/**
 * 菜单实体对象
 *
 */
type Node struct {
	// 主键
	ID uint64 `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT"`
	// 监控点名称
	NodeIP string `gorm:"Column:node_ip"`
	// 监控点名称
	NodeName string `gorm:"Column:node_name"`
	// 开关(1-启用；0-停用)
	Status *int `gorm:"Column:status"`
	// 排列顺序
	Sort *int `gorm:"Column:sort"`
	// 逻辑删除标识(0-未删除；1-已删除)
	Deleted int `gorm:"Column:deleted"`

	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`
}

// 指定表名
func (Node) TableName() string {
	return "sys_node"
}
