package vo

type IpsecVO struct {
	ID            uint64 `json:"id"`
	AgentHostname string `json:"agent_hostname"`
	Domain        string `json:"domain"`
	OnlineNum     *int   `json:"onlineNum"`
	AlterNum      *int   `json:"alterNum"`
	Status        *int   `json:"status"`
	CreateTime    string `json:"createTime"`
	Sort          *int   `json:"sort"`
}
