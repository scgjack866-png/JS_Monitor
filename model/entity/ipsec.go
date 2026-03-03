package entity

import (
	"time"
)

/**
 * 菜单实体对象
 *
 */
type Ipsec struct {
	// 主键
	ID uint64 `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT"`
	// 主机名
	AgentHostname string `gorm:"Column:agent_hostname;NOT NULL"`
	// 被墙域名
	Domain string `gorm:"Column:domain"`
	// 期望状态码
	OnlineNum *int `gorm:"Column:online_num;default:0"`
	// 关键词
	AlterNum *int `gorm:"Column:alter_num;default:30""`
	// 监控开关(1-启用；0-停用)
	Status *int `gorm:"Column:status"`
	// 排列顺序
	Sort *int `gorm:"Column:sort"`
	// 排列顺序
	RuleUID *string `gorm:"Column:rule_uid;default:'';NOT NULL"`

	// 逻辑删除标识(0-未删除；1-已删除)
	Deleted int `gorm:"Column:deleted"`
	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`
}

// 指定表名
func (Ipsec) TableName() string {
	return "sys_ipsec_domain"
}
