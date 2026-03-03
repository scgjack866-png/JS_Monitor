package form

type IpsecForm struct {
	ID            uint64 `json:"id"`
	AgentHostname string `json:"agent_hostname"`
	Domain        string `json:"domain"`
	AlterNum      *int   `json:"alterNum"`
	Sort          *int   `json:"sort"`
}
