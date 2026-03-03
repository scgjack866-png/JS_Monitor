package entity

/**
 * 菜单实体对象
 *
 */
type NodeDomain struct {
	// 节点ID
	NodeID uint64 `gorm:"Column:node_id;PRIMARY_KEY"`
	// 域名ID
	DomainID uint64 `gorm:"Column:domain_id;PRIMARY_KEY"`
	// 规则ID
	RuleID string `gorm:"Column:rule_id;PRIMARY_KEY"`
}

// 指定表名
func (NodeDomain) TableName() string {
	return "sys_node_domain"
}
