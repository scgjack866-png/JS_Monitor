package vo

type HostVO struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	IpAddr     string `json:"ipAddr"`
	Status     *int   `json:"status"`
	GroupName  string `json:"groupName"`
	Remark     string `json:"remark"`
	IsAlter    *int   `json:"isAlter"`
	CreateTime string `json:"createTime"`
	Sort       *int   `json:"sort"`
}
