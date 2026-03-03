package entity

import "time"

/**
 * 字典类型实体
 *
 */
type DictType struct {
	// 主键
	ID int `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT"`
	// 类型名称
	Name string `gorm:"Column:name"`
	// 类型编码
	Code string `gorm:"Column:code"`
	// 状态(1:正常;0:禁用)
	Status string `gorm:"Column:status"`
	// 备注
	Remark string `gorm:"Column:remark"`
	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`
}

// 指定表名
func (DictType) TableName() string {
	return "sys_dict_type"
}
