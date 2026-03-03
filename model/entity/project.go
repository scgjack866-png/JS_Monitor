package entity

import "time"

type Project struct {
	// 主键
	ID uint64 `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT"`
	// 部门名称
	Name string `gorm:"Column:name"`
	// 父节点id
	ParentId *uint64 `gorm:"Column:parent_id" mson:"ParentId"`
	// 父节点id路径
	TreePath string `gorm:"Column:tree_path"`
	// 显示顺序
	Sort int `gorm:"Column:sort"`
	// 状态(1:正常;0:禁用)
	Status *int `gorm:"Column:status"`
	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`
	// 逻辑删除标识(0:未删除;1:已删除)
	Deleted int `gorm:"Column:deleted"`
}

// 指定表名
func (Project) TableName() string {
	return "sys_project"
}
