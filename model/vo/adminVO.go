package vo

type Admin struct {
	// 主键
	ID uint64 `json:"id" mson:"UserId"`
	// 用户名
	UserName string `json:"username"`
	// 昵称
	NickName string `json:"nickname" mson:"NickName"`
	// 性别(1:男;2:女)
	Gender string `json:"genderLabel"`
	// 用户头像
	Avatar string `json:"avatar" mson:"Avatar"`
	// 联系方式
	Mobile string `json:"mobile"`
	// 用户状态(1:正常;2:禁用)
	Status int `json:"status" mson:"status"`
	// 邮箱
	Email string `json:"email"`
	// 部门名
	DeptName string `json:"deptName"`
	// 角色名
	RoleNames []string `json:"roleNames"`
	// 创建时间
	CreateTime string `json:"createTime" mson:"createTime"`
}
