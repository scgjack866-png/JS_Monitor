package entity

import "time"

/**
 * 字典数据表
 *
 */
type Dict struct {
	// 主键
	ID int `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT"`
	// 字典项名称
	Name string `gorm:"Column:name"`
	// 字典项值
	Value string `gorm:"Column:value"`
	// 显示顺序
	Sort string `gorm:"Column:sort"`
	// 状态(1:正常;0:禁用)
	Status string `gorm:"Column:status"`
	// 是否默认(1:是;0:否)
	Defaulted int `gorm:"Column:defaulted"`
	// 备注
	Mark string `gorm:"Column:mark"`
	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`
}

// 指定表名
func (Dict) TableName() string {
	return "sys_dict"
}
