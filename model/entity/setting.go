package entity

import "time"

type Setting struct {
	ID         uint64    `gorm:"Column:id;type:tinyint(11);PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name       string    `gorm:"Column:name;type:varchar(100);NOT NULL"`
	Value      string    `gorm:"Column:value;NOT NULL"`
	CreateTime time.Time `gorm:"Column:create_time;type:datetime;NOT NULL"`
	UpdateTime time.Time `gorm:"Column:update_time;type:datetime;NOT NULL"`
}

// 指定表名
func (Setting) TableName() string {
	return "sys_setting"
}
