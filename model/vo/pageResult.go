package vo

type PageResult struct {
	Data  interface{} `json:"list"`
	Total int64       `json:"total"`
}
