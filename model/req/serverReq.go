package req

type ServerReq struct {
	PageNum   int    `json:"pageNum" uri:"pageNum" form:"pageNum"`
	PageSize  int    `json:"pageSize" uri:"pageSize" form:"pageSize"`
	ZoneID    uint   `json:"zoneId" uri:"zoneId" form:"zoneId"`
	ProjectID uint   `json:"projectId" uri:"projectId" form:"projectId"`
	RoomID    uint   `json:"roomId" uri:"roomId" form:"roomId"`
	Status    string `json:"status" uri:"status" form:"status"`
	Keywords  string `json:"keywords" uri:"keywords" form:"keywords"`
}
