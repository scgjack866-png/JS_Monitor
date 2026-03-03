package vo

/**
 * 菜单实体对象
 *
 */
type NodeVO struct {
	// 主键
	ID uint64 `json:"id"`
	// 监控点名称
	NodeIP string `json:"nodeIP"`
	// 监控点名称
	NodeName string `json:"nodeName"`
	// 开关(1-启用；0-停用)
	Status int `json:"status"`
	// 排列顺序
	Sort int `json:"sort"`

	// 创建时间
	CreateTime string `json:"createTime"`
}
