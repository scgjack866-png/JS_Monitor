package vo

type RoomVO struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	ParentId   uint64 `json:"parentId" mson:"ParentId"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`

	Children []RoomVO `json:"children"`
}
