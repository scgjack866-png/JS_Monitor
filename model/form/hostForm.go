package form

type HostForm struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	IpAddr      string `json:"ipAddr"`
	GroupID     uint64 `json:"groupId"`
	FlowLimit   int    `json:"flowLimit"`
	LoadLimit   int    `json:"loadLimit"`
	DiskLimit   int    `json:"diskLimit"`
	CpuLimit    int    `json:"cpuLimit"`
	MemLimit    int    `json:"memLimit"`
	NetworkName string `json:"networkName"`
	DelayTime   string `json:"delayTime"`
	Remark      string `json:"remark"`
	Sort        *int   `json:"sort"`
	MachineCode string `json:"machine_code"`
	AllIp       string `json:"all_ip"`
	IsAlter     *int   `json:"is_alter"`
}

type HostFormNoGroup struct {
	ID        uint64 `json:"id"`
	IpAddr    string `json:"ipAddr"`
	FlowLimit int    `json:"flowLimit"`
	LoadLimit int    `json:"loadLimit"`
	DiskLimit int    `json:"diskLimit"`
	CpuLimit  int    `json:"cpuLimit"`
	MemLimit  int    `json:"memLimit"`
}
