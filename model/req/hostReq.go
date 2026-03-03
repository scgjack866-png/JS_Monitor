package req

type HostReq struct {
	PageNum  int    `json:"pageNum" uri:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" uri:"pageSize" form:"pageSize"`
	GroupID  uint   `json:"groupId" uri:"groupId" form:"groupId"`
	Status   string `json:"status" uri:"status" form:"status"`
	Keywords string `json:"keywords" uri:"keywords" form:"keywords"`
}
