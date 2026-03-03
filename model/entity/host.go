package entity

import "time"

type Host struct {
	ID          uint64    `gorm:"Column:id;type:tinyint(11);PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UID         string    `gorm:"Column:uid;type:char(36);NOT NULL;default:0"`
	RuleUID     *string   `gorm:"Column:rule_uid;type:char(36);NOT NULL;default:0"`
	SilenceUID  *string   `gorm:"Column:silence_uid;type:char(36)"`
	Name        string    `gorm:"Column:name;type:char(20);default:''"`
	IpAddr      string    `gorm:"Column:ip_addr;type:char(15);NOT NULL"`
	Status      *int      `gorm:"Column:status;type:tinyint(1);NOT NULL;default:0"`
	GroupID     uint64    `gorm:"Column:group_id;type:tinyint(11)"`
	Remark      *string   `gorm:"Column:remark;type:char(255)"`
	IsAlter     *int      `gorm:"Column:is_alter;type:tinyint(1);NOT NULL;default:0"`
	FlowLimit   *int      `gorm:"Column:flow_limit;type:tinyint(11);NOT NULL;default:5"`
	LoadLimit   *int      `gorm:"Column:load_limit;type:tinyint(11);NOT NULL;default:5"`
	DiskLimit   *int      `gorm:"Column:disk_limit;type:tinyint(11);NOT NULL;default:80"`
	CpuLimit    *int      `gorm:"Column:cpu_limit;type:tinyint(11);NOT NULL;default:80"`
	MemLimit    *int      `gorm:"Column:mem_limit;type:tinyint(11);NOT NULL;default:80"`
	NetworkName string    `gorm:"Column:network_name;type:char(20)"`
	DelayTime   time.Time `gorm:"Column:delay_time;type:datetime"`
	Sort        *int      `gorm:"Column:sort;default:1"`
	MachineCode string    `gorm:"Column:machine_code"`
	AllIp       string    `gorm:"Column:all_ip"`
	CreateTime  time.Time `gorm:"Column:create_time;type:datetime;NOT NULL"`
	UpdateTime  time.Time `gorm:"Column:update_time;type:datetime;NOT NULL"`
	Deleted     int       `gorm:"Column:deleted;type:tinyint(1);NOT NULL;default:0"`
}

// 指定表名
func (Host) TableName() string {
	return "sys_host"
}
