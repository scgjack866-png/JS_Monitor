package vo

type ServerVO struct {
	ID          uint64   `json:"id"`
	MainIp      string   `json:"mainIp"`
	IpAddr      []string `json:"ipAddr"`
	IpNum       int      `json:"ipNum"`
	Status      *int     `json:"status"`
	ZoneName    string   `json:"zoneName"`
	ProjectName string   `json:"projectName"`
	RoomName    string   `json:"roomName"`
	OrderName   string   `json:"orderName"`
	Remark      string   `json:"remark"`
	CreateTime  string   `json:"createTime"`
	Sort        *int     `json:"sort"`
}
