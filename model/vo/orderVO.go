package vo

type OrderVO struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	ParentId   uint64 `json:"parentId" mson:"ParentId"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`

	Children []OrderVO `json:"children"`
}
