package form

type DeptForm struct {
	ID       uint64 `json:"id"`
	ParentId uint64 `json:"parentId" mson:"ParentId"`
	Status   int    `json:"status"`
	Sort     int    `json:"sort"`
	Name     string `json:"name"`
}
