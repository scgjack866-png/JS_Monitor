package req

type FlowReq struct {
	Ips []string `json:"ips" uri:"ips" form:"ips"`
}

type FlowSnapshotReq struct {
	Ip string `json:"ip" uri:"ip" form:"ip"`
}
