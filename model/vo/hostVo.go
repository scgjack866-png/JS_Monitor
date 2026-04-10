package vo

type HostVO struct {
	ID         uint64   `json:"id"`
	Name       string   `json:"name"`
	IpAddr     string   `json:"ipAddr"`
	OtherIp    []string `json:"otherIp"`
	IpNum      int      `json:"ipNum"`
	Status     *int     `json:"status"`
	GroupName  string   `json:"groupName"`
	Remark     string   `json:"remark"`
	IsAlter    *int     `json:"isAlter"`
	FlowIn     float64  `json:"flowIn"`
	FlowOut    float64  `json:"flowOut"`
	FlowInUnit  string   `json:"flowInUnit"`
	FlowOutUnit string   `json:"flowOutUnit"`
	CreateTime string   `json:"createTime"`
	Sort       *int     `json:"sort"`
}
