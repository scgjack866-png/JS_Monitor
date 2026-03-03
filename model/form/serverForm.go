package form

type ServerForm struct {
	ID        uint64 `json:"id"`
	IpAddr    string `json:"ipAddr"`
	ProjectID uint64 `json:"projectId"`
	RoomID    uint64 `json:"roomId"`
	ZoneID    uint64 `json:"zoneId"`
	OrderID   uint64 `json:"orderId"`
	Remark    string `json:"remark"`
	Sort      *int   `json:"sort"`
	Status    *int   `json:"status"`
}
