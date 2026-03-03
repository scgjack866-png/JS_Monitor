package base

import "time"

type Model struct {
	// 主键
	ID uint64 `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT" mson:"UserId"`
	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`
}
