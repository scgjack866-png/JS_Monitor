package entity

import "time"

type Server struct {
	ID         uint64    `gorm:"Column:id;type:tinyint(11);PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	IpAddr     string    `gorm:"Column:ip_addr;type:text;NOT NULL"`
	Status     *int      `gorm:"Column:status;type:tinyint(1);NOT NULL;default:0"`
	ZoneID     uint64    `gorm:"Column:zone_id;NOT NULL;type:tinyint(11)"`
	ProjectID  uint64    `gorm:"Column:project_id;NOT NULL;type:tinyint(11)"`
	RoomID     uint64    `gorm:"Column:room_id;NOT NULL;type:tinyint(11)"`
	OrderID    uint64    `gorm:"Column:order_id;NOT NULL;type:tinyint(11)"`
	Remark     *string   `gorm:"Column:remark;type:text"`
	Sort       *int      `gorm:"Column:sort;default:1"`
	AllIp      string    `gorm:"Column:all_ip;type:text"`
	CreateTime time.Time `gorm:"Column:create_time;type:datetime;NOT NULL"`
	UpdateTime time.Time `gorm:"Column:update_time;type:datetime;NOT NULL"`
	Deleted    int       `gorm:"Column:deleted;type:tinyint(1);NOT NULL;default:0"`
}

// 指定表名
func (Server) TableName() string {
	return "sys_server"
}
