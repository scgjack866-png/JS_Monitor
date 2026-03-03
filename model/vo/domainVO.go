package vo

type DomainVO struct {
	// 主键
	ID uint64 `json:"id"`
	// 被墙域名
	Domain string `json:"domain"`
	// 监控开关(1-正常；0-停用)
	Status int `json:"status"`
	// 期望状态码
	Code int `json:"code"`
	// 关键词
	Keyword string `json:"keyword"`
	// 排列顺序
	Sort int `json:"sort"`

	// 创建时间
	CreateTime string `json:"createTime"`
}
